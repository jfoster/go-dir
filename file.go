package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// File ...
type File struct {
	Name    string
	Ext     string
	Content []byte
	Parent  *Directory
}

// NewFile returns a pointer to a File with a specified name
func NewFile(name string) *File {
	ext := filepath.Ext(name)
	return &File{
		Name: strings.TrimSuffix(name, ext),
		Ext:  ext,
	}
}

// NewFileContent returns a pointer to a File with a specified name and contents
func NewFileContent(name string, content []byte) (f *File) {
	f = NewFile(name)
	f.Content = content
	return f
}

// String returns a formatted string of the file's name and extension
func (f File) FullName() string {
	return f.Name + f.Ext
}

// Read populates the content of the file from the path
func (f File) Read() (err error) {
	f.Content, err = ioutil.ReadFile(f.Path())
	return err
}

// Write writes contents of the file to disk
func (f File) Write() error {
	return f.WritePerm(0644)
}

// WritePerm writes contents of the file to disk with specified permissions
func (f File) WritePerm(perm os.FileMode) error {
	return ioutil.WriteFile(f.Path(), []byte(f.Content), perm)
}

// Parents impements Child and returns all the parents of a file
func (f File) Parents() (s []*Directory) {
	for parent := f.Parent; parent != nil; parent = parent.Parent {
		s = append([]*Directory{parent}, s...)
	}
	return s
}

// Path impements Child and returns the full path to a file
func (f File) Path() string {
	var s = []string{}
	for _, v := range f.Parents() {
		s = append(s, v.Name)
	}
	s = append(s, f.FullName())
	return filepath.Join(s...)
}

// String implements Stringer and return the full name of the file
func (f File) String() string {
	return f.FullName()
}
