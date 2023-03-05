package hypolashlckhelpers

import (
	logg "github.com/hypolas/hypolaslogger"
)

// InitVars modules (http, MySQL, PostgreSQL ...)
type InitHlckCustom struct {
	ID *string
}

var (
	logf = logg.NewLogger("")
)
