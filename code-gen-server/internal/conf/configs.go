package conf

import "github.com/google/wire"

// 除了http的，其他的都要提供
var ProviderSet = wire.NewSet(
	NewGin,
	ZapProvider,
	GormProvider,
)
