package cmd

import (
	"github.com/Hexilee/javabeat/beater"

	cmd "github.com/elastic/beats/libbeat/cmd"
)

// Name of this beat
const Name = "javabeat"
const Version = "0.1.0-beta"

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmd(Name, Version, beater.New)
