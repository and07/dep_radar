// Code generated by go-bindata.
// sources:
// html/templates/apps.html
// html/templates/libs.html
// html/templates/main.html
// DO NOT EDIT!

package templates

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

var _htmlTemplatesAppsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x24\x12\xd7\x2d\x6c\xa1\xe5\xf8\xb3\x6d\xe9\x9e\xee\x85\x13\x9b\xc4\xe0\xd8\xa9\x3d\x50\x38\x94\xef\x7e\x4a\xa0\xb4\xbb\x6a\xa5\xbb\x7b\x03\x9e\x67\x3c\xf6\x6f\x46\x8f\x1c\xef\x37\xaa\x22\x58\x67\x0c\x25\x90\x8a\xc0\xf2\x8a\x3f\x24\x88\x8c\x7d\x9b\x49\xbb\x10\x18\xa1\x81\x75\xe0\xa5\x0c\x08\x8a\x12\xa2\x0d\x03\xdf\x5e\xc0\xd4\xb9\xb0\xf7\xba\x24\x29\xf3\xed\x25\x67\xcf\x99\xd2\x60\xa3\x48\x49\x60\x12\x7c\xfb\x99\x53\x48\x7c\xca\x96\x3c\x62\x4e\x19\x7c\x46\x5c\x72\xe0\x44\x38\x26\x22\x82\xf9\xd5\xcf\xc8\x24\x9a\xcb\xb9\x03\xca\x99\x72\xf0\xa5\x2a\xcf\x15\x5c\xce\x91\x66\xc2\xb7\x0d\xac\x05\x33\x09\x63\x60\xa3\x44\xb3\xa9\x6f\x27\x00\x99\xf9\x82\x71\x4a\x56\x11\x95\x6e\xa8\x14\x18\xd0\x24\x2b\x82\x48\xa5\x78\x2f\xe0\xba\x5b\x71\x2b\x38\x32\xe6\x55\x73\x53\x2e\xdd\xc8\x18\x1b\x71\x09\x2c\xd6\x1c\xd6\xbe\x6d\x12\x52\xbb\xa8\x3b\x1d\x79\x56\xbb\xa8\xaf\x9e\x86\x55\xa2\x1e\x26\xad\x93\xca\xd9\xc5\x68\x32\x58\x0d\xe2\xc6\x74\x5d\xbf\x7e\x58\x7e\xef\x27\x95\xab\xd3\x46\x6d\x92\xb6\xa3\xae\x18\xb7\x9e\x79\x27\x6e\xb7\x1e\x30\x6d\xf1\x71\xa3\x3b\x49\x6d\x14\x69\x65\x8c\xd2\x3c\xe6\xd2\xb7\x89\x54\x72\x9d\xaa\x85\x29\x06\x89\xb7\x93\xf4\x42\x45\xd7\x81\xe5\x01\x09\x05\x43\x91\x20\xc6\xf8\xf6\x36\x28\x7f\x9d\x50\x69\xca\x34\xa3\xbb\xd0\xa4\x45\x31\x14\xc5\xfb\xdd\x45\xe0\x50\xa2\xe7\xe5\xa4\x40\x07\x1e\x24\x81\xe0\xa1\x26\x9a\x33\xe3\x61\x48\x82\xcd\x46\x13\x19\x33\xe4\xb6\xb2\xec\x9e\xb3\xe7\x3c\xf7\x4a\xd5\xed\x93\x94\xe5\xf9\x6e\x0f\x93\xb4\x5c\xeb\x82\x0f\x76\x80\xb0\x25\xdc\x6c\x1c\x74\x48\xb6\xd5\xe8\x8b\xbf\x3f\x09\xe5\x79\x99\xdb\x9e\x7f\x28\x78\x58\x66\x6f\x78\xf8\x92\x2d\x88\x0a\x2e\xfa\x53\x7b\x3b\x60\x2f\x0c\x36\x9b\xb2\xac\x44\x41\xc5\xfd\x61\x80\x3c\x93\x12\x21\xf6\xa9\x6f\x5c\x82\x29\x73\x5b\xdd\xc3\x50\xb8\xf0\xcd\xbd\x24\xcb\x8a\x7b\xf7\x84\x79\x6e\x1d\x1c\xbc\x30\xdf\x6c\xa1\xb8\xa4\x6c\x55\x0a\x05\x9e\xf9\xe9\x56\xeb\xe0\x2d\xe1\x66\xb3\x2b\x73\x2f\x0b\x21\xcf\x6d\x44\x09\x10\x07\x54\x1c\x0b\xe6\xdb\xa0\x94\x00\x9e\xd9\x08\x38\x14\xf1\xeb\xfe\xef\x85\x90\xe7\x76\xf0\x2a\xdd\x33\x6d\xb8\x92\xe5\x64\x5f\xa8\x99\xa4\xe5\x68\xca\x59\xff\xa2\xec\x1c\x81\xcb\x31\x05\x96\x67\x22\xcd\x33\x40\x46\x47\xaf\x56\x8f\x14\x65\xee\xec\x69\xc1\xf4\xba\xb4\xf8\x76\xe9\xd4\xdc\x53\xb7\xea\x1a\xc1\xd3\xd2\xd6\xb3\x77\x5d\xdd\xeb\xd6\xd4\xe9\xd7\x1e\x5c\xcf\x97\x8f\xd7\xbd\xda\xdd\x55\xff\xef\xf4\xf6\xbc\x77\x39\x1f\x69\xac\xaf\x9a\x78\x98\xc5\x0d\xd2\xfa\xd1\xe9\x3e\xb7\xbf\xde\xde\xf7\x5b\xb8\x93\x75\xda\xed\x66\x2d\x99\x64\x9d\xb3\xde\xbc\xff\xb1\xab\x3d\xbc\x65\xfd\x08\x9a\xca\x99\x71\x23\xa1\x16\x74\x2a\x88\x66\x25\x39\x99\x91\x15\x16\x3c\x34\x38\x53\x59\xc6\xb4\x3b\x33\xb8\xea\x56\x4f\xdd\x26\x5e\xa4\xf4\x45\xfc\xb8\x9b\x56\xd6\x0f\xe3\xa4\xf9\xfb\xc9\x63\x75\xd8\x83\x65\x6d\x24\xcf\x1f\x6a\x69\x3c\x58\x25\x77\xcd\x1e\x1e\x47\x43\xd3\x1a\x9c\x27\x77\x3c\x9c\xd4\x9a\xb3\xf3\x29\x99\xb7\x07\x66\xbe\x9c\x2c\xcc\x72\x4a\x2a\x61\x7d\xf8\xbf\xbb\xf9\xb7\xaf\xcd\xec\xd7\xc7\xe6\xfd\x3e\xba\x3f\x46\x8d\x71\xc6\x66\x49\xfd\xae\x72\x4a\x2f\x66\x7f\x40\x63\x79\x73\xf5\x6d\xca\x70\x77\xd8\xe1\xa3\xd1\x78\x38\x5c\x8d\xa7\xed\x87\x8c\x57\x6f\x9f\x16\xf7\xb4\xb5\x9e\xdd\x11\x7d\x76\x72\xde\x18\xdc\x5f\xa6\x8f\xe2\x3f\xf4\x11\x58\x87\x47\xd3\x85\x8c\x80\x2b\x89\x8e\x8e\xd1\xc6\x42\xe8\xf0\xe8\xd3\x9f\xef\x5a\xfc\xaf\x4f\xc7\xee\x6e\x7d\x74\x6c\xe5\xc7\xd6\x9b\xe3\xf6\x16\xc5\x2f\x9e\xdd\x7e\x36\xac\x7f\x02\x00\x00\xff\xff\x95\x2a\x20\x56\x48\x06\x00\x00")

func htmlTemplatesAppsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesAppsHtml,
		"html/templates/apps.html",
	)
}

func htmlTemplatesAppsHtml() (*asset, error) {
	bytes, err := htmlTemplatesAppsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/apps.html", size: 1608, mode: os.FileMode(420), modTime: time.Unix(1525722917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesLibsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x04\x89\xeb\x16\xb6\xd0\x72\xfc\xd9\xb6\x74\x4f\xf7\xc2\x89\x4d\x62\x70\xec\xd4\x1e\x28\x1c\xca\x77\x3f\x25\xa1\xb4\xb7\x6a\xa5\xbb\x7d\x03\x99\x67\x3c\x93\xdf\x8c\x1e\x39\xde\x6f\x54\x85\xb0\x49\x19\x8a\x21\x11\x2d\xcb\xcb\xff\x90\x20\x32\xf2\x6d\x26\xed\x5c\x60\x84\xb6\xac\x03\x2f\x61\x40\x50\x18\x13\x6d\x18\xf8\xf6\x12\x66\xce\x85\xbd\xd7\x25\x49\x98\x6f\xaf\x38\x7b\x4e\x95\x06\x1b\x85\x4a\x02\x93\xe0\xdb\xcf\x9c\x42\xec\x53\xb6\xe2\x21\x73\x8a\xe0\x33\xe2\x92\x03\x27\xc2\x31\x21\x11\xcc\xaf\x7e\x46\x26\xd6\x5c\x2e\x1c\x50\xce\x8c\x83\x2f\x55\xd1\x57\x70\xb9\x40\x9a\x09\xdf\x36\xb0\x11\xcc\xc4\x8c\x81\x8d\x62\xcd\x66\xbe\x1d\x03\xa4\xe6\x0b\xc6\x09\x59\x87\x54\xba\x81\x52\x60\x40\x93\x34\x0f\x42\x95\xe0\xbd\x80\xeb\x6e\xc5\xad\xe0\xd0\x98\x57\xcd\x4d\xb8\x74\x43\x63\x6c\xc4\x25\xb0\x48\x73\xd8\xf8\xb6\x89\x49\xed\xa2\xee\x74\xe5\x59\xed\xa2\xbe\x7e\x1a\x55\x89\x7a\x98\xb6\x4f\x2a\x67\x17\xe3\xe9\x70\x3d\x8c\x1a\xb3\x4d\xfd\xfa\x61\xf5\x7d\x10\x57\xae\x4e\x1b\xb5\x69\xd2\x09\x7b\x62\xd2\x7e\xe6\xdd\xa8\xd3\x7e\xc0\xb4\xcd\x27\x8d\xde\x34\xb1\x51\xa8\x95\x31\x4a\xf3\x88\x4b\xdf\x26\x52\xc9\x4d\xa2\x96\x26\x5f\x24\x2e\x37\xe9\x05\x8a\x6e\x5a\x96\x07\x24\x10\x0c\x85\x82\x18\xe3\xdb\x65\x50\xfc\x3a\x81\xd2\x94\x69\x46\x77\xa1\x49\xf2\x62\xc8\x8b\xf7\xa7\xf3\xc0\xa1\x44\x2f\x8a\x4d\x81\x6e\x79\x10\xb7\x04\x0f\x34\xd1\x9c\x19\x0f\x43\xdc\xda\x6e\x35\x91\x11\x43\xee\x0d\x0f\xee\x39\x7b\xce\x32\xaf\x50\xdd\x01\x49\x58\x96\x21\xe4\x99\x84\x08\xd1\xda\x6e\x91\xfb\x8d\x4b\x30\x28\xcb\x3c\x5c\x6a\xbb\x06\x4c\xd2\x5c\x03\x9d\xc3\xc3\x8e\x1e\x4a\xfc\xed\xd6\x41\x87\xa2\x6c\x8d\xbe\xf8\xfb\xd7\xa0\x2c\x2b\x72\xe5\xcb\x0f\x49\x9a\x16\xd9\x76\x9a\xbe\x64\x73\xdc\x1c\x9a\xfe\x6b\xf6\xdd\x34\x5e\x90\x03\xe5\x65\x05\x67\xc1\x14\xe4\x3c\xb9\x05\xdf\xf4\x15\x3c\xc8\xfb\xee\x09\xb2\xcc\x3a\x28\xf2\x79\xe9\x4d\x99\xe4\x92\xb2\x75\xd9\xeb\x86\x07\xa6\x38\xfc\xd2\xd5\x3a\x78\x4b\xb0\xdd\xee\xca\xdc\xcb\x5c\xc8\x32\x1b\x51\x02\xc4\x01\x15\x45\x82\xf9\x36\x28\x25\x80\xa7\x36\x02\x0e\x79\xfc\x7a\xfe\x7b\x2e\x64\x99\xdd\x7a\x95\xee\x99\x36\x5c\xc9\x62\x73\x2f\xd4\x4c\xd2\x62\xf4\x62\x97\x3f\x29\x3b\x3b\xe0\x62\x0d\x2d\xcb\x33\xa1\xe6\x29\x20\xa3\xc3\x57\x9f\x87\x8a\x32\x77\xfe\xb4\x64\x7a\x53\xf8\xbb\x7c\x74\x6a\xee\xa9\x5b\x75\x8d\xe0\x49\xe1\xe9\xf9\xbb\x96\xee\xf7\x6a\xea\xf4\x6b\x1f\xae\x17\xab\xc7\xeb\x7e\xed\xee\x6a\xf0\x77\x72\x7b\xde\xbf\x5c\x8c\x35\xd6\x57\x4d\x3c\x4a\xa3\x06\x69\xff\xe8\xf6\x9e\x3b\x5f\x6f\xef\x07\x6d\xdc\x4d\xbb\x9d\x4e\xb3\x16\x4f\xd3\xee\x59\x7f\x31\xf8\xd8\xd2\x1e\x2e\x59\x3f\x82\xa6\x72\x6e\xdc\x50\xa8\x25\x9d\x09\xa2\x59\x41\x4e\xe6\x64\x8d\x05\x0f\x0c\x4e\x55\x9a\x32\xed\xce\x0d\xae\xba\xd5\x53\xb7\x89\x97\x09\x7d\x11\x3f\x9e\xa6\x9d\x0e\x82\x28\x6e\xfe\x7e\xf2\x58\x1d\xf5\x61\x55\x1b\xcb\xf3\x87\x5a\x12\x0d\xd7\xf1\x5d\xb3\x8f\x27\xe1\xc8\xb4\x87\xe7\xf1\x1d\x0f\xa6\xb5\xe6\xfc\x7c\x46\x16\x9d\xa1\x59\xac\xa6\x4b\xb3\x9a\x91\x4a\x50\x1f\xfd\xf2\x34\xff\xf5\xaa\x99\xff\x7c\xd3\xbc\x3f\x47\xef\xc7\xb8\x31\x49\xd9\x3c\xae\xdf\x55\x4e\xe9\xc5\xfc\x0f\x68\xac\x6e\xae\xbe\xcd\x18\xee\x8d\xba\x7c\x3c\x9e\x8c\x46\xeb\xc9\xac\xf3\x90\xf2\xea\xed\xd3\xf2\x9e\xb6\x37\xf3\x3b\xa2\xcf\x4e\xce\x1b\xc3\xfb\xcb\xe4\x51\xfc\x8f\x39\x5a\xd6\xe1\xd1\x6c\x29\x43\xe0\x4a\xa2\xa3\x63\xb4\xb5\x10\x3a\x3c\xfa\xf4\xe7\xbb\x16\xff\xeb\xd3\xb1\xbb\x7b\x3e\x3a\xb6\xb2\x63\xeb\x4d\xbb\xbd\x45\xf1\x8b\x67\xcb\x6f\x86\xf5\x4f\x00\x00\x00\xff\xff\x8e\x65\x9b\xee\x45\x06\x00\x00")

func htmlTemplatesLibsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesLibsHtml,
		"html/templates/libs.html",
	)
}

func htmlTemplatesLibsHtml() (*asset, error) {
	bytes, err := htmlTemplatesLibsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/libs.html", size: 1605, mode: os.FileMode(420), modTime: time.Unix(1525722917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesMainHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\xeb\x8e\xdb\x36\xf6\xff\x3c\x7e\x8a\x13\x06\x68\x6d\x8c\x25\xcd\x34\x93\x4b\x3d\x92\xff\xff\x36\xc9\x74\x93\x26\x9d\x34\x69\x2e\x6d\x51\xa0\x94\x78\x2c\x71\x2c\x91\x0c\x49\xd9\x71\x0d\x03\xfb\x34\xfb\x60\xfb\x24\x0b\xea\x66\xf9\x32\x59\x64\x17\xd8\x2f\x33\xe4\xe1\xe1\xe1\xef\x77\x2e\xe4\x91\xc3\x3b\x4f\xae\x1f\xff\xf2\xeb\xab\xa7\x90\xd9\x22\x9f\x0e\x42\xf7\x0f\x72\x2a\xd2\x88\xa0\x20\xd3\x01\x40\x98\x21\x65\x6e\x00\x10\x16\x68\x29\x24\x19\xd5\x06\x6d\x44\xde\xfe\x72\xe5\x3d\x22\xfd\x25\x41\x0b\x8c\xc8\x82\xe3\x52\x49\x6d\x09\x24\x52\x58\x14\x36\x22\x4b\xce\x6c\x16\x31\x5c\xf0\x04\xbd\x6a\x32\x06\x2e\xb8\xe5\x34\xf7\x4c\x42\x73\x8c\xce\xfd\xb3\x1d\x53\x99\xb5\xca\xc3\x8f\x25\x5f\x44\xe4\x83\xf7\xf6\x3b\xef\xb1\x2c\x14\xb5\x3c\xce\xb1\x67\x97\x63\x84\x2c\xc5\x76\xa7\xe5\x36\xc7\xe9\x13\x54\xa0\x29\xa3\x3a\x0c\x6a\x41\xbd\x68\xec\xaa\x1d\x03\xf8\x4c\x5a\x8b\x6c\x6d\xf1\x93\xf5\x18\x26\x52\x53\xcb\xa5\x98\x80\x90\x02\x2f\x21\x96\x9a\xa1\xf6\x62\x69\xad\x2c\x26\xe7\xea\x13\xd4\xfa\x70\xf7\xec\xec\x61\x3c\x9b\x5d\x6e\x76\xed\x4c\x32\xb9\x40\x0d\x6b\xf8\x0f\xec\xdd\x7f\x10\xdf\xbb\x84\xda\x60\x18\xf4\x50\x86\x39\x17\x73\xd0\x98\x47\xa4\x92\x9a\x0c\xd1\x12\xc8\x34\xce\x22\xe2\xfc\x63\x26\x41\x60\x2c\x4d\xe6\x8a\xda\xcc\x8f\xa5\xb4\xc6\x6a\xaa\x12\x26\xfc\x44\x16\x41\x27\x08\x2e\xfc\x73\xff\x2c\x48\x8c\xd9\xca\xfc\x82\x0b\x3f\x31\x86\x00\x17\x16\x53\xcd\xed\x2a\x22\x26\xa3\xf7\x1e\x5d\x78\xdf\xa6\xef\x7e\xbe\x60\xbf\x5e\x2d\x97\xef\xdf\xdc\x3c\x7b\xf2\x9b\x78\xf1\xf4\xbd\xf8\xf4\xf8\x06\xdf\xbc\xbf\x52\xd9\x73\xbe\xfc\xe1\xd5\x07\x7d\x7e\xc3\xd8\xb3\xec\x1a\x53\x5e\x9e\x5f\x2d\xaf\xef\x7f\x7c\xfd\xc3\xe2\xea\xc3\x35\x7b\xfe\xdb\x05\x81\x44\x4b\x63\xa4\xe6\x29\x17\x11\xa1\x42\x8a\x55\x21\x4b\x43\xbe\x80\x57\x69\xd0\x9f\x49\x61\xe9\x12\x8d\x2c\xb0\x22\xa4\x31\x47\x6a\xd0\x04\x8b\xfb\xfe\x99\x7f\x5e\x33\xa2\x79\x7e\x1b\x8f\x53\x76\xf6\xea\xd1\x3d\xf1\xed\x9c\xfe\xfc\xf2\xf1\xf2\xe6\xd1\xd5\xc5\xeb\xe7\xdf\x3f\x78\x60\xff\x7a\xb6\xbc\xfe\xb1\xd0\x2c\xbe\x78\x70\xaa\xa4\x7e\x12\x5c\x2f\xf4\xf3\xd3\x7b\x0f\xdf\x7f\x7c\xf6\xf2\xe1\x5b\xf9\xbd\x5d\xfe\xed\xfa\xc1\x4f\x79\xfa\x59\x1e\x61\x50\x17\x46\x18\x4b\xb6\x9a\x0e\x06\x21\xe3\x0b\xe0\x2c\x22\x54\xa9\x5a\x81\xa1\xf2\xaa\x3c\x0c\xa6\x83\x30\x60\x7c\xe1\xb4\x4c\xa2\xb9\xb2\x60\x74\xb2\xe5\xea\x02\x76\x63\x18\xe6\x7c\xa1\x7d\x81\x36\x10\xaa\x08\x16\x25\x06\x8c\x1b\xeb\x06\xfe\x8d\x21\xd3\x30\xa8\xf7\x4e\x07\xe1\x1d\xcf\x83\x2f\xb2\xd4\xdb\x0d\x9e\x37\x3d\x0e\xa3\x14\x6a\x9e\x56\x8e\xa6\x9f\xb8\x34\xf5\xe9\xd5\xb0\x4a\x96\x5d\x0c\x1d\x95\xe9\x60\xf0\xae\xac\xe2\xa3\xa4\x40\x61\x87\x5f\x77\xbc\xbf\x1e\xc3\x7a\x00\xc0\xa8\xa5\x13\x98\x95\x22\x71\x25\x01\xc3\x51\x25\x05\xd0\x68\x4b\x2d\x9a\x49\xab\xb6\xde\x8c\x9b\x79\x4e\x8d\xbd\xd6\xe9\x04\x08\x69\x45\x72\x77\x9a\x4b\xca\xb8\x48\x27\x30\xa3\xb9\xc1\x56\x8a\x5a\x4b\xdd\x57\x9b\xf1\xdc\xa2\xfe\x4e\xa9\x43\xe1\x0b\x1e\xf7\x85\xf8\x89\x16\x2a\x47\x33\x81\xdf\x49\xc7\x82\x8c\x81\x64\xa8\x2c\x97\x6e\xf4\x92\x1a\x8b\xba\xe0\x82\x19\x37\x2d\x63\xd4\x5e\x5a\xad\xdc\x50\x4c\x51\x5b\x4d\x13\x2e\x52\xf2\x47\x6d\xd3\x95\x75\xc5\x28\xd1\x48\x5d\xc1\x57\xec\x4f\x4e\xf8\x0c\x86\x36\xe3\xc6\x97\x3a\x85\x3b\x11\x10\x52\xcb\x4f\x2a\xe1\x0c\x6d\x92\x3d\xa1\x96\x0e\x47\x83\x93\x93\xce\x46\x81\x36\x93\xcc\x4c\x1a\x97\x75\x5a\x5b\x97\x02\x74\x86\x1b\xe7\x6c\x57\x5a\x87\x37\xd3\xcd\xfe\x06\x87\x24\x6a\x91\xb4\x5b\xaa\x95\xca\xa3\x10\x01\x79\x5a\x28\xbb\xaa\x6e\x79\xf2\x79\xa3\x7d\x04\x10\x81\xd5\x25\x0e\x8e\x19\x24\x3b\xfa\x75\xc0\x9d\x7e\x83\xa7\x59\xad\x72\xb0\x3b\xd0\x4f\xd1\x0e\x49\x40\x15\x0f\xfe\xaf\x7e\x70\x4e\x5b\xfd\xd1\x56\xc9\x66\x28\x86\x1a\x8d\x92\xc2\x20\x44\xd3\x1e\xa7\xe6\x38\x97\x6f\x10\x41\xab\x53\xcd\x2f\x7b\x4a\x9d\x63\xdc\x82\x4f\x95\x32\x7e\x8e\x22\xb5\x99\xf3\xd2\xd9\x68\xc7\xe0\x3e\xaf\x9f\x24\x68\x54\xd2\xc0\x92\xdb\x0c\x18\x2a\x43\x7a\xda\x9b\x6e\xbc\xe9\x21\x4e\xa8\x4d\xb2\x61\x63\x62\x17\x6f\x22\x85\x91\x39\xfa\xb9\x4c\x6b\x85\xd1\x3e\x99\xee\xe4\xd7\x0d\x9d\x09\x10\x38\xad\x6b\xc1\xdf\xa1\x08\xa7\x40\xc6\x60\x2c\xb5\xa5\x81\x44\xb2\xe3\x9a\xf5\xfa\x51\x9c\x33\x2e\x68\x9e\xaf\x86\xc3\x91\x83\xb9\x17\xe9\xaa\x12\x6b\xe5\xa6\x90\xeb\x42\xeb\xd7\xbf\x8c\x6f\x30\xb1\x66\x0c\xa6\x8c\x8d\xd5\xbb\xb9\x5b\xcb\x0e\x13\xb1\xb9\x2c\x9a\xbd\x97\x7b\xf9\xb6\xbb\xea\xd7\x67\x0e\xdd\xdc\x61\x94\xf1\x8d\xef\x32\xc5\xe7\x22\xc9\x4b\x86\xa6\x39\x65\x34\xba\x3c\x44\xfa\x23\xae\xcc\xff\x06\xed\x75\xb5\xea\xcf\x71\x65\xda\x53\x76\xdc\x5c\x71\x98\xe3\xca\x51\x98\xe3\xea\x10\xfd\x56\x59\x23\x2b\x13\x1c\x3a\x33\x63\xa7\x3b\xda\xcf\x20\x19\xdf\xfc\x3e\xc7\xd5\x1f\x10\xb5\xa0\xaa\x69\x3f\xdf\xb7\x98\xb7\xd2\xcd\x18\xd6\x9b\xbe\x97\xaa\x3f\x16\x0b\x95\x53\x8b\x13\xf8\xb3\x7a\xf5\x92\x9c\x1a\x13\x91\xc2\xfb\x86\xb4\xfd\x55\x38\x93\xba\x68\x17\xdc\xd8\xe3\x22\xe7\x02\xa1\x88\xbd\x6f\x08\x2c\x3c\x29\x26\xa6\x8c\x0b\x6e\x7d\xa5\x71\x51\x35\x73\xdd\x7d\xd6\x59\x71\x8f\xe8\xd6\x3e\x17\xaa\xb4\x5e\xaa\x65\xa9\xa0\xd0\xbd\xc3\x3e\xa3\xe8\x29\x8d\x0a\x05\xdb\xd1\x75\xcd\xa0\xa2\xe2\x98\xba\xeb\xdf\xc8\x34\xe5\x36\x2b\xe3\xea\x29\x0c\x03\xa7\xba\x73\x52\xfd\x94\xf7\x04\x95\x01\xb0\x2b\x85\x11\xa9\x0c\xec\x10\x77\xcd\xaa\x96\xb9\x23\x5d\x48\xe6\xba\x1e\xa9\x53\x02\x2a\xa7\x09\x66\x32\x67\xa8\x2b\x09\x15\xfc\xaf\xaa\x67\x1c\x43\x69\x50\x83\xd4\xd5\x15\x42\x80\x96\x56\xce\x64\x52\x9a\x9e\x57\x6a\x08\x27\x27\x27\x27\x61\x5c\x5a\x2b\x3b\x2e\xb1\x15\x10\x5b\xe1\x29\xcd\x0b\xaa\x57\x04\x26\x8c\x1b\x1a\xe7\xc8\x22\xd2\x54\x29\x99\x86\xbc\x55\x67\x6d\x58\x66\xd4\xc0\x8c\x7a\x46\x71\x21\x50\xbb\xa1\x2a\x73\x83\x0e\x34\x9f\xf5\xb7\x06\x7c\x0a\x6f\x12\x2a\xc2\xa0\x3e\xb8\x0b\x78\xe0\xc8\x76\xb3\x9e\x8b\xc2\x32\x6f\x8f\xcb\xb9\xb1\xcd\x89\xad\xe5\xf6\xc9\x6d\xef\xd6\x29\x9c\xed\xc6\x35\xe7\x47\x76\x7b\xdc\x62\xe1\x4c\xcc\xa4\xee\x6c\x00\x17\xdd\x0b\xbe\x1f\x6f\xda\xb4\x97\x77\xbb\xd8\xd4\x0d\x38\x81\xff\x4f\x72\x9e\xcc\xb7\x59\x28\x75\x1a\x35\x56\xf6\x30\xba\x77\xda\x85\x6e\xba\x5e\x37\x92\xcd\x26\x0c\xe8\x7e\x66\x61\x7d\x71\x2c\x3c\xcc\x0d\xee\xea\x36\x6b\xbb\xd9\x94\xf3\xad\x0f\xcb\x7c\x3b\xee\xbb\xb0\x97\xdc\x34\x47\x6d\xa1\xfa\xeb\x31\x2a\x52\xd4\x1d\x4c\x77\x85\xf7\x88\xaf\xd7\xf5\xad\x0e\x9b\xcd\x51\xa3\x7b\x48\x77\x6b\x6e\x2f\xee\x3d\xc4\x2f\x6a\xd9\x3f\xff\xfe\x8f\xfd\x7c\xbc\xc5\x07\x95\xa5\xee\x11\x85\xaf\xbe\x82\x83\x17\xf5\x20\xea\x3d\xc2\xd6\xa5\xaf\xd7\xbc\x4c\x7c\x81\x5e\x71\x50\xcd\x95\xca\x8e\x3e\xd4\xbb\xea\xef\x2e\x64\xcd\xd4\x14\x7b\x3b\xdd\x5e\xd7\xc6\x77\x7b\xdd\xc4\x63\x54\xcf\x0f\x14\x9d\xaa\x3e\x14\x3a\x31\x3b\x26\x3e\x7a\x2f\xec\x14\x7d\x7d\xc3\x83\xcd\x10\x9c\x2f\x7a\x17\x44\xd7\xb1\xb6\xa1\xfd\x77\xde\x3b\x3f\x02\xf7\x0b\x21\xe4\x3c\x3e\x84\xf0\x82\xc7\x47\x2d\x87\xc1\x71\xd2\xa1\xcd\xda\xa2\xcc\x79\x3c\x06\xee\x6a\xb2\x79\xc9\x2a\xd8\xee\x94\xf1\xb6\xf9\x1e\xb9\x13\x63\x2e\xd8\x64\x8e\xab\x6a\x4f\xf5\x50\xdf\xc2\xc6\xa5\x74\xab\xb2\xd9\x1c\x87\x95\x1d\x89\x5b\x70\x18\x38\xa7\xd9\xfd\xae\xd1\xc7\x5f\x7f\xcd\x1d\xf2\xd2\x2d\x2f\xaa\xd4\x11\x5e\x2e\x1c\xe3\xed\x97\xc6\x1e\x2f\xaa\xd4\xed\xbc\x42\xcb\x76\x53\xbd\x4e\xbf\xf5\xba\xdd\xe5\x6e\x8e\xdb\xdc\xcd\x6e\x73\xb7\x6b\x63\x86\xce\xc2\x7f\xe5\xf1\xda\xdf\x0b\xd4\x86\x4b\x71\xc4\xe5\xc7\x81\xdd\xe2\xf0\x43\xd7\x86\x41\xc5\xf8\x33\xaf\xeb\xc1\x9d\xb9\x27\xa8\xd5\xff\x1c\x6c\x46\x83\x81\xeb\x92\xad\x2b\x25\x88\x40\xe0\x12\xde\x95\x38\x74\x3d\x10\xe6\x13\x20\x77\xdd\x67\xf9\x78\xe0\x9a\x99\x41\xef\x2b\x36\x0c\x6a\x54\x61\x50\xff\x00\xf6\xaf\x00\x00\x00\xff\xff\x34\xbd\x1a\x67\x11\x13\x00\x00")

func htmlTemplatesMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesMainHtml,
		"html/templates/main.html",
	)
}

func htmlTemplatesMainHtml() (*asset, error) {
	bytes, err := htmlTemplatesMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/main.html", size: 4881, mode: os.FileMode(420), modTime: time.Unix(1525864255, 0)}
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
	"html/templates/apps.html": htmlTemplatesAppsHtml,
	"html/templates/libs.html": htmlTemplatesLibsHtml,
	"html/templates/main.html": htmlTemplatesMainHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"apps.html": &bintree{htmlTemplatesAppsHtml, map[string]*bintree{}},
			"libs.html": &bintree{htmlTemplatesLibsHtml, map[string]*bintree{}},
			"main.html": &bintree{htmlTemplatesMainHtml, map[string]*bintree{}},
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

