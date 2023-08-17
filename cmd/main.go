package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/Kawaii-Konnections-KK-Limited/skraper/models"
	"github.com/Kawaii-Konnections-KK-Limited/skraper/pkg/telegram"
)

func main() {
	models.InitDB()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := telegram.Run(ctx); err != nil {
		if errors.Is(err, context.Canceled) && ctx.Err() == context.Canceled {
			fmt.Println("\rClosed")
			os.Exit(0)
		}
		_, _ = fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Done")
		os.Exit(0)
	}
}
