package apollostats

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

type Instance struct {
	Debug bool
	DB    *DB

	addr   string
	router *gin.Engine
}

func (i *Instance) Init() {
	if i.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO: replace Default with New and use custom logger and stuff
	i.router = gin.Default()

	// Custom functions for the templates
	funcmap := template.FuncMap{
		"pretty_time": func(t time.Time) string {
			return t.Format("2006-01-02 15:04 MST")
		},
		"year": func() int {
			return time.Now().Year()
		},
		"commas": func(i int64) string {
			return humanize.Comma(i)
		},
		"default_job": func(s string) string {
			if len(strings.TrimSpace(s)) < 1 {
				return "Unknown"
			}
			return s
		},
	}

	// Load templates
	templatebox := rice.MustFindBox("templates")
	templates := template.New("ServerTemplates").Funcs(funcmap)
	// Iterate over all templates and mash them together
	templatebox.Walk("", func(p string, i os.FileInfo, e error) error {
		if i.IsDir() {
			return nil
		}
		s, e := templatebox.String(p)
		if e != nil {
			log.Fatalf("Failed to load template: %s\n%s\n", p, e)
		}
		template.Must(templates.New(p).Parse(s))
		return nil
	})
	i.router.SetHTMLTemplate(templates)

	// And static files
	static := rice.MustFindBox("static")
	i.router.StaticFS("/static/", static.HTTPBox())

	// Setup all views
	i.router.GET("/", i.index)
	i.router.GET("/favicon.ico", i.favicon)
	i.router.GET("/robots.txt", i.robots)
	i.router.GET("/bans", i.bans)
	i.router.GET("/account_items", i.account_items)
	i.router.GET("/rounds", i.rounds)
	i.router.GET("/round/:round_id", i.round_detail)
	i.router.GET("/characters", i.characters)
	i.router.GET("/character/:char_id", i.character_detail)
}

func (i *Instance) Serve(addr string) error {
	i.addr = addr
	return i.router.Run(i.addr)
}

func (i *Instance) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pagetitle": "Index",
		"Round":     i.DB.GetLatestRound(),
		"Stats":     i.DB.GetStats(),
	})
}

func (i *Instance) favicon(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/static/favicon.ico")
}

func (i *Instance) robots(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/static/robots.txt")
}

func (i *Instance) bans(c *gin.Context) {
	c.HTML(http.StatusOK, "bans.html", gin.H{
		"pagetitle": "Bans",
		"Bans":      i.DB.AllBans(),
	})
}

func (i *Instance) account_items(c *gin.Context) {
	c.HTML(http.StatusOK, "account_items.html", gin.H{
		"pagetitle":    "Account Items",
		"AccountItems": i.DB.AllAccountItems(),
	})
}

func (i *Instance) rounds(c *gin.Context) {
	c.HTML(http.StatusOK, "rounds.html", gin.H{
		"pagetitle": "Rounds",
		"Rounds":    i.DB.AllRounds(),
	})
}

func (i *Instance) round_detail(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("round_id"), 10, 0)
	if e != nil {
		id = -1
	}
	round := i.DB.GetRound(id)

	c.HTML(http.StatusOK, "round_detail.html", gin.H{
		"pagetitle": fmt.Sprintf("Round #%d", round.ID),
		"Round":     round,
		"Antags":    i.DB.GetAntags(id),
		"AILaws":    i.DB.GetAILaws(id),
		"Deaths":    i.DB.GetDeaths(id),
	})
}

func (i *Instance) characters(c *gin.Context) {
	ckey := c.Query("ckey")
	name := c.Query("name")
	chars := i.DB.SearchCharacter(ckey, name)

	c.HTML(http.StatusOK, "characters.html", gin.H{
		"pagetitle": "Characters",
		"Chars":     chars,
	})
}

func (i *Instance) character_detail(c *gin.Context) {
	id, e := strconv.ParseInt(c.Param("char_id"), 10, 0)
	if e != nil {
		id = -1
	}
	char := i.DB.GetCharacter(id)

	c.HTML(http.StatusOK, "character_detail.html", gin.H{
		"pagetitle": fmt.Sprintf("%v by %v", char.NiceName(), char.CKey),
		"Char":      char,
	})
}
