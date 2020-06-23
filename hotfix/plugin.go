package hotfix

// Config is the plugin config
type Config struct {
	// ImportPath specifies the import path
	ImportPath string
	// NewFunc creates an instance of the plugin
	NewFunc interface{}
}

var (
	// Default plugin loader
	defaultPlugin = &plugin{}
)

// Generate a plugin file in pluginPath.
// name: the name of .so file.
func Build(name string, c *Config) error {
	return defaultPlugin.Build(name, c)
}

func load(name string) (*Config, error) {
	return defaultPlugin.Load(name)
}

func runNewFunc(c *Config) error {
	return defaultPlugin.Init(c)
}
