package hypolashlckhelpers

import (
	"fmt"
	logg "github.com/hypolas/hypolaslogger"
	"strings"
)

// NewResult this variable is used for HYPOLAS HEALTHCHECK
//
// for interpret your returned value
//
// # Result.IsUP
//
// must be true if your healthcheck is OK
//
// # Result.Output
//
// must return value of your healthcheck if user want
//
// do other thing with this
//
// # Result must be returned at end of check
func NewResult() Result {
	return Result{}
}

// NewLogger is for unified logs output,
//
// logs all you think is needed for debugging
//
// Please not use other logs system
func NewLogger() logg.HypolasLogger {
	return logg.HypolasLogger{}
}

// InitEnvVars will resolve command output from variable
//
// Exemple env var with value: http://#CMDSTART# hostname -i #CMDEND#:8082/ping
//
// will be interpreted to : http://XX.XXX.XX.X:8082/ping
func (init InitHlckCustom) NewEnvVars(environmentVariable, defaultValue string) string {
	logf.Info.Println(environmentVariable)
	if init.ID != "" {
		newVar := strings.Replace(environmentVariable, "HYPOLAS_HEALTHCHECK_", fmt.Sprintf("HYPOLAS_HEALTHCHECK_%s_", init.ID), -1)
		environmentVariable = newVar

	}
	logf.Info.Println(environmentVariable)
	return getEnv(environmentVariable, defaultValue)
}
