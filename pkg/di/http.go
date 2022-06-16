package di

import (
	"github.com/google/wire"
	api "github.com/thnkrn/go-gin-template/pkg/api"
)

var HTTPSet = wire.NewSet(
	api.NewServerHTTP,
	wire.Struct(new(api.Handlers), "*"),
)
