// Code generated by go-bindata.
// sources:
// prom-templates/deployment.yaml
// prom-templates/kustomization.yaml
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

var _promTemplatesDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xb1\x6e\x02\x31\x0c\x86\xf7\x7b\x0a\xbf\xc0\xc1\x1d\x52\x07\x32\xb3\x22\x31\x75\x37\xc7\x2f\x88\x1a\x27\x69\xe2\xd0\xa2\xaa\xef\x5e\xa5\x85\x23\x48\x9d\x5a\x4f\x77\xff\x67\x7f\x56\xcc\xd1\x3e\x23\x65\x1b\xbc\x21\xbc\x2b\x7c\xfd\xcc\xcb\xf3\xb8\x87\xf2\xd8\xbd\x58\x7f\x30\xb4\x41\x74\xe1\x22\xf0\xda\x09\x94\x0f\xac\x6c\x3a\x22\xcf\x02\x43\x31\x05\xb9\xfe\xe4\xc8\xd3\x35\xe9\x15\x59\xbb\x1c\x31\xd5\xce\x84\xe8\xec\xc4\xd9\xd0\xd8\x11\x29\x24\x3a\x56\x54\x42\xd4\x1a\x6b\x3d\x58\x6b\x39\xde\xc3\xe5\x1b\x26\xe2\x18\x1b\x7e\x5b\x51\x6b\x0a\x5e\xd9\x7a\xa4\xb9\xbb\x6f\x74\xd0\x13\x4a\x9e\x35\x56\xf8\x78\x25\xcb\x3b\x36\xe7\xd5\x62\x5c\x2d\x86\xb9\x2d\x86\xa4\xcd\xf2\x9b\xf0\x0d\xfb\x62\xe7\xb4\x59\xbd\x0b\x49\x0d\xad\x87\xf5\x5d\x91\x90\x43\x49\x13\x1a\x4d\x0d\x5f\x0b\xb2\x3e\x64\x44\x53\x2c\x86\x9e\x86\x41\x1e\x52\x81\x84\x74\xf9\x06\xdb\x06\x38\x2b\xf6\x1f\x82\x73\x70\x45\xb0\x0d\xc5\xff\xf6\xc0\xfb\x49\xfa\xac\x21\xf1\x11\xfd\xcf\x40\x23\x96\x3a\xbb\x63\x3d\x19\x6a\x4e\xb8\xec\x5a\xff\x1f\xd5\x90\xa8\x97\x8d\x4d\x86\x3e\x3e\xbb\xaf\x00\x00\x00\xff\xff\x24\xc1\x25\xfa\xa5\x02\x00\x00")

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

	info := bindataFileInfo{name: "prom-templates/deployment.yaml", size: 677, mode: os.FileMode(420), modTime: time.Unix(1570184474, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesKustomizationYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xc8\xc1\x09\x80\x30\x0c\x05\xd0\x7b\xa7\xc8\x02\x75\x80\x6e\x53\xe2\x3f\x08\x4d\x13\x92\x54\xe8\xf6\x82\x08\x1e\xdf\x9b\x5d\x10\xd6\x19\x8d\xcc\x55\x6a\x22\xb2\x38\x42\x97\x33\xa2\x15\xa2\x4a\x27\x6c\xe8\x16\xcc\x3c\x76\x97\xf1\x5e\xc0\xef\x8b\xf1\x87\xeb\xca\x8f\x4f\x00\x00\x00\xff\xff\x39\xd8\xfe\xde\x54\x00\x00\x00")

func promTemplatesKustomizationYamlBytes() ([]byte, error) {
	return bindataRead(
		_promTemplatesKustomizationYaml,
		"prom-templates/kustomization.yaml",
	)
}

func promTemplatesKustomizationYaml() (*asset, error) {
	bytes, err := promTemplatesKustomizationYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "prom-templates/kustomization.yaml", size: 84, mode: os.FileMode(420), modTime: time.Unix(1570185135, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesRouteYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xc1\x4e\x03\x31\x0c\x44\xef\xf9\x0a\xff\x00\x65\x7b\x6c\xee\x5c\x51\x55\x10\x77\x2b\x19\xb6\x16\x9b\x38\x72\xdc\x56\xfc\x3d\xca\xee\x22\xc4\xc5\xd2\xcc\x1b\x8f\xfd\x25\x35\x47\xba\xe8\xcd\x11\xb8\xc9\x07\xac\x8b\xd6\x48\x36\x9c\x83\x36\xd4\x7e\x95\x4f\x3f\x88\x3e\xdf\x8f\xa1\xc0\x39\xb3\x73\x0c\x44\x95\x0b\x22\x35\xd3\xb2\x8b\xde\x38\xed\xce\x93\xa3\x7b\xe8\x0d\x69\x24\x5d\xc7\x24\xda\x6e\xbd\xc1\xee\x92\xb0\x3a\xff\x3a\x88\x1e\x90\xf9\xea\x91\x8e\xd3\x14\x88\x9a\x9a\x6f\x8b\xce\x36\xc3\xcf\x43\xd3\x69\x3a\x0d\xe8\x4b\xdf\x19\xac\x48\x65\x5f\xbf\x46\x9e\xb7\x62\xa9\x1d\xe9\x66\x78\xc9\x33\xde\xff\x12\x67\x5d\x24\x7d\x47\xba\x20\x8b\x21\x79\x20\x7a\xc8\x92\x13\x5b\xfe\x45\xaf\x5a\x11\x7e\x02\x00\x00\xff\xff\xfc\xa8\x43\x2e\x16\x01\x00\x00")

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

	info := bindataFileInfo{name: "prom-templates/route.yaml", size: 278, mode: os.FileMode(420), modTime: time.Unix(1570184236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _promTemplatesServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x31\x0e\xc2\x30\x0c\x45\xf7\x9c\xe2\x5f\xa0\x52\x19\xeb\x95\x0b\x20\x81\xd8\xdd\xd4\x43\x44\x5a\x5b\xb1\x29\xd7\x47\x2d\x99\x18\xdf\xf3\x93\x3f\x5b\x79\x4a\xf3\xa2\x1b\x61\xbf\xa4\x57\xd9\x16\xc2\x5d\xda\x5e\xb2\xa4\x55\x82\x17\x0e\xa6\x04\x6c\xbc\x0a\xc1\x9a\xae\x1d\xdc\x38\x77\x33\x84\x78\x24\xa0\xf2\x2c\xd5\x8f\x1a\x60\xb3\x9e\xbb\x49\x3e\x9c\x69\x8b\x7e\x1c\x4e\x20\x4c\xe3\x34\x9e\x02\x47\x1a\x9a\xb5\x12\x1e\xd7\x5b\x77\xbf\xcd\x8f\xcc\xef\x92\x00\x97\x2a\x39\xb4\xfd\xff\xff\x06\x00\x00\xff\xff\x85\x74\xfb\xbd\xc3\x00\x00\x00")

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

	info := bindataFileInfo{name: "prom-templates/service.yaml", size: 195, mode: os.FileMode(420), modTime: time.Unix(1570184199, 0)}
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
	"prom-templates/kustomization.yaml": promTemplatesKustomizationYaml,
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
		"kustomization.yaml": &bintree{promTemplatesKustomizationYaml, map[string]*bintree{}},
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
