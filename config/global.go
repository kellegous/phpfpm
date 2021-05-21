package config

import (
	"errors"
	"fmt"
	"io"
	"time"
)

type LogLevel string

const (
	LevelAlert   = "alert"
	LevelError   = "error"
	LevelWarning = "warning"
	LevelNotice  = "notice"
	LevelDebug   = "debug"
)

type Global struct {
	PIDFile                   string
	ErrorLog                  string
	LogLevel                  LogLevel
	LogLimit                  int
	LogBuffering              *bool
	EmergencyRestartThreshold int
	ProcessControlTimeout     time.Duration
	Process                   *Process
	RLimitFiles               int
	RlimitCore                int
	Pools                     []*Pool
}

func New() *Global {
	return &Global{}
}

func (g *Global) WithPIDFile(v string) *Global {
	g.PIDFile = v
	return g
}

func (g *Global) WithErrorLog(v string) *Global {
	g.ErrorLog = v
	return g
}

func (g *Global) WithLogLevel(v LogLevel) *Global {
	g.LogLevel = v
	return g
}

func (g *Global) WithLogLimit(v int) *Global {
	g.LogLimit = v
	return g
}

func (g *Global) WithLogBuffering(v bool) *Global {
	g.LogBuffering = &v
	return g
}

func (g *Global) WithEmergencyRestartTheshold(v int) *Global {
	g.EmergencyRestartThreshold = v
	return g
}

func (g *Global) WithProcessControlTimeout(v time.Duration) *Global {
	g.ProcessControlTimeout = v
	return g
}
func (g *Global) WithPool(
	name string,
	addr string,
	user string,
	pmType ProcessManagerType,
	maxChildren int,
) *Pool {
	p := &Pool{
		Name: name,
		Listen: Listen{
			Addr: addr,
		},
		User: user,
		ProcessManager: ProcessManager{
			Type:        pmType,
			MaxChildren: maxChildren,
		},
	}
	g.Pools = append(g.Pools, p)
	return p
}

func (g *Global) validate() error {
	if len(g.Pools) == 0 {
		return errors.New("must have at least one pool")
	}

	seen := map[string]bool{}
	for _, pool := range g.Pools {
		if seen[pool.Name] {
			return fmt.Errorf("duplicate pool: %s", pool.Name)
		}
		seen[pool.Name] = true

		if err := pool.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Global) Write(w io.Writer) error {
	return g.write(&writer{w})
}

func (g *Global) write(w *writer) error {
	if err := g.validate(); err != nil {
		return err
	}

	if v := g.PIDFile; v != "" {
		if err := w.writeString("pid", v); err != nil {
			return err
		}
	}

	if v := g.ErrorLog; v != "" {
		if err := w.writeString("error_log", v); err != nil {
			return err
		}
	}

	if v := g.LogLevel; v != "" {
		if err := w.writeString("log_level", string(v)); err != nil {
			return err
		}
	}

	if v := g.LogLimit; v != 0 {
		if err := w.writeInt("log_limit", v); err != nil {
			return err
		}
	}

	if v := g.LogBuffering; v != nil {
		if err := w.writeBool("log_buffering", *v); err != nil {
			return err
		}
	}

	if v := g.EmergencyRestartThreshold; v != 0 {
		if err := w.writeInt("emergency_restart_threshold", v); err != nil {
			return err
		}
	}

	if v := g.ProcessControlTimeout; v != 0 {
		if err := w.writeDuration("process_control_timeout", v); err != nil {
			return err
		}
	}

	if p := g.Process; p != nil {
		if err := p.write(w); err != nil {
			return err
		}
	}

	if v := g.RLimitFiles; v != 0 {
		if err := w.writeInt("rlimit_files", v); err != nil {
			return err
		}
	}

	if v := g.RlimitCore; v != 0 {
		if err := w.writeInt("rlimit_core", v); err != nil {
			return err
		}
	}

	for _, pool := range g.Pools {
		if err := w.writeSection(pool.Name); err != nil {
			return err
		}

		if err := pool.write(w); err != nil {
			return err
		}
	}

	return nil
}
