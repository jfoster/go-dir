package dir

import (
	"os"
	"path/filepath"
)

// Directory ...
type Directory struct {
	Name     string
	Parent   *Directory
	Children []Child
}

// CurrentDirectory returns a pointer to the current directory
func CurrentDirectory() (d *Directory, err error) {
	d = new(Directory)
	d.Name, err = os.Getwd()
	return d, err
}

// NewDirectory returns a pointer to a directory object with a specified name
func NewDirectory(name string) *Directory {
	return &Directory{
		Name: name,
	}
}

// AddFile adds a file to a directory
func (d *Directory) AddFile(f *File) {
	d.Children = append(d.Children, f)
	f.Parent = d
}

// AddDirectory adds a child directory to a directory
func (d *Directory) AddDirectory(dir *Directory) *Directory {
	d.Children = append(d.Children, dir)
	dir.Parent = d
	return dir
}

// AddDirectoryName adds a child directory with a specified name to a directory
func (d *Directory) AddDirectoryName(s string) *Directory {
	return d.AddDirectory(NewDirectory(s))
}

// Write writes a directory to disk
func (d Directory) Write() error {
	return os.MkdirAll(d.Path(), 0755)
}

// WriteAll recurses a directory and all it's children and writes it to disk
func (d Directory) WriteAll() error {
	for _, v := range d.AllChildren() {
		if d, ok := v.(*Directory); ok {
			err := d.Write()
			if err != nil {
				return err
			}
		} else if f, ok := v.(*File); ok {
			err := f.Write()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AllChildren recurses the a directories children and returns a slice of children
func (d *Directory) AllChildren() (c []Child) {
	recurseChild(d, &c)
	return c
}

// Parents impements Child and returns all the parents of a directory
func (d Directory) Parents() (s []*Directory) {
	s = []*Directory{&d}
	for parent := d.Parent; parent != nil; parent = parent.Parent {
		s = append([]*Directory{parent}, s...)
	}
	return s
}

// Path impements Child and returns the full path to a directory
func (d Directory) Path() string {
	var s = []string{}
	for _, v := range d.Parents() {
		s = append(s, v.Name)
	}
	return filepath.Join(s...)
}

// String implements Stringer and returns the name of a directory
func (d Directory) String() string {
	return d.Name
}

func recurseChild(c Child, s *[]Child) {
	if d, ok := c.(*Directory); ok {
		if len(d.Children) > 0 {
			for _, v := range d.Children {
				*s = append(*s, v)
				recurseChild(v, s)
			}
		}
	}
}
