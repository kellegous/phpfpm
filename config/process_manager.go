package config

import (
	"errors"
	"time"
)

type ProcessManagerType string

const (
	TypeStatic ProcessManagerType = "static"

	TypeOnDemand ProcessManagerType = "ondemand"

	TypeDynamic ProcessManagerType = "dynamic"
)

type ProcessManager struct {
	Type               ProcessManagerType
	MaxChildren        int
	StartServers       int
	MinSpareServers    *int
	MaxSpareServers    *int
	ProcessIdleTimeout time.Duration
	MaxRequests        int
	StatusPath         string
}

func (p *ProcessManager) validate() error {
	if p.Type == "" {
		return errors.New("process manager type required")
	}

	if p.MaxChildren == 0 {
		return errors.New("max_children is required")
	}

	if p.Type == TypeDynamic {
		if p.MinSpareServers == nil {
			return errors.New("min_spare_servers required for dynamic")
		}

		if p.MaxSpareServers == nil {
			return errors.New("max_spare_servers required for dynamic")
		}

		if *p.MinSpareServers > *p.MaxSpareServers {
			return errors.New("min_spare_servers must be <= max_spare_servers")
		}
	}

	return nil
}

func (p *ProcessManager) write(w *writer) error {
	if err := p.validate(); err != nil {
		return err
	}

	if err := w.writeString("pm", string(p.Type)); err != nil {
		return err
	}

	if err := w.writeInt(
		"pm.max_children",
		p.MaxChildren); err != nil {
		return err
	}

	if v := p.StartServers; v != 0 {
		if err := w.writeInt(
			"pm.start_servers",
			v); err != nil {
			return err
		}
	}

	if v := p.MinSpareServers; v != nil {
		if err := w.writeInt(
			"pm.min_spare_servers",
			*v); err != nil {
			return err
		}
	}

	if v := p.MaxSpareServers; v != nil {
		if err := w.writeInt(
			"pm.max_spare_servers",
			*v); err != nil {
			return err
		}
	}

	if v := p.ProcessIdleTimeout; v != 0 {
		if err := w.writeDuration("pm.process_idle_timeout", v); err != nil {
			return err
		}
	}

	if v := p.MaxRequests; v != 0 {
		if err := w.writeInt("pm.max_requests", v); err != nil {
			return err
		}
	}

	if v := p.StatusPath; v != "" {
		if err := w.writeString("pm.status_path", v); err != nil {
			return err
		}
	}

	return nil
}
