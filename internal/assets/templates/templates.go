// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/midoks/novelsearch/internal/render"
	"github.com/midoks/novelsearch/internal/tools"
)

//go:generate go-bindata -nomemcopy -nometadata -ignore="\\.DS_Store" -pkg=templates -prefix=../../../templates -debug=false -o=templates_gen.go ../../../templates/...

// NewTemplateFileSystem returns a macaron.TemplateFileSystem instance for embedded assets.
// The argument "dir" can be used to serve subset of embedded assets. Template file
// found under the "customDir" on disk has higher precedence over embedded assets.
func NewTemplateFileSystem(dir, customDir string) render.FileSystem {

	return render.NewCbFS(func() []render.File {
		if dir != "" && !strings.HasSuffix(dir, "/") {
			dir += "/"
		}

		var files []render.File
		names := AssetNames()
		for _, name := range names {
			if !strings.HasPrefix(name, dir) {
				continue
			}

			var err error
			var data []byte
			fpath := path.Join(customDir, name)
			if tools.IsFile(fpath) {
				data, err = ioutil.ReadFile(fpath)
			} else {
				data, err = Asset(name)
			}
			if err != nil {
				panic(err)
			}

			name = strings.TrimPrefix(name, dir)
			ext := path.Ext(name)
			name = strings.TrimSuffix(name, ext)
			files = append(files, render.NewFile(name, data, ext))
		}

		return files
	}, false)
}
