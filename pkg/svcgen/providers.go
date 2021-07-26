package svcgen

import (
	"github.com/google/wire"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
	"github.com/Ryo-not-rio/grapi/pkg/svcgen/params"
	_ "github.com/Ryo-not-rio/grapi/pkg/svcgen/template"
)

func ProvideParamsBuilder(rootDir cli.RootDir, protocCfg *protoc.Config, grapiCfg *grapicmd.Config) params.Builder {
	return params.NewBuilder(
		rootDir,
		protocCfg.ProtosDir,
		protocCfg.OutDir,
		grapiCfg.Grapi.ServerDir,
		grapiCfg.Package,
	)
}

var Set = wire.NewSet(
	ProvideParamsBuilder,
	App{},
)
