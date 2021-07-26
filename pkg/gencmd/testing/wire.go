//+build wireinject

package testing

import (
	"github.com/google/wire"

	"github.com/Ryo-not-rio/grapi/pkg/cli"
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
)

func NewTestApp(*gencmd.Command, cli.UI) (*gencmd.App, error) {
	wire.Build(
		gencmd.Set,
	)
	return nil, nil
}
