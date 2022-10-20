package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	_ "github.com/benthosdev/benthos/v4/public/components/io"
	_ "github.com/benthosdev/benthos/v4/public/components/jaeger"
	_ "github.com/benthosdev/benthos/v4/public/components/prometheus"

	_ "github.com/disintegrator/benthos-pglisten/internal/postgres"
)

func main() {
	service.RunCLI(context.Background())
}
