package render

// File represents a interface of template file that has name and can be read.
type File interface {
	Name() string
	Data() []byte
	Ext() string
}

// TplFile implements TemplateFile interface.
type file struct {
	name string
	data []byte
	ext  string
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Data() []byte {
	return f.data
}

func (f *file) Ext() string {
	return f.ext
}

// NewTplFile cerates new template file with given name and data.
func NewFile(name string, data []byte, ext string) *file {
	return &file{name, data, ext}
}
