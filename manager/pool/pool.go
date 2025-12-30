package pool

import "github.com/google/uuid"

type Pool struct {
	PoolId        uuid.UUID
	PoolTag       string
	PoolPort      int
	PoolSubdomain string
	Upstreams     map[string]*Upstream
}

type Upstream struct {
	UpstreamID      uuid.UUID
	UpstreamTag     string
	UpstreamAddress string
	UpstreamHost    string
	UpstreamPort    int
	Weight          int
}

func NewPool(poolId uuid.UUID, poolTag string, poolPort int, poolSubdomain string, upstreams []Upstream) *Pool {
	return &Pool{
		PoolId:        poolId,
		PoolTag:       poolTag,
		PoolPort:      poolPort,
		PoolSubdomain: poolSubdomain,
		Upstreams:     make(map[string]*Upstream),
	}
}
