package plants

import "github.com/google/wire"

// Set of providers for channels components.
var Householders = wire.NewSet(NewGreenHouse, HttpHandler)
