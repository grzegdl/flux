// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 6865,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\x5b\x6f\x1b\x37\x16\x7e\xf7\xaf\x38\x50\x16\x48\x0c\x48\x23\xbb\x6e\x8b\xc5\x74\x5d\x6c\x9a\x8b\xeb\x4d\xe3\x18\x71\xbc\x8b\x3c\xd5\x14\xe7\x48\x43\x88\x43\xce\xf2\x70\xa4\x0a\x46\xff\xfb\xe2\x90\x73\xe1\xc8\xb2\x5d\xe4\x6d\xfb\x50\xdb\x9c\x73\xbf\x7e\x64\x66\xb3\xd9\x91\xa8\xd5\xbf\xd1\x91\xb2\x26\x07\x51\xd7\x34\xdf\x9c\x1e\xad\x95\x29\x72\x78\x8b\xb5\xb6\xbb\x0a\x8d\x3f\xaa\xd0\x8b\x42\x78\x91\x1f\x01\x18\x51\x61\x0e\x4b\xdd\xfc\x71\x7f\x0f\x6a\x09\xd9\x95\xa8\x90\x6a\x21\x11\xfe\xfc\xb3\xfd\x1e\xfe\xcc\xe1\xfe\x7e\xfc\xf5\xfe\x1e\xd0\x14\x4c\x46\x35\x4a\x16\xe6\xb0\xd6\x4a\x0a\xca\xe1\xf4\x08\x80\x50\xa3\xf4\xd6\xf1\x17\x80\x4a\x78\x59\xfe\x26\x16\xa8\x29\x1e\xa4\xba\x99\xda\x3b\xe1\x71\xb5\x8b\x1f\xfd\xae\xc6\x1c\x3e\xa3\x74\x28\x3c\x1e\x01\x78\xac\x6a\x2d\x3c\xb6\xc2\x12\x0f\xf8\x3f\x61\x8c\xf5\xc2\x2b\x6b\x7a\xe1\x00\xb5\xb3\x15\xfa\x12\x1b\xca\x94\x9d\xd7\xd6\xf9\x1c\x26\x67\x27\x67\xa7\x13\x78\x01\x1e\xb5\x4e\x28\xc0\x5b\x20\xe9\x44\x8d\x30\xaf\xd0\x3b\x25\x89\x9d\xab\xad\x32\xfe\x25\x01\x33\x67\xad\x60\x3d\xf2\x61\xcf\x0b\x80\x2e\x16\xe1\x77\x74\x1b\x25\xf1\xb5\x94\xb6\x31\xfe\x6a\x4c\x08\xb0\xb1\xba\xa9\xb0\x17\x35\x6b\x45\xad\x94\x9f\xad\x71\xd7\x2b\x20\x8e\x82\x1f\x14\x76\x27\x83\xbc\x19\xb3\x14\x21\xc1\x09\x55\x81\x4b\xd1\x68\xff\xd1\x16\x98\xc3\xc9\xf7\x27\x27\xf0\x02\xb6\x25\x1a\xa8\xd8\x1a\x2c\xc0\xa1\x28\x66\xd6\xe8\xdd\x14\xb6\x08\x5b\x6b\x5e\x7a\x58\x20\x88\x85\x46\x8e\x87\x2c\x2b\x5b\x1c\xb5\x02\x5f\xc0\x97\x52\x11\x28\x02\x01\xbe\xaa\x97\x04\x0d\x61\x01\x4b\xeb\x60\x85\x06\x9d\xf0\xca\xac\xe0\xe6\xe6\x57\x58\xe3\x8e\x32\xb8\x34\xf0\xe1\xef\x04\x3f\x9f\xc3\x69\x76\x7a\x32\xed\xa5\x74\xba\xa3\x0b\x04\xc2\x61\x6a\x07\x59\x36\xc5\x20\x16\x20\x80\xb0\x16\x5c\x14\x6d\xa0\x60\x8b\xbd\x18\x29\x0c\x6c\x9d\xf2\x6c\x68\x76\x38\x7e\x2b\x34\x7d\x30\xb0\xaa\xfd\xee\xad\x72\x69\x10\x2b\x2c\x54\x53\xe5\xf0\x11\x2b\xeb\x76\xa9\x9f\x08\x4b\xab\xb5\xdd\xb2\x47\xad\x6a\x45\xc1\xd5\x86\xf8\x4c\x80\x6c\xc8\xdb\x4a\x71\x04\xd6\xc6\x6e\xcd\xef\xa5\x25\x4f\xbd\x88\xa5\xd2\x38\x85\x6d\xa9\x64\x09\x3b\xdb\xc0\x56\x69\x1d\x9d\xf2\x16\x0a\xcb\x7d\xc6\xc7\xcc\xc4\xbf\x38\xb0\x5b\xc3\x66\xf7\x02\x1c\xd6\x16\x9c\xf0\x25\x3a\xf0\xa5\x30\xad\xe2\x95\xf2\x65\xb3\x00\xcb\x87\x08\x5a\xad\x31\x83\xaf\xb6\x79\xa9\x35\x08\x4d\xb6\x53\x31\x0e\x36\x28\x0f\xca\x78\x1b\x78\xa4\x35\x5e\x28\x83\x6e\x0a\x0b\xd4\x76\x9b\xc1\x0d\x0e\x51\x2d\xbd\xaf\x29\x9f\xcf\x0b\x2b\x29\xe3\xc2\x92\x05\xb7\x0e\x9a\x39\xb7\x1e\xf9\xf9\xaa\x51\x05\xd2\xbc\x21\x9c\xd5\x4e\x6d\x84\xc7\x50\x7a\xec\x48\x56\xfa\x4a\xf7\x92\xba\x5c\x10\x95\x33\x69\xcd\x52\xad\xfa\x4f\x00\xf1\xe0\xa3\xa8\xf3\xe4\x30\x6d\xa4\x59\xc2\xf6\xad\x79\xc9\xd6\xcd\x02\xe7\x51\xc8\x50\x7e\xcf\xe6\x64\xab\xa8\xe4\x93\x52\x6c\x10\x04\x14\x6a\xb9\x44\xc7\x43\xb3\x93\xd0\x76\xd5\x30\x18\x43\x0a\xa2\xb8\x34\x09\x3c\x5c\x36\xaa\xc0\x2e\xec\x4b\xb5\xaa\x44\x3d\x18\xa2\x7c\x09\xc2\x00\x1a\xef\x76\xc1\x87\xbb\x48\x74\x37\x05\x61\x0a\x68\x8c\xb4\x15\x4f\xeb\xc0\x1f\xbd\xfd\x18\xd2\x29\x4c\xd1\x4b\x41\xb3\x09\x12\x14\x52\x9b\xcf\x07\x19\xe0\x30\x7c\x43\x06\x12\xb6\x67\x33\x10\x26\x81\xb7\xa0\x2a\x9e\x93\x70\x71\x7d\x11\x86\x00\xbc\x62\xb7\x48\xad\x8c\x32\x83\x72\x76\x6e\x83\x4e\x2d\x95\x0c\x03\x1b\xea\xc6\xd5\x96\x90\x8e\xff\x42\x20\x7b\x29\x71\x7c\xc4\x28\x72\x80\x58\xdf\x5f\x08\x1c\x08\xb7\x1a\xda\xf4\x91\x88\xad\xea\x15\xcf\x0f\x4a\x42\x33\x1e\xc1\x2f\x1e\x19\xc2\x0f\xf9\x0e\x0c\xe1\x2e\x9c\x7d\x27\x3e\x98\xff\xc9\x86\x68\xa3\xee\x30\xcc\x49\x63\x61\x92\xc7\x4e\x9c\x80\xaa\xc4\x0a\x63\xf5\x33\x43\x06\xef\x95\x29\x82\xcf\x15\x8f\x15\x87\x72\xa8\xda\x38\x52\x34\x0a\x42\x1e\x1e\x81\x95\x93\xc0\x38\x01\x84\xef\xfb\xbe\x6c\x16\x59\x61\xe5\x1a\x5d\x26\x6d\x35\x77\xf3\x38\x03\xc2\x8f\xb9\x17\x7d\xe8\xba\x3c\xf2\xbe\x67\x2c\xc0\x5a\xbd\x58\x01\x5b\x9a\xf5\x34\x41\x4d\x84\x0d\x97\xfc\xeb\x27\xf7\xb6\xed\x9d\x80\x2c\x12\xa2\xeb\x46\xeb\x6b\xab\x95\xdc\xe5\x70\xb9\xbc\xb2\xfe\xda\x21\xa5\xe6\x3b\x24\xdb\x38\x89\x94\x0e\x70\x87\xff\x6d\x90\xfc\xe8\x0c\x40\xd6\x4d\x0e\x3f\x9c\x54\xa3\xc3\x2a\xcc\xf8\x1c\x7e\xfc\xfe\xa3\x1a\xf0\x81\x75\x29\xf3\x6c\x48\xc9\x75\xc0\x0a\x67\x27\x67\xbc\x32\x95\x59\x5a\x57\x85\x5a\x15\xba\xa7\xd6\x6a\x83\x06\x89\xae\x9d\x5d\x60\x6a\x01\xc7\xf2\x62\xbc\xae\xa3\xaa\x28\x70\x7c\x2c\x7c\x99\xc3\x5c\xd4\x2a\x86\x78\xf3\xe3\x5c\x15\x68\xbc\xf2\xbb\xac\x6e\x16\x09\xad\x32\xca\x2b\xa1\xdf\xa2\x16\xbb\x1b\x6e\xcc\x82\x72\xf8\x21\x21\xf0\xaa\x42\xdb\xf8\x03\xdf\x78\xbb\xaa\xff\x0f\x53\x93\x6e\x1d\x25\xe6\x30\x2e\x82\xb8\xdf\xae\xa3\x65\xe8\x65\xb0\xac\x98\x13\x95\x0c\xf0\x6c\x84\x9c\xa0\x6d\x3b\x68\x56\x9c\x32\x50\x26\xd6\xdc\x4b\x8a\x3c\x44\xe5\x7c\x34\x1f\xbb\x98\x7d\x32\x7a\x97\x83\x77\x0d\xb2\x34\x06\x3f\x61\x34\x2d\xda\x89\xce\xbd\x54\xa3\x5b\x5a\x27\x91\x85\x46\xb4\xc3\x60\xe7\x31\xc3\x53\x40\x32\xb6\x7d\x23\x5c\x6b\x7b\x24\xfb\x36\xf3\x93\xe6\xbc\x34\x52\x37\x61\x64\x32\x66\x8b\x9b\xad\x1b\xa7\x11\x14\x3c\x83\x61\x3a\x14\xf3\x13\xb3\xee\xe1\x8b\x7e\xac\x42\x81\x52\x0b\xc7\x58\x6d\x61\x37\x49\xe7\x3f\xb1\xff\xe3\x5c\x4c\x9d\x77\xd6\xfa\x79\x46\x54\x3e\xea\x80\x30\x23\xad\x93\x61\x37\x4d\xa2\xe6\x69\x47\x92\x48\x40\xb3\x51\xce\x9a\xb0\x09\xe2\x92\x9d\x7c\xb8\xfd\xe5\xdd\x9b\x4f\x57\xef\x2f\x2f\x26\x71\xf6\x4f\x39\x1e\x76\x83\xce\x8d\x17\x75\x22\x26\xec\xb6\xc5\x2e\xae\x51\xaf\x0f\xf9\xf8\x60\xc3\x3e\xf4\x71\x28\x4e\x26\x7e\xd4\x51\x5e\x76\x7c\xe3\xe8\xb4\xf1\x6c\x4e\x30\x48\x6b\x5d\xc8\x49\x22\x62\x1f\xc9\xa4\x49\x0f\x30\xa6\xc3\xdc\xc2\x80\xd0\x1e\x9d\x61\x4c\xfd\xc0\xe2\xa5\xb3\x15\x97\x45\x07\x55\xa6\x20\x88\xcb\xad\x5d\xa7\x1c\x06\x6d\xe5\x9a\x1e\x26\x1b\xcd\x26\x3f\x10\x97\x21\xdc\xa3\xb8\x6c\x84\x6e\xf0\x41\x4c\x9e\x2b\xe2\xfd\x1a\xe8\x96\xed\x13\x15\xc0\xbb\x7e\xbc\xe3\x9f\xd8\xf2\x8f\xd4\x25\x53\x45\x58\x33\xa2\x1b\xcf\x87\xe7\x3a\x6f\x2b\x18\x8d\x58\xa0\xa6\xae\xf5\x0e\x7e\xfd\xf2\xe5\x1a\x16\x82\x94\x04\xd1\xf8\x12\xa4\xc3\x30\x49\x85\x8e\xeb\x7c\xb8\x08\xb0\xc0\x8d\x12\xc1\xf1\xbb\x8b\xcb\x2f\xbf\xbf\xbe\xfd\xf2\xeb\xed\xcd\xbb\xcf\x77\xc1\xdd\xfe\xe8\xc3\xbb\xaf\x77\xa3\x82\xdf\x08\xa7\xf8\x1a\x47\x1d\x32\x4e\x04\x46\xdc\xb2\x97\xbf\xf7\xce\x56\xe3\x1c\x46\xb2\xcf\xb8\xcc\x47\x9e\x8f\x40\x22\x0f\x36\x76\x61\x08\x00\xc7\x3c\x1f\xc5\x23\x86\x20\x5e\x4e\xb1\xe0\x4d\x2c\x85\x2c\xb1\xe0\xd2\x4a\x6b\xbb\xc7\xd3\x1c\x29\x96\x3e\x4d\xa4\x58\xd7\x02\xe6\x84\xa1\xbd\x5c\x07\xc6\x69\x50\xc2\x97\xc2\x36\xc6\xbe\x44\x4a\x6b\x61\x80\xad\x7e\x6b\xd9\xca\x86\xe3\x14\x3a\x2e\xbc\x04\x84\x42\x84\xd2\x6e\xc3\xc5\xd7\x1a\x83\x32\xa4\x4c\xf9\x71\xed\xcc\x66\xbd\x03\xe1\xd6\xc3\xca\xcf\xfb\xa3\xac\x45\x7b\x19\x6d\x64\x26\x75\x43\x1e\x5d\xc6\x03\x5c\xa7\x21\xb9\xa5\x38\x6b\x86\x50\xbc\x89\xa4\x97\xd7\x23\xa7\x78\xec\x10\xfa\x70\xb1\x1e\x57\xf6\x60\x43\x47\xcf\xd5\xe5\x1d\x53\x86\xab\x6e\xb2\x82\x52\x8b\x5b\xea\xf3\xa3\x11\xbc\x54\x04\x55\x43\xe1\xea\x1f\xa2\xa7\xb0\x88\xed\xb4\x08\x8b\x2d\x80\xbb\x70\xe3\x7f\xd5\x5d\xa3\x8f\x53\x5b\xba\xe1\x12\xdb\x90\x0b\x38\xb9\xf8\x8f\x0c\xe1\x65\x10\x17\xdc\xac\x50\xee\xfc\xc1\xda\x4b\xcd\xfa\x9c\x40\xcb\x21\x79\xb7\x9f\x7f\x8b\x2f\x13\xc2\xac\xe2\xb7\x0b\xe5\xc3\x6d\x99\x94\xb7\x6e\xd7\x8f\xeb\xf7\x0c\x89\x13\x71\x4f\xf5\x1c\x97\x4d\xe2\x7b\xdb\x32\x07\xdb\x29\xed\x85\x0e\x34\xff\xed\x55\xda\x99\xc7\xf9\xf0\xf7\x87\x77\x5f\x8f\xff\x19\xef\xec\x01\x4f\x37\x84\x6e\x3e\x18\x9b\xa5\x8d\xce\xf1\xe1\x76\x6a\x9c\x3e\x67\xc4\x7c\xa1\x3c\x3b\x9b\x20\xe5\x8e\x62\xe1\x84\x91\x65\x47\xf4\x4b\xf8\x2b\xbe\xc6\xa9\x65\x38\xe2\xf9\x45\x87\x38\x19\xc3\x31\xdf\x4d\xa8\x14\xfa\x97\x55\x26\x61\x98\x4c\x27\xed\xa3\x9e\x26\x4c\xd9\x9f\x1e\x6a\x0e\xb9\xf0\x64\xbc\x6e\x55\xc2\xa8\x25\x63\x72\xee\x21\x52\x05\xba\x98\x8e\xbd\x2b\x4d\x78\x8c\xb0\x84\xd0\x98\x02\xdd\x5e\x8e\x1d\x6a\xe1\xd5\x06\x03\xe4\xa4\xae\x02\x57\xa3\x3c\xef\xf5\x64\xef\x1c\x35\x8b\x42\xb9\xd3\x69\xfc\xf9\x5d\xff\x42\x39\x04\x27\xbc\x40\x1e\x0a\x4e\x78\xd6\xeb\xa2\xda\x51\x1d\x10\x70\x4b\xe8\x0e\xf1\x73\x72\xfb\xcc\x31\x0d\x1c\xe6\x7f\x57\x09\x75\xd0\x00\xe4\x0f\x9d\x84\x8e\x6a\x78\x63\x3d\x98\x0e\xe4\x51\xb2\xb5\x1c\x50\x34\xe1\xdd\x8e\xe3\xc4\x1b\x5b\xf9\xbd\x9b\x77\x1a\xab\x76\xf7\xb5\x9b\xed\xfc\x89\x55\xd7\x71\xb4\xb2\x98\xeb\xfc\x1f\x6b\xdc\x81\x2a\x7e\xee\xc9\x9e\x80\x33\x89\x55\x2c\x42\xf8\xc6\xe1\xe8\xfa\x7f\x40\x57\xf8\xbc\x9b\xf5\xf4\x34\x1a\x57\xdd\xb4\x06\xe5\xa1\x14\x14\x56\xb1\x35\x7a\x07\x42\x4a\xa4\x38\xd1\x4b\x8c\x2f\x68\xaf\xba\xc7\x9a\xbb\xa5\xd0\x84\x77\xc7\x07\xb4\x75\xfc\xe3\x00\x93\x77\x8d\xf4\x51\xd1\x36\x5c\xc0\x19\x9b\x35\x1e\x68\x67\x24\x2c\xac\x5d\xaf\x11\x6b\x2e\xd7\x5e\xc7\x64\xa5\xfc\x64\x0a\x15\x0a\x8e\x14\x4f\x22\x10\xe1\x56\xdc\x56\x70\x53\x93\x77\x28\xaa\xbe\x94\xf7\xad\x61\xd1\x33\xf2\xc2\xe3\x39\x4f\x86\x47\x13\x6e\xf0\x0f\xdf\x65\x3d\x59\x55\xc2\xc0\xa4\xd3\x31\xe9\x16\x49\x22\xe4\x15\x66\xab\x6c\x0a\xff\x41\x86\x84\x6f\xb4\x6d\x8a\xe3\x2c\x3c\xe9\x78\xbb\xe6\x8b\x05\x41\x2d\x9c\x57\xb2\xd1\xc2\x75\x51\x6c\xa5\xec\xef\xc0\x56\xeb\xf9\x96\x78\x00\x4a\x96\x95\x6d\x59\x6e\xb6\xb5\x6e\x4d\xfd\x2d\x71\x8f\x2d\x28\x3a\x17\x0b\x79\xfa\xdd\xd9\xc3\xff\xa7\x0e\xdf\xa0\xdb\x1c\x78\x89\x67\x3c\x3c\x00\x00\x2e\xd5\x9f\xd2\x4d\x24\xd6\x3c\xc5\x63\xae\x08\x7d\xf2\xbc\xff\x32\xf9\x17\x82\xe4\xa9\x9f\x5d\x0c\x4f\x56\x01\x93\x66\xa3\x96\xd4\x8a\x3c\x9a\x59\x6b\xc2\x79\x7e\x76\x72\x76\x7a\xd4\xb6\xf1\xeb\xa2\x50\xf1\x3d\x80\xf7\xcc\x6b\xc6\x99\xa3\x79\x39\x7c\x1f\xa0\xc6\xfd\x3d\xb8\xb0\xb5\x9e\xe1\x9e\x85\x07\x93\x51\xeb\x0f\xbf\x75\x0a\x3e\xd5\xad\xf8\xb7\x57\x37\x1d\x46\xa0\x69\x8b\xdd\x1b\xd7\x22\x06\x30\x85\xf5\x04\x36\x10\x43\x25\x76\xe1\x1d\x45\x6f\x86\x67\x34\x43\xda\xda\x75\x53\x83\x22\x6a\x90\xc0\x1a\x20\x5b\x21\x7c\x68\x16\xe8\x0c\x7a\x24\x96\xde\xd4\x34\xbc\x92\x15\x86\xba\xa7\x9a\xc9\x95\x35\x38\x49\xbf\xbc\x09\x06\xa4\xef\x64\x51\x39\x8d\x9f\xce\x3a\x0c\x1e\xec\x1b\x7d\xe9\xaf\x07\x93\xd3\xc9\xd1\xff\x02\x00\x00\xff\xff\xe9\x49\x07\x02\xd1\x1a\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
