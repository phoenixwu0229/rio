// Code generated by go-bindata.
// sources:
// stacks/coredns-stack.yml
// stacks/istio-crd-stack.yml
// stacks/istio-stack.yml
// stacks/local-storage-stack.yml
// stacks/storageclasses-stack.yml
// DO NOT EDIT!

package stacks

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

var _stacksCorednsStackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xcd\x6e\xdb\x3c\x10\xbc\xfb\x29\xf6\x01\x2c\x4a\x82\x3e\x7f\x75\x79\x6c\xea\x43\x2e\x39\xb4\x45\xaf\x06\x43\xae\x2d\xc2\xfc\xc3\x2e\x95\xc6\x68\xfb\xee\x05\x4d\xcb\x50\xe3\x46\x17\x2d\x77\x86\xcb\xc5\xcc\xe8\x18\x0e\xf6\xc8\x72\x05\x50\xca\xf2\xbf\x54\x19\x43\x96\xf0\xab\xb9\x9c\x01\x84\xdc\x0c\xf0\xf3\x7a\x28\x1f\x12\x45\xe2\x45\x63\x44\xe5\xf2\xb8\x68\x9c\xa6\x67\xa4\x80\x19\x19\xb4\x9b\x38\x23\x09\x17\xb5\x72\x60\x43\xa3\x8c\x21\xa1\x28\x29\xb0\xe9\xff\x5a\x2c\x87\x97\x2f\x45\xc3\x60\x03\xa3\x9e\x08\xdf\x60\x53\xe2\x4c\xa8\xfc\x9b\xf6\x41\x39\x97\x47\x8a\xd3\x71\xfc\xf7\x23\x0b\xfe\xef\x45\x9d\x28\x7a\xcc\x23\x4e\x0c\xf2\x63\xbf\x19\xfe\x86\x5e\xcf\x20\xa0\xc5\xac\x5b\x42\x8e\xee\x45\x14\x9d\x16\x14\xad\xf4\x88\x30\x74\xab\x79\x30\x23\xbd\x58\x8d\x57\x4d\x09\x4d\xe0\x2a\x6b\x29\x60\x8c\x9c\xab\xc8\x2a\xed\x95\x31\x12\x9e\x76\xdf\xf6\x9f\x1e\x9f\x3e\xef\xbf\xee\xbe\x7c\x7f\x7c\xd8\xdd\x50\x43\x31\x49\x50\xce\x5d\x4d\xf1\x5e\x05\x53\x47\x35\xd0\xdc\xd6\x68\xea\x76\xd7\xa7\xda\x87\x48\x78\xb0\x0e\x67\x27\x67\x7b\x0b\xf1\x62\xf1\xfb\x6c\x7c\x4d\x91\x71\x26\xf7\x9d\xf8\x6f\x10\x9d\xe8\x3b\xb9\x19\xe4\x66\x68\x27\x93\xd6\x26\xf0\x3b\x70\x81\x9a\xac\xd3\x15\x2e\x42\xae\x3d\x66\xb2\xba\xde\x38\xba\xf8\xac\xdc\x3e\x21\x79\xcb\x6c\x63\xb8\x6d\xe5\x2c\xe7\xf5\x0f\x95\xf5\x08\x18\x4c\x8a\x36\x64\xbe\x87\x66\x59\xef\x91\x12\x95\xfb\x6e\x50\x1e\x39\xa9\xf9\x46\xcd\xa7\x1e\x51\x9f\x24\x8c\x39\x27\xd9\xb6\x97\x3c\x16\x43\xe4\xb6\xdb\x76\xed\x22\xc2\xd6\xab\x23\x4a\x38\x6d\x59\x1c\x35\x09\x1b\x67\xc1\x64\x2f\x7a\x51\x13\xe2\xd1\x47\x3a\x4b\xf8\xd0\xd5\x20\x12\x2a\xb3\x8f\xc1\x9d\x25\x64\x9a\xaa\xa2\xac\x95\x43\x09\xfd\xea\x4f\x00\x00\x00\xff\xff\xd5\x57\x8c\x0c\x68\x03\x00\x00")

func stacksCorednsStackYmlBytes() ([]byte, error) {
	return bindataRead(
		_stacksCorednsStackYml,
		"stacks/coredns-stack.yml",
	)
}

func stacksCorednsStackYml() (*asset, error) {
	bytes, err := stacksCorednsStackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stacks/coredns-stack.yml", size: 872, mode: os.FileMode(420), modTime: time.Unix(1534733485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stacksIstioCrdStackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd1\x41\x4b\xc3\x40\x10\x05\xe0\x7b\x7e\xc5\x42\xce\x46\xaa\x37\x6f\x8a\x45\xbd\xc5\xb6\x78\x0d\xe3\xe6\x35\x1d\x9a\xce\x84\xdd\xd9\x96\xfe\x7b\x71\x5b\x14\x2d\x24\x3d\xbf\xef\x0d\xcb\xdb\x6d\xfa\x44\x10\x18\xe2\x43\xe1\x9c\xd0\x0e\x71\x20\x8f\xb6\xf1\x29\x9a\xee\x9a\x80\xa8\x29\x78\x34\x2d\xd6\x2c\x6c\xac\x92\x65\xe9\xbc\xca\x9a\xbb\xc2\xb9\x1b\xf7\x8c\x68\x2c\xf4\x1d\xd6\xda\xb3\x3f\x56\xa7\xb0\xe2\x68\xac\x15\xeb\xed\x7e\x46\xfd\xb0\xa1\xbb\xcc\xe7\x5d\x40\x8c\x8b\xd4\x63\xdc\x2d\x34\x19\xa6\xd9\xeb\x6a\x55\x3f\xd6\x6f\xcb\x01\xfe\x89\xa5\x65\xe9\xae\xf6\xe3\xf0\x3d\xa9\xd1\xd5\x67\x7f\xf4\x18\x2b\x9d\xc0\x0e\x1a\xb6\x2c\xa7\xe9\x3e\x38\x58\xa2\x7e\x89\xb0\x67\x8f\xea\x37\xbd\xa8\xdf\xff\x9f\x3a\x2f\x33\x59\x38\x5f\x9e\x8b\x85\xe3\xb4\x7e\x21\xc3\x81\x26\x61\xe9\x28\xd9\x06\x62\xec\xf3\x53\x72\xf7\xfc\xf5\x7f\x93\x8b\xfa\xac\xf8\x0a\x00\x00\xff\xff\xbd\xa2\x21\x22\x73\x02\x00\x00")

func stacksIstioCrdStackYmlBytes() ([]byte, error) {
	return bindataRead(
		_stacksIstioCrdStackYml,
		"stacks/istio-crd-stack.yml",
	)
}

func stacksIstioCrdStackYml() (*asset, error) {
	bytes, err := stacksIstioCrdStackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stacks/istio-crd-stack.yml", size: 627, mode: os.FileMode(420), modTime: time.Unix(1539128986, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stacksIstioStackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x51\x6f\xdb\x38\x12\x7e\xcf\xaf\x20\x74\x05\x72\x2d\x42\xcb\x69\x9b\xbb\x56\x80\x1f\x0c\xc7\x45\x03\xa4\xb1\x2f\x76\x71\xd8\x27\x83\x16\xc7\x36\x11\x8a\x24\xc8\x91\x53\x6f\xb7\xff\x7d\x41\x4a\xb6\xa8\x44\xae\xb3\xbb\x58\xec\x4b\x12\xf1\xfb\x66\x38\xe4\xcc\x7c\x9c\xe4\x5a\xad\xc4\xda\x65\x67\x84\x14\xe0\x36\xfe\x37\x21\xb9\x56\x08\x0a\x33\xf2\x1b\x0d\xdf\x84\x70\xe1\xd8\x52\xc2\x54\x4b\x91\xef\x46\x1b\xc8\x1f\x5c\x46\xd0\x96\x50\xe3\x42\xad\x2d\x38\x37\xd2\x0a\xad\x96\x12\xec\x17\xcd\x21\x23\xc9\xe4\xd3\xa7\xa4\xa6\xb0\x12\x37\x95\x7d\x46\xee\x26\x77\xe3\x7a\xd9\x72\x77\x0f\x2b\x0b\x6e\x73\x0d\x92\xed\x32\x72\xd9\x77\x35\xa4\x4b\x5c\xea\x52\xf1\xb9\x65\xab\x95\xc8\x6b\xe3\x1a\x24\xa4\x08\x5b\x0c\x6f\x6f\x27\xff\x5f\x0c\xef\x7e\xd9\x47\x0a\x2b\x56\x4a\x1c\x85\x73\x35\x64\x2e\x5c\xae\xb7\x60\x77\xc7\x36\x0b\xa7\x56\x90\xe3\x5c\x14\xa0\x4b\xcc\xc8\xbb\x36\xb6\x12\xeb\x29\xc3\x4d\x46\x92\x14\x30\x4f\x85\x43\xa1\x53\x63\xf5\xb7\x5d\x72\xa0\x2d\x85\x62\x76\xb7\xa7\x95\xce\xa6\x52\xe7\x4c\xa6\x4b\xa1\x52\x50\x5b\x1d\x51\x1d\xd8\xad\xc8\x61\x24\x4b\x87\x60\x33\x12\xfc\xd1\xe0\xaf\x09\xda\x32\xa1\xae\x4b\xcb\x50\x68\x95\x91\xf7\x57\x4d\x40\x86\x59\x50\x38\xdb\x94\xc8\xf5\x63\xc4\xb9\x2c\xa2\xa8\x85\x42\xb0\x39\x18\x0f\x55\x09\xb9\x1f\x5f\xdf\xdc\x8f\x47\xf3\xc6\x8f\xdf\x70\xc8\x0b\xa1\xa6\xda\x62\x46\x2e\xaf\xfa\xfd\x7e\x7c\x6c\x9f\xcf\xa9\x64\x0a\x86\xdd\xf9\x8b\xae\x76\xc8\xb9\x2f\x82\xc3\x59\x84\xd4\xd8\x7b\xf5\xfd\x6e\xf8\x65\x3c\x9b\x0e\x47\xe3\x1f\x99\xf7\xfe\xdf\xb3\xb3\xfa\xec\xa1\xea\x22\x6e\x95\x2d\xc9\x96\x20\xdd\x3e\x73\x01\xce\x48\xc0\xc3\xd2\x5a\xea\x25\x93\x0b\x03\xb6\x10\xce\x09\xad\x6a\x2a\x25\xe7\x6f\xea\x34\xf5\x82\x51\x4f\xe8\xf4\xcd\x79\x83\x29\xc0\x47\x6d\x1f\x84\x3a\x82\xfb\xfa\x04\x85\x22\x0f\x37\x79\x84\x63\x04\x7c\x43\x50\x61\xdf\xde\xc3\x07\xe7\x09\x79\xe9\x50\x17\x16\x9c\x2e\x6d\x0e\x1c\x56\x42\x09\xef\xc2\x45\x86\x8d\x55\x8a\x1b\x61\xb9\x61\x16\x77\x7b\x93\x97\x13\x7b\x0d\xdc\x6d\x53\xf7\xe1\x31\x97\x07\x38\x75\xc8\xb0\x3c\xb0\xec\x92\xe5\xd1\x89\xeb\xd5\xdc\x02\x43\xb8\x58\x03\x5e\x48\xe1\xf0\xe2\x91\x61\xbe\xb9\x28\x0d\x67\x08\xf5\x55\x17\xcc\xb8\x9a\x0d\x8a\x1b\x2d\x14\xee\xbf\x8d\xe6\xfb\x3f\xf7\xf9\xae\x3f\x15\x2b\xc0\x19\x16\x2d\x68\x0e\x0d\x37\xb7\x50\x3b\x71\x82\xc3\x83\xf0\x5a\x73\xf6\xa4\xd4\xb2\xa8\x44\x8b\x82\x29\x9e\x35\x18\xa1\x54\xea\xf5\x42\x97\x68\x4a\x5c\x48\xd8\x82\x1c\x70\x58\x96\x6b\x42\xa9\x05\xb7\x53\xf9\xe0\xed\xe6\x49\x67\xbb\xc6\x21\xad\x84\x30\x6a\xf2\x8a\x92\xfa\xe5\x03\x0b\xd4\x56\x58\xad\x0a\x2f\x93\x91\xe9\x74\x72\xbd\xf0\xf5\x3e\x78\xf5\x6f\x07\x72\x95\xfa\xa3\xbe\xee\xc0\x43\x3f\xc4\xa4\x70\x1f\x2d\xe6\xcd\xed\x64\xbe\x98\x7f\xbe\x9f\xcc\xe7\xb7\xe3\xc1\x55\xd4\x95\x7b\x70\x34\x1c\x7d\x1e\x2f\x66\xff\xfb\x3a\x9c\x7d\x1e\x5c\x35\x5d\x5f\xb0\x35\xd4\x4d\x98\x56\x8d\x75\xd9\xeb\xf7\xde\x46\xd2\x13\xae\x38\x23\x82\xfb\x82\xc7\x5d\x75\xd8\x1c\x2c\xba\xb8\xed\x2a\x35\x6a\x8e\x07\xdf\x8c\x76\x10\x1f\x37\xf4\x73\xba\x41\x34\xed\xc5\xcb\x7e\xba\xb6\x26\x8f\x16\x3f\xf6\x3f\xbe\x0b\xc4\x0b\xff\x83\x16\x5a\x09\xd4\x56\xa8\xf5\x91\xb0\xfd\xce\xdb\xb7\x4f\x02\xdf\xe7\x3a\xf2\xdb\x16\x4c\x4a\x28\x6d\x0b\x6b\x04\x45\x4a\xd3\x32\x40\x28\x8c\x64\x08\x9f\x84\x84\x08\x78\xaa\xf2\x95\x80\x2f\x2a\x55\xdb\xb1\x42\xf6\xb0\x30\xb2\xe5\xa9\x5b\x2e\x23\x4a\x4b\x37\xff\xf6\x12\xba\xb9\x9b\xcd\x87\x77\xa3\xf1\xe2\x66\xba\xe7\x09\xf3\xfa\x65\x65\x70\x50\xe6\x5c\x20\xe3\x20\xab\x08\xeb\x1c\x25\x75\x5f\xd4\x50\x48\x52\x72\xf6\x2c\x43\xfe\x4a\x98\x31\xa0\x38\xe5\xca\xd1\x10\xe3\xe0\x30\x34\x78\xd4\x17\x09\x35\xda\xe2\xe0\x43\xff\x3f\xfd\xf6\xf2\x46\x3b\xf4\x26\x83\x7a\x9b\xb3\x26\xc1\x72\x45\x9d\x58\x2b\xe0\x34\x67\x6d\x87\x35\x97\x3a\xd4\x96\xad\x81\x1e\xee\x65\x60\x85\xa6\x6e\xe7\x10\x8a\x13\xcf\xc8\xa3\x15\x08\x2d\x19\x3a\x08\x18\xcb\x73\x5d\xaa\xa7\xab\xd5\x67\x87\x27\x2f\x9c\x1c\x24\x20\x10\x0e\x46\xea\x9d\x4f\xb5\x7b\x0e\x76\x7b\x7f\x4e\xe8\x40\xba\xfd\x06\xb5\xae\x74\x7a\x4f\x8c\x1e\x81\x3f\x63\x63\xc1\x48\x91\x33\x77\x50\xe6\x17\xd5\xce\x9a\x21\x3c\xb2\x5d\xd7\xbb\x9e\xd4\x58\x92\x91\xc4\xef\x63\x15\x93\xc9\x89\xcc\x24\x3e\xce\xea\x11\x8a\xa2\x3d\xf1\x64\x26\x3f\x37\x7e\x93\x6e\x85\xc5\x92\xc9\xfd\x1d\x9f\xe4\x73\x70\x28\x54\x18\x12\x6c\x29\x5f\x60\x50\x9f\xb4\x26\x56\xcb\x0b\x6d\xb9\x9f\xfa\x1c\x6a\x43\x57\xc2\x3a\xec\x7e\xf0\x3a\x75\xd8\x37\x8c\xcb\xc8\xab\xef\xd3\xc9\xfd\x7c\xf6\xe3\xa9\x82\x26\x1d\x12\x9a\xfc\x11\x0d\xb5\xba\x6c\x2b\x27\xdd\x46\x1f\x49\xe4\xcc\xb7\x5c\xe7\x64\x1d\x31\xce\x2f\xdd\x39\xf9\xd7\x29\x16\xa5\xad\x61\x37\xb6\x7f\x7f\x15\x1c\x1c\x81\x29\xed\x1e\x85\x5b\x11\x14\x7d\xef\xe2\x24\x31\x28\x78\xf4\x3f\x40\xcb\x47\x70\x71\x14\x3f\xfd\xee\x3c\x7b\xa8\x7e\x15\xe6\x41\xa8\x7a\x68\x8e\x2f\xb8\x7d\xbf\x7e\x52\x73\xfc\x2b\x37\x27\x99\xed\x59\x3e\xe6\x85\xa9\x3e\xf9\x0b\x6f\x55\x2b\xd1\xcf\x03\x89\x87\xf8\x6a\xc6\xdf\x63\xa0\xb6\xff\xc8\xe3\x46\xc9\xcd\x6c\x7e\x33\x59\x7c\x19\xcf\x87\x8b\x9f\xef\xfa\x33\x2d\xfb\x3d\x00\x00\xff\xff\x1b\x4c\xc5\x09\x1f\x0f\x00\x00")

func stacksIstioStackYmlBytes() ([]byte, error) {
	return bindataRead(
		_stacksIstioStackYml,
		"stacks/istio-stack.yml",
	)
}

func stacksIstioStackYml() (*asset, error) {
	bytes, err := stacksIstioStackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stacks/istio-stack.yml", size: 3871, mode: os.FileMode(420), modTime: time.Unix(1540013546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stacksLocalStorageStackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x8f\xd3\x30\x10\x85\xef\xfd\x15\x56\x84\xb4\x80\x36\xf1\x8a\x13\x8a\xc4\xa1\xa2\x59\x81\xb4\xdb\x56\xb4\x9c\x10\x8a\xa6\xce\x6c\x62\xb0\x3d\x96\xc7\x9b\x0b\xf0\xdf\x51\x12\x27\x6c\x2b\x95\x93\xe7\xcd\x7b\x63\xd9\xfe\xac\xc8\x3d\xe9\x96\xcb\x95\x10\x53\x39\x54\x63\x1d\xd1\xc5\x52\xfc\xce\x47\x2d\xc4\xaf\xb4\x0a\x91\x39\x6a\x70\x0f\xb1\x7b\x04\x9f\x95\xe2\xdb\x62\xbc\x0c\x2d\xc1\xac\x14\xd9\xa6\xba\x5f\x7f\x7d\x38\xd6\xfb\xf5\xf1\x53\x7d\xbf\xfb\x52\x6f\x77\xdb\xfa\xe1\xf3\xe1\x58\x6d\xea\xed\x6e\x53\x1d\xb2\xdb\xf3\x41\x0f\xb1\xe3\x8b\xbd\x47\x43\xf6\x10\xa4\xd1\x27\x19\xc0\xa9\x0e\x83\x0c\x9a\x64\x4f\xe6\xd9\x22\x67\x67\xe1\xef\x2f\xd4\x9f\xd5\x65\x77\xea\x18\x38\xa1\xe1\x32\xf5\x14\xc4\x68\xb0\xd0\x24\x55\x40\x88\x14\x4a\xe1\x28\x58\x70\x2b\xc6\xd0\x6b\x85\x63\xd2\x90\x02\x93\x0f\x07\xcc\x7d\xa0\x5e\xb3\x26\x87\x61\x7e\x35\x6b\xc1\x35\x93\xc8\xaf\x44\x93\x99\xe7\x0d\x9e\x9e\xdb\xa4\x38\x42\x88\x8b\x33\x91\x48\x52\x62\x54\x72\xea\xa4\xa5\xf8\xc1\xe4\x66\x4c\x23\xbd\x99\xdd\xff\xb2\xe8\x7a\x1d\xc8\xd9\x01\x6b\xda\x7a\xbf\xdb\xd4\xdb\xf5\x63\x75\xd8\xaf\x3f\x56\x1f\x5e\xbd\x66\x34\x4f\xd2\x81\x45\xf6\xa0\xf0\xcd\x98\x6a\x0d\x9d\xc0\xd4\x1e\x83\xd5\x3c\xdc\x80\xe7\xe9\x81\x2e\xa7\xda\x63\x60\xcd\xc3\x97\x99\x60\x28\x03\xda\xce\xe6\xcd\x5b\x81\xae\xf1\xa4\x5d\xe4\x9b\x7f\xbd\xcb\x99\x33\x8f\x9a\x45\x8e\x34\xf0\xd6\x43\x54\x9d\xc0\x1e\x5d\xe4\xe5\xd5\x28\x40\x8b\xc5\xcf\xf7\x3c\x60\x4b\x52\x19\x60\x4e\x27\xd3\x16\x5a\x2c\xc5\xfc\x5b\xae\xc0\xeb\xef\x8a\xbb\xe2\xdd\xea\x6f\x00\x00\x00\xff\xff\x4c\x0a\x8d\x92\x0c\x03\x00\x00")

func stacksLocalStorageStackYmlBytes() ([]byte, error) {
	return bindataRead(
		_stacksLocalStorageStackYml,
		"stacks/local-storage-stack.yml",
	)
}

func stacksLocalStorageStackYml() (*asset, error) {
	bytes, err := stacksLocalStorageStackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stacks/local-storage-stack.yml", size: 780, mode: os.FileMode(420), modTime: time.Unix(1540013894, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stacksStorageclassesStackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x31\x6b\x33\x31\x0c\x86\xf7\xfc\x0a\x91\xdd\xf9\xf8\xb6\xe2\xb1\x29\xd9\x0a\x85\x42\x3b\x16\xf5\xac\x24\xe2\x6c\xe9\x90\x74\x07\x85\xfe\xf8\x72\xbe\x86\x0e\x9d\x3b\xfa\x7d\x1f\xeb\xb1\x35\xce\xef\x64\x42\x41\x9e\x77\x00\x82\x8d\x7c\xc2\x81\xca\x5b\x43\xe1\x33\x79\x64\xf8\x4c\x3b\x00\x00\x9c\xf8\x85\xcc\x59\x25\x83\x87\x1a\x5e\xe8\x30\xde\xf9\x81\xf5\xdf\xf2\xbf\x13\x23\x4b\xc9\xf0\xbc\x75\xc7\x8a\xee\x3d\x6e\x14\x58\x30\x30\xf7\xd3\x26\xc9\x20\xe7\xad\x9d\x4c\x17\x5e\xa7\x92\x6d\x61\x4f\x53\x4a\xbb\xbf\xb0\x56\x1d\xb0\x7e\x27\x28\xa2\x81\xc1\x2a\x7e\x83\xe0\xe6\x18\xd6\x39\x87\x9f\xe5\xac\x3e\xf6\x54\xe8\x8c\x73\x8d\xd4\xeb\x0c\xfb\xb0\x99\xf6\xbf\xbf\x61\x28\xc3\x95\x6c\xbd\xd4\x85\x69\xc2\xb8\x76\x6c\xd1\x3a\x37\xba\x67\x29\x2c\x97\x47\x2d\x94\xe1\x15\x39\x4e\x6a\x27\x36\x8f\xa3\x8a\xcf\x8d\xac\xb3\xb6\xbe\x82\xdb\x93\x56\x1e\x3e\x32\x3c\x50\xa5\xa0\xaf\x00\x00\x00\xff\xff\x0a\xc5\x63\xf0\xb0\x01\x00\x00")

func stacksStorageclassesStackYmlBytes() ([]byte, error) {
	return bindataRead(
		_stacksStorageclassesStackYml,
		"stacks/storageclasses-stack.yml",
	)
}

func stacksStorageclassesStackYml() (*asset, error) {
	bytes, err := stacksStorageclassesStackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stacks/storageclasses-stack.yml", size: 432, mode: os.FileMode(420), modTime: time.Unix(1540013894, 0)}
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
	"stacks/coredns-stack.yml":        stacksCorednsStackYml,
	"stacks/istio-crd-stack.yml":      stacksIstioCrdStackYml,
	"stacks/istio-stack.yml":          stacksIstioStackYml,
	"stacks/local-storage-stack.yml":  stacksLocalStorageStackYml,
	"stacks/storageclasses-stack.yml": stacksStorageclassesStackYml,
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
	"stacks": &bintree{nil, map[string]*bintree{
		"coredns-stack.yml":        &bintree{stacksCorednsStackYml, map[string]*bintree{}},
		"istio-crd-stack.yml":      &bintree{stacksIstioCrdStackYml, map[string]*bintree{}},
		"istio-stack.yml":          &bintree{stacksIstioStackYml, map[string]*bintree{}},
		"local-storage-stack.yml":  &bintree{stacksLocalStorageStackYml, map[string]*bintree{}},
		"storageclasses-stack.yml": &bintree{stacksStorageclassesStackYml, map[string]*bintree{}},
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
