package local

// Dependencies define what can be injected into Local
type Dependencies struct {
	FileSystem FileSystem
}

// GetFileSystem returns the injected filesystem
// or generates a default implementation using os
func (dependencies *Dependencies) GetFileSystem() FileSystem {
	if dependencies != nil && dependencies.FileSystem != nil {
		return dependencies.FileSystem
	}

	return &OSFileSystem{}
}
