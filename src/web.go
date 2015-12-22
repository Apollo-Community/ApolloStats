package apollostats

import (
	"html/template"
	"net/http"
	"time"

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

	// Load templates
	funcmap := template.FuncMap{
		"pretty_time": func(t time.Time) string {
			return t.Format("2006-01-02 15:04 MST")
		},
		"year": func() int {
			return time.Now().Year()
		},
	}
	tmpl := template.Must(template.New("ServerTemplates").Funcs(funcmap).ParseGlob("templates/*"))
	i.router.SetHTMLTemplate(tmpl)

	// Setup all URLS
	i.router.Static("/static", "./static")

	i.router.GET("/", i.index)
	i.router.GET("/bans", i.bans)
	i.router.GET("/account_items", i.account_items)
	i.router.GET("/deaths", i.deaths)
}

func (i *Instance) Serve(addr string) error {
	i.addr = addr
	return i.router.Run(i.addr)
}

func (i *Instance) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pagetitle": "Index",
	})
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

func (i *Instance) deaths(c *gin.Context) {
	c.HTML(http.StatusOK, "deaths.html", gin.H{
		"pagetitle": "Deaths",
		"Deaths":    i.DB.AllDeaths(),
	})
}
