package render

import (
	"bytes"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
	"path/filepath"
)

type CALLBACK func() []File

// FileSystem represents a interface of template file system that able to list all files.
type FileSystem interface {
	Files() []File
	Get(string) (io.Reader, error)
}

// TplFS implements TemplateFileSystem interface.
type fileSystem struct {
	files []File
}

func (fs *fileSystem) Files() []File {
	return fs.files
}

func (fs *fileSystem) Get(name string) (io.Reader, error) {
	for i := range fs.files {
		if fs.files[i].Name()+fs.files[i].Ext() == name {
			return bytes.NewReader(fs.files[i].Data()), nil
		}
	}
	return nil, fmt.Errorf("file '%s' not found", name)
}

func NewCbFS(cb CALLBACK, omitData bool) FileSystem {

	fs := fileSystem{}
	fs.files = cb()

	return &fs
}

// NewTemplateFileSystem creates new template file system with given options.
func NewFS(opt Options, omitData bool) FileSystem {
	var err error
	fs := fileSystem{}

	// // Directories are composed in reverse order because later one overwrites previous ones,
	// // so once found, we can directly jump out of the loop.
	dirs := make([]string, 0, len(opt.AppendDirectories)+1)
	for i := len(opt.AppendDirectories) - 1; i >= 0; i-- {
		dirs = append(dirs, opt.AppendDirectories[i])
	}
	dirs = append(dirs, opt.Directory)

	// var err error
	for i := range dirs {
		// Skip ones that does not exists for symlink test,
		// but allow non-symlink ones added after start.
		if !IsExist(dirs[i]) {
			continue
		}

		dirs[i], err = filepath.EvalSymlinks(dirs[i])
		if err != nil {
			panic("EvalSymlinks(" + dirs[i] + "): " + err.Error())
		}
	}

	lastDir := dirs[len(dirs)-1]

	// We still walk the last (original) directory because it's non-sense we load templates not exist in original directory.
	if err := filepath.Walk(lastDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		r, err := filepath.Rel(lastDir, path)
		if err != nil {
			return err
		}

		ext := GetExt(r)
		for _, extension := range opt.Extensions {
			if ext != extension {
				continue
			}

			var data []byte
			if !omitData {
				// Loop over candidates of directory, break out once found.
				// The file always exists because it's inside the walk function,
				// and read original file is the worst case.
				for i := range dirs {
					path = filepath.Join(dirs[i], r)

					if !IsFile(path) {
						continue
					}

					data, err = os.ReadFile(path)
					if err != nil {
						return err
					}

					break
				}
			}

			name := filepath.ToSlash((r[0 : len(r)-len(ext)]))
			fs.files = append(fs.files, NewFile(name, data, ext))
		}

		return nil
	}); err != nil {
		panic("NewFS: " + err.Error())
	}

	return &fs
}
