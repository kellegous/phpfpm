package config

type PoolProcess struct {
	Priority *int
	Dumpable bool
}

func (p *PoolProcess) write(w *writer) error {
	if v := p.Priority; v != nil {
		if err := w.writeInt("process.priority", *v); err != nil {
			return err
		}
	}

	if v := p.Dumpable; v {
		if err := w.writeBool("process.dumpable", v); err != nil {
			return err
		}
	}

	return nil
}
