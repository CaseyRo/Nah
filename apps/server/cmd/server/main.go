// Command server runs a Nah? server.
//
//	server           serve
//	server invite    print a one-time invitation for the first person here
//
// Nah? is by invitation, so the first person on a server needs one from the
// operator; everyone after them joins through someone already here.
//
// Configuration is environment variables read into a struct at start, per
// ADR-0014. There is no configuration file and there is no flag parsing.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CaseyRo/Nah/apps/server/internal/nah"
)

type config struct {
	addr    string
	dataDir string
}

func load() config {
	return config{
		addr:    env("NAH_ADDR", ":8080"),
		dataDir: env("NAH_DATA_DIR", "./data"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	if len(os.Args) > 2 || (len(os.Args) == 2 && os.Args[1] != "invite") {
		fmt.Fprintln(os.Stderr, "usage: server           serve\n       server invite    print a one-time invitation for the first person here")
		os.Exit(2)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := load()

	store, err := nah.NewStore(cfg.dataDir)
	if err != nil {
		log.Error("cannot open the data directory", "dir", cfg.dataDir, "err", err)
		os.Exit(1)
	}
	defer store.Close()

	if len(os.Args) == 2 {
		invite, err := store.OperatorInvite()
		if err != nil {
			log.Error("cannot make an invitation", "err", err)
			os.Exit(1)
		}
		// The app reads an invitation as person#invite; an operator's names nobody.
		fmt.Println("#" + invite)
		return
	}

	if err := serve(cfg, store, log); err != nil {
		log.Error("server stopped", "err", err)
		os.Exit(1)
	}
}

func serve(cfg config, store *nah.Store, log *slog.Logger) error {
	// Sessions outlive the process because this key does (ADR-0015). Deploys
	// are going to be frequent and nobody should be able to tell.
	key, err := nah.SessionKey(cfg.dataDir)
	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:              cfg.addr,
		Handler:           nah.NewServer(store, nah.NewAuth(key), log).Handler(),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       2 * time.Minute,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	errc := make(chan error, 1)
	go func() {
		log.Info("listening", "addr", cfg.addr, "data", cfg.dataDir)
		errc <- srv.ListenAndServe()
	}()

	select {
	case err := <-errc:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		log.Info("shutting down")
		shutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return srv.Shutdown(shutdown)
	}
}
