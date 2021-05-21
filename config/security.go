package config

import "strings"

type Security struct {
	LimitExtensions []string
}

func (s *Security) write(w *writer) error {
	if v := s.LimitExtensions; len(v) > 0 {
		if err := w.writeString(
			"security.limit_extensions",
			strings.Join(v, " ")); err != nil {
			return err
		}
	}

	return nil
}
