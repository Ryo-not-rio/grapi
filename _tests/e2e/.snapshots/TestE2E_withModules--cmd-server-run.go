package main

import (
	"github.com/srvc/appctx"

	"github.com/Ryo-not-rio/grapi/pkg/grapiserver"
)

func run() error {
	// Application context
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
		// TODO
		),
	)
	return s.Serve(ctx)
}
