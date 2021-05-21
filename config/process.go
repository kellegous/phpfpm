package config

import "io"

type Process struct {
	Max      int
	Priority *int
}

func (p *Process) write(w io.Writer) error {
	if v := p.Max; v != 0 {
		if err := writeInt(w, "process.max", v); err != nil {
			return err
		}
	}

	if v := p.Priority; v != nil {
		if err := writeInt(w, "process.priority", *v); err != nil {
			return err
		}
	}

	return nil
}
