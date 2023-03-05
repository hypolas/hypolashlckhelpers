package hypolashlckhelpers

import (
	logg "github.com/hypolas/hypolaslogger"
)

// InitHlckCustom modules (http, MySQL, PostgreSQL ...)
type InitHlckCustom struct {
	ID string
}

var (
	logf = logg.NewLogger("")
)
