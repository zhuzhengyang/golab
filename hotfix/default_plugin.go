package hotfix

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	pg "plugin"
	"text/template"
)

type plugin struct{}

// runNewFunc sets up the plugin
func (p *plugin) Init(c *Config) error {
	c.NewFunc.(func())()
	return nil
}

// load loads a plugin created with `go build -buildmode=plugin`
func (p *plugin) Load(name string) (*Config, error) {
	path := filepath.Join(getPluginPath(), name+".so")
	plugin, err := pg.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := plugin.Lookup("Plugin")
	if err != nil {
		return nil, err
	}
	pl, ok := s.(*Config)
	if !ok {
		return nil, errors.New("could not cast Plugin object")
	}
	return pl, nil
}

// Generate creates a go file at the specified path.
// You must use `go build -buildmode=plugin`to build it.
func (p *plugin) Generate(path string, name string, c *Config) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(f, c)
}

// Build generates a dso plugin using the go command `go build -buildmode=plugin`
func (p *plugin) Build(name string, c *Config) error {
	path := filepath.Join(getPluginPath(), name)

	// create go file in current path
	goFile := filepath.Join(getPluginPath(), name+".go")

	// generate .go file
	if err := p.Generate(goFile, name, c); err != nil {
		return err
	}
	// remove .go file
	// defer os.Remove(goFile)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("Failed to create dir %s: %v", filepath.Dir(path), err)
	}
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-gcflags=-l", "-o", path+".so", goFile)
	return cmd.Run()
}
