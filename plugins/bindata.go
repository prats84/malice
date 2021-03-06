// Code generated by go-bindata.
// sources:
// plugins/.DS_Store
// plugins/plugins.toml
// plugins/templates/.DS_Store
// plugins/templates/python/.DS_Store
// plugins/templates/python/.dockerignore
// plugins/templates/python/.gitignore.template
// plugins/templates/python/CHANGELOG.md
// plugins/templates/python/Dockerfile
// plugins/templates/python/LICENSE.md
// plugins/templates/python/README-short.txt
// plugins/templates/python/README.md
// plugins/templates/python/circle.yml
// plugins/templates/python/plugin.toml
// plugins/templates/python/requirements.txt
// plugins/templates/python/scan.py
// DO NOT EDIT!

package plugins

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

var _pluginsDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x4f\x4a\xc5\x30\x10\xc6\xbf\x89\x5d\x44\xdc\x64\xe9\x32\x57\xf0\x06\xe1\xf1\x3c\xc1\xbb\x80\x82\xe2\xa6\x56\xf1\xdf\x3a\x2b\xcf\xe5\xd1\x6c\x9c\x4f\x2c\xb4\x5d\x08\x42\x45\xbf\x1f\x0c\xbf\x07\x6f\x26\x6d\x37\x99\x4c\x00\xd8\xee\xf9\xea\x0c\x48\xe3\xcf\x08\x37\x5e\xb1\x48\x64\xcc\x08\xb4\x79\x8c\x6b\xf4\xb8\xc4\x13\xae\xf1\xf8\x72\x58\x5e\x6b\x46\xab\x3d\xfe\xa8\xb9\xc5\xfd\xb4\xfe\x61\xe8\xef\x86\x1b\x66\x08\x21\x84\x10\xe2\xa7\x60\x67\x8d\x27\xdb\xbe\x86\x10\xe2\x17\xd2\xf6\x87\x4c\x17\xba\xba\x8d\xff\x07\xba\x9b\xd4\x24\x3a\xd3\x85\xae\x6e\x63\x5e\xa0\x3b\x3a\xd2\x89\xce\x74\xa1\xab\x9b\x9b\x96\x71\xf8\x30\x3e\xd9\x38\xa1\x58\xa2\x33\x5d\xbe\xf9\xd1\x42\xfc\x13\x8e\x5c\xa9\xf5\xff\x73\xac\xce\xff\x42\x88\x3f\x8c\x75\xfb\xc3\x7e\x87\xf5\xab\xb6\xd6\x6b\xf3\x18\x17\xcc\x79\xfb\x2c\x5c\x39\x08\x04\xbf\x30\x3c\xc5\x57\x5e\xa6\x0b\x5d\xdd\x3a\x0c\x08\xb1\x05\xef\x01\x00\x00\xff\xff\xaf\x38\x06\x73\x04\x18\x00\x00")

func pluginsDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_pluginsDs_store,
		"plugins/.DS_Store",
	)
}

func pluginsDs_store() (*asset, error) {
	bytes, err := pluginsDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1475951010, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsPluginsToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x98\x5d\x6f\xdb\x36\x14\x86\xef\xfd\x2b\x0e\x9c\x9b\x6d\xa8\xe3\xb5\x43\x3b\xa0\xc0\x2e\x3c\xc7\x69\x93\x39\x89\x61\xb9\x01\xba\x20\x48\x69\x89\x92\xb9\x48\x22\x41\x52\x4e\x9c\x5f\x3f\x1e\xfa\x43\xfe\x60\x1a\x51\x56\x2e\x22\x9b\x87\x7c\x0f\x9f\xf7\x50\x14\xe5\x13\xe8\x73\xb1\x90\x2c\x99\x69\xf8\x25\xfc\x15\x3e\xfc\xfe\xfe\x0f\xe8\xe0\xe5\x13\x4c\x53\x12\x3e\x6a\x2e\xe0\x92\xab\x59\x41\xe0\x8a\xb0\x9c\xbe\x83\x5e\x9a\xc2\x18\x07\x28\x18\x53\x45\xe5\x9c\x46\xa7\xad\x13\x08\x28\x85\xe1\x45\x7f\x70\x1d\x0c\x20\xe6\x12\x52\x16\xd2\x5c\x51\x60\xb9\xf9\x96\x11\xcd\x78\x7e\xda\x6a\x9d\x34\xf3\x67\xf2\x8d\x86\xdf\xbe\x5c\x5c\x9b\xe9\xe7\x31\x4b\x0a\x69\x13\x80\xbf\x4e\x43\xf3\x69\x69\xa6\x53\x0a\x7f\x41\xfb\x8a\x20\x39\x8c\xd2\x22\x61\xf9\xee\xf4\x54\xbb\xd5\xba\xbb\x13\x36\x72\x7f\xdf\x02\xa0\x39\x99\xa6\x34\x32\xc3\x62\x92\x2a\x6a\x5a\x72\x92\x59\x95\x5c\xc9\xb4\x6d\xbe\x47\x54\x85\x92\x09\x0b\x67\x9a\xaf\x83\xf1\x10\xce\x88\x26\x53\x62\xac\xfd\x4a\xd4\xcc\xf8\x4e\x64\x38\xc3\xbe\x21\xd1\x34\xe1\x72\x81\x1d\x59\xae\xa9\x15\x60\x19\x49\xac\x62\x66\xe7\xd5\x5d\x0b\x4b\x2a\xb8\x62\x7a\xd5\x7d\xa6\xb5\x50\x9f\xbb\xdd\x84\xe9\x59\x31\x3d\x0d\x79\xd6\x5d\xf6\x67\x7c\xf5\xa1\x83\x03\x4f\x4d\x1c\x07\x4f\x0b\x96\x6e\xcf\x3a\x63\xcb\x59\xcf\xcc\x84\x30\x8e\x57\xbd\x10\x54\x99\xc6\x3b\x93\x3a\xfa\xd8\x86\xfb\x57\xd9\xb5\x2c\xb6\xd1\xe7\x4c\x16\x4a\x73\x4d\x5c\x06\xdc\x62\x70\x82\x41\xb3\x4a\x63\x96\x9a\x14\x2a\x24\x39\x90\x3c\xb2\x59\x21\xe5\xfc\xb1\x10\x95\xed\xd8\x4d\xe6\x6d\x4a\x39\xfc\x35\x6b\x0a\x91\x72\xb2\x69\x00\xbc\x4f\xf4\xba\x51\x73\xf8\x81\xf0\x3f\x80\xc5\xb0\xe0\x05\x3c\x91\x5c\x63\xeb\x2a\xae\x48\x26\xcc\xa2\x32\x0d\x5b\x79\xcc\x34\x8c\x2c\x11\xec\x91\xda\x59\x5a\xd4\x0c\x33\xb4\x4b\xf6\x0a\x05\x79\x07\x6d\x35\x23\xef\x57\xd7\x0f\x1f\x3f\x61\x89\xb0\x2c\x73\xec\xd3\xbe\xea\xe1\xcd\xfc\x70\x3b\x79\xe8\x8d\x2e\xda\x95\x8b\x67\xa4\x22\xfe\xd4\xb1\xdb\x82\x74\xd4\x2f\xb0\xf1\xc0\x86\x4d\x05\xeb\x94\xec\x20\x85\x77\xd5\x76\x14\x8e\x58\xd3\x6b\x0b\x7f\xb2\xb6\xf7\xef\x6b\x4d\x49\xd6\x09\x17\x99\x2c\x1c\xe6\x4c\x4c\xb0\x8f\xb1\x9a\xce\xec\x8a\x7b\xdb\x52\x0e\x6f\xc6\x93\x83\xc5\x34\xf9\xfa\xf0\x2d\x18\x8c\xb1\x4f\xd9\xf2\xcf\xe0\x7b\xf5\xe5\x85\x77\x3c\x3e\x50\x1c\xe6\x29\x15\x51\x2a\xba\x93\x31\x3b\xeb\xd2\x67\x16\x6b\xce\xd3\x7d\xe7\x32\xaa\x49\x64\x76\x4f\x97\x79\xdb\xd2\xde\xd6\xad\x07\xbf\x65\xdc\x6f\xaf\x3f\x01\xf6\x48\x17\x44\x12\x07\xe5\xf7\xde\xb8\x07\x81\xd9\xf1\xf6\xd1\xc8\xdc\x05\xb5\x56\xf1\x06\xc2\x81\x47\xc0\xec\x2f\x7b\x32\x27\x4a\x3b\x70\x7a\xd8\x0e\xbd\x5c\x33\xbb\xb1\x57\x83\xda\x88\x79\x53\x91\x39\x32\x9d\x7c\xde\x28\x1c\x5f\x27\x32\x4f\x5c\x5c\xb7\x5f\xfc\xa9\x92\x23\x99\x92\x66\x88\xa6\x4c\x47\x34\xa6\x79\xe4\xdc\xc0\xff\x2e\xa3\xbe\x84\x7b\xc2\xb5\x49\xf7\x74\x8e\x27\x0e\x53\x92\x2d\xa7\xbc\x07\xdb\x37\x81\xde\x6d\x35\xba\x52\xa4\x36\x58\x29\xd1\x00\x13\xcf\x78\xe4\xda\x27\xfb\x36\xe0\x5b\xbb\x52\xae\x3e\xdd\x46\xe2\x78\xba\x58\x48\xee\xda\x4f\xce\x3b\xa3\xf1\xcd\xc4\x17\x6e\xa3\x56\x9b\x6d\xa3\xd0\x00\x9a\x39\x97\x84\x85\xa4\x4e\xba\xc0\x86\xbc\xf9\xb6\x24\xeb\x23\x6e\x89\x1c\x4f\xa9\xb8\x98\x71\xe5\x3a\x20\xda\x80\x2f\x61\x29\x57\x9b\xaf\x94\x68\xe0\x79\x27\x5c\xe5\x1b\x0d\xcc\xb9\x0e\x8f\x25\x78\xa2\xd7\x92\x21\x82\xe0\x52\xa3\x08\xd0\x67\xe3\xae\xfd\x78\x80\x6c\x42\x2e\x66\x51\xaf\x9e\x82\xbe\xf5\x5c\x27\x42\x98\x9e\xf6\x8d\xb5\xfb\xdc\x89\xb8\xc2\xb9\x55\x5f\xbf\x29\x57\xae\xc2\x9e\x33\x49\x07\x0b\x0a\x43\x32\x55\x70\x33\x8d\x0b\x85\x8c\x11\x04\xc6\x89\x3c\x81\x80\xa7\xab\xe3\x7c\x05\xf4\x4d\x0e\xff\x83\x1a\x8e\x6c\xd8\x80\xfd\xe2\xf3\x38\x36\xa9\x1c\x16\xdc\xd8\xc0\xe1\x22\xb8\x19\x0e\xba\xe3\xc9\x39\x44\x3c\x2c\x32\x9a\xeb\x83\x15\xb0\x0e\xb8\xbc\x28\xb3\x79\x9b\xb1\x1c\xda\xe0\x31\x4f\x44\xb1\x6b\xdd\x9f\x9d\x1f\x32\x63\x63\x2d\xde\x55\x0e\xff\x75\x1f\xc5\x3e\x75\xb7\x69\xaa\x72\xff\x87\x07\x4a\x8b\xec\xc0\xbf\xdc\x04\x0f\x5d\xb8\x0c\x60\x19\xf2\xb2\x60\x37\x9d\xb7\x13\xe5\x70\x1f\x43\xb6\x93\x56\xf5\xe5\x85\x09\x87\x21\xff\x32\xb1\xe5\x44\x91\xe3\xef\x57\x6c\x4e\x61\x75\x3d\xdc\xf2\x97\xed\x2e\x27\x56\x09\xbc\x2d\x30\xe3\xde\x64\x2f\xb3\x9a\xde\x5b\x6f\xb9\x9a\x48\x7c\x83\x4d\x5e\xf0\xbf\xf8\xf3\x05\x5f\x75\xff\x0f\x00\x00\xff\xff\x0b\x2a\xaf\x31\x47\x15\x00\x00")

func pluginsPluginsTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsPluginsToml,
		"plugins/plugins.toml",
	)
}

func pluginsPluginsToml() (*asset, error) {
	bytes, err := pluginsPluginsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/plugins.toml", size: 5447, mode: os.FileMode(420), modTime: time.Unix(1475679267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x41\x4a\x04\x31\x10\x45\x7f\xc5\x46\x02\x6e\xb2\x74\x99\x2b\x78\x83\x30\x8c\x27\x98\x0b\xb8\x10\xda\x85\x74\x8b\xa8\xe0\x2e\x2b\xcf\xe5\xd1\xec\x50\x5f\x6c\xb5\x7b\xa1\x9b\x16\xe7\x3f\x08\x2f\x90\x54\xcd\xf4\x26\x95\x0a\x00\xdb\x3d\x5e\x5f\x00\x69\x9a\x46\xb8\xf1\x82\x45\x22\xc7\x37\xc2\xcc\xd6\x72\xdc\x8e\x43\xef\x73\x9c\x2e\xe7\x5a\xcc\x11\xd0\x63\x7c\x3a\xdc\x0f\x9f\xe2\xef\xf0\x8c\x07\xdc\x60\xc4\xf0\x65\x4d\x08\x21\x84\x10\xbf\x87\xb5\x34\x9e\x6d\xfb\x37\x84\x10\x7f\x90\x76\x3e\x64\xba\xd0\xd5\x6d\x5c\x0f\x74\x37\x8b\x49\x74\xa6\x0b\x5d\xdd\xc6\x7d\x81\xee\xe8\x48\x27\x3a\xd3\x85\xae\x6e\x1e\x5a\xc6\xe6\xc3\xf8\xcb\xc6\x0e\xc5\x12\x9d\xe9\xf2\xc3\x8f\x16\xe2\x48\x38\x71\xa5\x56\xff\x2f\xb1\xda\xff\x0b\x21\xfe\x31\xd6\xed\x0f\xfb\x1d\xd6\x1f\xd7\x5a\xad\xcd\xd3\xb8\xe2\x9e\xd7\xf7\xc0\x95\x8b\x40\xf0\x07\xc3\x73\x7c\xec\xcb\x74\xa1\xab\x5b\x97\x01\x21\xb6\xe0\x2d\x00\x00\xff\xff\x42\xfd\xbc\x7e\x04\x18\x00\x00")

func pluginsTemplatesDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesDs_store,
		"plugins/templates/.DS_Store",
	)
}

func pluginsTemplatesDs_store() (*asset, error) {
	bytes, err := pluginsTemplatesDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1475951231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDs_store = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\x3b\x0e\xc2\x30\x10\x44\x77\x8c\x0b\x4b\x34\x2e\x29\xdd\x70\x00\x6e\x60\x45\xe1\x04\x5c\x80\x82\x2b\xd0\xfb\xe8\x24\xda\x11\xb2\x14\x52\x50\x25\x82\x79\x92\xf5\x56\x8a\x9d\x4f\xe3\xec\xd8\xcc\x30\x3c\x1f\x17\xb3\x3c\x95\xc9\xdc\x76\xb6\x8f\x24\x8e\x05\xa1\xab\xc1\x7b\x08\x21\x84\x10\x62\xdf\xc0\x95\x8e\xdb\xbe\x86\x10\x62\x87\xcc\xfb\x43\xa1\x2b\xdd\xdc\xe0\xf5\x40\xc7\x6e\x4d\xa6\x0b\x5d\xe9\xe6\x06\xe7\x05\x3a\xd2\x89\xce\x74\xa1\x2b\xdd\xdc\xdc\xb4\xc0\xf0\x01\x3e\x19\x4c\x28\x60\x0a\x41\xa1\xeb\x97\x1f\x2d\xc4\x9f\x70\x70\xe5\xf9\xff\x7f\xb5\xd5\xfc\x2f\x84\xf8\x61\x10\xc7\xdb\x38\xd8\x3b\x10\x2c\x27\x4c\xe3\xde\xd5\xcd\xd6\x9b\x80\xe0\x87\x85\xa7\x6e\x6d\xa1\x2b\xdd\xdc\x6a\x04\x84\xd8\x8a\x57\x00\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func pluginsTemplatesPythonDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDs_store,
		"plugins/templates/python/.DS_Store",
	)
}

func pluginsTemplatesPythonDs_store() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1475951710, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xf0\x4c\xcf\xcb\x2f\x4a\x55\xd0\x4b\xcf\x2c\x51\x48\xcb\xcf\x49\x49\x2d\xe2\x02\xb1\xb5\xb8\xc0\x54\x26\x58\x96\x2b\xc8\xd5\xd1\xc5\xd7\x55\x2f\x37\x85\xcb\xc7\xd3\xd9\xd5\x2f\x18\xcc\x74\xf6\x70\xf4\x73\x77\xf5\xf1\x77\x07\x71\x92\x4a\x33\x73\x52\xb8\x8a\x52\x73\x52\x13\x8b\x53\xb9\x7c\x13\xb3\x53\xd3\x32\x73\x52\xb9\x92\x33\x8b\x92\x73\x52\xf5\x2a\x73\x73\x00\x01\x00\x00\xff\xff\xba\xc0\x2a\xa8\x6a\x00\x00\x00")

func pluginsTemplatesPythonDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerignore,
		"plugins/templates/python/.dockerignore",
	)
}

func pluginsTemplatesPythonDockerignore() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.dockerignore", size: 106, mode: os.FileMode(420), modTime: time.Unix(1475951558, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonGitignoreTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x52\x3d\x8f\xdb\x30\x0c\xdd\xf5\x2b\x08\x74\x0b\x2e\xf6\x52\x74\xe8\x78\xcd\x52\xe0\x86\x03\xd2\x9b\x8a\xc2\x50\x64\xda\xe6\x55\x16\x05\x91\xc9\xc5\xfd\xf5\xa5\x94\xb4\x77\x0b\x4d\x3e\xfa\x51\x8f\x1f\xdd\xe1\x38\x1c\x95\x0b\x3a\xf7\x09\x1e\x37\xc5\x7d\xe0\x35\x53\xc4\x11\x7a\xe0\xac\xb4\xd2\x9f\xe6\x1f\x9e\x9e\x60\x32\x5c\xdc\x30\xe4\x2d\xf8\xb0\xe0\x30\xf4\x6e\xd7\xe5\xed\x67\xe0\xf1\x57\xe5\x7f\x03\xbc\x2a\x26\x21\x4e\x62\x19\xe1\x0a\x1e\x48\xb4\xd0\xe9\xac\x86\x5a\x9d\xec\xc3\x6f\x3f\x53\x9a\x5d\xf7\xbc\xe9\xc2\xc9\x61\xba\xf4\xee\x74\xa6\x38\xf6\x6e\xc4\x0b\x46\xce\x7b\x9c\x67\xb1\xc8\xa8\x66\xf9\x2d\x45\xf6\xa3\x01\x37\xb8\xbb\x7d\x22\x9d\x9a\xf9\xf2\xb9\x77\xd9\x17\x35\x48\x6e\x8c\x8b\x2f\x55\x99\xfd\xb6\xa7\x34\xb1\x31\x28\x89\xfa\x68\x5d\x75\x61\x9a\x6f\xa9\xaa\xed\x79\xfb\x7e\x4f\x14\x8b\xe0\x45\xce\xe6\x6f\xa0\x0b\x0a\xde\xba\x05\x5f\x10\xde\x0a\xa9\xf5\x05\xa7\x0d\x3c\xe4\xa6\x1a\x24\x14\xca\x0a\x53\xe1\xd5\x40\xc5\x35\x47\xaf\x58\xab\x9c\x70\xb2\x79\x7e\xac\x0d\xad\x3b\xa9\x75\x6d\x42\xf8\x00\xc2\xe0\x2d\x64\xa0\xf4\x8a\x41\x61\x34\x6a\xcf\x96\x2e\x50\x05\x8b\xd9\x9a\xd4\xce\xa4\xae\x3e\xd1\x84\xa2\x75\xa0\x19\x43\x95\xfd\x5e\x38\xf2\x2c\x2e\x53\xde\x9b\xd3\xe9\x55\x9b\x3f\x62\x44\x5b\xa4\x2e\x24\xfb\x91\x8a\x3d\xc0\x65\x6b\x59\xe3\xbe\x24\x52\x93\x2b\x6a\xbb\x08\x7c\xc1\xe2\x67\x84\x82\x99\x6d\x80\x6e\xd1\x35\x1a\x68\x03\x53\xbe\x9a\xfd\xf7\xc3\xbb\xd7\xed\xcc\xaf\xcb\x77\x89\x05\x6b\x1d\xe9\xae\x6b\x74\xff\xf3\x35\xd8\x3d\xb4\xb0\x3e\xf7\xa3\xf8\x24\x36\x99\xfb\x49\xac\x5c\x2f\x86\x9b\x92\xc3\xab\x4f\x33\x83\xe8\x79\x9a\xbe\x1a\x6c\x2d\x54\xf8\x98\x17\x4a\x57\x18\x39\x9c\x57\x4c\xda\xa8\x76\x03\x41\xfa\xe1\x7e\x24\x6d\x71\x8f\xd5\xb7\x37\xd4\x97\x19\xb5\xff\x1b\x00\x00\xff\xff\x27\xc3\x36\xdb\xc8\x02\x00\x00")

func pluginsTemplatesPythonGitignoreTemplateBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonGitignoreTemplate,
		"plugins/templates/python/.gitignore.template",
	)
}

func pluginsTemplatesPythonGitignoreTemplate() (*asset, error) {
	bytes, err := pluginsTemplatesPythonGitignoreTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/.gitignore.template", size: 712, mode: os.FileMode(493), modTime: time.Unix(1475952964, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonChangelogMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\x56\x30\xd0\x33\xd4\x33\x50\xd0\x55\x70\xcb\x2c\x2a\x2e\x51\x08\x4a\xcd\x49\x4d\x2c\x4e\xe5\xd2\x52\x70\x2d\x4b\x2d\xaa\x54\x48\x4b\x4d\x2c\x29\x2d\x4a\x55\x48\x4c\x49\x49\x4d\x81\x8b\x26\x95\xa6\x2b\xa4\x65\x56\x00\x45\x00\x01\x00\x00\xff\xff\x82\x96\x39\xb5\x41\x00\x00\x00")

func pluginsTemplatesPythonChangelogMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonChangelogMd,
		"plugins/templates/python/CHANGELOG.md",
	)
}

func pluginsTemplatesPythonChangelogMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonChangelogMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/CHANGELOG.md", size: 65, mode: os.FileMode(493), modTime: time.Unix(1473195758, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x4f\x6b\xea\x40\x14\xc5\xf7\xf3\x29\x2e\x2e\x5c\x3c\x98\x64\xff\xc0\x85\x2f\xe6\xf1\x82\xcf\x24\x44\xdb\x22\xad\x84\x31\xb9\xc6\xa1\x93\xc9\xed\x64\x82\x8a\xf8\xdd\x3b\xb1\x8d\x14\x2a\x5d\x04\x32\xe7\xfe\x0e\xf7\xcf\xf9\x9b\x25\x0b\xa8\x85\x92\x05\xfa\x42\x91\xd4\xf8\xdb\x4a\x2d\x19\x5b\x4c\xa3\x78\xe5\xbe\x30\x83\xf3\x19\x0a\x83\xc2\x36\x06\x2e\x17\xc6\x82\x24\x5d\x83\x07\xbe\xad\xc9\x6f\x4d\xe1\xbb\x32\xa9\xae\x92\x3a\xd7\xa2\xc6\x1e\xc9\x1e\x62\x10\xf4\xca\xa5\x6e\xad\x50\x0a\xe8\x64\xf7\x8d\xfe\x26\x73\x0b\xde\xb6\x93\xaa\xe4\x25\x52\x0b\x95\xb4\x50\xa3\x29\x3a\x23\x45\xef\xe1\x24\x09\x5e\x18\xc0\x78\x0c\x2d\x5a\xe0\xc7\xe1\x55\x94\x3f\x74\x1f\x20\x3c\x52\x63\x2c\xa4\x51\x9a\xc7\x49\x1e\x4c\x83\x7f\x61\x3e\x8b\xb2\x49\xb3\xdb\xdd\x41\x66\xd1\x72\xfa\xe7\x7f\x98\xf7\xff\x8f\x61\xb6\x8c\x92\x38\x77\x8e\x60\x3e\x69\xf4\x80\xf7\xf3\xdc\x66\xe7\x1d\x55\x46\x94\x78\x55\x0f\x7b\x44\x75\x17\x33\x60\xf0\xad\x93\x06\x6b\xd4\xb6\xf5\xec\xd1\xde\xb0\xeb\x51\xfa\xcd\x3a\xf2\xe8\x74\xb3\x7c\x96\x4d\xed\xcc\xbb\x8f\x3d\x7f\x0d\xa2\xbb\x1e\x94\xd8\x77\xa7\xce\x54\xf8\xf5\x7c\x8c\x3d\x25\xd9\xdc\x2d\x08\xbe\x8b\xf3\x20\x0c\x32\x16\xc6\xab\x6c\x9d\x26\x2e\x48\x78\x1e\xf9\x5b\xa9\xfd\xb6\x10\x7a\xb4\x71\x19\x2e\x66\x4e\xe2\x7c\x8f\x8a\x46\x9b\xf7\x00\x00\x00\xff\xff\x22\x74\x40\xdc\x05\x02\x00\x00")

func pluginsTemplatesPythonDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonDockerfile,
		"plugins/templates/python/Dockerfile",
	)
}

func pluginsTemplatesPythonDockerfile() (*asset, error) {
	bytes, err := pluginsTemplatesPythonDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/Dockerfile", size: 517, mode: os.FileMode(420), modTime: time.Unix(1475953290, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonLicenseMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x51\xcf\x6e\xf3\x28\x10\xbf\xfb\x29\x46\x39\xb5\x92\xd5\xbd\xef\x8d\xd8\xa4\x41\xeb\x18\x0b\x93\x66\x73\x5a\x11\x9b\xc4\xac\x1c\x13\x01\xde\x28\xaa\xfa\xee\x3b\x90\xf4\x6b\xfb\x49\x96\xac\x81\xf9\xfd\xa5\xb0\x97\x9b\x33\xa7\x21\xc0\x53\xf7\x0c\xef\xef\xd0\xcd\xce\xe9\x29\xfc\x73\xd3\xca\xc1\xc7\x47\x3a\x72\x5a\x05\x1b\xa7\x2c\x6b\xb4\x3b\x1b\xef\x8d\x9d\xc0\x78\x18\xb4\xd3\x87\x1b\x9c\x9c\x9a\x82\xee\x73\x38\x3a\xad\xc1\x1e\xa1\x1b\x94\x3b\xe9\x1c\x82\x05\x35\xdd\xe0\xa2\x9d\x47\x80\x3d\x04\x65\x26\x33\x9d\x32\x05\x1d\xea\xc6\xcd\x30\x20\x8d\xb7\xc7\x70\x55\x4e\xe3\x72\x0f\xca\x7b\xdb\x19\x85\x7c\xd0\xdb\x6e\x3e\xa3\x19\x15\xa2\xde\xd1\x8c\xda\xc3\x53\x18\x74\xb6\x68\x1f\x88\xc5\x73\x12\xe9\xb5\x1a\xc1\x4c\xc8\xa6\xe1\xf3\x0a\xae\x26\x0c\x76\x0e\xe0\xb4\x0f\xce\x74\x91\x23\xc7\xa5\x6e\x9c\xfb\xe8\xe1\xf3\x7a\x34\x67\xf3\x50\x88\xf0\x54\x86\x8f\xa4\xb3\xc7\x04\xd1\x67\x0e\x67\xdb\x9b\x63\xfc\xeb\x14\xeb\x32\x1f\x46\xe3\x87\x3c\xeb\x4d\xa4\x3e\xcc\x01\x0f\x7d\x3c\xec\xf4\x14\x51\x98\xe3\x0f\x2c\xcc\xeb\x71\x8c\x0c\x06\x7d\xa7\xac\x5f\xee\xd2\x0e\xaa\x64\x97\x58\x68\x78\x54\x94\x74\xaf\x83\x3d\xff\x4c\x82\x15\x1d\x67\x37\xa1\xa4\xee\x53\x5c\x8b\x95\x25\xc5\x7f\x75\x17\x22\x4b\x5c\x3f\xda\x71\xb4\x57\x8c\x86\x92\x53\x6f\x62\x22\xff\x67\x96\x49\xbc\x52\x07\xfb\x9f\x4e\x59\xee\x6f\x3d\xd9\x80\x56\xef\x16\xe2\x03\x5c\xbe\x5e\xf5\x71\xe5\x07\x85\xde\x0f\x3a\xbb\x17\x86\xba\x58\xaf\xfa\x16\xc7\x45\x79\x1f\xf0\xe1\x0d\x76\x7f\xb1\x2e\xe9\xfd\x1e\xf3\x05\xf5\xd7\x14\x5a\xbe\x92\x3b\x22\x28\xb0\x16\x1a\xc1\xdf\x58\x49\x4b\x58\x90\x16\xe7\x45\x0e\x3b\x26\xd7\x7c\x2b\x01\x37\x04\xa9\xe5\x1e\xf8\x0a\x48\xbd\x87\xbf\x58\x5d\xe6\x19\xfd\xbb\x11\xb4\x6d\x81\x0b\x60\x9b\xa6\x62\xb4\xcc\x81\xd5\x45\xb5\x2d\x59\xfd\x0a\x4b\xc4\xd5\x5c\x42\xc5\x36\x4c\x22\xa9\xe4\x10\x05\x1f\x54\x8c\x22\x6e\x95\x6d\xa8\x28\xd6\x38\x92\x25\xab\x98\xdc\xe7\xb0\x62\xb2\x8e\x9c\x2b\x24\x25\xd0\x10\x21\x59\xb1\xad\x88\x80\x66\x2b\x1a\xde\x52\x94\x2f\xb3\x9a\xd7\xac\x5e\x09\x54\xa1\x1b\x5a\xcb\x17\x54\x45\x29\xa0\x6f\x38\x40\xbb\x26\x55\x95\xa4\xc8\x16\xdd\x8b\xe4\xaf\xe0\xcd\x5e\xb0\xd7\xb5\x84\x35\xaf\x4a\x8a\x87\x4b\x9a\x55\x8c\x2c\x2b\x7a\x97\xc2\x50\x45\x45\xd8\x26\x87\x92\x6c\xc8\x2b\x4d\x28\x8e\x2c\x02\xe2\xda\xc3\xdd\x6e\x4d\xd3\x11\xea\x11\xfc\x0a\xc9\x78\x9d\x61\x27\x05\xaf\xa5\xc0\x31\xc7\x94\x42\xfe\x82\xee\x58\x4b\x73\x20\x82\xb5\xb1\x90\x95\xe0\x48\x1f\xeb\x44\x04\x4f\x24\x88\xab\xe9\x9d\x25\x56\x0d\x3f\x5e\x04\x57\xe2\xbc\x6d\xe9\x97\x97\x92\x92\x0a\xb9\xda\x08\xfe\xbe\xfc\x92\xfd\x1f\x00\x00\xff\xff\xe4\x3f\x0e\xa2\x2f\x04\x00\x00")

func pluginsTemplatesPythonLicenseMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonLicenseMd,
		"plugins/templates/python/LICENSE.md",
	)
}

func pluginsTemplatesPythonLicenseMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonLicenseMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/LICENSE.md", size: 1071, mode: os.FileMode(493), modTime: time.Unix(1475964853, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeShortTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x56\x28\xc8\x29\x4d\xcf\xcc\x8b\x4f\x49\x2d\x4e\x2e\xca\x2c\x28\xc9\xcc\xcf\x53\xa8\xad\x05\x04\x00\x00\xff\xff\xf7\x36\xbe\xfa\x18\x00\x00\x00")

func pluginsTemplatesPythonReadmeShortTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeShortTxt,
		"plugins/templates/python/README-short.txt",
	)
}

func pluginsTemplatesPythonReadmeShortTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeShortTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README-short.txt", size: 24, mode: os.FileMode(420), modTime: time.Unix(1475958868, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcc\xbd\xae\x83\x20\x1c\x05\xf0\x9d\xa7\xe0\xe6\x2e\xf7\x0e\x80\x68\x44\xe9\xd6\x07\x68\xd2\xbd\x69\x1a\x3e\xfe\x28\x89\x82\x11\x18\x1a\xe3\xbb\xd7\x34\x1d\xba\x9e\xdf\x39\xe7\x17\x6f\x1b\x5e\xa6\x32\xf8\xf0\x08\x6a\x06\xbc\xef\x08\x5d\xd4\xe4\x0d\x7c\x89\x85\x64\x56\xbf\x64\x1f\xc3\x51\xc0\xd7\x77\x4a\x11\xfa\xb9\x9d\xf1\x21\x00\x21\x8d\x31\xe3\xe8\xf0\x33\x96\xf5\xb3\xba\xff\x8d\x39\x2f\xe9\xc4\x98\xa3\x66\x8a\xc5\xd2\xc1\xe7\xb1\x68\x6a\xe2\xcc\x54\x4a\x90\x13\x13\x92\x0b\xc9\xea\x5a\x56\x75\x5b\x31\xd3\xb4\xb6\x17\x9d\x22\xaa\xe2\x1d\xe1\x1c\x1a\xd2\x0b\x0d\xc4\xd8\xce\xb4\xda\x35\xce\x49\x7d\x9c\xb8\x7f\xf4\x0a\x00\x00\xff\xff\x6d\x30\xd9\xc7\xb7\x00\x00\x00")

func pluginsTemplatesPythonReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonReadmeMd,
		"plugins/templates/python/README.md",
	)
}

func pluginsTemplatesPythonReadmeMd() (*asset, error) {
	bytes, err := pluginsTemplatesPythonReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/README.md", size: 183, mode: os.FileMode(493), modTime: time.Unix(1475954829, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonCircleYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xc1\x6e\x02\x31\x0c\x44\xef\xfb\x15\x16\xed\xa1\x3d\x2c\x7b\x07\xa9\x3f\x82\x10\x32\x89\x01\x8b\xc4\x59\x39\xd9\x95\x2a\x44\xbf\xbd\x61\x49\x28\x48\x94\x5b\x34\xf6\x9b\x99\xd8\xa3\x39\xb0\xd0\xa2\x01\x88\xa4\x23\x1b\x8a\x97\x37\x40\x0b\x36\x98\x23\x69\xd3\x58\xea\x49\x2c\x89\xe1\xeb\xcc\x64\x84\x36\x96\x95\x4c\x0a\xca\x7f\xc0\xec\xa7\xbb\x32\xb3\x2c\x84\x91\x54\xd9\xd2\xa3\x1b\xb0\xec\x42\x51\x78\x07\xab\x15\xb4\x04\x15\xeb\xd8\xe3\x9e\xe6\x09\x15\xd6\xeb\x25\xa4\x03\x49\xc5\x5c\x40\x0b\x6d\xcb\xd2\x0f\xe9\xc9\xfe\x12\x76\xfc\x98\xb3\x1d\xd8\x65\x22\x81\x61\x35\x8e\x0c\x77\xa7\x13\xf4\x6e\xd8\xb3\x6c\x04\x3d\xc1\xf9\x0c\xf3\x82\xf8\x63\xfe\x0c\xb4\xfd\xcd\x78\x59\x5d\x22\x8e\xf4\xca\xe1\xeb\x49\x97\xa6\x49\x14\xd3\xe2\xff\x13\xe8\x20\xaf\x3c\x2f\xf4\x74\x74\x17\xbe\x3d\xc9\xe4\xf4\x06\x1e\x63\x22\xbd\xbe\x01\xb6\x8a\x62\x0e\x8b\xa2\x16\xd1\x04\xef\x51\x6c\xac\x4b\x97\x48\x1a\xd1\xc1\xfb\xc7\x5d\xf2\xde\xe5\x46\xea\x70\x1b\xbb\x3e\x37\x9e\xee\x94\x47\xb5\x50\x6b\xbc\xfd\xcc\x06\x4a\x8e\x30\x96\xea\x35\xae\x88\x93\x76\x9f\x56\xce\x88\x47\xba\xad\xfc\x06\x00\x00\xff\xff\x5b\x80\xf9\x05\x58\x02\x00\x00")

func pluginsTemplatesPythonCircleYmlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonCircleYml,
		"plugins/templates/python/circle.yml",
	)
}

func pluginsTemplatesPythonCircleYml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonCircleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/circle.yml", size: 600, mode: os.FileMode(420), modTime: time.Unix(1475953367, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonPluginToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x91\xbd\x6a\xf3\x30\x14\x86\xf7\x5c\xc5\x41\x59\x3f\x3e\x68\x21\xdd\x3a\x94\x50\x4a\x20\x2d\x5d\x3a\x99\x60\x94\xe8\xc4\x16\xd5\x1f\x92\xec\x62\x42\xee\xbd\xfa\xb1\x53\xdb\xe9\x52\x48\x96\x90\xe7\x7d\x74\xf4\xe6\x68\x79\x9b\xcf\x62\x09\xef\xdb\x8f\x97\xcd\x1b\xac\xb5\x3a\xf2\xaa\xb1\xd4\x73\xad\xe0\xef\x73\x6e\xd4\x67\x51\x14\x46\x34\x15\x57\xbb\xdd\x02\x40\x51\x89\xf0\x08\xe4\x74\x82\x4c\xcb\x44\xce\x67\x12\xc2\x16\xad\x8b\x5d\x27\xf9\x00\xb3\xc2\xd0\x1d\x2c\x37\xfe\x4a\x1b\x07\x59\x3d\x50\x8f\x95\xb6\xdd\xd4\xbb\xd0\x2c\x7d\x62\xf7\xa5\x2d\x73\x41\x2a\x46\xd6\x05\x07\x2b\xd6\xe6\x92\x56\xb3\xde\x19\xe5\x29\x16\x8d\x76\xdc\x5f\x5d\x36\xe2\x59\x14\xfc\x80\xca\xcd\x26\x0d\x30\x2b\xa8\x02\x43\xd7\x2b\x92\xc6\xb4\xec\x61\xe9\x1a\x63\xb4\xf5\xc8\x7a\xb9\x5f\xed\x7f\xc3\x8d\x88\x42\xec\x0a\xe0\xf5\xf4\x82\x90\x86\xb3\x81\xe6\x43\x00\x47\xab\xe5\x6f\x4a\xe2\x59\xda\x37\x5c\xb0\xe0\x1c\xa9\x70\x18\x7e\x53\xc3\xc3\x52\xe2\xa1\xb4\x5b\xc9\x66\x6b\x95\x43\x25\xc9\xe7\x2f\x9c\x48\x0e\x6b\xea\x6a\xdf\x99\xf4\xf7\x0a\x20\x92\xad\xc8\x3f\x20\xae\xa6\x77\xfd\xf7\xfd\xea\x81\xc0\x2e\xad\xa1\x4d\x6f\xf2\xfa\xb4\xdd\xac\x9f\xcb\x9f\x69\x21\x28\x5b\x6a\x87\x97\x41\x45\xf7\x02\x63\x1b\x6f\x1b\xfc\x0e\x00\x00\xff\xff\xab\x02\x2d\x63\x45\x03\x00\x00")

func pluginsTemplatesPythonPluginTomlBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonPluginToml,
		"plugins/templates/python/plugin.toml",
	)
}

func pluginsTemplatesPythonPluginToml() (*asset, error) {
	bytes, err := pluginsTemplatesPythonPluginTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/plugin.toml", size: 837, mode: os.FileMode(420), modTime: time.Unix(1475953976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonRequirementsTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\x28\xc8\x2c\x50\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\xd0\x2d\x52\x28\x4a\x2d\x2c\xcd\x2c\x4a\xcd\x4d\xcd\x2b\x29\xd6\x2b\xa9\x28\x01\x04\x00\x00\xff\xff\x5d\x24\x95\x7b\x21\x00\x00\x00")

func pluginsTemplatesPythonRequirementsTxtBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonRequirementsTxt,
		"plugins/templates/python/requirements.txt",
	)
}

func pluginsTemplatesPythonRequirementsTxt() (*asset, error) {
	bytes, err := pluginsTemplatesPythonRequirementsTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/requirements.txt", size: 33, mode: os.FileMode(420), modTime: time.Unix(1475951657, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pluginsTemplatesPythonScanPy = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\xcd\x6a\xdc\x30\x10\xbe\xeb\x29\xa6\x0a\xc1\xde\x66\xed\x6d\x0b\x29\xc5\xa0\x43\x61\x53\xe8\x21\xd0\x43\xef\x42\x2b\xcf\xda\x02\x5b\x52\x47\xf2\xa6\x26\xa4\xcf\x5e\x49\x6b\x9a\x42\xa9\x2e\x62\x46\xdf\xcf\xfc\xe8\xe6\xcd\x61\x09\x74\x38\x19\x7b\x40\x7b\x01\xbf\xc6\xd1\x59\x76\x03\xcd\xdb\x06\xb4\xeb\x8d\x1d\x3a\x58\xe2\xb9\xf9\x94\x33\x8c\x73\xce\x82\x56\xb6\xf5\x2b\xfb\xb5\x1d\xf6\xfc\x0c\x3d\x06\x4d\xc6\x47\xe3\x2c\xbc\xbc\xb4\xac\xd3\xce\xaf\x64\x86\x31\x76\x50\xeb\x1d\x7c\x78\xf7\xfe\x23\x9c\x56\x48\x50\x4d\xa8\xa2\xa3\x2b\x6c\x32\x1a\x6d\xc0\x0e\x1e\xbf\x7e\x2f\xe2\xcc\xcc\xde\x51\x04\x45\x83\x57\x14\x90\x31\xd6\xe3\x19\x66\x65\x6c\xbd\xeb\x18\xa4\x53\xf2\x04\xe2\x0f\xa6\xfd\x4c\xc3\x32\xa3\x8d\xdf\xca\x4b\xed\xc9\x0d\xa2\x4a\x56\x7e\x5a\x06\x63\xa5\x55\x33\x26\xbb\x6a\xf7\x17\xbd\x55\x7d\x2f\xd5\xc6\xab\x79\x73\xe1\x7b\xe0\x4d\x73\x41\x3a\xb9\x80\x29\x18\x71\xf2\x82\x1f\x4d\xf0\x93\x5a\x61\xcb\x83\x5b\xa2\x5f\x22\xcc\x18\x82\x1a\x32\x4c\xe9\xdc\xb3\xe0\x21\xb5\x84\x32\xd2\x92\x93\x84\x3f\x16\x43\xd8\x8b\x2f\x6a\x0a\xf8\x7f\xdb\x6a\x54\x61\xac\xf6\x49\x2e\xaa\x8b\x22\x51\x3d\x1e\xef\x53\x18\x57\x8f\x22\x44\xda\x83\x4d\xd0\x20\xaa\xbb\x6a\xab\xa7\x52\x30\xf7\xf7\x90\x69\x10\x1d\x04\x54\xa4\x47\x38\x3b\x6a\xb7\xe6\x32\x3e\x4d\x66\x33\x2b\x57\xb6\x0b\xf5\x8e\x95\x77\x73\x2e\x90\x36\x2b\x5c\xa7\x59\x6a\x23\x63\x23\xf0\x63\x52\x74\xa9\x96\x31\x2d\x1d\xd4\x13\xe6\x00\x9e\x4c\x1c\x8b\x61\x07\x1c\xee\x5e\xd9\x85\x4c\x18\x17\xb2\x69\x47\x49\x57\x96\x39\x4b\x09\x42\x00\x97\x32\x6f\x4c\x4a\x7e\x35\x89\xb4\xbe\xba\x5d\x77\x59\x42\xfc\xa9\xd1\x47\x78\x28\x57\xfe\x3b\x2a\x00\xfe\x53\xd7\x03\x91\xa3\x0e\x6e\x03\x87\x5b\x40\xf6\x3b\x00\x00\xff\xff\xe2\x82\x12\x89\xb2\x02\x00\x00")

func pluginsTemplatesPythonScanPyBytes() ([]byte, error) {
	return bindataRead(
		_pluginsTemplatesPythonScanPy,
		"plugins/templates/python/scan.py",
	)
}

func pluginsTemplatesPythonScanPy() (*asset, error) {
	bytes, err := pluginsTemplatesPythonScanPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "plugins/templates/python/scan.py", size: 690, mode: os.FileMode(420), modTime: time.Unix(1475952110, 0)}
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
	"plugins/.DS_Store": pluginsDs_store,
	"plugins/plugins.toml": pluginsPluginsToml,
	"plugins/templates/.DS_Store": pluginsTemplatesDs_store,
	"plugins/templates/python/.DS_Store": pluginsTemplatesPythonDs_store,
	"plugins/templates/python/.dockerignore": pluginsTemplatesPythonDockerignore,
	"plugins/templates/python/.gitignore.template": pluginsTemplatesPythonGitignoreTemplate,
	"plugins/templates/python/CHANGELOG.md": pluginsTemplatesPythonChangelogMd,
	"plugins/templates/python/Dockerfile": pluginsTemplatesPythonDockerfile,
	"plugins/templates/python/LICENSE.md": pluginsTemplatesPythonLicenseMd,
	"plugins/templates/python/README-short.txt": pluginsTemplatesPythonReadmeShortTxt,
	"plugins/templates/python/README.md": pluginsTemplatesPythonReadmeMd,
	"plugins/templates/python/circle.yml": pluginsTemplatesPythonCircleYml,
	"plugins/templates/python/plugin.toml": pluginsTemplatesPythonPluginToml,
	"plugins/templates/python/requirements.txt": pluginsTemplatesPythonRequirementsTxt,
	"plugins/templates/python/scan.py": pluginsTemplatesPythonScanPy,
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
	"plugins": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{pluginsDs_store, map[string]*bintree{}},
		"plugins.toml": &bintree{pluginsPluginsToml, map[string]*bintree{}},
		"templates": &bintree{nil, map[string]*bintree{
			".DS_Store": &bintree{pluginsTemplatesDs_store, map[string]*bintree{}},
			"python": &bintree{nil, map[string]*bintree{
				".DS_Store": &bintree{pluginsTemplatesPythonDs_store, map[string]*bintree{}},
				".dockerignore": &bintree{pluginsTemplatesPythonDockerignore, map[string]*bintree{}},
				".gitignore.template": &bintree{pluginsTemplatesPythonGitignoreTemplate, map[string]*bintree{}},
				"CHANGELOG.md": &bintree{pluginsTemplatesPythonChangelogMd, map[string]*bintree{}},
				"Dockerfile": &bintree{pluginsTemplatesPythonDockerfile, map[string]*bintree{}},
				"LICENSE.md": &bintree{pluginsTemplatesPythonLicenseMd, map[string]*bintree{}},
				"README-short.txt": &bintree{pluginsTemplatesPythonReadmeShortTxt, map[string]*bintree{}},
				"README.md": &bintree{pluginsTemplatesPythonReadmeMd, map[string]*bintree{}},
				"circle.yml": &bintree{pluginsTemplatesPythonCircleYml, map[string]*bintree{}},
				"plugin.toml": &bintree{pluginsTemplatesPythonPluginToml, map[string]*bintree{}},
				"requirements.txt": &bintree{pluginsTemplatesPythonRequirementsTxt, map[string]*bintree{}},
				"scan.py": &bintree{pluginsTemplatesPythonScanPy, map[string]*bintree{}},
			}},
		}},
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

