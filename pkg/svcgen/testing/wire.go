//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
	"github.com/Ryo-not-rio/grapi/pkg/svcgen"
)

func NewTestApp(*gencmd.Command, protoc.Wrapper, cli.UI) (*svcgen.App, error) {
	wire.Build(
		gencmd.Set,
		svcgen.ProvideParamsBuilder,
		svcgen.App{},
	)
	return nil, nil
}
