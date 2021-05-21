package config

type Ping struct {
	Path     string
	Response string
}

func (p *Ping) validate() error {
	// TODO(knorton): path needs to begin with a /
	return nil
}

func (p *Ping) write(w *writer) error {

	if v := p.Path; v != "" {
		if err := w.writeString("ping.path", v); err != nil {
			return err
		}
	}

	if v := p.Response; v != "" {
		if err := w.writeString("ping.response", v); err != nil {
			return err
		}
	}

	return nil
}
