package hypolashlckhelpers

import (
	logg "github.com/hypolas/hypolaslogger"
)

// InitHlckCustom modules (http, MySQL, PostgreSQL ...)
//
// Exemple:
//
// healthcheck --id MYID while transform
//
// HYPOLAS_HEALTHCHECK_XXX_XXX to HYPOLAS_HEALTHCHECK_MYID_XXX_XXX
//
// so user can chain healthcheck
type InitHlckCustom struct {
	ID string // Used for chain healthcheck
}

var (
	log = logg.NewLogger("")
)

// NewResult this variable is used for HYPOLAS HEALTHCHECK
//
// for interpret your returned value
//
// # Result.IsUP
//
// must ne true if your healthcheck is OK
//
// # Result.Output
//
// must return value of your healthcheck if user want
//
// do other thing with this
type Result struct {
	IsUP   bool
	Output string
}
