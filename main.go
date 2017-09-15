package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

var VERSION = "dev"

func handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"verions": VERSION,
		"path":    r.URL.Path,
		"status":  "ok",
	})
}

func main() {
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
						fmt.Println("start listen", c.String("address"))
						http.HandleFunc("/", handler)
						fmt.Println(http.ListenAndServe(c.String("address"), nil))

						return nil
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
