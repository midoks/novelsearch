package tmpl

import (
	"html/template"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/midoks/novelsearch/internal/conf"
)

func Year() int {
	return time.Now().Year()
}

func Unescape(s string) template.HTML {
	return template.HTML(s)
}

func AdminPath() string {
	return conf.Admin.AdminPath
}

func Safe(raw string) template.HTML {
	return template.HTML(raw)
}

func Str2HTML(raw string) template.HTML {
	return template.HTML(bluemonday.UGCPolicy().Sanitize(raw))
}

// NewLine2br simply replaces "\n" to "<br>".
func NewLine2br(raw string) string {
	return strings.Replace(raw, "\n", "<br>", -1)
}

// TODO: Use url.Escape.
func EscapePound(str string) string {
	return strings.NewReplacer("%", "%25", "#", "%23", " ", "%20", "?", "%3F").Replace(str)
}
