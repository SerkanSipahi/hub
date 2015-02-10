package commands

import (
	"fmt"
	"os"

	"github.com/github/hub/console"
	"github.com/github/hub/git"
	"github.com/github/hub/utils"
)

var Version = "2.2.0-rc1"

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "Show hub version",
	Long:  `Shows git version and hub client version.`,
}

func init() {
	CmdRunner.Use(cmdVersion)
}

func runVersion(cmd *Command, args *Args) {
	gitVersion, err := git.Version()
	utils.Check(err)

	ghVersion := fmt.Sprintf("hub version %s", Version)

	console.Infoln(gitVersion)
	console.Infoln(ghVersion)

	os.Exit(0)
}
