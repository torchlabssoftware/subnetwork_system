package manager

import (
	"github.com/google/uuid"
)

type Pool struct {
	PoolId        uuid.UUID
	PoolTag       string
	PoolPort      int
	PoolSubdomain string
	Upstreams     []Upstream
}

type Upstream struct {
	UpstreamID       uuid.UUID
	UpstreamTag      string
	UpstreamFormat   string
	UpstreamUsername string
	UpstreamPassword string
	UpstreamHost     string
	UpstreamPort     int
	UpstreamProvider string
	Weight           int
}

func NewPool(poolId uuid.UUID, poolTag string, poolPort int, poolSubdomain string, upstreams []Upstream) *Pool {
	pool := &Pool{
		PoolId:        poolId,
		PoolTag:       poolTag,
		PoolPort:      poolPort,
		PoolSubdomain: poolSubdomain,
		Upstreams:     upstreams,
	}

	return pool

}
