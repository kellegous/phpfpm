package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/kellegous/phpfpm"
	"github.com/kellegous/phpfpm/config"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	cfg := config.New().
		WithErrorLog(filepath.Join(wd, "error.log")).
		BuildPool(
			"www",
			"127.0.0.1:9696",
			"knorton",
			config.TypeDynamic,
			10,
			func(p *config.Pool) {
				minSpare := 2
				maxSpare := 4
				p.ProcessManager.StartServers = 2
				p.ProcessManager.MinSpareServers = &minSpare
				p.ProcessManager.MaxSpareServers = &maxSpare
			})

	p, err := phpfpm.Start(
		context.Background(),
		"/usr/sbin/php-fpm",
		cfg)
	if err != nil {
		log.Panic(err)
	}

	if err := p.Wait(); err != nil {
		log.Panic(err)
	}
}
