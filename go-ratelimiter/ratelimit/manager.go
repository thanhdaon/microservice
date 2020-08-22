package ratelimit

type Manager struct {
	errorChan    chan error
	releaseChan  chan *Token
	outChan      chan *Token
	inChan       chan struct{}
	needToken    int64
	activeTokens map[string]*Token
	limit        int
	makeToken    tokenFactory
}

func NewManager(conf *Config) *Manager {
	return &Manager{
		errorChan:    make(chan error),
		outChan:      make(chan *Token),
		inChan:       make(chan struct{}),
		activeTokens: make(map[string]*Token),
		releaseChan:  make(chan *Token),
		needToken:    0,
		limit:        conf.Limit,
		makeToken:    NewToken,
	}
}

func (m *Manager) Acquire() (*Token, error) {
	go func() {
		m.inChan <- struct{}{}
	}()

	select {
	case t := <-m.outChan:
		return t, nil
	case err := <-m.errorChan:
		return nil, err
	}
}

func (m *Manager) Release(t *Token) {
	go func() {
		m.releaseChan <- t
	}()

}

func (m *Manager) tryGenerateToken() {
	if m.makeToken == nil {
		panic("ErrTokenFactoryNotDefined")
	}

	go func() {
		m.outChan <- m.makeToken()
	}()
}
