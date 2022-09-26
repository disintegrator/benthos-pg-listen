package main

import (
	"context"

	"github.com/benthosdev/benthos/v4/public/service"

	_ "github.com/benthosdev/benthos/v4/public/components/all"

	_ "github.com/disintegrator/benthos-pglisten/internal/postgres"
)

func main() {
	service.RunCLI(context.Background())
}
