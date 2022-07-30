package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"todo/config"
)

func main() {
	if (len(os.Args)) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}

	p := os.Args[1]

	log.Println("server start...")

	log.Printf("listen port %s", p)

	if err := run(context.Background()); err != nil {
		fmt.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())

	log.Printf("start with %v", url)

	mux, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		return err
	}

	defer cleanup()

	s := NewServer(l, mux)

	return s.Run(ctx)
}
