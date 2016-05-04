package apollostats

import (
	"html"
	"strings"
)

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
