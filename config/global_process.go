package config

type GlobalProcess struct {
	Max      int
	Priority *int
}

func (p *GlobalProcess) write(w *writer) error {
	if v := p.Max; v != 0 {
		if err := w.writeInt("process.max", v); err != nil {
			return err
		}
	}

	if v := p.Priority; v != nil {
		if err := w.writeInt("process.priority", *v); err != nil {
			return err
		}
	}

	return nil
}
