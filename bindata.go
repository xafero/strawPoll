// Code generated by go-bindata.
// sources:
// view/index.html
// view/poll.html
// view/question.html
// view/results.html
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

var _viewIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x5f\x53\xdb\xb8\x17\x7d\xb6\x3f\x85\x7e\xfa\xcd\xce\xd0\xa1\xb1\x92\x86\xff\x6b\x33\x13\x20\xa1\x40\x17\x02\x2d\xa4\xf0\x26\x5b\x37\xb1\x1c\x59\x32\x92\x9c\xc4\x61\xf8\xee\x3b\xb6\x13\x08\x2c\x74\xdb\xbe\xed\x53\x24\xf9\xde\xab\x73\xae\xce\x91\xf2\xf0\xc0\x60\xc8\x25\x20\x9c\x52\x2e\xf1\xe3\xa3\xeb\xff\xef\xe8\xe2\xf0\xdb\x6d\xbf\x8b\x62\x9b\x8a\x7d\xd7\x5f\xfe\x00\x65\xfb\xae\xe3\x5b\x6e\x05\xec\x77\xd0\x7d\x0e\xc6\x72\x25\x7d\x52\xaf\xb8\x8e\x2f\xb8\x1c\x23\x0d\x22\xc0\xc6\x16\x02\x4c\x0c\x60\x31\x8a\x35\x0c\x03\x1c\x5b\x9b\x99\x3d\x42\x52\x3a\x8b\x98\xf4\x42\xa5\xac\xb1\x9a\x66\xe5\x24\x52\x29\x79\x5a\x20\x1b\x5e\xd3\x6b\x36\xa8\xc8\x62\xea\x6d\x91\xc8\x98\xe7\x6f\x5e\xca\xa5\x17\x19\x83\x11\x97\x16\x46\x9a\xdb\x22\xc0\x26\xa6\xed\x9d\x8d\x86\x9e\xaa\x93\x2b\x30\xc9\xf5\xa7\x22\x6a\xcf\x77\x8e\x6f\xc8\x79\x1f\xee\x06\x9d\xc9\xe6\x96\xfe\x9a\x7e\x11\xec\xb0\x7d\x45\x3a\x77\xf3\xe3\x2b\x79\x3c\xbb\xbc\x3c\x93\x67\x63\xd5\xbb\x89\x7b\x97\xf1\xf9\xf5\xb4\x5b\x9c\x62\x14\x69\x65\x8c\xd2\x7c\xc4\x65\x80\xa9\x54\xb2\x48\x55\x6e\x70\xc9\xcc\x44\x9a\x67\x16\x19\x1d\x3d\x33\x89\x14\x03\x2f\xb9\xcf\x41\x17\x15\x83\x7a\xd8\x68\x7b\x2d\xaf\xe5\x19\xc1\xd3\x0a\x6d\xf2\x26\xd8\xce\x76\xef\x2e\xd9\x9e\xac\x33\x62\x58\xfa\xd7\x7d\x46\xe4\xc5\xe5\x54\xf0\x2f\x93\x6b\x73\x3a\x3c\xfa\x3c\x58\x1f\xef\x5e\xa4\x23\x42\x49\x37\x86\x0e\x1b\xd9\xf9\xb9\x69\xc7\xd9\x90\x8e\xb6\xba\x6c\x77\xb3\x29\xdf\x07\xeb\x93\x1a\xeb\xbb\xa8\x99\x4c\x8c\x17\x09\x95\xb3\xa1\xa0\x1a\x2a\xe8\x34\xa1\x33\x22\x78\x68\x88\x05\x1b\x83\x26\x2d\x6f\xc3\x6b\x92\x64\x39\xff\x01\x93\xa3\xb9\x65\x9d\xfe\xc1\xa0\x7f\xf5\xfd\x6b\x87\xb4\xe1\xb6\xdb\xbd\x1e\xe8\xc1\x61\xb1\x7d\xbc\x79\xd6\x0b\x61\x67\xd8\x4b\xc6\x9b\xa7\x9d\x93\xd9\xf5\xed\xe7\xb3\xf1\xd1\x6c\xeb\x92\xcb\xd6\xd1\x78\x30\xdb\x6c\x85\x07\x3a\xfc\x7d\x26\xbf\xaa\xa4\xe4\xb5\x90\xde\x26\x34\x39\x18\x0c\xe6\xe2\xee\x74\x07\xe8\x2e\x3d\xfc\xbe\x91\x75\x07\x6d\x7d\xf3\x39\x19\x25\x76\x7b\x9e\x8d\xcf\xb3\xbb\xf1\x7a\xf3\xd3\xd1\x6e\x16\xcf\x0b\xb8\x19\x77\xd7\x13\xd5\xe4\x70\xcc\xe7\xf7\xfd\x2f\x3d\xa5\x7f\xed\x68\x6c\x91\x41\x80\x2d\xcc\x2c\x49\xe8\x84\xd6\xab\xa5\xe2\x9c\x61\x2e\xa3\xd2\x5e\x48\xc2\xf4\x22\x2b\x47\x6b\x1f\x1e\x5c\xc7\x71\x54\x35\x31\x28\x40\x4c\x45\x79\x0a\xd2\x7a\x23\xb0\x5d\x01\xe5\xf0\xa0\x38\x61\x6b\x78\x11\x82\x3f\x94\xf1\x54\xd6\xe9\xab\x09\x91\x06\x6a\x61\x91\xb3\x86\x05\xaf\x43\xb9\xcc\x72\xdb\x53\xfa\x5f\x13\xaa\xc0\xb7\x72\x3c\x03\xb6\x63\xad\xe6\x61\x6e\x61\x0d\x97\xf4\xf0\x47\x54\x11\xfc\x89\x70\x49\xd3\x2a\xbc\xc6\xff\x13\x09\x99\xa0\x11\xc4\x4a\x30\xd0\x65\x5e\x57\x5a\xd0\x28\x53\x42\xa0\xba\x84\xe7\x79\x6f\x56\x51\x72\xa8\xa2\xbc\xec\xe1\x53\x77\x57\x5b\xe5\xd1\x2c\x03\xc9\x0e\x63\x2e\xd8\xda\xcb\xd4\xaa\xda\xbb\x8d\x17\xd4\xd8\x93\xba\x37\xaf\x90\x2e\x76\xc4\x1f\x91\xcc\x85\xf8\xdd\x2a\x9c\xad\x16\xf8\x61\x6b\xaa\xd0\xd5\x52\x2b\xda\x79\xc1\x6f\x49\xfa\x83\xeb\x3a\xce\xa3\xeb\xbc\x50\x69\x79\x8d\xef\xbb\x8e\xb7\xbc\xec\x4d\x29\xc1\x94\xea\x11\x97\x7b\xa8\x89\x68\x6e\xd5\x9f\xae\xe3\x4c\x39\xb3\xf1\x1e\xda\x6c\xfe\x51\xce\x42\x1a\x8d\x47\x5a\xe5\x92\x35\x22\x25\x94\xde\x43\xff\x6f\xb5\xc2\x61\x6b\x67\x7b\xa7\xfa\xac\x34\x03\xdd\xd0\x94\xf1\xdc\xec\x21\xaf\xad\x21\x2d\xd7\x33\xca\x18\x97\xa3\x3d\xb4\xa1\x21\x45\x9f\xaa\x55\xd7\x79\x74\x5d\x27\x54\xac\x40\xab\x3b\xbf\xda\xb7\xd5\xac\x37\x8e\x81\x8f\x62\xfb\x34\xaf\xc9\xd4\x14\x7c\x52\xbf\x5f\x7e\x59\x6b\xdf\x75\x1d\x9f\xf1\x09\x8a\x04\x35\x26\xc0\x4f\xec\x50\x92\xa7\xa1\xb2\x5a\xc9\xca\x83\xfe\x50\xe9\x14\xd1\xca\x87\x01\x26\x12\xa6\xa5\xb6\x08\x46\x29\xd8\x58\xb1\x00\x67\xca\x58\x8c\x40\x46\xb5\x8f\xcb\xf0\x06\xa3\x96\x56\xd9\x8e\x5f\x1d\xd0\x8a\xc7\x31\x2a\x15\xfe\xbc\x1f\x46\x2b\x02\x0e\xf0\xb7\x22\x83\xa7\x67\x15\xc5\xa0\xe1\x23\x5a\xba\x3d\x04\xa1\xa6\xde\xa2\x6e\xa8\x49\x3d\xc8\x05\xe2\x2c\x78\xf2\x7b\xb5\x58\x3e\xc3\xf5\xe0\x7d\x00\x0b\x83\xbd\xdc\xfe\x4d\xf7\x2c\x4a\x92\x65\xcd\xff\x4e\x71\xb4\xb0\x5c\x80\x57\x6e\x50\x5c\xf5\xeb\xd9\x15\xaf\x10\xf8\x24\x17\xff\x3c\xba\x28\x86\x68\x1c\xaa\xd9\x12\x43\x1f\xf4\x49\x1f\x2f\x8e\x20\xcc\xad\x55\x72\xa9\xa4\xd0\x4a\x14\x5a\xd9\xc8\x34\x4f\xa9\x2e\xf0\xa2\x82\xc9\xc3\x94\x5b\xbc\x7f\x58\x5d\xa3\x15\x52\x9f\xd4\x99\x95\xce\x48\xa9\x9c\xd2\x70\x84\xf1\x49\x29\xd6\x5a\xa5\x3e\xa9\xff\x7b\x3d\x3c\x80\x64\x8f\x8f\xee\xdf\x01\x00\x00\xff\xff\xae\x3e\xce\xba\xad\x09\x00\x00")

func viewIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewIndexHtml,
		"view/index.html",
	)
}

func viewIndexHtml() (*asset, error) {
	bytes, err := viewIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/index.html", size: 2477, mode: os.FileMode(420), modTime: time.Unix(1516830274, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewPollHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\x4d\x6b\xe4\x30\x0c\x86\xcf\xc9\xaf\xd0\x98\x1c\x77\x63\xe6\xba\x38\x86\xa5\xd3\x43\x4f\xfd\xa0\x0c\xf4\xe8\x19\x2b\x8d\xc1\xb1\x43\xa2\x4c\x19\x84\xff\x7b\x31\x9e\x8f\x42\x4f\x12\x8f\x5e\xe9\xb5\x25\x66\x8b\xbd\x0b\x08\x62\x8a\xde\x8b\x94\x6a\xb5\xd9\x3d\x3f\xbc\x7f\xbc\x3c\xc2\x40\xa3\xd7\xb5\xba\x06\x34\x56\xd7\x95\x22\x47\x1e\x35\x33\xb4\xaf\x2b\x2e\xe4\x62\x80\x94\x94\x2c\xb8\x56\xb2\xe8\xd4\x21\xda\x73\x96\x0f\x5b\xcd\x7c\x93\x66\xe5\xb0\xcd\xbc\x8f\xf3\x08\xe6\x98\x61\x27\x64\x36\x97\xcc\xed\xd3\x2e\x25\x79\x8a\x84\x52\xc0\x88\x34\x44\xdb\x89\x29\x2e\x24\x00\xc3\x91\xce\x13\x76\x22\x37\xfe\xb5\x86\x8c\xd0\x75\x55\xa9\xd5\xe7\x50\x31\xbb\x1e\xda\xff\x61\xf9\xc2\x79\x49\x29\xa3\x0c\x67\x13\x3e\x11\x1a\xf7\x07\x1a\x84\x7f\xdd\x2f\x45\xa5\xbc\xd3\x97\xb4\x52\x2e\x4c\x2b\x41\xb1\x99\x8d\x75\x51\x40\x30\x23\x76\xa2\x74\x09\x38\x19\xbf\x62\x27\x98\x1b\x97\x92\x00\xa9\x01\x98\x1b\x6c\xf7\x99\xdf\x67\xca\xdb\x50\x66\x0c\xb6\x14\xee\xa9\x92\xe5\xd1\xea\xb0\x12\xc5\x70\x31\x5c\xd6\xc3\xe8\x48\xe8\x7d\x24\xdc\x28\x59\x6a\x79\x53\x32\xff\x38\x27\x06\x86\x19\xfb\x9f\xdb\xb2\x29\xc9\x59\x0a\xfd\x86\xcb\xea\x69\x51\xd2\xe4\x0b\x94\xd5\x2b\x59\x0e\x77\xf5\xfd\x0e\x00\x00\xff\xff\x7f\x0c\xc2\x64\xea\x01\x00\x00")

func viewPollHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewPollHtml,
		"view/poll.html",
	)
}

func viewPollHtml() (*asset, error) {
	bytes, err := viewPollHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/poll.html", size: 490, mode: os.FileMode(420), modTime: time.Unix(1516830274, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewQuestionHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8c\x31\x0a\xc3\x30\x10\x04\x6b\xf9\x15\x87\x1f\x20\xf7\x41\x18\xf2\x84\x34\xe9\x0d\xde\x84\x03\x71\x21\xbe\x88\x14\xcb\xfd\x3d\x28\xc2\xe5\xec\x32\x43\xee\x78\xa8\x41\xe6\x77\x83\xcf\x11\x53\x22\x25\xdf\x1a\xfc\xa3\x2f\x93\x3e\x94\x56\xd7\x29\x25\xf2\xd8\xec\x09\xc9\x57\xf3\x2f\x0e\xef\x57\x2a\x55\xd7\xc1\xa2\x7e\x11\x32\xdf\xb7\xda\x10\x21\x65\xa9\x3a\x34\xd8\xfe\xcf\x2c\xbd\x73\xe2\x2f\x00\x00\xff\xff\x3b\x5a\xca\xad\x79\x00\x00\x00")

func viewQuestionHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewQuestionHtml,
		"view/question.html",
	)
}

func viewQuestionHtml() (*asset, error) {
	bytes, err := viewQuestionHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/question.html", size: 121, mode: os.FileMode(420), modTime: time.Unix(1516830274, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewResultsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x4d\x6b\xe3\x3a\x14\x5d\x3b\xbf\xe2\x3e\x6f\x9c\x40\x90\xda\xc5\x5b\xbc\xd4\x0e\xb4\x69\x17\x81\xc7\x34\x03\xa1\x50\x4a\x17\x37\xd6\x8d\xad\x56\xb1\x8c\x24\xc7\xcd\x08\xfd\xf7\x41\x76\x3d\x5d\xcc\xc0\x78\x63\xdf\x8f\x73\x7d\xce\xd1\x95\xf7\x82\x8e\xb2\x21\x48\x0d\xd9\x4e\x39\x9b\x86\x30\xcb\xff\xb9\x7f\xdc\xec\x9f\x77\x0f\x50\xbb\x93\x5a\xcf\xf2\xe9\x45\x28\xd6\xb3\x24\x77\xd2\x29\x5a\x7b\x0f\xec\x7b\x47\xd6\x49\xdd\x40\x08\x39\x1f\xd3\xb3\x24\xc9\x6d\x69\x64\xeb\xc0\x5d\x5a\x2a\x52\x47\x1f\x8e\xbf\xe1\x19\xc7\x6c\x0a\xd6\x94\x45\x5a\x3b\xd7\xda\x15\xe7\x7d\xdf\xb3\xca\x3a\x74\xb2\x64\xa5\x3e\xf1\xb2\x46\xe3\x2c\x57\x1a\x05\x19\xf6\x66\xd3\x75\xce\x47\xe0\xdf\x27\xc7\x8e\x4a\xeb\x4a\x11\x1b\xc7\xb0\x38\x66\x9e\x95\x9d\x31\xd4\xb8\x6c\x09\x3e\x6b\xb1\x7c\xc7\x8a\x6c\xb6\x7a\xc9\x4a\x6d\x68\x68\xcc\x5e\xc3\xe2\xe6\x37\xb0\x25\xf7\xd8\xfc\xaf\x51\x6c\x50\xa9\x03\x96\xef\x73\x61\xb0\xdf\xc4\xe2\xd0\x7d\xec\x9a\x72\x10\xff\x2b\x3d\x5f\x80\x9f\x25\x49\x72\x46\x03\x02\x1d\x42\x01\x9f\x23\xcf\xd2\x76\xa8\xe4\x0f\x8c\x00\x86\xc6\xe0\x65\xaf\xef\xd1\xe1\x1e\x0f\x8a\xe6\x2f\x11\x05\xf0\x92\x3d\x69\x47\x36\x5b\x42\xf6\xad\x3b\x1d\xc8\x80\x3e\xc2\x98\x7a\x5d\xc6\x16\xef\x0d\x36\x15\x01\xbb\x6d\x6c\x4f\xc6\x86\x30\x01\xbd\x67\x4f\xa8\x3a\x0a\x21\x5b\x42\x7c\xbc\x67\x7b\xed\x50\x85\x30\x40\x01\x12\xef\xa9\x11\x23\xe2\x75\x71\x33\x9b\x88\xea\x36\x72\xb2\x50\x8c\xdc\x01\x86\x73\x5c\x41\x1c\x39\x9d\x6f\x08\x59\xac\x85\x2f\xd4\xe0\x11\x14\xd0\x50\xff\x67\x8d\x3b\x49\xa3\x27\x42\x97\xdd\x89\x1a\xc7\x2a\x72\x0f\x8a\xe2\xe7\xdd\x65\x2b\xe6\x59\x2b\x3f\xdd\x5f\x7c\xb2\x19\x22\x16\xdd\x9c\x47\xf7\x96\x13\xb5\xc1\xed\x30\x8b\xaa\xbe\x96\x21\xe7\xe3\x3a\xe6\x07\x2d\x2e\xf1\xe4\xbd\x97\x47\x60\xb7\xca\x10\x8a\x4b\x34\x6d\xd4\x9a\xd7\xd7\xeb\x67\xdd\x01\x8e\x05\x38\xc7\x0a\x83\x0d\x36\x8d\x76\x43\x04\xae\x97\x25\xc1\xd1\xe8\x13\x58\x3c\x11\x6c\x77\x2c\xda\xb7\xdd\xc5\xad\xae\xaf\xc7\xe1\x93\x77\xb9\x90\x67\x90\xa2\x48\x27\xfa\x29\x58\x77\x51\x54\xa4\xbd\x14\xae\x5e\xc1\x7f\x57\x57\xed\xc7\x0d\xd4\x24\xab\xda\xad\xe0\xdf\x21\x8c\x6b\x2c\xe4\x79\xd8\x61\x84\xda\xd0\xb1\x48\x79\xab\x95\xe2\xf1\x47\xf7\x21\xf0\x74\x1d\x39\xe7\x1c\x87\x9e\x83\xe1\xeb\x41\x70\xb2\x23\xb3\xdd\xdd\x29\x5d\xbe\xaf\x22\xa9\x21\x8c\x77\x94\x8f\xba\x73\x3e\x5e\xce\x89\xe0\xcf\x00\x00\x00\xff\xff\x98\x98\xb1\x62\xd1\x03\x00\x00")

func viewResultsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewResultsHtml,
		"view/results.html",
	)
}

func viewResultsHtml() (*asset, error) {
	bytes, err := viewResultsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/results.html", size: 977, mode: os.FileMode(420), modTime: time.Unix(1516830274, 0)}
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
	"view/index.html": viewIndexHtml,
	"view/poll.html": viewPollHtml,
	"view/question.html": viewQuestionHtml,
	"view/results.html": viewResultsHtml,
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
	"view": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewIndexHtml, map[string]*bintree{}},
		"poll.html": &bintree{viewPollHtml, map[string]*bintree{}},
		"question.html": &bintree{viewQuestionHtml, map[string]*bintree{}},
		"results.html": &bintree{viewResultsHtml, map[string]*bintree{}},
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
