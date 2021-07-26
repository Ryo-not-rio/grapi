//+build wireinject

package svcgen

import (
	"github.com/google/wire"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
)

func NewApp(*gencmd.Command) (*App, error) {
	wire.Build(
		Set,
		gencmd.Set,
		cli.UIInstance,
		protoc.WrapperSet,
	)
	return nil, nil
}
