package files

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

// Service holds references to any necessary dependencies for the files Service
type Service struct{}

// New creates a new Service
func New() *Service {
	return &Service{}
}

// Monitor wraps the given context and starts the long running goroutine
func (s *Service) Monitor(ctx context.Context, path string) (context.CancelFunc, error) {
	if ctx == nil {
		panic("expected context to be passed to service")
	}
	ctx, cancel := context.WithCancel(ctx)

	go s.doMonitor(ctx, path)

	return cancel, nil
}

func (s *Service) doMonitor(ctx context.Context, path string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("shutting down file path monitor goroutine\n")
			if ctx.Err() != nil {
				fmt.Printf("an error occured while shutting down file path monitor: %v\n", ctx.Err())
			}
			return
		case <-time.After(5 * time.Second):
			fmt.Printf("checking path %s for files\n", path)

			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatalf("failed to read given path %s - due to error: %v", path, err)
			}

			for _, f := range files {
				fmt.Printf("Found file with name: %s\n", f.Name())
			}

		}
	}
}
