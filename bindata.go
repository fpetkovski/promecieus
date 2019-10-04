// Code generated by go-bindata.
// sources:
// prom-templates/deployment.yaml
// prom-templates/route.yaml
// prom-templates/service.yaml
// DO NOT EDIT!

package main

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

var _promTemplatesDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xcb\x6e\xdb\x3a\x10\xdd\xeb\x2b\xe6\xe6\x76\x2b\x3f\x52\x74\x11\xee\x02\xa4\x40\x17\x71\x6b\x04\x4d\xb6\xc5\x98\x1e\x5b\xac\xf9\xea\x70\xa8\xd8\x4d\xf2\xef\x85\x22\x4b\x66\x9a\xac\x5a\x54\x2b\xea\xbc\x66\x70\x48\x8c\xe6\x8e\x38\x99\xe0\x15\xd0\x5e\xc8\x77\xc7\x34\x6d\xe7\x2b\x12\x9c\x57\x3b\xe3\xd7\x0a\xae\x28\xda\x70\x70\xe4\xa5\x72\x24\xb8\x46\x41\x55\x01\x78\x74\xa4\x20\x72\x70\x55\x8a\xa4\x3b\x88\x29\x5a\xa3\x31\x29\x98\x57\x00\x42\x2e\x5a\x14\xea\x18\x80\xd2\xda\x7d\x85\xbd\x07\x2c\xae\xc8\xa6\x81\x06\xc0\x18\x0b\x7e\x18\xd1\x7d\x3a\x78\x41\xe3\x89\x47\x75\x7d\x8c\xd3\xa6\xde\x90\xe8\x86\x78\x8c\x31\x0e\xb7\xa4\x80\x69\x6b\x92\xf0\x61\xb2\xa1\x75\x60\x8c\x1c\xbe\x93\x96\x49\xe0\xed\xb4\x47\xd4\xfb\xd9\xe8\xd1\xc1\x39\xf4\xeb\xd3\x2e\x35\x4c\x57\xc6\x4f\x57\x98\x9a\x02\xab\x75\xf1\xf3\x38\x9e\x01\xfe\xff\xef\xb5\x1c\x20\x91\x40\x4d\xfb\x1c\x20\x9a\x48\x1b\x34\xb6\x20\x75\x66\x0b\xf5\x35\xbc\x7b\x58\xde\x7c\x59\x7c\xbd\xbc\x79\x82\x47\x10\x64\xd8\xb7\x3f\x0b\x59\xdc\x19\x6b\xa1\xfe\x74\xbb\xec\x37\xea\xfa\x21\x69\x28\xa7\x51\x44\xbe\x55\x85\x63\xa8\xe6\x18\x5b\x30\x00\x2d\xda\x4c\x0a\xce\x8e\xdc\xb7\xbb\xcb\xeb\xdb\x8f\x67\xa3\xe4\x3e\xf0\xce\xf8\xed\x95\x61\x05\xc5\xa4\xe9\x28\x68\x83\xcd\x8e\x16\x21\x7b\x49\x65\x59\xa7\xbb\xed\x1d\x75\x92\xc0\xb8\xa5\xba\x37\x14\x3b\xb8\xce\xbb\x44\x69\xde\x9c\xf0\x3a\xe9\xf7\x6b\xed\x98\xc2\xa8\xda\xf3\xc9\xfc\x7c\x72\xba\xc9\x18\xf8\xad\xd5\xee\x69\x95\x4d\xd9\xfe\xf0\xa2\x96\x81\x45\xc1\xc5\xec\xe2\x14\xc1\x94\x42\x66\x4d\xa9\x6c\x95\xe9\x47\xa6\x24\x2f\x30\x00\x1d\xb3\x82\x0f\xb3\x99\x7b\x81\x3a\x72\x81\x0f\xcf\xc4\xa2\x20\xac\x71\xe6\x2f\x02\xfe\x75\xf7\xa9\x41\xa6\x25\x07\x4d\x29\x7d\x46\x47\x29\xa2\x26\x05\xc2\x79\xc8\xe8\x03\xff\x70\x38\xb9\x28\x87\xe7\x97\xf5\xf0\x54\xfd\x0a\x00\x00\xff\xff\x03\x0f\xb5\xc1\x87\x04\x00\x00")

func promTemplatesDeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesDeploymentYaml,
		"prom-templates/deployment.yaml",
	)
}

func promTemplatesDeploymentYaml() (*asset, error) {
	bytes, err := promTemplatesDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/deployment.yaml", size: 1159, mode: os.FileMode(420), modTime: time.Unix(1570192159, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesRouteYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x4e\x03\x31\x0c\x44\xef\xf9\x0a\x7f\x41\xd9\x1e\x9b\x3b\x57\x54\x15\xc4\xdd\x4a\x86\xd4\x62\x13\x47\x8e\xdb\x8a\xbf\x47\xd9\x05\xc1\x25\xd2\xbc\x37\x19\xf9\x53\x5a\x8e\x74\xd1\x9b\x23\x70\x97\x77\xd8\x10\x6d\x91\x6c\x92\x83\x76\xb4\x71\x95\x0f\x3f\x88\x3e\xdd\x8f\xa1\xc2\x39\xb3\x73\x0c\x44\x8d\x2b\x22\x75\xd3\x1a\x46\x47\x9a\xc8\x75\xbe\x44\xfb\xe8\x2b\xec\x2e\x09\x1b\xf9\x57\x9e\xf1\x01\x29\x57\x8f\x74\x5c\x96\x40\xd4\xd5\x7c\xff\xe8\x6c\x05\x7e\x9e\x99\x4e\xcb\x69\x4a\x5f\xc7\x8f\x83\x55\x69\xec\xdb\x79\xc8\x65\x1f\x96\x36\x90\x6e\x86\xe7\x5c\xf0\xf6\xd7\x38\xeb\x2a\xe9\x2b\xd2\x05\x59\x0c\xc9\x03\xd1\x43\xd6\x9c\xd8\xf2\xaf\x7a\xd1\x86\xf0\x1d\x00\x00\xff\xff\x84\xf0\x9a\x83\xff\x00\x00\x00")

func promTemplatesRouteYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesRouteYaml,
		"prom-templates/route.yaml",
	)
}

func promTemplatesRouteYaml() (*asset, error) {
	bytes, err := promTemplatesRouteYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/route.yaml", size: 255, mode: os.FileMode(420), modTime: time.Unix(1570192158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcc\x31\xae\xc2\x30\x0c\xc6\xf1\x3d\xa7\xf8\x2e\xf0\xa4\xbe\xb1\x5e\xb9\x00\x12\x88\xdd\x4d\x3d\x44\x24\xb1\x95\x98\x72\x7d\x94\x92\x89\xd1\x3f\x7f\xfa\xb3\xa5\x87\xb4\x9e\xb4\x12\x8e\xff\xf0\x4c\x75\x27\xdc\xa4\x1d\x29\x4a\x28\xe2\xbc\xb3\x33\x05\xa0\x72\x11\x82\x35\x2d\x01\xc8\xbc\x49\xee\x83\x01\x36\x9b\xde\x4d\xe2\x30\xd3\xe6\xf3\xf9\x77\x1e\x84\x75\x59\x97\x13\x30\xa6\xae\x51\x33\xe1\x7e\xb9\x4e\xfb\xc6\xdf\xb2\xbd\x52\x00\xba\x64\x89\xae\xed\xb7\xff\x09\x00\x00\xff\xff\x6d\xe1\xdb\x34\xac\x00\x00\x00")

func promTemplatesServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesServiceYaml,
		"prom-templates/service.yaml",
	)
}

func promTemplatesServiceYaml() (*asset, error) {
	bytes, err := promTemplatesServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/service.yaml", size: 172, mode: os.FileMode(420), modTime: time.Unix(1570192162, 0)}
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
	"prom-templates/deployment.yaml": promTemplatesDeploymentYaml,
	"prom-templates/route.yaml": promTemplatesRouteYaml,
	"prom-templates/service.yaml": promTemplatesServiceYaml,
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
	"prom-templates": &bintree{nil, map[string]*bintree{
		"deployment.yaml": &bintree{promTemplatesDeploymentYaml, map[string]*bintree{}},
		"route.yaml": &bintree{promTemplatesRouteYaml, map[string]*bintree{}},
		"service.yaml": &bintree{promTemplatesServiceYaml, map[string]*bintree{}},
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

