package dir

// Child is an interface for the child of a directory
type Child interface {
	Parents() []*Directory
	Path() string
}
