package service

import "github.com/silenceper/pool"

// NewHeadlessClientPool creates HeadlessClientPool
func NewHeadlessClientPool(cfg *HeadlessConfig) (pool.Pool, error) {
	factory := func() (interface{}, error) {
		headlessClient, err := NewHeadlessClient(cfg)
		if err != nil {
			return nil, err
		}

		return headlessClient, nil
	}

	close := func(v interface{}) error {
		return v.(*HeadlessClient).Close()
	}

	poolConfig := &pool.Config{
		InitialCap: 15,
		MaxCap:     30,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 0,
	}

	return pool.NewChannelPool(poolConfig)
}