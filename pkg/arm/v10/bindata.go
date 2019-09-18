// Package arm Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

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

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x6f\x57\x1b\xb7\xd2\x7f\x5d\x7d\x8a\xe9\x9a\x36\x0d\x45\x5e\x48\xdb\xe7\xe9\x71\x4b\xce\x21\x40\x7a\xb8\xa5\x81\x0b\xc9\xed\x8b\x34\x27\x47\x5e\xcd\xda\x8a\xb5\xd2\x56\x7f\x6c\x5c\xe2\xef\x7e\xcf\x68\xd7\xc6\xc6\x86\x24\x25\xbd\x79\x41\x6c\x69\x34\x33\x1a\xfd\x34\xbf\x19\xb9\xf3\x65\xde\x57\x26\xef\x0b\x3f\x04\x8e\x57\x8c\x75\xe0\xb9\x75\x10\xd0\x07\x65\x06\x3d\xd0\x76\x00\xc2\x48\x90\xce\xd6\x20\xb4\x86\xe0\x44\x59\xaa\x02\xc2\x50\x04\x98\xd8\xa8\x25\x38\x1b\x03\xc2\x58\x09\x08\x43\x84\x4a\xf8\x80\x0e\x8e\x4f\x9f\xb1\x0e\x5c\x1c\x5f\x9e\xbd\xba\x38\x3c\xfe\xe5\xe2\xec\xd5\xf9\x7e\x36\xb5\xd1\x71\x87\xde\x46\x57\x20\x1f\x38\x1b\xeb\x8c\x75\xe0\xec\xf2\xed\xf3\x7f\x1f\xbd\xd8\xcf\x6c\x8d\xc6\x0f\x55\x19\xba\x5b\x2b\x2b\xbb\xd6\x0b\x89\xe3\x6e\xa1\x6d\x94\x19\xeb\xb0\x0e\xa8\x3a\x88\xbe\x46\x0f\xfc\x04\x4e\x5e\x9c\xbf\x7a\x09\xdc\xc3\xd6\x37\x52\x0d\xe0\x5b\x3f\xb4\x2e\x40\xb6\xd5\xea\xcd\xe0\x3d\x04\xa1\x34\xf0\xbd\xc7\xc0\xdf\xc1\xe9\xd9\x2f\xc0\xb9\xb6\x03\x5e\x3b\x2c\xd5\x15\x64\xbf\xbe\x7a\x76\x0c\x24\x0a\x47\x17\x67\xe7\xbd\xec\x61\xfa\x49\x07\x63\xd7\xd7\xa0\x4a\xe8\x1e\x5a\x53\xaa\x41\xf7\x12\x8b\xe8\x54\x98\x9e\x8b\x50\x0c\xcf\x45\x31\x12\x03\xf4\x30\x9b\x31\x6d\x07\x03\x74\xc0\x43\x1b\x38\xee\x83\x70\x21\xd6\x5d\x3f\x84\x4c\x19\x1f\x84\xd6\xca\x0c\xc0\xa1\x04\x0a\x79\x21\x0d\x14\x49\x67\x74\x22\x28\x6b\xc0\x1a\xd8\xfa\x66\x68\x7d\x30\xa2\xc2\xc7\x19\x2b\x44\x80\xa7\xf9\x58\xb8\x5c\xab\x7e\x3e\x8d\x55\x5e\x68\x85\x26\xf0\x02\x5d\xe8\xd6\x58\xc1\xcf\x3f\x3f\x3a\x3e\x7b\xfe\x88\x5c\x3c\x44\x17\x0e\xfc\xb3\x69\x40\xbf\xf0\x95\xc6\x54\xa9\x0a\x11\xd0\x77\x5b\x5f\x2f\xb0\xb6\x5e\x05\xeb\xa6\x69\x1a\xde\xc3\x65\x70\xe4\xd7\x6c\xc6\x8e\xcf\x9e\xdf\x6d\x74\x84\xd3\xdb\x36\xcf\x9d\x1a\x8b\x80\xbf\xe2\xf4\x13\x2d\xff\x8a\xd3\x35\xc3\x1f\x1d\xc0\x83\x8b\x33\xf0\xed\x29\x40\xac\x25\xd9\x80\xd7\xd7\xd7\xad\x3e\xff\x2f\xab\xcc\x07\x8e\x2b\xdb\x81\x0c\x66\xb3\x37\x6b\x21\x2f\xad\x03\x11\x02\x56\x75\x00\x65\xe0\x7a\xaf\xdb\xfd\x61\xf6\x13\x48\xcb\x00\xa6\xb1\x82\xd6\x0d\xe0\x53\xe0\x7f\xc2\xa7\xd9\x4c\x26\xe1\xeb\xaf\xa1\xef\x50\x8c\x18\xc0\xbd\x1b\x7e\x3d\x77\x63\xeb\xba\xfd\x34\x7b\xb3\x79\xeb\xad\x4f\x0d\x86\x4a\xa1\x34\xca\x8c\x01\x61\xf6\xf5\xeb\xa5\xd5\xc0\x75\x80\x1f\xe0\xcd\x9b\x9f\xe8\x76\x1b\xf0\x1a\xb1\x86\xbd\x9f\x00\xb5\x47\xc0\x2b\x15\xe8\x4b\xa9\x98\xb4\x06\x3f\x70\x1a\x0e\x2b\x3b\xfe\x34\x30\x53\xf4\x0a\x8d\xc2\x50\xf2\x61\xae\x02\xee\x4a\xb8\x17\xdc\xf7\x80\x90\x5d\x5f\xa3\x91\xb3\x19\x65\xb9\xc2\xa1\x08\x48\xd6\x83\x50\x06\x1d\xd4\x51\x6b\x8a\x92\xc3\x00\xac\x1a\x49\xe5\x80\xd7\x37\xda\xac\x53\x03\x65\xf2\xae\xb4\xc5\x08\xdd\x2d\xbc\xaf\x4e\xe6\xcd\x96\xba\xef\xbc\x35\xcb\xb8\xef\x1e\xa1\x53\x63\x94\xdd\x43\x5b\xf5\x95\x41\x79\x52\x89\x01\x9e\x47\xad\x2f\x1b\xb3\x2d\x12\xd6\x30\xae\x0d\x25\x9f\x3b\xac\x41\xee\xac\x0d\x39\xed\xe9\xe5\xd9\xd1\x59\x0f\x24\x6a\x0c\x98\x72\x71\x69\xb5\xb6\x13\xd2\x94\x72\x6d\xb3\x69\x0a\xb3\x28\x29\x47\xab\x00\xca\x43\x5f\x8c\x50\x82\x32\xc1\x82\x8d\x0e\xfe\xf3\x1b\x28\xf2\xcb\xb3\xb4\x46\x48\x09\xbc\x84\x76\xdb\x4c\x95\xf0\x25\x0c\x1c\x2e\x45\x66\xee\x06\x86\x22\x2f\x7d\x10\xfd\x06\x29\x0c\xc0\x4f\x7d\xc0\xaa\x08\x1a\x7c\xb0\x75\xab\x83\xa7\xe3\x8c\x75\x37\xa8\x0a\xdd\x07\xa5\x3c\xba\xb1\x2a\xf0\x2e\xb9\xa5\xf9\x6a\x54\xfa\xee\x55\xe9\xc9\xdd\x5c\xe2\x38\x97\xca\x8f\x72\xf1\x57\x74\x98\x2f\x38\xa7\x16\x2e\xec\x31\x00\x2c\x86\x16\x1e\xdd\x2f\x06\x6b\x7b\x04\x52\x0f\x03\x57\xff\x19\x6d\x10\x00\xbb\xb0\xfb\x08\x9e\x3e\xbd\xd9\x3a\xb9\x61\xa3\x09\xb7\x57\x32\x00\x87\x3e\x58\x87\x85\x35\xc0\x2f\x36\xcc\x37\x88\x22\x4d\x2d\x8a\xa4\xc0\xca\x9a\x5b\x28\x62\x00\x19\x31\x97\x24\x24\xb9\xac\x07\xd9\x3b\x1b\x9d\x11\x5a\x66\x3b\x34\x27\x95\x27\xda\xe2\x1a\x07\xa2\x98\x72\x87\x03\xe5\x83\x9b\x66\x3d\x08\x2e\x22\x6b\xf0\xb4\x1a\x4b\xe1\xc2\x7a\x30\x37\x0b\xdc\x3a\xbb\x52\x31\xd6\x46\x26\xdd\x1e\xc2\x78\x9b\xcc\x12\xb4\x7d\xf7\x85\x95\x98\xd2\xd7\xd3\x14\x6a\x43\x52\x5f\x6f\x44\x11\x86\x42\x6e\xc2\xd0\xe2\x54\x6f\x9f\x95\x2f\xbc\xda\xcb\x75\x34\xbb\xf0\xfe\x7d\xb3\xbb\xbb\x8e\x75\x49\xf4\x96\xc1\xe6\x40\x25\x96\x22\xea\xe0\x3f\xea\x40\x69\xdd\xdd\xc7\x99\x66\x29\x2e\x44\x09\xd2\x27\x3a\x08\x45\xbd\xf3\xe3\xf7\xdf\x7f\x9f\x08\xe1\x8b\xda\xd9\x60\xf7\xb7\xae\xa5\x0f\x5f\x7d\xb5\xb3\x3d\x63\x5f\xd4\xd6\x85\x66\xa0\xd3\xd9\xde\x99\xb1\x2f\x6e\x6a\x8f\x83\x54\x1b\x9d\x5c\x1c\xff\x7e\x70\x7a\xfa\xf6\xe0\xf4\xf4\xec\x77\xca\x4a\x5b\x49\x09\xf0\x8a\x4e\x27\x20\x70\xde\xfc\xff\xe2\xf8\x77\x1a\x9c\x4f\x73\x49\xaa\x61\x2b\xfd\xe5\xef\xe0\xe0\xf0\xf0\xf8\xfc\x25\xf0\x49\x9b\xac\xe7\x76\xb8\x17\x63\x6c\xc1\xe7\xa7\xbe\x49\x5f\xf9\x7c\x96\x32\xcb\x24\xa5\x7e\x42\x02\x25\x13\x43\xa7\x3a\x11\x62\x80\x26\xa4\xea\xd0\x60\x98\x58\x37\x82\x18\x94\x56\x41\xa1\x87\x81\x4d\x14\x13\x2c\x38\x51\xa4\x34\x2b\x15\x65\x9e\x2e\x95\x56\xe5\x62\xb1\x8b\xc6\x43\x1f\x4b\xeb\x10\xa4\xf1\x94\x8e\x46\xc6\x4e\x0c\x04\x9b\x12\x58\x63\x09\x01\x8d\x84\x58\xc3\x44\x85\x21\x10\x2d\x4d\xc1\xa7\x0c\xc9\x26\x43\xa5\x31\x31\xd6\x82\x35\x80\xcb\xc7\xb0\xbf\x0f\x59\x96\x58\x4b\xda\x1b\xce\xfa\x08\x8e\x22\x20\xd3\x1e\xd7\xb1\x7c\xd9\x48\xc1\x6c\x76\x3f\xe1\xdf\x7f\x23\x6e\xb4\x3c\x8c\xd3\x3f\xda\xca\xa7\x52\xfb\xff\xed\xde\xc5\xed\x54\x76\xbf\x38\x7b\x79\xdc\x83\x13\x03\x65\x0c\xd1\xe1\x0e\x54\x76\x8c\x4d\x33\xa0\x4c\x69\x5d\xd5\xd2\x78\x0c\x5e\x49\x04\x5b\x02\x9a\xb1\x72\xd6\x54\x74\xdc\x63\xe1\x54\x83\xa9\x0e\xf3\x18\xe0\xdb\x2b\x86\x57\x09\x9d\x97\x07\x97\xaf\x2e\x4e\xf6\x1f\x2d\x6d\xe5\xb7\x14\x89\x76\x27\xcd\x3c\xcc\x66\x8f\xd2\x42\x7e\x35\x4f\x3c\x2e\x1a\xe0\xbc\x76\x6a\xac\x34\x0e\x50\x02\xe7\x54\x25\xf0\x39\x24\x09\x15\xc0\xc7\x90\xf7\x72\xfa\xd8\xfb\x0b\x38\xb6\xd6\xee\x8f\x5b\x7b\x02\x2c\x1a\x32\xd8\xac\x60\xac\xa9\x9e\x78\x21\x78\x70\xd1\x07\xba\x1b\x1e\x43\xba\x15\xb1\x86\x01\x1a\x1c\x8b\x74\x9a\x34\xe2\x83\x28\x46\x20\x3c\x78\x4b\x9c\xeb\x13\xa4\x57\xeb\x1d\xe5\x41\x0b\x25\x29\x60\xd0\x9f\xb2\x4e\x12\x69\x4d\xdf\x14\x27\x3b\xcd\x4a\x6d\x3d\x3a\x08\x43\x95\x2e\x4a\x7b\x45\xee\x10\xae\xac\x43\xd6\x21\x57\x3c\x94\xce\x56\x2b\xb2\xb5\xb3\x05\x7a\x4f\x37\x6b\xa2\xa8\xec\x19\xaa\x9a\xf4\x35\xfe\xb3\xc6\x0d\x8f\xe0\x87\x4d\x83\x17\xa9\x30\x2b\x10\x04\x48\x31\x05\x6b\xf4\x94\x76\x53\x27\x67\x90\xa0\xe8\x59\x1e\xbd\xcb\xb5\x2d\x84\x4e\x0d\xa5\xf8\xcb\x63\x21\xdb\xcd\x52\xf5\xd2\x17\x1e\xb5\x32\x74\x3b\xe1\x7c\xef\xe8\x83\xf2\xde\x96\x61\x22\xdc\x47\xcb\x17\x5a\x54\x62\x3c\x97\x66\x1d\x40\x43\x48\x4b\xe9\xa9\xa1\xb0\xd5\x53\x69\xa9\xce\xb3\x1b\xa6\x8b\xa6\x12\x7e\x04\x95\xf4\x72\xce\x84\xd0\xd8\x59\xfd\x5a\x59\x73\x33\x52\xea\x88\x26\x2c\xbe\x2f\xa9\x6b\x1d\xf8\x5c\xea\x9a\x4d\x3c\x4c\x1b\xeb\xc0\xb9\x32\x30\x8a\x7d\x6c\x22\x97\x50\x14\x3d\x42\x8a\x2c\x88\x5a\x71\x92\x45\xc7\x3c\x5d\x25\x05\xdc\x21\x64\xbe\xf3\x0d\x6c\x37\xe3\x3d\x78\xdc\xdd\xee\xfc\xb1\x37\x0c\xa1\xf6\xbd\x3c\x5f\x2a\xd6\x3b\x59\xc3\xdf\x6d\x79\xda\x24\xb2\x7c\xd1\xd2\xf3\x66\xa0\x7b\x63\xfc\xe1\x36\x88\x1e\xd2\x9f\xcf\xaf\xd5\x4b\xf3\x70\xa5\xa9\x32\x4f\x6a\xda\x9e\x80\xb1\xeb\x6b\x4e\x89\xd8\x20\x6c\x75\x9f\x89\x62\x14\xeb\x67\xda\xf6\x5f\x10\x6f\x65\xd9\x07\x1f\x04\x16\x14\x4c\x95\xc7\x18\xdd\x74\xad\x61\xa2\x84\x14\x88\xed\x60\x80\x21\x5d\xcf\x7e\xb2\x92\x7a\xa7\x8b\x72\xb5\x52\xc9\xb7\x19\x51\x01\xf9\x71\xa4\xdc\xfe\xea\x5c\xbb\xae\x69\x85\xb6\x96\xe4\xfe\x36\x83\x1e\x87\x42\x36\x7b\x7e\x20\x89\xae\x28\xfa\x27\x79\x74\xd5\xd0\xe7\xa3\xd2\x7b\xfd\x94\x76\x62\xb4\x15\x92\x82\xd8\x1c\x42\xb6\xca\x76\xeb\x04\xf7\x07\x83\x44\x72\x6b\xf7\xaf\xb7\x3e\xb4\x49\x38\x3d\xac\xd5\xce\x8e\x95\x44\x97\xf7\xf2\xb7\x52\x04\x91\xbf\x25\x56\x6a\xa5\x97\x01\xd0\xcb\x6d\x24\x26\xa5\xa9\x0f\xc5\x8c\xa0\xd4\x6c\xa2\xd1\xc4\xfb\x2d\xdc\xf7\x69\xe5\xad\x1b\x30\x9b\xb5\x42\x32\xbd\x3f\x26\x8a\xdc\x27\x63\x2d\x18\xbb\xb2\xdf\x0a\x88\x22\xcd\xcd\x43\x75\x7f\x40\x5b\xfb\x73\x61\x3a\xc2\xf9\x35\x79\x32\x2f\xe4\xff\x2e\xa6\x9b\x2a\x85\xf6\xfc\x40\x4c\xaf\x28\xfa\x27\x31\xbd\x6a\xe8\x7f\x84\xe9\x26\xca\x89\x7e\x8d\xa8\xfd\xd0\x86\x4f\xc2\x34\xa1\xa8\xb7\xf8\xb4\x98\x5a\xce\x57\xbd\xd5\x6f\x0d\x3a\x39\xc2\xf1\xcb\xc3\xa3\xc3\x97\xa7\x6f\x0f\xce\x4f\xf6\xb3\xef\xb2\x3b\x40\xbb\x1a\x14\x92\x21\x2d\x89\x77\x5b\x7f\xe7\x40\x59\xb9\x09\x6b\xb8\xa4\x7b\xc3\x29\x61\xae\xe6\x52\x83\x93\x56\x20\x75\x28\x4b\x19\xbb\x1d\x56\x46\x05\x25\x34\x2f\x74\x4c\x77\x34\x6b\x63\xb8\x9b\xfe\xed\xcf\xf9\x65\x65\xb4\xf7\xe4\xbb\x1f\x77\x77\x96\x87\xf6\x36\x0a\xee\xad\x0b\x3e\xd9\x28\xf8\x24\x09\x66\x9b\x5d\xe2\xc1\x8e\xd0\xa4\xb0\xf0\xd2\x3a\x9e\x5a\xeb\x5b\xa2\x42\x8e\xd1\x05\xe5\x91\xd7\x88\x8e\x47\xa7\x3d\x6c\xa0\xc6\x64\x86\xb1\x6a\xbc\x1e\xa5\x7c\xfb\xd6\xd8\xda\x9b\xdf\x22\x9e\x2b\x94\xb4\xd2\x8e\xdf\xd2\xfb\x31\xc8\xc4\xd4\x1b\x66\x89\x9e\xa9\xd9\x9c\xcd\x18\x0b\xd1\xa0\xe4\x42\x56\x54\x2f\x97\xd4\x67\xde\x14\x33\x54\x6d\x3b\xab\x79\xad\x45\x6a\x8d\xa8\x94\x16\xda\x5b\x30\x88\xf2\x46\xae\x9b\xea\xaa\xee\xd8\xea\x58\xa1\x07\x02\x46\xf3\xf0\x28\xe7\x5d\xef\x55\xe9\xa1\x79\x4d\x2a\xa8\xd7\xa5\x86\x78\xfe\xfa\x58\xc1\xee\xff\xff\xb0\xbb\xe9\x15\xf2\x0e\xfd\xe4\x47\xf3\x00\x94\x4a\x04\x3f\xf5\xda\x0e\xc0\x2b\x2a\xdd\x27\x08\x95\x30\x62\x80\x80\x54\x37\x84\x21\x89\x84\xa1\xb3\x71\x30\x84\xf9\x1b\xd2\x52\xb9\xd9\x3e\x24\xcd\xb5\x6c\x2c\x48\x6d\xbd\x36\xcd\x3a\x60\x6c\xc0\x1e\x88\x60\x2b\x55\xf0\x9b\x88\xa5\x56\xbe\x70\xc2\x0f\x41\x5b\x5b\x7b\x88\x26\x28\x3d\xff\xb9\x48\x79\x88\xf5\x7a\xf1\xbc\x51\xcb\xc2\xd8\xe7\xf8\x89\xc5\x17\x43\x94\x31\x05\x6c\xf9\x56\x3a\xec\x5b\x1b\xa8\x3a\x2e\x6c\x55\xa7\x07\xd5\x4d\xaf\xe8\x19\xf3\xc3\x18\x88\x58\x28\x85\x35\x6b\xbe\x7d\xc2\xae\xaf\x29\x45\xce\x66\x6b\xe5\xfb\xbd\xfb\x59\xbc\x63\x2d\x3d\x53\x13\x11\x3a\x3b\x85\xa9\x8d\xce\xa3\x2e\xd7\xae\x43\xfb\x92\x92\x17\xd1\x07\x5b\x71\x5f\x38\x55\x87\x7c\x4e\x75\xf9\x76\xde\x8c\x74\xfd\xf0\xbf\x01\x00\x00\xff\xff\xf4\xcd\xb8\x76\xeb\x1b\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\xdd\x6e\xdb\x48\xd2\xbd\xef\xa7\x38\x23\x0d\xe2\x6f\xbe\x09\xc9\x64\x81\xec\x02\xce\x24\x40\x36\xc9\x00\xd9\xc1\x8e\x0d\x7b\x76\xf7\x22\xc8\x45\x8b\x5d\x24\x7b\xd4\xec\x62\xfa\x47\x32\xa3\xe8\xdd\x17\x4d\x52\xb1\x6c\xd9\x72\x82\xec\x9d\xc4\xae\xae\xdf\x53\x75\xaa\xe7\x3f\x14\x0b\x6d\x8b\x85\xf4\x0d\x32\xba\x12\x62\xb3\x81\xae\x90\xbf\x66\x5b\xe9\x3a\xbf\xa4\x32\x3a\x1d\xfa\x73\x19\xca\xe6\x5c\x96\x4b\x59\x93\xc7\x76\x2b\x0c\xd7\x35\x39\x64\x01\x96\x15\x65\x3e\x48\x17\x62\x97\xfb\x06\x33\x6d\x7d\x90\xc6\x68\x5b\xc3\x91\x42\x23\x03\x4a\x65\x51\x0e\x1a\xa3\x93\x41\xb3\x05\x5b\xfc\xf8\x7f\x0d\xfb\x60\x65\x4b\x3f\xcd\x44\x29\x03\x5e\x16\x2b\xe9\x0a\xa3\x17\x45\x1f\xdb\xa2\x34\x9a\x6c\xc8\x4a\x72\x21\xef\xa8\xc5\x2f\xbf\x9c\xbc\x3d\xfb\xf5\x24\x39\xf8\x9a\x5c\x78\xe5\xff\xde\x07\xf2\x5f\x3c\x4d\xdf\x74\xa5\x4b\x19\xc8\xe7\x93\xa7\x17\xd4\xb1\xd7\x81\x5d\x3f\x1c\xe3\x33\x2e\x83\x4b\x7e\x6d\xb7\xe2\xed\xd9\xaf\xf7\x1b\x5d\x52\x7f\xdb\xe6\xb9\xd3\x2b\x19\xe8\x37\xea\xbf\xd1\xf2\x6f\xd4\x1f\x18\xfe\xca\xf4\xbd\xba\x38\x83\x9f\x2a\x80\xd8\xa9\x64\x01\xef\x37\x9b\x49\x9b\xff\x07\x6b\xfb\x40\xa9\x66\x8f\x31\xc3\x76\xfb\xe1\x20\xe1\x15\x3b\xc8\x10\xa8\xed\x02\xb4\xc5\xe6\x69\x9e\x3f\xdb\x3e\x87\x62\x01\xf4\xb1\xc5\xe4\x06\xb2\x1e\xd9\x47\x7c\x9b\xcd\xc1\x24\x1e\x3d\xc2\xc2\x91\x5c\x0a\xe0\x48\xb8\xef\x77\x4e\xfc\xb8\x99\x7e\x6d\x3f\xdc\x1d\xf8\xe4\xd1\x88\x9f\x4a\x6a\x43\x6a\x26\x90\xd0\xfa\xfe\xfd\xde\x6d\x64\x26\xe0\x19\x3e\x7c\x78\x8e\xd0\x90\x85\x37\x44\x1d\x9e\x3e\x07\x19\x4f\xa0\x2b\x1d\xd2\x9f\x4a\x0b\xc5\x96\x8e\x56\xc2\x51\xcb\xab\x6f\x83\x71\xca\x5c\x69\x48\x5a\x48\x63\x84\x6b\x91\xb9\x0a\x47\x61\x7d\x04\x7e\x62\xb3\x21\xab\xb6\x5b\x21\xe6\x28\x1d\xc9\x40\xc9\x7a\x90\xda\x92\x43\x17\x8d\x49\x39\x72\x14\x20\xda\xa5\xd2\x0e\x59\x77\xad\x8d\x9d\xae\xb5\x2d\x72\xc5\xe5\x92\xdc\x2d\xa4\xdf\x3c\x2c\xc6\x90\xf2\x3f\x3d\xdb\x7d\xc4\xe7\x6f\xc8\xe9\x15\xa9\xfc\x35\xb7\x0b\x6d\x49\xbd\x6b\x65\x4d\xe7\xd1\x98\xcb\xd1\xec\x84\x82\x03\x74\x1b\x8b\xcc\xdf\xe7\x0a\x0a\xc7\x1c\x8a\x14\xd3\x1f\x67\x6f\xce\x4e\xa1\xc8\x50\xa0\x54\x2b\x54\x6c\x0c\xaf\x93\xa6\xda\x71\xec\xc6\xa0\x53\x9a\x65\x15\xc8\x41\x07\x68\x8f\x85\x5c\x92\x82\xb6\x81\xc1\xd1\xe1\xdf\xff\x84\x4e\x7e\x79\x31\xdc\x91\x4a\x21\xab\x30\x85\x2d\x74\x85\x1f\x50\x3b\xda\xcb\xcc\xce\x0d\x0a\x65\x51\xf9\x20\x17\x23\x4e\x04\xe0\x7b\x1f\xa8\x2d\x83\x81\x0f\xdc\x4d\x3a\xb2\xa1\x9c\xb1\xcb\x83\x6e\xc9\x3d\x28\xe5\xc9\xad\x74\x49\xf7\xc9\xed\x9d\xb7\xcb\xca\xe7\x57\x95\x4f\xee\x16\x8a\x56\x85\xd2\x7e\x59\xc8\x4f\xd1\x51\xe1\xc8\x73\x74\x25\x65\x9d\x74\xe1\xa9\x00\xa8\x6c\x18\x27\xc7\xc5\x70\x10\x23\x92\x7a\xd4\xae\xfb\x18\x39\x48\xe0\x09\x9e\x9c\xe0\xe5\xcb\xeb\xd0\x93\x1b\x1c\x6d\xb8\x7d\x53\x00\x8e\x7c\x60\x47\x25\x5b\x64\x17\x07\xe7\x9b\x4d\x96\x1a\x8f\x3e\x22\xbf\x60\x43\x69\x6a\x55\x4e\xa6\xb6\x17\xc0\x08\xb6\x64\x64\x02\x98\x92\xd4\xb2\xbd\x05\x30\x01\xcc\x0c\xd7\x99\x4a\x20\x73\xb3\x53\xcc\xfe\xe4\xe8\xac\x34\x6a\xf6\x38\x9d\x29\xed\xe5\xc2\x50\x66\xa8\x96\x65\x9f\x39\xaa\xb5\x0f\xae\x9f\x9d\x22\xb8\x48\x62\x84\x5a\xf2\x83\xac\x1a\xed\xee\x67\x5c\xba\x70\x98\xf2\xbb\x05\x6e\x55\xb8\xd2\x42\x4c\xf9\x1b\x7a\x2c\x75\xc2\x34\xee\x86\x06\xf0\xf9\xef\xac\x68\x18\x70\x2f\x87\x82\xd8\x24\xf5\x28\x01\x7a\x3d\xcc\x9b\xa4\x3a\x61\x38\xcd\x13\xac\xa5\xac\xc9\x06\x48\xab\x60\x29\xac\xd9\x2d\x11\x83\x36\x3a\x68\xf2\xa8\x79\x98\x6b\x81\xe1\x64\x39\x74\xb7\xd2\x09\xf0\xb9\x98\xa7\xf4\xee\x2e\xbb\x68\x3d\x16\x54\xb1\x23\x28\xeb\x53\x17\x2c\x2d\xaf\x2d\x02\x0f\x7d\x33\x5a\xa2\x21\x13\xb1\xc3\x5a\x87\x06\x69\x16\xf6\xf0\x43\x63\x8a\x75\xa3\x0d\x0d\x63\xf2\xcb\xb0\x42\xa6\x7e\xc2\x8b\x17\x98\xcd\x86\x51\xa9\xf8\x7a\x50\x3e\x38\x18\x53\x5e\x52\x84\x87\xa9\xb9\x1c\xa5\xb0\xdd\x1e\x67\x98\xe3\x09\xbe\xd6\xf2\x3d\x24\xf2\xd5\x36\xbe\x95\x4b\xfe\xfa\xe4\x3e\x32\x99\x8b\x39\x7e\x3f\xfb\xe3\xed\x29\xde\x59\x54\x31\x44\x47\x8f\xd1\xf2\x2a\x8d\x37\x99\xb2\x50\xb1\x6b\x27\xe6\x88\xc1\x6b\x45\xe0\x0a\x64\x57\xda\xb1\x6d\x53\xa9\x57\xd2\xe9\x04\x7b\x2f\xe6\xc2\x53\xc0\xcf\x57\x82\xae\x3a\x76\x01\x97\xaf\x2e\xff\x75\xf1\xee\xc5\xc9\x5e\x28\xff\x61\xb7\x24\x37\x45\x32\x9e\x63\xbb\x3d\x19\x2e\x66\x57\x3b\x14\xbb\x68\x91\x65\x9d\xd3\x2b\x6d\xa8\x26\x85\x2c\x4b\xc4\x94\xed\xe0\x98\x10\x81\x6c\x85\xe2\xb4\x48\x3f\x4f\x3f\x21\xa3\xc9\xda\xf1\xbc\x4d\x15\x10\xd1\x26\x83\xe3\x0d\x21\x46\xba\xce\x4a\x99\x05\x17\x7d\x48\x7d\xe1\x29\x0c\x1d\x11\x3b\xd4\x64\x69\x25\x87\x5a\xa6\x2f\x3e\xc8\x72\x09\xe9\xe1\x39\x8d\x79\x3f\xc0\xf9\x26\xc5\x6a\x0f\x23\xb5\x4a\x09\xc3\xa2\x17\xf3\x41\x64\x32\x7d\xcd\x87\x8f\xc7\x9b\x86\x3d\x39\x84\x46\x0f\x4d\x32\xb5\xc7\x3d\xc2\x2d\x3b\x12\xf3\xe4\x8a\x47\xe5\xb8\xbd\x21\xdb\x39\x2e\xc9\xfb\xd4\x55\x6b\x9d\x98\xb6\xd1\x5d\xd2\x37\xfa\x2f\x46\x37\x3c\xc1\x37\x1c\x8d\x1a\x72\xcc\xb6\x24\x48\x28\xd9\x83\xad\xe9\x53\x34\xdd\xe0\x0c\x25\x28\x7a\x51\x44\xef\x0a\xc3\xa5\x34\xc3\xca\x2d\x3f\x79\x2a\xd5\x14\x6c\x22\xcc\x85\xf4\x64\xb4\x4d\x9d\x89\xf3\xa7\x6f\x1e\x94\xf7\x5c\x85\xb5\x74\x5f\x2d\x5f\x1a\xd9\xca\xd5\x4e\x5a\xcc\x41\x36\x21\x6d\x18\x4d\xe3\x3c\xbc\x59\x95\x69\x6e\x7a\x71\x3d\x36\xa3\x6d\xa5\x5f\xa2\x55\x5e\xed\xc6\x2a\x46\x3b\x37\xff\xb6\x6c\xaf\xbf\x54\x26\x92\x0d\x5f\xfe\xef\xa9\x9b\x1c\xf8\x5f\xa9\x1b\x83\xf8\x3e\x6d\xe2\x18\xb5\x85\x68\x49\x65\x52\xb5\x09\x1e\x55\x1a\xa9\xdc\x91\xf5\x8d\xae\x42\x96\xc0\xe5\xd8\x64\x9d\x91\x96\x46\x5e\x4a\x23\xe2\x81\x5b\x69\x92\xed\x93\x58\x22\x11\x82\x34\x9e\x61\x89\xd4\xb5\x64\x3e\x14\x36\x5f\xb1\x89\x2d\x79\xa4\x75\x6f\x5c\x09\xd5\x8e\x18\x12\xd7\x8f\x3c\x5f\x26\x3a\x48\x9c\xb1\xdb\x0b\x5b\x3c\xf9\xdb\xb3\x27\x77\xed\x87\xf7\xe8\x4f\x7e\x8c\xfc\x3b\xac\xbf\xbe\xf7\x86\x6b\x78\x9d\x10\xbe\x26\xb4\xd2\xca\x9a\x40\x2b\x72\x7d\x68\x92\x48\x68\x1c\xc7\xba\xc1\x8e\xc2\xf7\xaa\x32\xf1\xf8\x4e\xcb\x9d\x75\xe3\xee\xe0\x58\xcc\x61\x39\xd0\x29\x64\xe0\x56\x97\xd9\xcd\x9c\xa1\x74\xe9\xc5\x6a\x98\x3b\x8f\x68\x83\x36\x68\xa5\x1f\x96\x44\x8f\xd8\x1d\x62\xec\x4e\x2d\x5f\x8c\x7d\xff\xa3\xd7\x97\x0d\xa9\x38\xa4\x6b\xef\x41\x00\x47\x0b\xe6\x90\x06\x47\xc9\x6d\x37\x2c\xba\x77\xbd\x6d\x66\xc2\x37\x31\xa8\x44\xe9\x59\x36\xdd\xf9\xf9\x2f\x69\xff\x37\x9e\xb6\xdb\x03\x8c\x1f\x8d\x06\x9f\x3f\x8f\x2b\xd2\xde\xf3\x41\x91\x0f\x8e\x7b\xf4\x1c\x9d\x27\x53\x1d\x3c\x4d\xa6\x55\xa3\x28\xa3\x0f\xdc\x66\xbe\x74\xba\x0b\x45\xf2\xc8\xb0\x54\xc5\xff\x17\xe3\x97\xdc\x37\xff\x0d\x00\x00\xff\xff\xad\x36\x79\x0c\x32\x10\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
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
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
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
