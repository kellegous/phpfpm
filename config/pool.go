package config

import (
	"errors"
	"time"
)

type Pool struct {
	Name                    string
	Listen                  Listen
	User                    string
	Group                   string
	ProcessManager          ProcessManager
	Ping                    Ping
	Process                 PoolProcess
	Prefix                  string
	RequestTerminateTimeout time.Duration
	RequestSlowLogTimeout   time.Duration
	SlowLog                 string
	RlimitFiles             int
	RlimitCore              int
	Chroot                  string
	CatchWorkersOutput      bool
	DecorateWorkersOutput   *bool
	ClearEnv                *bool
	Security                Security
	Access                  Access
}

func (p *Pool) WithGroup(v string) *Pool {
	p.Group = v
	return p
}

func (p *Pool) WithPrefix(v string) *Pool {
	p.Prefix = v
	return p
}

func (p *Pool) WithRequestTerminateTimeout(v time.Duration) *Pool {
	p.RequestSlowLogTimeout = v
	return p
}

func (p *Pool) WithRequestSlowLogTimeout(v time.Duration) *Pool {
	p.RequestSlowLogTimeout = v
	return p
}

func (p *Pool) WithSlowLog(v string) *Pool {
	p.SlowLog = v
	return p
}

func (p *Pool) WithRlimitFiles(v int) *Pool {
	p.RlimitFiles = v
	return p
}

func (p *Pool) WithRlimitCore(v int) *Pool {
	p.RlimitCore = v
	return p
}

func (p *Pool) WithChroot(v string) *Pool {
	p.Chroot = v
	return p
}

func (p *Pool) WithCatchWorkersOutput(v bool) *Pool {
	p.CatchWorkersOutput = v
	return p
}

func (p *Pool) WithDecorateWorkersOutput(v bool) *Pool {
	p.DecorateWorkersOutput = &v
	return p
}

func (p *Pool) WithClearEnv(v bool) *Pool {
	p.ClearEnv = &v
	return p
}

func (p *Pool) validate() error {
	if err := p.Listen.validate(); err != nil {
		return err
	}

	if p.User == "" {
		return errors.New("user is required")
	}

	if err := p.ProcessManager.validate(); err != nil {
		return err
	}

	if err := p.Ping.validate(); err != nil {
		return err
	}

	return nil
}

func (p *Pool) write(w *writer) error {
	if err := p.Listen.write(w); err != nil {
		return err
	}

	if err := w.writeString("user", p.User); err != nil {
		return err
	}

	if v := p.Group; v != "" {
		if err := w.writeString("group", v); err != nil {
			return err
		}
	}

	if err := p.ProcessManager.write(w); err != nil {
		return err
	}

	if err := p.Ping.write(w); err != nil {
		return err
	}

	if err := p.Process.write(w); err != nil {
		return err
	}

	if v := p.Prefix; v != "" {
		if err := w.writeString("prefix", v); err != nil {
			return err
		}
	}

	if v := p.RequestTerminateTimeout; v != 0 {
		if err := w.writeDuration("request_terminate_timeout", v); err != nil {
			return err
		}
	}

	if v := p.RequestSlowLogTimeout; v != 0 {
		if err := w.writeDuration("request_slowlog_timeout", v); err != nil {
			return err
		}
	}

	if v := p.SlowLog; v != "" {
		if err := w.writeString("slowlog", v); err != nil {
			return err
		}
	}

	if v := p.RlimitFiles; v != 0 {
		if err := w.writeInt("rlimit_files", v); err != nil {
			return err
		}
	}

	if v := p.RlimitCore; v != 0 {
		if err := w.writeInt("rlimit_core", v); err != nil {
			return err
		}
	}

	if v := p.Chroot; v != "" {
		if err := w.writeString("chroot", v); err != nil {
			return err
		}
	}

	if v := p.CatchWorkersOutput; v {
		if err := w.writeBool("catch_workers_output", v); err != nil {
			return err
		}
	}

	if v := p.DecorateWorkersOutput; v != nil {
		if err := w.writeBool("decorate_workers_output", *v); err != nil {
			return err
		}
	}

	if v := p.ClearEnv; v != nil {
		if err := w.writeBool("clear_env", *v); err != nil {
			return err
		}
	}

	if err := p.Security.write(w); err != nil {
		return err
	}

	if err := p.Access.write(w); err != nil {
		return err
	}

	return nil
}
