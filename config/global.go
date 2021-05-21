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
	Pools                     map[string]*Pool
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
		Listen: Listen{
			Addr: addr,
		},
		User: user,
		ProcessManager: ProcessManager{
			Type:        pmType,
			MaxChildren: maxChildren,
		},
	}

	g.Pools[name] = p

	return p
}

func writeString(w io.Writer, key, val string) error {
	_, err := fmt.Fprintf(w, "%s = %s\n", key, val)
	return err
}

func writeInt(w io.Writer, key string, val int) error {
	_, err := fmt.Fprintf(w, "%s = %d\n", key, val)
	return err
}

func writeBool(w io.Writer, key string, val bool) error {
	s := "no"
	if val {
		s = "yes"
	}

	_, err := fmt.Fprintf(w, "%s = %s\n", key, s)
	return err
}

func writeSection(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "[%s]\n", name)
	return err
}

func (g *Global) validate() error {
	if len(g.Pools) == 0 {
		return errors.New("must have at least one pool")
	}

	for _, pool := range g.Pools {
		if err := pool.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (g *Global) Write(w io.Writer) error {
	if err := g.validate(); err != nil {
		return err
	}

	if v := g.PIDFile; v != "" {
		if err := writeString(w, "pid", v); err != nil {
			return err
		}
	}

	if v := g.ErrorLog; v != "" {
		if err := writeString(w, "error_log", v); err != nil {
			return err
		}
	}

	if v := g.LogLevel; v != "" {
		if err := writeString(w, "log_level", string(v)); err != nil {
			return err
		}
	}

	if v := g.LogLimit; v != 0 {
		if err := writeInt(w, "log_limit", v); err != nil {
			return err
		}
	}

	if v := g.LogBuffering; v != nil {
		if err := writeBool(w, "log_buffering", *v); err != nil {
			return err
		}
	}

	if v := g.EmergencyRestartThreshold; v != 0 {
		if err := writeInt(w, "emergency_restart_threshold", v); err != nil {
			return err
		}
	}

	if v := g.ProcessControlTimeout; v != 0 {
		if err := writeString(
			w,
			"process_control_timeout",
			fmt.Sprintf("%ds", int(v.Seconds())),
		); err != nil {
			return err
		}
	}

	if p := g.Process; p != nil {
		if err := p.write(w); err != nil {
			return err
		}
	}

	if v := g.RLimitFiles; v != 0 {
		if err := writeInt(w, "rlimit_files", v); err != nil {
			return err
		}
	}

	if v := g.RlimitCore; v != 0 {
		if err := writeInt(w, "rlimit_core", v); err != nil {
			return err
		}
	}

	for name, pool := range g.Pools {
		if err := writeSection(w, name); err != nil {
			return err
		}

		if err := pool.write(w); err != nil {
			return err
		}
	}

	return nil
}
