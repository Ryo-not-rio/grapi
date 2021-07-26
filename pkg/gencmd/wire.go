//+build wireinject

package gencmd

import (
	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/google/wire"
)

func newApp(*Command) (*App, error) {
	wire.Build(
		Set,
		cli.UIInstance,
	)
	return nil, nil
}
