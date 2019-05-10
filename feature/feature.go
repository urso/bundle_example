package feature

import "github.com/elastic/beats/libbeat/common"

type Info struct {
	Name         string
	Version      string
	ShortDoc     string
	Doc          string
	ConfigSchema Schema
}

type Schema interface {
	Validate(config *common.Config) error
}

func Namespaced(name string, schema Schema) Schema {
	return nil
}

func OneOf(names []string, schemas []Schema) Schema {
	return nil
}
