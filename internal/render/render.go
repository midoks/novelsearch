package render

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
)

const (
	_TMPL_DIR = "templates"

	_CONTENT_TYPE    = "Content-Type"
	_CONTENT_BINARY  = "application/octet-stream"
	_CONTENT_JSON    = "application/json"
	_CONTENT_HTML    = "text/html"
	_CONTENT_PLAIN   = "text/plain"
	_CONTENT_XHTML   = "application/xhtml+xml"
	_CONTENT_XML     = "text/xml"
	_DEFAULT_CHARSET = "UTF-8"
)

var (
	// Provides a temporary buffer to execute templates into and catch errors.
	bufpool = sync.Pool{
		New: func() interface{} { return new(bytes.Buffer) },
	}

	// Included helper functions for use when rendering html
	helperFuncs = template.FuncMap{
		"yield": func() (string, error) {
			return "", fmt.Errorf("yield called with no layout defined")
		},
		"current": func() (string, error) {
			return "", nil
		},
	}
)

type (

	// Delims represents a set of Left and Right delimiters for HTML template rendering
	Delims struct {
		// Left delimiter, defaults to {{
		Left string
		// Right delimiter, defaults to }}
		Right string
	}

	// RenderOptions represents a struct for specifying configuration options for the Render middleware.
	Options struct {
		// Directory to load templates. Default is "templates".
		Directory string
		// Addtional directories to overwite templates.
		AppendDirectories []string
		// Layout template name. Will not render a layout if "". Default is to "".
		Layout string
		// Extensions to parse template files from. Defaults are [".tmpl", ".html"].
		Extensions []string
		// Funcs is a slice of FuncMaps to apply to the template upon compilation. This is useful for helper functions. Default is [].
		Funcs []template.FuncMap
		// Delims sets the action delimiters to the specified strings in the Delims struct.
		Delims Delims
		// Appends the given charset to the Content-Type header. Default is "UTF-8".
		Charset string
		// Outputs human readable JSON.
		IndentJSON bool
		// Outputs human readable XML.
		IndentXML bool
		// Prefixes the JSON output with the given bytes.
		PrefixJSON []byte
		// Prefixes the XML output with the given bytes.
		PrefixXML []byte
		// Allows changing of output to XHTML instead of HTML. Default is "text/html"
		HTMLContentType string
		// FileSystem is the interface for supporting any implmentation of template file system.
		FileSystem
	}
)

var tmpl *template.Template

func prepareRenderOptions(options []Options) Options {
	var opt Options
	if len(options) > 0 {
		opt = options[0]
	}

	// Defaults.
	if len(opt.Directory) == 0 {
		opt.Directory = _TMPL_DIR
	}
	if len(opt.Extensions) == 0 {
		opt.Extensions = []string{".tmpl", ".html"}
	}
	if len(opt.HTMLContentType) == 0 {
		opt.HTMLContentType = _CONTENT_HTML
	}

	return opt
}

func Renderer(opt Options) {
	opt = prepareRenderOptions([]Options{opt})

	tmpl = template.New(opt.Directory)
	tmpl.Delims(opt.Delims.Left, opt.Delims.Right)

	if opt.FileSystem == nil {
		opt.FileSystem = NewFS(opt, false)
	}

	for _, f := range opt.FileSystem.Files() {
		tmpl = tmpl.New(f.Name())

		for _, funcs := range opt.Funcs {
			tmpl.Funcs(funcs)
		}

		// Bomb out if parse fails. We don't want any silent server starts.
		template.Must(tmpl.Funcs(helperFuncs).Parse(string(f.Data())))
	}
}

func HTML(name string, data interface{}) ([]byte, error) {
	var err error
	buf := bufpool.Get().(*bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, name, data)

	r := buf.Bytes()
	buf.Reset()
	bufpool.Put(buf)

	return r, err
}
