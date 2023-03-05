package hypolashlckhelpers

import (
	"fmt"
	"os"
	"strings"

	logg "github.com/hypolas/hypolaslogger"

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
	return logg.NewLogger("")
}

// NewEnvVars will resolve command output from variable
//
// Exemple env var with value: http://#CMDSTART# hostname -i #CMDEND#:8082/ping
//
// will be interpreted to : http://XX.XXX.XX.X:8082/ping
func NewEnvVars(environmentVariable, defaultValue string) string {
	environmentVariable = GetUpdatedEnvVarName(environmentVariable)
	log.Info.Println("environmentVariable", environmentVariable)
	return getEnv(environmentVariable, defaultValue)
}

// GetRUpdatedEnvVarName change the environnment variable if ID used
func GetUpdatedEnvVarName(environmentVariable string) string {
	customID := os.Getenv("HYPOLAS_HEALTHCHECK_ID")
	log.VarDebug(customID, "customID")
	if customID != "" {
		return strings.Replace(environmentVariable, "HYPOLAS_HEALTHCHECK_", fmt.Sprintf("HYPOLAS_HEALTHCHECK_%s_", customID), -1)
	}
	return environmentVariable
}
