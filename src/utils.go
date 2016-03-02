package apollostats

import (
	"html"
	"strings"
)

// Func to take care of garbled text data.
func filter_string(s string) string {
	// Fuck it, might aswell assume all text has been escaped.
	tmp := html.UnescapeString(s)

	// Usually seen in the room names in the death table.
	tmp = strings.Trim(tmp, "ÿ")

	tmp = strings.TrimSpace(tmp)
	return tmp
}
