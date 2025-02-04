package main

import (
	"fmt"
	"os"

	"github.com/izumin5210/clig/pkg/clib"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd/cmd"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.NewGrapiCommand(&grapicmd.Ctx{
		IO:      clib.Stdio(),
		RootDir: cli.RootDir{clib.Path(cwd)},
		Build: clib.Build{
			AppName:   name,
			Version:   version,
			Revision:  revision,
			BuildDate: buildDate,
		},
	}).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
