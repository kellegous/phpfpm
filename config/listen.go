package config

import (
	"errors"
	"strings"
)

type Listen struct {
	Addr           string
	Backlog        *int
	AllowedClients []string
	Owner          string
	Group          string
	Mode           string
	ACLUsers       []string
	ACLGroups      []string
}

func (l *Listen) validate() error {
	if l.Addr == "" {
		return errors.New("Listen.Addr is required")
	}

	return nil
}

func (l *Listen) write(w *writer) error {
	if err := w.writeString("listen", l.Addr); err != nil {
		return err
	}

	if v := l.Backlog; v != nil {
		if err := w.writeInt("listen.backlog", *v); err != nil {
			return err
		}
	}

	if v := l.AllowedClients; len(v) > 0 {
		if err := w.writeString(
			"listen.allowed_clients",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}

	if v := l.Owner; v != "" {
		if err := w.writeString("listen.owner", v); err != nil {
			return err
		}
	}

	if v := l.Group; v != "" {
		if err := w.writeString("listen.group", v); err != nil {
			return err
		}
	}

	if v := l.Mode; v != "" {
		if err := w.writeString("listen.mode", v); err != nil {
			return err
		}
	}

	if v := l.ACLUsers; len(v) > 0 {
		if err := w.writeString(
			"listen.acl_users",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}

	if v := l.ACLGroups; len(v) > 0 {
		if err := w.writeString(
			"listen.acl_groups",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}
	return nil
}
