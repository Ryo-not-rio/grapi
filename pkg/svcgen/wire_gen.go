// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package svcgen

import (
	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
)

import (
	_ "github.com/Ryo-not-rio/grapi/pkg/svcgen/template"
)

// Injectors from wire.go:

func NewApp(command *gencmd.Command) (*App, error) {
	ctx := gencmd.ProvideCtx(command)
	grapicmdCtx := gencmd.ProvideGrapiCtx(ctx)
	config := grapicmd.ProvideProtocConfig(grapicmdCtx)
	fs := grapicmd.ProvideFS(grapicmdCtx)
	executor := grapicmd.ProvideExec(grapicmdCtx)
	io := grapicmd.ProvideIO(grapicmdCtx)
	ui := cli.UIInstance(io)
	rootDir := grapicmd.ProvideRootDir(grapicmdCtx)
	gexConfig := protoc.ProvideGexConfig(fs, executor, io, rootDir)
	repository, err := protoc.ProvideToolRepository(gexConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, executor, ui, repository, rootDir)
	grapicmdConfig := grapicmd.ProvideConfig(grapicmdCtx)
	builder := ProvideParamsBuilder(rootDir, config, grapicmdConfig)
	app := &App{
		ProtocWrapper: wrapper,
		ParamsBuilder: builder,
	}
	return app, nil
}
