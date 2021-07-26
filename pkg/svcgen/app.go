package svcgen

import (
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
	"github.com/Ryo-not-rio/grapi/pkg/svcgen/params"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	ProtocWrapper protoc.Wrapper
	ParamsBuilder params.Builder
}
