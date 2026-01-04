package bumpcmd

import (
	"github.com/indaco/sley/cmd/sley/flags"
	"github.com/indaco/sley/internal/config"
	"github.com/urfave/cli/v3"
)

// Run returns the "bump" parent command.
func Run(cfg *config.Config) *cli.Command {
	cmdFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "pre",
			Usage: "Optional pre-release label",
		},
		&cli.StringFlag{
			Name:  "meta",
			Usage: "Optional build metadata",
		},
		&cli.BoolFlag{
			Name:  "preserve-meta",
			Usage: "Preserve existing build metadata when bumping",
		},
	}
	cmdFlags = append(cmdFlags, flags.MultiModuleFlags()...)

	return &cli.Command{
		Name:      "bump",
		Usage:     "Bump semantic version (patch, minor, major)",
		UsageText: "sley bump <subcommand> [--flags]",
		Flags:     cmdFlags,
		Commands: []*cli.Command{
			patchCmd(cfg),
			minorCmd(cfg),
			majorCmd(cfg),
			preCmd(cfg),
			releaseCmd(cfg),
			autoCmd(cfg),
		},
	}
}
