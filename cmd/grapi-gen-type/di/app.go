package di

import (
	"github.com/Ryo-not-rio/grapi/pkg/gencmd"
	"github.com/Ryo-not-rio/grapi/pkg/protoc"
)

type CreateAppFunc func(*gencmd.Command) (*App, error)

type App struct {
	Protoc protoc.Wrapper
}
