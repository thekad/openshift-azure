// Package proxyinfrastructure Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/start.sh
package proxyinfrastructure

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

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _startSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\x4f\x4f\x1b\x31\x10\xc5\xef\xfe\x14\x53\x5a\x95\x93\xd7\x4a\x8b\x50\x85\x08\x12\xa4\x41\x54\xa8\x02\x29\xfd\x73\x40\x1c\x26\xf6\xe0\x58\x78\xed\x95\x3d\xa6\xbb\x21\xf9\xee\xd5\x66\x29\xd0\xa0\x56\x5c\xd6\x96\xdf\xcc\xdb\xb7\xbf\x1d\xbf\x7d\xa3\xe6\x2e\xa8\x39\xe6\x05\x48\x6a\x85\xa0\x96\x34\x7c\x38\x7a\x3f\xea\xb7\x4d\x4c\x0c\x67\x17\x5f\xa7\x63\x95\x62\x64\xa1\x8d\x10\x5d\xa9\xc1\x85\xcc\xe8\x3d\x58\xc7\x20\x3b\x21\x74\x49\x1e\x64\x86\x05\x73\x93\x0f\x94\xca\x1c\x13\x5a\xaa\x6c\x8c\xd6\x13\x36\x2e\x57\x3a\xd6\xca\x46\x8f\xc1\x2a\x1b\x47\xd5\x68\x54\xed\x57\xde\x85\xd2\x4a\xac\xcd\xfe\x5e\xc5\x98\x2a\xbb\x84\x15\x30\x26\x90\x13\x50\x25\x27\xe5\xa3\x46\x0f\xb2\x5d\x8a\xcb\xe3\x6f\x67\xe3\x77\xfd\xf3\xe0\x49\x51\x36\xf6\xe9\x85\xb0\x11\x2c\x31\xc8\xd2\x27\x5a\x94\xf9\xe6\x6d\xb1\xa1\x90\x17\xee\x86\x9f\x76\x12\x97\x25\x91\xd2\xb5\x51\x4d\x8a\x6d\x27\x44\x7d\x6b\x5c\x02\x83\x8c\x42\x23\xc3\xd1\xe6\x58\x6a\x4a\x5c\x35\x54\xc3\xe1\xe1\xee\xf4\xe2\x74\x57\xdc\xdf\x43\x35\x89\xe1\xc6\xd9\x6a\x46\xe9\x8e\xd2\x84\x12\xc3\x0a\xfa\xe5\x38\x9f\x74\x4c\x19\x56\x30\xe3\xe4\x82\x85\xf5\x5a\x4c\x2f\x4e\xc5\x73\xc3\x5b\xea\xfe\xef\x77\x4e\x1d\xac\xe0\x32\xb9\x3b\x64\x3a\xa7\xee\x35\xa6\x1a\xff\xe5\x39\xc1\x57\x65\x53\xc4\x5a\xe5\x2e\x33\xd5\xe6\x61\x55\x98\xa2\xdc\xd8\x57\x99\xd2\x9d\xd3\xf4\xe8\x7f\xf5\x3d\x38\xbe\x16\x9f\x29\xeb\xe4\x1a\x76\x31\x8c\x1f\x8b\x85\xb8\x9a\x0d\xe5\xd7\x62\xda\x92\x9e\x31\x26\x1e\x86\xe6\xe1\x27\x0d\xc0\x61\x83\x16\x06\x61\x8b\x75\x0f\xe9\x2f\xe5\x0f\x34\xa9\xf1\x65\xd7\xf0\xed\x32\x97\x79\x20\x86\x9d\xe7\x40\x37\x47\xb9\xfa\xd1\x0b\xeb\xf5\x8e\x10\x57\x5f\x86\x79\xbd\x16\x3f\x31\x30\x99\x93\x6e\x5c\x17\xcf\x4e\x96\x4c\xa9\x9f\x3c\x4b\x3c\x60\xc9\xc5\x44\x18\x48\x68\xf6\x60\x90\xea\x18\x64\x22\x1f\xd1\xbc\x50\x29\xe0\xdc\x13\x3c\x41\xd8\xd2\x73\x0f\xe1\x99\x2c\x6e\x5c\xa2\x5f\xe8\xbd\xd4\xb5\x01\x29\x97\x31\xd0\xb8\x29\x73\xef\x34\x48\x89\xc6\xc8\xfe\xbe\x8d\x3f\xed\xed\x7d\x54\xac\x1b\x90\xb2\xa1\x54\x63\xa0\xc0\xdb\xad\x0f\x89\x7e\x07\x00\x00\xff\xff\xee\x67\xf7\x3c\xbe\x03\x00\x00")

func startShBytes() ([]byte, error) {
	return bindataRead(
		_startSh,
		"start.sh",
	)
}

func startSh() (*asset, error) {
	bytes, err := startShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "start.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"start.sh": startSh,
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
	"start.sh": {startSh, map[string]*bintree{}},
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