// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package testing

import (
	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
	"github.com/Ryo-not-rio/grapi/pkg/svcgen"
)

// Injectors from wire.go:

func NewTestApp(command *gencmd.Command, wrapper protoc.Wrapper, ui cli.UI) (*svcgen.App, error) {
	ctx := gencmd.ProvideCtx(command)
	grapicmdCtx := gencmd.ProvideGrapiCtx(ctx)
	rootDir := grapicmd.ProvideRootDir(grapicmdCtx)
	config := grapicmd.ProvideProtocConfig(grapicmdCtx)
	grapicmdConfig := grapicmd.ProvideConfig(grapicmdCtx)
	builder := svcgen.ProvideParamsBuilder(rootDir, config, grapicmdConfig)
	app := &svcgen.App{
		ProtocWrapper: wrapper,
		ParamsBuilder: builder,
	}
	return app, nil
}
