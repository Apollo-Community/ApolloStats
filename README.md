ApolloStats
--------------------------------------------------------------------------------

Webpage for showing various stats from the Apollo Station SS13 game database.
With heavy inspiration from other servers' webpages such as [Goon](http://goonhub.com/) and [/vg/station](http://ss13.pomf.se/index.php/bans).

Requires a running Mysql/MariaDB server with the game's database.

Installation
--------------------------------------------------------------------------------

Compile time dependencies:

- Go v1.5+ (Unknown if older versions works)
- [cli](https://github.com/codegangsta/cli)
- [gin](https://github.com/gin-gonic/gin)
- [go-humanize](https://github.com/dustin/go-humanize)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [go.rice](https://github.com/GeertJohan/go.rice)
- [gorm](https://github.com/jinzhu/gorm)

Compilation:

Download the source code (via `go get`, `git` or from a [zipped archive](https://github.com/Apollo-Community/ApolloStats/archive/v0.1.zip)).
Open the directory with the source code and run `go build` to compile the code.
You will now have a `ApolloStats` binary in the dir, which you can run to start
the web server.

Stand alone binary:

TODO

Usage
--------------------------------------------------------------------------------

```
$ ApolloStats -h
NAME:
   main - Run a web server, serving stats for Apollo.

USAGE:
   main [global options] command [command options] [arguments...]
   
VERSION:
   0.1
   
COMMANDS:
   run          Run the web server
   help, h      Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --addr, -a "127.0.0.1:8000"                  serve web pages on this address [$APOLLOSTATS_ADDR]
   --database, -d "user:password@/database"     database authentication string [$APOLLOSTATS_DBAUTH]
   --debug                                      run in debug mode
   --help, -h                                   show help
   --version, -v                                print the version
```

Environment variables:

TODO

TODO
--------------------------------------------------------------------------------

Templates:
- Show some nice error pages.
- Needs some fallback for when we can't load css from forums?
- Change the colors used in css.

Database:
- Timeout and show error when we can't connect to the ext. db.

Account items:
- Would be nice to show why a player got an item too.

Heat maps:
- Really nice if we could show a heatmap of deaths.
- Ask @HiddenKn how he made his python version.

Game map:
- Huge, zoomable map of the main station (only).
- Store the map as picture tiles?
- Investigate how goon made their map.
- Need to rebuild the map after any new map changes from a commit.
