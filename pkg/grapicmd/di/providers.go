package di

import (
	"net/http"

	"github.com/google/wire"
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/izumin5210/execx"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/afero"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd/internal/module"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd/internal/module/script"
	"github.com/Ryo-not-rio/grapi/pkg/grapicmd/internal/usecase"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
)

func ProvideScriptLoader(ctx *grapicmd.Ctx, io *clib.IO, exec *execx.Executor) module.ScriptLoader {
	return script.NewLoader(ctx.FS, io, exec, ctx.RootDir.String())
}

func ProvideGenerator(ctx *grapicmd.Ctx, ui cli.UI, fs afero.Fs, tmplFs http.FileSystem, outDir clib.Path) gencmd.Generator {
	return gencmd.NewGenerator(
		fs,
		ui,
		outDir,
		tmplFs,
		nil,
	)
}

var Set = wire.NewSet(
	grapicmd.CtxSet,
	protoc.WrapperSet,
	cli.UIInstance,
	ProvideScriptLoader,
	ProvideGenerator,
	fs.New,
	usecase.NewInitializeProjectUsecase,
)
