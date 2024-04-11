package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Ricky004/microservice-server/app"
)

func main() {
	app := app.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}

