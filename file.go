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

// NewFilePath returns a pointer to a File from a full filepath
func NewFilePath(path string) (f *File) {
	file := filepath.Base(path)
	dir := filepath.Dir(path)
	f = NewFile(file)
	f.Parent = NewDirectory(dir)
	return f
}

// WithContent returns a pointer to a File from an existing file with specified contents
func (f *File) WithContent(content []byte) *File {
	f.Content = content
	return f
}

// WithParent returns a pointer to a File from an existing file with a specified parent
func (f *File) WithParent(parent *Directory) *File {
	f.Parent = parent
	return f
}

// String returns a formatted string of the file's name and extension
func (f *File) FullName() string {
	return f.Name + f.Ext
}

// Read populates the content of the file from the path
func (f *File) Read() (err error) {
	f.Content, err = ioutil.ReadFile(f.Path())
	return err
}

// Write writes contents of the file to disk
func (f *File) Write() error {
	return f.WritePerm(0644)
}

// WritePerm writes contents of the file to disk with specified permissions
func (f *File) WritePerm(perm os.FileMode) error {
	return ioutil.WriteFile(f.Path(), []byte(f.Content), perm)
}

// Parents impements Child and returns all the parents of a file
func (f *File) Parents() (s []*Directory) {
	for parent := f.Parent; parent != nil; parent = parent.Parent {
		s = append([]*Directory{parent}, s...)
	}
	return s
}

// Path impements Child and returns the full path to a file
func (f *File) Path() string {
	var s = []string{}
	for _, v := range f.Parents() {
		s = append(s, v.Name)
	}
	s = append(s, f.FullName())
	return filepath.Join(s...)
}

// String implements Stringer and return the full name of the file
func (f *File) String() string {
	return f.FullName()
}
