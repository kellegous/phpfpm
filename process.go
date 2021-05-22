package phpfpm

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/kellegous/phpfpm/config"
)

type Process struct {
	p   *os.Process
	cfg *config.Global
	cf  string
}

func (p *Process) cleanup() error {
	return os.Remove(p.cf)
}

func (p *Process) Wait() error {
	defer p.cleanup()
	if _, err := p.p.Wait(); err != nil {
		return err
	}
	return nil
}

func Start(
	ctx context.Context,
	path string,
	cfg *config.Global,
) (*Process, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// TODO(knorton): Delete the temp file if we can't start

	if err := cfg.Write(f); err != nil {
		return nil, err
	}

	c := exec.Command(
		path,
		"--fpm-config", f.Name(),
		"--nodaemonize")
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		return nil, err
	}

	return &Process{
		p:   c.Process,
		cfg: cfg,
		cf:  f.Name(),
	}, nil
}
