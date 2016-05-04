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
	g_verbose  bool
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
			Name:        "verbose",
			Usage:       "run with verbose output",
			Destination: &g_verbose,
		},
	}
	app.Commands = []cli.Command{
		{Name: "run", Usage: "Run the web server", Action: run_server},
	}
	app.Run(os.Args)
}

func run_server(c *cli.Context) {
	db, e := apollostats.OpenDB(g_database, g_verbose)
	if e != nil && g_verbose {
		log.Printf("Failed to connect to the database:\n%s\n", e.Error())
	}

	i := apollostats.Instance{
		Verbose: g_verbose,
		DB:      db,
	}

	e = i.Init()
	if e != nil {
		log.Panic(e)
	}
	e = i.Serve(g_addr)
	if e != nil {
		log.Panic(e)
	}
}
