package config

import (
	"errors"
	"io"
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

func (l *Listen) write(w io.Writer) error {
	if v := l.Addr; v != "" {
		if err := writeString(w, "listen", v); err != nil {
			return err
		}
	} else {
		return errors.New("Listen.Addr is required")
	}

	if v := l.Backlog; v != nil {
		if err := writeInt(w, "listen.backlog", *v); err != nil {
			return err
		}
	}

	if v := l.AllowedClients; len(v) > 0 {
		if err := writeString(
			w,
			"listen.allowed_clients",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}

	if v := l.Owner; v != "" {
		if err := writeString(w, "listen.owner", v); err != nil {
			return err
		}
	}

	if v := l.Group; v != "" {
		if err := writeString(w, "listen.group", v); err != nil {
			return err
		}
	}

	if v := l.Mode; v != "" {
		if err := writeString(w, "listen.mode", v); err != nil {
			return err
		}
	}

	if v := l.ACLUsers; len(v) > 0 {
		if err := writeString(
			w,
			"listen.acl_users",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}

	if v := l.ACLGroups; len(v) > 0 {
		if err := writeString(
			w,
			"listen.acl_groups",
			strings.Join(v, ",")); err != nil {
			return err
		}
	}
	return nil
}
