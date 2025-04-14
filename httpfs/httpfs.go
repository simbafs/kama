// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpfs

import (
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/afero"
)

type httpDir struct {
	basePath string
	fs       HTTPFs
}

func (d httpDir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	dir := string(d.basePath)
	if dir == "" {
		dir = "."
	}

	f, err := d.fs.Open(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
	if err != nil {
		return nil, err
	}
	return f, nil
}

type HTTPFs struct {
	source afero.Fs
}

func NewHTTPFs(source afero.Fs) *HTTPFs {
	return &HTTPFs{source: source}
}

func (h HTTPFs) Dir(s string) *httpDir {
	return &httpDir{basePath: s, fs: h}
}

func (h HTTPFs) Name() string { return "h HttpFs" }

func (h HTTPFs) Create(name string) (afero.File, error) {
	return h.source.Create(name)
}

func (h HTTPFs) Chmod(name string, mode os.FileMode) error {
	return h.source.Chmod(name, mode)
}

func (h HTTPFs) Chown(name string, uid, gid int) error {
	return h.source.Chown(name, uid, gid)
}

func (h HTTPFs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return h.source.Chtimes(name, atime, mtime)
}

func (h HTTPFs) Mkdir(name string, perm os.FileMode) error {
	return h.source.Mkdir(name, perm)
}

func (h HTTPFs) MkdirAll(path string, perm os.FileMode) error {
	return h.source.MkdirAll(path, perm)
}

func (h HTTPFs) Open(name string) (http.File, error) {
	name = strings.TrimLeft(name, "/")
	if name == "" {
		name = "."
	}
	f, err := h.source.Open(name)
	if err == nil {
		if httpfile, ok := f.(http.File); ok {
			return httpfile, nil
		}
	}
	return nil, err
}

func (h HTTPFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return h.source.OpenFile(name, flag, perm)
}

func (h HTTPFs) Remove(name string) error {
	return h.source.Remove(name)
}

func (h HTTPFs) RemoveAll(path string) error {
	return h.source.RemoveAll(path)
}

func (h HTTPFs) Rename(oldname, newname string) error {
	return h.source.Rename(oldname, newname)
}

func (h HTTPFs) Stat(name string) (os.FileInfo, error) {
	return h.source.Stat(name)
}
