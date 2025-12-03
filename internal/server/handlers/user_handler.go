package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/torchlabssoftware/subnetwork_system/internal/db/repository"
	functions "github.com/torchlabssoftware/subnetwork_system/internal/server/functions"
	server "github.com/torchlabssoftware/subnetwork_system/internal/server/models"
)

type UserHandler struct {
	queries *repository.Queries
	db      *sql.DB
}

func NewUserHandler(q *repository.Queries, db *sql.DB) *UserHandler {
	return &UserHandler{
		queries: q,
		db:      db,
	}
}

func (h *UserHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.CreateUser)
	r.Get("/", h.GetUserbyId)
	return r
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	//begin transaction
	ctx, err := h.db.Begin()
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}
	defer func() {
		_ = ctx.Rollback()
	}()

	qtx := h.queries.WithTx(ctx)

	//get responce
	var req server.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "bad request in create user", err)
		return
	}

	//validate email
	var email sql.NullString
	if req.Email != nil && *req.Email != "" {
		mail, err := mail.ParseAddress(*req.Email)
		if err != nil {
			functions.RespondwithError(w, http.StatusBadRequest, "bad request in create user", err)
			return
		}
		email = sql.NullString{String: mail.String(), Valid: true}
	} else {
		email = sql.NullString{Valid: false}
	}

	//validate datalimit
	dataLimit := int64(0)
	if req.DataLimit != nil && *req.DataLimit >= int64(0) {
		dataLimit = *req.DataLimit
	} else {
		functions.RespondwithError(w, http.StatusBadRequest, "send valid data limit", fmt.Errorf("send valid data limit"))
		return
	}

	//insert user data
	createUserParams := repository.CreateUserParams{
		Email:     email,
		Username:  uuid.New().String()[:8],
		Password:  uuid.New().String()[:8],
		DataLimit: dataLimit,
	}

	user, err := qtx.CreateUser(r.Context(), createUserParams)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	//insert allow_pool data
	var allowPools []string
	if req.AllowPools != nil && len(*req.AllowPools) > 0 {
		allowPools = *req.AllowPools
	}

	pools, err := qtx.GetPoolsbyTags(r.Context(), allowPools)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	insertUserPoolParams := repository.InsertUserPoolParams{
		UserID:  user.ID,
		Column2: pools,
	}

	_, err = qtx.InsertUserPool(r.Context(), insertUserPoolParams)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	//insert ipwhilist data
	var ipWhitelist []string
	if req.IpWhiteList != nil && len(*req.IpWhiteList) > 0 {
		ipWhitelist = *req.IpWhiteList
	}

	userIpWhitelistParams := repository.InsertUserIpwhitelistParams{
		UserID:  user.ID,
		Column2: ipWhitelist,
	}

	_, err = qtx.InsertUserIpwhitelist(r.Context(), userIpWhitelistParams)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	if err := ctx.Commit(); err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	responce := server.CreateUserResponce{
		Id:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		DataLimit:   user.DataLimit,
		IpWhitelist: ipWhitelist,
		AllowPools:  allowPools,
	}

	functions.RespondwithJSON(w, http.StatusCreated, responce)
}

func (h *UserHandler) GetUserbyId(w http.ResponseWriter, r *http.Request) {
	//begin transaction
	ctx, err := h.db.Begin()
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}
	defer func() {
		_ = ctx.Rollback()
	}()
	qtx := h.queries.WithTx(ctx)

	//get user params and get user data
	queryParams := r.URL.Query()
	userId := queryParams.Get("user-id")
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "incorrect user id", err)
		return
	}

	user, err := qtx.GetUserbyId(r.Context(), userIdUUID)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "cant get user by id", err)
		return
	}

	userPools, err := qtx.GetUserPoolByUserId(r.Context(), user.ID)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "cant get user by id", err)
		return
	}

	ipWhitelist, err := qtx.GetIpWhitelistByUserId(r.Context(), user.ID)
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "cant get user by id", err)
		return
	}

	resp := server.GetUserByIdResponce{
		Id:          user.ID,
		Email:       user.Email.String,
		Username:    user.Username,
		Password:    user.Password,
		Data_limit:  user.DataLimit,
		Data_usage:  user.DataUsage,
		Status:      user.Status,
		IpWhitelist: ipWhitelist,
		UserPool:    userPools,
		Created_at:  user.CreatedAt,
		Updated_at:  user.UpdatedAt,
	}

	if err := ctx.Commit(); err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}

	functions.RespondwithJSON(w, http.StatusInternalServerError, resp)
}
