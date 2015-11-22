package local

// Dependencies define what can be injected into Local
type Dependencies struct {
	fs FileSystem
}

// NewDependencies constructs a new dependencies instance
func NewDependencies(fs FileSystem) *Dependencies {
	return &Dependencies{fs: fs}
}

// GetFS returns the injected fileSystem
// or generates a default implementation using os
func (dependencies *Dependencies) GetFS() FileSystem {
	if dependencies != nil && dependencies.fs != nil {
		return dependencies.fs
	}

	return NewFileSystem()
}
