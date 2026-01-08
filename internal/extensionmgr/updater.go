package extensionmgr

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/indaco/sley/internal/config"
)

var (
	marshalFunc            = yaml.Marshal
	AddExtensionToConfigFn = AddExtensionToConfig
)

// AddExtensionToConfig appends an extension entry to the YAML config at the given path.
// It avoids duplicates and preserves existing fields.
func AddExtensionToConfig(path string, extension config.ExtensionConfig) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config %q: %w", path, err)
	}

	var cfg config.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse config %q: %w", path, err)
	}

	// Avoid duplicates
	for _, ext := range cfg.Extensions {
		if ext.Name == extension.Name {
			return nil
		}
	}

	cfg.Extensions = append(cfg.Extensions, extension)

	out, err := marshalFunc(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, out, config.ConfigFilePerm); err != nil {
		return fmt.Errorf("failed to write config %q: %w", path, err)
	}
	return nil
}
