package apollostats

import (
	"html"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// Thanks to gin for overriding the standard flags...
	log.SetFlags(log.LstdFlags)
}

// Func to take care of garbled text data.
func filter_string(s string) string {
	// Fuck it, might aswell assume all text has been escaped.
	tmp := html.UnescapeString(s)
	// And there are cases where someone's escaped the data at least twice,
	// turning already escaped text like she&#39;s into she&amp;#39;s ...
	// Ridiculus.
	tmp = html.UnescapeString(tmp)
	// Usually seen in the room names in the death table.
	tmp = strings.Trim(tmp, "ÿ")
	tmp = strings.TrimSpace(tmp)
	return tmp
}

func (i *Instance) logMsg(format string, args ...interface{}) {
	if i.Verbose {
		log.Printf(format+"\n", args...)
	}
}

// Simple logging middleware, for replacing Gin's fancy shit.
func (i *Instance) logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		stop := time.Now()

		i.logMsg("%s\t%s\t%d\t%s\t%s",
			c.ClientIP(),
			stop.Sub(start).String(),
			c.Writer.Status(),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
