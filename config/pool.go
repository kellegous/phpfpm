package config

import (
	"errors"
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

func (p *Pool) validate() error {
	// TODO(knorton): validate listen

	if err := p.ProcessManager.validate(); err != nil {
		return err
	}

	if p.User == "" {
		return errors.New("user is required")
	}

	return nil
}

func (p *Pool) write(w io.Writer) error {
	if err := p.Listen.write(w); err != nil {
		return err
	}

	if err := writeString(w, "user", p.User); err != nil {
		return err
	}

	if v := p.Group; v != "" {
		if err := writeString(w, "group", v); err != nil {
			return err
		}
	}

	return nil
}
