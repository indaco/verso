package plugins

// PluginType represents the type of a plugin.
type PluginType string

const (
	TypeCommitParser       PluginType = "commit-parser"
	TypeTagManager         PluginType = "tag-manager"
	TypeVersionValidator   PluginType = "version-validator"
	TypeDependencyChecker  PluginType = "dependency-check"
	TypeChangelogParser    PluginType = "changelog-parser"
	TypeChangelogGenerator PluginType = "changelog-generator"
	TypeReleaseGate        PluginType = "release-gate"
	TypeAuditLog           PluginType = "audit-log"
)

// PluginInfo is a common interface that all plugins should implement.
// It provides basic metadata about a plugin.
type PluginInfo interface {
	Name() string
	Description() string
	Version() string
}

// PluginMetadata contains metadata about a built-in plugin.
type PluginMetadata struct {
	// Type is the plugin type identifier.
	Type PluginType `json:"type"`

	// Name is the unique plugin name.
	Name string `json:"name"`

	// Description explains what the plugin does.
	Description string `json:"description"`

	// Version is the plugin version (semver format).
	Version string `json:"version"`

	// ConfigPath is the YAML path for configuration (e.g., "plugins.tag-manager").
	ConfigPath string `json:"config_path,omitempty"`
}

// builtinPlugins contains metadata for all built-in plugins.
var builtinPlugins = []PluginMetadata{
	{
		Type:        TypeCommitParser,
		Name:        "commit-parser",
		Description: "Parses conventional commits to infer bump type",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.commit-parser",
	},
	{
		Type:        TypeTagManager,
		Name:        "tag-manager",
		Description: "Manages git tags synchronized with version bumps",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.tag-manager",
	},
	{
		Type:        TypeVersionValidator,
		Name:        "version-validator",
		Description: "Enforces versioning policies and constraints",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.version-validator",
	},
	{
		Type:        TypeDependencyChecker,
		Name:        "dependency-check",
		Description: "Syncs version across dependent files",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.dependency-check",
	},
	{
		Type:        TypeChangelogParser,
		Name:        "changelog-parser",
		Description: "Infers bump type from CHANGELOG.md",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.changelog-parser",
	},
	{
		Type:        TypeChangelogGenerator,
		Name:        "changelog-generator",
		Description: "Generates changelog from git commits",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.changelog-generator",
	},
	{
		Type:        TypeReleaseGate,
		Name:        "release-gate",
		Description: "Pre-bump validation for release readiness",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.release-gate",
	},
	{
		Type:        TypeAuditLog,
		Name:        "audit-log",
		Description: "Records version history for audit trail",
		Version:     "v0.1.0",
		ConfigPath:  "plugins.audit-log",
	},
}

// GetBuiltinPlugins returns metadata for all built-in plugins.
func GetBuiltinPlugins() []PluginMetadata {
	result := make([]PluginMetadata, len(builtinPlugins))
	copy(result, builtinPlugins)
	return result
}
