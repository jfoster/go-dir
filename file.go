package file

import (
	"path/filepath"
	"strings"
)

// File ...
type File struct {
	Dir       string
	Name      string
	Extension string
}

// New creates a new File instance from a filepath string
func New(file string) File {
	dir := filepath.Dir(file)
	ext := filepath.Ext(file)
	name := filepath.Base(strings.TrimSuffix(file, ext))
	return File{Dir: dir, Name: name, Extension: ext}
}

// FullName returns a string of the filename + extension
func (f File) FullName() string {
	fe := f.Name + f.Extension
	return filepath.Clean(fe)
}

// FullPath returns a string of the filepath + filename + extension
func (f File) FullPath() string {
	fp := filepath.Join(f.Dir, f.FullName())
	return filepath.Clean(fp)
}
