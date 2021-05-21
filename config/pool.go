package config

import (
	"io"
	"time"
)

type Pool struct {
	Listen                  Listen
	User                    string
	Group                   string
	ProcessManager          ProcessManager
	Ping                    Ping
	Process                 *Process
	Prefix                  string
	RequestTerminateTimeout time.Duration
	RequestSlowLogTimeout   time.Duration
	SlowLog                 string
	RlimitFiles             int
	RlimitCore              int
	Chroot                  string
	CatchWorkersOutput      bool
	DecorateWorkersOutput   bool
	ClearEnv                bool
	Security                Security
	Access                  Access
}

func (p *Pool) write(w io.Writer) error {
	return nil
}
