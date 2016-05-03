ApolloStats
================================================================================

Webpage for showing various stats from the Apollo Station SS13 game database.
With heavy inspiration from other servers' webpages such as [Goon](http://goonhub.com/) and [/vg/station](http://ss13.pomf.se/index.php/bans).

Requires a running Mysql/MariaDB server with the game's database.

Installation
================================================================================

Compile time dependencies
--------------------------------------------------------------------------------

- Go v1.5+ (Unknown if older versions works)
- [cli](https://github.com/codegangsta/cli)
- [gin](https://github.com/gin-gonic/gin)
- [go-humanize](https://github.com/dustin/go-humanize)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [gorm](https://github.com/jinzhu/gorm)

Build tool for embedding assets:
- [yaber](https://github.com/lmas/yaber)

Compilation
--------------------------------------------------------------------------------

Download the source code

    go get -u github.com/Apollo-Community/ApolloStats

Or from a [zipped archive](https://github.com/Apollo-Community/ApolloStats/releases)).

Go to the directory with the unpacked source code and run `make build` to compile
the code to a stand alone binary, called `ApolloStats64`.

See the `Makefile` for more options.

Usage
================================================================================

```
$ ApolloStats -h
NAME:
   main - Run a web server, serving stats for Apollo.

USAGE:
   main [global options] command [command options] [arguments...]
   
VERSION:
   0.3
   
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

Environment variables
--------------------------------------------------------------------------------

You can set some environment variables instead of using the command line flags.
```
$ export APOLLOSTATS_ADDR="127.0.0.1:8000"
$ export APOLLOSTATS_DBAUTH="user:password@/database"
$ ./ApolloStats --debug run
```

License
================================================================================

MIT License, see the LICENSE file for details.

TODO
================================================================================

Info:
- Clarify that there are never info from a currently running round.
- Clarify total deaths.

Templates:
- Need to remake the custom CSS and make the pages prettier.
- Change the colors used in css.

Database:
- Fix duplicate entries of drones.
- Fix duplicate entries of AI laws.

Account items:
- Would be nice to show why a player got an item too.

Rounds:
- Show player death graph over round's duration.

Heat maps:
- Really nice if we could show a heatmap of deaths.
- Ask @HiddenKn how he made his python version.

Game map:
- Huge, zoomable map of the main station (only).
- Store the map as picture tiles?
- Investigate how goon made their map.
- Need to rebuild the map after any new map changes from a commit.

Tests:
- Must do unit tests.
