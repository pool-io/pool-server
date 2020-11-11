package simple

import (
	ppool "github.com/pool-beta/pool-server/pool"
	. "github.com/pool-beta/pool-server/pool/types"
)

/*
	Implements POOL (simple pools)
*/

type pools struct {
	pf ppool.PoolFactory
}

// Implements Pool/Drain/Tank
type pool struct {
	pool ppool.Pool
}

func InitPools() (Pools, error) {
	pf, err := ppool.NewPoolFactory()
	if err != nil {
		return nil, err
	}

	return &pools{
		pf: pf,
	}, nil
}

func (ps *pools) CreatePool(user User, name string) (Pool, error) {
	p, err := ps.pf.CreatePool(name, user.ID(), POOL)
	if err != nil {
		return nil, err
	}
	
	return &pool{
		pool: p,
	}, nil
}

func (ps *pools) CreateDrainPool(user User, name string) (Drain, error) {
	d, err := ps.pf.CreatePool(name, user.ID(), DRAIN)
	if err != nil {
		return nil, err
	}

	return &pool{
		pool: d,
	}, nil
}

func (ps *pools) CreateTankPool(user User, name string) (Tank, error) {
	t, err := ps.pf.CreatePool(name, user.ID(), TANK)
	if err != nil {
		return nil, err
	}

	return &pool{
		pool: t,
	}, nil
}

// Returns the pool with the pid; could be pool, drain, or tank
func (ps *pools) GetPool(pid PoolID) (Pool, error) {
	return nil, nil
}

func (ps *pools) RemovePool(pid PoolID) error {
	
	return nil
}

func (ps *pools) CleanUp() error {
	return nil
}

// -------------------------------------------------------------------------------------------------

func (p *pool) CreateStream(puller Pool) (Stream, error) {
	return nil, nil
}

func (p *pool) GetStream(sid StreamID) (Stream, error) {
	return nil, nil
}

func (p *pool) RemoveStream(sid StreamID) error {
	return nil
}

func (p *pool) CreateFlow() (Flow, error) {
	return nil, nil
}

func (p *pool) CleanUp() error {
	return nil
}