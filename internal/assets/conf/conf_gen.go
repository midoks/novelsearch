// Code generated for package conf by go-bindata DO NOT EDIT. (@generated)
// sources:
// ../../../conf/app.conf
package conf

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _confAppConf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x5d\x4f\x1a\x4b\x18\xc7\xaf\xcf\x7c\x8a\xc9\x72\x67\x22\x0e\x22\x1c\xb8\xe0\x24\x5e\x9c\xfb\x93\x9c\x4b\x62\xc8\xb0\x33\xba\x1b\xd9\x9d\xcd\xee\x52\x6d\xaf\xa8\x51\x4b\x5f\x90\x17\x23\xd6\x97\xd4\xd0\xaa\xd8\x68\x81\x0b\x4b\x2b\x88\xfd\x32\xcc\x2c\x5c\xf5\x2b\x34\xb3\x4b\x2d\x36\xa6\xbb\xc9\xec\x3e\xcf\x3c\xff\xff\xfc\xe6\xc9\x0c\xb6\x2c\x13\x1b\x14\xa6\xa0\xc9\x9e\xd0\x9c\x43\xb1\xad\x6a\x40\x73\x5d\xcb\x62\xb6\x0b\x53\x30\x89\x92\x08\xd8\x79\xd3\x60\x44\x56\x59\x36\x23\x00\xa4\x09\xcd\xe6\x57\x96\xc0\xa4\x26\x8e\xe2\x08\x80\x34\x26\x86\x6e\x2e\x81\x10\xf4\x2e\x0f\xbc\x8d\x1b\x5e\x29\x79\xcd\x0e\x70\x74\x97\x86\x27\x6b\xf0\x4e\x79\xd4\xbe\x16\x95\x63\xef\xfa\xfd\x2c\xaf\xec\xf0\x72\xc7\x6b\x35\xbc\xca\x36\x08\x05\xd1\xe8\x4b\x9b\xdf\x6d\x02\xdf\x29\x63\x61\x57\x83\x29\xe8\x07\x72\xcd\xac\xf4\x16\x7b\x1d\x51\x6a\xf1\xde\xee\x78\xab\xe4\x0d\x5a\x80\x64\xc3\x1a\x73\x24\x45\x64\xfe\xef\x30\x0a\xa3\x70\x44\xe6\xf2\x0e\xb5\x61\x0a\xda\x8c\xb9\x32\xb4\xb0\xe3\xac\x31\x9b\xc0\x14\x54\x64\x4e\xf1\x93\x01\x7d\x34\x8a\xe2\x32\x7c\xa4\x0f\xb2\xc8\xa6\xcb\xfa\xba\xc4\xb0\xac\x8c\x4c\xb8\xba\x41\x9f\x31\x53\x96\x2e\x3a\x3a\x9e\xfb\x5f\xc3\xe6\x8a\x86\x75\x20\x37\x7e\x54\x15\x3b\xcd\x09\xd9\xbf\x26\xce\xe6\xe8\xa2\xa4\x87\x29\xe8\xda\x79\x0a\xfc\x60\x91\x10\x89\xa6\xdc\xf3\x2a\x41\xfe\xbf\x00\x27\x81\x12\x09\x00\x40\x5a\xb5\x99\xdf\xcc\x51\xab\xc1\x8b\xb5\x61\xbf\xcf\x5f\x35\x40\xe8\x2f\xe8\x50\x95\x99\x04\x1a\xba\x99\x77\x29\xd4\x58\xde\x86\x04\x3f\x85\x06\x33\x5d\x0d\xae\x51\xba\x0a\x21\x54\x99\x61\x60\x93\x80\x10\x1c\x37\x7a\xbc\x57\xfe\x7e\x7b\xe8\x35\x6b\xd0\x7f\x78\x71\x5b\x7e\xc4\x7e\xd7\x1f\xcf\xe4\x78\x5c\xf4\x67\xaa\x17\x93\x92\xea\x60\xd8\x3f\x95\xea\xf3\xfa\xb8\xf1\xd9\x2b\x5e\xf1\x72\x1d\x8a\x76\x99\x9f\x7e\x5c\xf0\x36\x6e\x62\x49\x7e\x5b\xe0\xcd\xd7\xc1\x04\x90\xa4\x61\xdd\x24\x74\x3d\xe3\x58\x3a\xf1\x1b\xaf\x20\x18\x4b\xc2\x05\x38\x23\x5f\x45\x6e\xe3\xdb\xbb\xd1\x60\x20\x8e\xae\x45\xbd\x23\xad\xa2\x5e\xb3\x26\x5e\x36\x47\x8d\x37\xc3\xaf\x05\x71\xd5\x08\x5c\xfc\xe6\x67\x7e\xf7\x9a\x99\x8b\xa2\xc0\xe9\xa7\x1b\xdf\xba\xf0\x2e\x0f\x02\xb7\xd9\x7f\x44\xbb\x1c\xe3\xc5\xed\x71\xed\x64\xda\x0b\xe7\x72\x0f\x78\x66\xe6\x22\x68\xda\x23\x38\x86\x81\x47\xa0\x70\x98\xc3\xa6\x24\xb1\xa4\xdc\xc3\x7c\xf4\x97\xc4\x3b\xb8\xe3\xbd\x73\x71\x7e\xcc\x4b\x7b\xd3\xc2\x2c\xd6\x49\x3e\x8b\xcd\x95\x3f\xaa\xc7\x1b\xad\x61\xbf\xcb\x5b\x87\x62\xbf\xcb\xcb\xd5\x71\xe1\xf9\x64\x59\x6a\x12\x03\xeb\xb9\xc7\x54\x92\xb3\xfe\x62\xd8\xef\x7a\xb7\xbb\xfc\xd3\x5b\x7e\xba\xe9\xdf\x96\x07\xc9\xc9\xa5\x59\xd6\x73\x34\xa3\x62\x55\x93\x47\x53\x09\xfb\x7f\xca\x54\x36\x43\xd7\x2d\xdd\x96\x93\x89\xf8\x02\x42\x08\xdd\xdb\x88\x0f\x05\x71\x72\x16\xc0\xa8\x1a\x55\x57\x33\x52\x15\x74\x0d\xc1\xc8\x3d\x4b\x5a\x52\x2e\xfd\x08\x00\x00\xff\xff\x32\x0f\x9c\xa4\x33\x04\x00\x00"

func confAppConfBytes() ([]byte, error) {
	return bindataRead(
		_confAppConf,
		"conf/app.conf",
	)
}

func confAppConf() (*asset, error) {
	bytes, err := confAppConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "conf/app.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"conf/app.conf": confAppConf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"conf": &bintree{nil, map[string]*bintree{
		"app.conf": &bintree{confAppConf, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
