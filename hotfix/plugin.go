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

func Build(path string, c *Config) error {
	return defaultPlugin.Build(path, c)
}

func Load(name string) (*Config, error) {
	return defaultPlugin.Load(name)
}

func Init(c *Config) error {
	return defaultPlugin.Init(c)
}
