package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/aranw/graceful-context-example/files"
)

func main() {
	path := os.Args[1]

	s := files.New()
	cancel, err := s.Monitor(context.Background(), path)
	if err != nil {
		fmt.Printf("failed to start file path monitor due to error: %v\n", err)
		os.Exit(1)
	}
	defer cancel()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill)

	sig := <-sigquit

	cancel()

	fmt.Printf("caught sig: %+v\n", sig)
	fmt.Printf("Gracefully shutting down file monitor...\n")

	<-time.After(5 * time.Second)
}
