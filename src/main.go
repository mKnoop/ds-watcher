package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const defaultTick = 60 * time.Second

type config struct {
	statusCode int
	tick       time.Duration
	urls       string
}

func (c *config) init(args []string) error {
	var (
		statusCode = flag.Int("status", 200, "Response HTTP status code")
		tick       = flag.Duration("tick", defaultTick, "Ticking interval")
		urls       = flag.String("urls", "", "Request URLs")
	)

	flag.Parse()

	c.statusCode = *statusCode
	c.tick = *tick
	c.urls = *urls

	return nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)

	c := &config{}

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					c.init(os.Args)
				case os.Interrupt:
					cancel()
					os.Exit(1)
				}
			case <-ctx.Done():
				log.Printf("Done.")
				os.Exit(1)
			}
		}
	}()

	if err := run(ctx, c, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, c *config, stdout io.Writer) error {
	c.init(os.Args)
	log.SetOutput(os.Stdout)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.Tick(c.tick):
			fmt.Printf("Check urls")
		}
	}
}
