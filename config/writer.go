package config

import (
	"fmt"
	"io"
	"time"
)

type writer struct {
	io.Writer
}

func (w *writer) writeString(
	key string,
	val string,
) error {
	_, err := fmt.Fprintf(w, "%s = %s\n", key, val)
	return err
}

func (w *writer) writeInt(
	key string,
	val int,
) error {
	_, err := fmt.Fprintf(w, "%s = %d\n", key, val)
	return err
}

func boolToString(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func (w *writer) writeBool(
	key string,
	val bool,
) error {
	return w.writeString(key, boolToString(val))
}

func (w *writer) writeSection(key string) error {
	_, err := fmt.Fprintf(w, "[%s]\n", key)
	return err
}

func (w *writer) writeDuration(
	key string,
	val time.Duration,
) error {
	return w.writeString(key, fmt.Sprintf("%ds", int(val.Seconds())))
}
