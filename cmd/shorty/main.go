package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	amartha "github.com/amartha-shorty/internal/app/shorty"
)

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "verbose",
		Usage:   "Verbosity level",
		Value:   false,
		EnvVars: []string{"VERBOSE"},
	},
}

func action(ctx *cli.Context) error {
	if ctx.Bool("verbose") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	options := amartha.Options{}

	return amartha.Start(options)
}

func main() {
	app := &cli.App{
		Name:    "amartha-shorty",
		Usage:   "ntar ah mikir",
		Version: "1.0.0",
		Flags:   flags,
		Action:  action,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
