package main

import (
	"log"
	"os"

	"github.com/Apollo-Community/ApolloStats/src"
	"github.com/codegangsta/cli"
)

var (
	g_addr     string
	g_database string
	g_debug    bool
)

func main() {
	app := cli.NewApp()
	app.Version = apollostats.VERSION
	app.Usage = "Run a web server, serving stats for Apollo."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "addr, a",
			Value:       "127.0.0.1:8000",
			Usage:       "serve web pages on this address",
			EnvVar:      "APOLLOSTATS_ADDR",
			Destination: &g_addr,
		},
		cli.StringFlag{
			Name:        "database, d",
			Value:       "user:password@/database",
			Usage:       "database authentication string",
			EnvVar:      "APOLLOSTATS_DBAUTH",
			Destination: &g_database,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "run in debug mode",
			Destination: &g_debug,
		},
	}
	app.Commands = []cli.Command{
		{Name: "run", Usage: "Run the web server", Action: run_server},
	}
	app.Run(os.Args)
}

func run_server(c *cli.Context) {
	db, e := apollostats.OpenDB(g_database)
	if e != nil && g_debug {
		log.Printf("Failed to connect to the database:\n%s\n", e.Error())
	}

	i := apollostats.Instance{
		Debug: g_debug,
		DB:    db,
	}

	i.Init()
	i.Serve(g_addr)
}
