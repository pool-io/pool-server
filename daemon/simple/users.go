package simple

import (
	. "github.com/pool-beta/pool-server/types"
	puser "github.com/pool-beta/pool-server/user"
	. "github.com/pool-beta/pool-server/user/types"
)

/*
	Implements Simple Users
*/

type users struct {
	pools Pools
	uf    puser.UserFactory
}

type user struct {
	pools Pools
	user  puser.User
}

func InitUsers(pools Pools) (Users, error) {
	// Start UserFactory
	uf, err := puser.InitUserFactory()
	if err != nil {
		return nil, err
	}

	return &users{
		pools: pools,
		uf:    uf,
	}, nil
}

func (us *users) CleanUp() error {
	return nil
}

func (us *users) CreateUser(name UserName, password string, amount USDollar) (User, error) {
	u, err := us.uf.CreateUser(name, password, amount)
	if err != nil {
		return nil, err
	}

	return &user{
		pools: us.pools,
		user:  u,
	}, nil
}

func (us *users) GetUser(name UserName, password string) (User, error) {
	u, err := us.uf.RetrieveUser(name, password)
	if err != nil {
		return nil, err
	}

	return &user{
		pools: us.pools,
		user:  u,
	}, nil
}

func (us *users) RemoveUser(name UserName, password string) error {
	err := us.uf.RemoveUser(name, password)
	return err
}

// Testing
func (us *users) GetAllUserNames() ([]UserName, error) {
	return us.uf.RetreieveAllUserNames()
}

func (u *user) ID() UserID {
	return u.user.GetID()
}

func (u *user) UserName() UserName {
	return u.user.GetUserName()
}

/* Add Pool */
func (u *user) AddTank(tank Tank) error {
	return u.user.AddTank(tank.Name(), tank.ID())
}

func (u *user) AddPool(pool Pool) error {
	return u.user.AddPool(pool.Name(), pool.ID())
}

func (u *user) AddDrain(drain Drain) error {
	return u.user.AddDrain(drain.Name(), drain.ID())
}

/* Get Pool */
func (u *user) GetTank(name string) (Tank, error) {
	return u.pools.GetPool(u.user.GetTank(name))
}

func (u *user) GetPool(name string) (Pool, error) {
	return u.pools.GetPool(u.user.GetPool(name))
}

func (u *user) GetDrain(name string) (Drain, error) {
	return u.pools.GetPool(u.user.GetDrain(name))
}

/* Get Pools */

func (u *user) GetTanks() ([]Tank, error) {
	pids := u.user.GetTanks()

	// Create simple tanks to return to return
	tanks := make([]Tank, 0)

	for _, pid := range pids {
		tank, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		tanks= append(tanks, tank)
	}

	return tanks, nil
}

func (u *user) GetPools() ([]Pool, error) {
	pids := u.user.GetPools()

	// Create simple tanks to return to return
	pools := make([]Pool, 0)

	for _, pid := range pids {
		pool, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		pools = append(pools, pool)
	}

	return pools, nil
}

func (u *user) GetDrains() ([]Drain, error) {
	pids := u.user.GetDrains()

	// Create simple tanks to return to return
	drains := make([]Drain, 0)

	for _, pid := range pids {
		drain, err := u.pools.GetPool(pid)
		if err != nil {
			return nil, err
		}

		drains = append(drains, drain)
	}

	return drains, nil
}

func (u *user) GetFlows() []Flow {
	return nil
}

func (u *user) CleanUp() error {
	return nil
}
