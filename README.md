ApolloStats
================================================================================

Webpage for showing various stats from the Apollo Station SS13 game database.
With heavy inspiration from other servers' webpages such as [Goon](http://goonhub.com/) and [/vg/station](http://ss13.pomf.se/index.php/bans).

Requires a running Mysql/MariaDB server with the game's database.

Installation
================================================================================

Compile time dependencies:
--------------------------------------------------------------------------------

- Go v1.5+ (Unknown if older versions works)
- [cli](https://github.com/codegangsta/cli)
- [gin](https://github.com/gin-gonic/gin)
- [go-humanize](https://github.com/dustin/go-humanize)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [go.rice](https://github.com/GeertJohan/go.rice)
- [gorm](https://github.com/jinzhu/gorm)

Compilation:
--------------------------------------------------------------------------------

Download the source code (via `go get`, `git` or from a [zipped archive](https://github.com/Apollo-Community/ApolloStats/archive/v0.1.zip)).
Open the directory with the source code and run `go build` to compile the code.
You will now have a `ApolloStats` binary in the dir, which you can run to start
the web server.

Please note that this binary still depends on the source templates and static
files, found inside the `src/` directory.

Stand alone binary:
--------------------------------------------------------------------------------

To truly make a stand alone binary without having to depend on the `src/` directory,
you will have to append the templates and static files to the binary itself.
This step depends on the `rice` command, so you will have to install it too using
`go get github.com/GeertJohan/go.rice/rice`.

Then you can compile the binary like before. After that, change into the `src/`
directory and run `rice append --exec ../ApolloStats` to append the templates
and static filesto the final binary.

Makefile:
--------------------------------------------------------------------------------

There is a `Makefile` that will do all these steps for you too. Make sure you
have `make` installed and then simply run `make` and it will build the stand
alone binary for you.

Usage
================================================================================

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

Templates:
- Change the colors used in css.

Database:
- Timeout and show error when we can't connect to the ext. db.

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
