package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gebv/golang_with_deploy/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	cli "gopkg.in/urfave/cli.v2"
)

var (
	VERSION = "dev"
	log     *zap.SugaredLogger
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"verions": VERSION,
		"path":    r.URL.Path,
		"status":  "ok",
	})
}

func main() {
	log = logger.NewLogger(zap.DebugLevel).Sugar()
	log.Info("APP", VERSION)

	app := &cli.App{}
	app.Name = "audiocrm-backend"
	app.Version = VERSION
	app.Commands = []*cli.Command{
		{
			Name: "api",
			Subcommands: []*cli.Command{
				{
					Name: "run",
					Flags: []cli.Flag{
						&cli.StringFlag{Name: "address", Value: ":8944"},
					},
					Action: func(c *cli.Context) error {
						e := echo.New()
						e.Use(middleware.Logger())
						e.Use(middleware.Recover())
						e.GET("/stats", stats)

						var appSignal = make(chan struct{}, 2)
						go func() {
							log.Info("API listen", c.String("address"))
							log.Error("API listen", e.Start(c.String("address")))
							appSignal <- struct{}{}
						}()

						osSignal := make(chan os.Signal, 2)
						close := make(chan struct{})
						signal.Notify(
							osSignal,
							os.Interrupt,
							syscall.SIGTERM,
						)

						go func() {

							defer func() {
								close <- struct{}{}
							}()

							select {
							case <-osSignal:
								log.Error("signal completion of the process: OS")
							case <-appSignal:
								log.Error("signal completion of the process: internal (http server, etc..)")
							}

							// TODO: destroy services

							ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
							defer cancel()
							e.Shutdown(ctx)
						}()

						<-close
						os.Exit(0)
						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func stats(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"verions": VERSION,
	})
}
