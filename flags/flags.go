package flags

import (
	"fmt"
	"os"
	"path/filepath"

	goFlags "github.com/jessevdk/go-flags"
	"github.com/senorprogrammer/wtf/wtf"
)

type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Module  string `short:"m" long:"module" optional:"yes" description:"Display info about a specific module, i.e.: 'wtf -m=todo'"`
	Version bool   `short:"v" long:"version" description:"Show version info"`
}

func NewFlags() *Flags {
	flags := Flags{}
	return &flags
}

/* -------------------- Exported Functions -------------------- */

func (flags *Flags) HasConfig() bool {
	return len(flags.Config) > 0
}

func (flags *Flags) HasModule() bool {
	return len(flags.Module) > 0
}

func (flags *Flags) HasVersion() bool {
	return flags.Version == true
}

func (flags *Flags) Parse(version string) {
	parser := goFlags.NewParser(flags, goFlags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*goFlags.Error); ok && flagsErr.Type == goFlags.ErrHelp {
			os.Exit(0)
		}
	}

	if !flags.HasConfig() {
		homeDir, err := wtf.Home()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		flags.Config = filepath.Join(homeDir, ".wtf", "config.yml")
	}
}
