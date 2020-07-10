// Code generated by forge bindata.
// sources:
// pg/001_add_ports_table.sql
// DO NOT EDIT!

package migrations

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _pg001_add_ports_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd1\x41\x4b\xf3\x30\x00\xc6\xf1\x7b\x3e\xc5\x73\x4b\xc3\xbb\x5d\x5e\xd8\x69\xa7\x6c\xcb\x34\x5a\xdb\x99\xa6\xe2\x4e\x23\xb4\x41\x0b\x36\x29\x31\x65\x4e\xf1\xbb\x8b\xab\xcc\xde\x36\xf0\xfe\xe3\xcf\x03\xcf\x74\x8a\x7f\x6d\xf3\x14\x4c\xb4\x28\x3b\xb2\x54\x82\x6b\x01\xcd\x17\xa9\x80\x5c\x23\xcb\x35\xc4\xa3\x2c\x74\x81\xce\x87\xf8\x4a\x12\x02\x00\xbd\x7b\xf1\x95\xaf\x2d\x80\x07\xae\x96\xd7\x5c\x25\x33\x86\x32\x93\xf7\xa5\xc0\x46\xc9\x3b\xae\xb6\xb8\x15\xdb\xc9\x51\x3b\xd3\x7e\x4b\x8c\xf4\xff\xd9\x8c\x1d\xe3\x59\x99\xa6\x83\xaa\x9a\x78\xb8\x40\xf9\xde\xc5\x70\x38\xa3\x62\xd3\xda\x77\xef\xec\xb9\x56\x7d\xc1\x2e\xfb\x16\x83\x19\xd4\x4d\x91\x67\x0b\x60\x25\xd6\xbc\x4c\x35\xe8\xc7\x27\xfd\x29\x05\x6b\xa2\xad\x77\x26\x62\x21\xaf\x64\xa6\x4f\xa6\x36\xd1\xee\x3a\x13\x62\x42\x6d\xe7\xab\x67\x3a\x39\x8d\x4b\x68\x1f\x2b\x3a\x81\xf3\xfb\x84\x31\x36\x94\xfa\xae\xfe\x73\x89\xb0\x39\x21\x64\xfc\xec\xca\xef\x1d\x59\xa9\x7c\xf3\xfb\xec\xf8\xd5\x39\xf9\x0a\x00\x00\xff\xff\x18\xcd\x3d\x6e\x07\x02\x00\x00")

func pg001_add_ports_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_pg001_add_ports_tableSql,
		"pg/001_add_ports_table.sql",
	)
}

func pg001_add_ports_tableSql() (*asset, error) {
	bytes, err := pg001_add_ports_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pg/001_add_ports_table.sql", size: 519, mode: os.FileMode(420), modTime: time.Unix(1594331026, 0)}
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
	"pg/001_add_ports_table.sql": pg001_add_ports_tableSql,
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
	"pg": &bintree{nil, map[string]*bintree{
		"001_add_ports_table.sql": &bintree{pg001_add_ports_tableSql, map[string]*bintree{}},
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
