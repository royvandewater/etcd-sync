package local

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"strings"
)

// Service implements an Etcd.Service and represents a collection
// of key/value pairs
type Service interface {
	Name() string
	Path() string
	Records() (map[string]string, error)
}

type localService struct {
	dirname  string
	basename string
	fs       FileSystem
}

// NewService generates a new Service instance
func NewService(dirname, basename string, dependencies *Dependencies) Service {
	fs := dependencies.GetFS()
	return &localService{dirname, basename, fs}
}

func (service *localService) Name() string {
	return service.basename
}

func (service *localService) Path() string {
	return path.Join(service.dirname, service.basename)
}

func (service *localService) Records() (map[string]string, error) {
	contents, err := service.fileContents()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(contents, "\n")

	records := make(map[string]string, len(lines))
	for i, line := range lines {
		if service.skipLine(line) {
			continue
		}

		if !service.isValidLine(line) {
			message := fmt.Sprintf("Malformed line %v: '%v'", i, line)
			return nil, errors.New(message)
		}

		key, value := service.parseLine(line)
		records[key] = value
	}
	return records, nil
}

func (service *localService) fileContents() (string, error) {
	contents, err := service.fs.ReadFile(service.Path())
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (service *localService) isValidLine(line string) bool {
	expression := regexp.MustCompile("\".*?\"")
	line = expression.ReplaceAllString(line, "x")

	parts := strings.Split(line, " ")
	return len(parts) == 2
}

func (service *localService) parseLine(line string) (key, value string) {
	parts := strings.SplitN(line, " ", 2)
	return parts[0], parts[1]
}

func (service *localService) skipLine(line string) bool {
	return len(line) == 0
}
