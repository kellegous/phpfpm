package phpfpm

import (
	"context"

	"github.com/kellegous/phpfpm/config"
)

type Process struct {
	pid int
}

func Start(
	ctx context.Context,
	path string,
	cfg *config.Global,
) (*Process, error) {
	return nil, nil
}
