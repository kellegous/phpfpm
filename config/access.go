package config

type Access struct {
	Log    string
	Format string
}

func (a *Access) write(w *writer) error {
	if v := a.Log; v != "" {
		if err := w.writeString("access.log", v); err != nil {
			return err
		}
	}

	if v := a.Format; v != "" {
		if err := w.writeString("access.format", v); err != nil {
			return err
		}
	}

	return nil
}
