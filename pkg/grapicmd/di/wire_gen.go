// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/izumin5210/grapi/pkg/cli"
	"github.com/izumin5210/grapi/pkg/excmd"
	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/module"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/usecase"
	"github.com/izumin5210/grapi/pkg/protoc"
)

// Injectors from wire.go:

func NewUI(ctx *grapicmd.Ctx) cli.UI {
	io := grapicmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	return ui
}

func NewCommandExecutor(ctx *grapicmd.Ctx) excmd.Executor {
	io := grapicmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	return executor
}

func NewGenerator(ctx *grapicmd.Ctx) module.Generator {
	io := grapicmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	generator := ProvideGenerator(ctx, ui)
	return generator
}

func NewScriptLoader(ctx *grapicmd.Ctx) module.ScriptLoader {
	io := grapicmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	scriptLoader := ProvideScriptLoader(ctx, executor)
	return scriptLoader
}

func NewProtocWrapper(ctx *grapicmd.Ctx) (protoc.Wrapper, error) {
	config := grapicmd.ProvideProtocConfig(ctx)
	fs := grapicmd.ProvideFS(ctx)
	execInterface := grapicmd.ProvideExecer(ctx)
	io := grapicmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	rootDir := grapicmd.ProvideRootDir(ctx)
	gexConfig := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(gexConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, execInterface, ui, repository, rootDir)
	return wrapper, nil
}

func NewInitializeProjectUsecase(ctx *grapicmd.Ctx) usecase.InitializeProjectUsecase {
	fs := grapicmd.ProvideFS(ctx)
	execInterface := grapicmd.ProvideExecer(ctx)
	io := grapicmd.ProvideIO(ctx)
	rootDir := grapicmd.ProvideRootDir(ctx)
	config := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	ui := cli.UIInstance(io)
	generator := ProvideGenerator(ctx, ui)
	initializeProjectUsecase := ProvideInitializeProjectUsecase(ctx, config, ui, generator)
	return initializeProjectUsecase
}
