package hypolashlckhelpers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

// InitEnvVars will resolve command output from variable
//
// Exemple env var with value: http://#CMDSTART# hostname -i #CMDEND#:8082/ping
//
// will be interpreted to : http://XX.XXX.XX.X:8082/ping
func (init InitHlckCustom) InitEnvVars(environmentVariable, defaultValue string) string {
	if init.ID != "" {
		environmentVariable = strings.Replace(environmentVariable, "HYPOLAS_HEALTHCHECK_", fmt.Sprintf("HYPOLAS_HEALTHCHECK_%s", init.ID), -1)
	}
	return getEnv(environmentVariable, defaultValue)
}

func getEnv(enVar string, fallback string) string {
	if value, ok := os.LookupEnv(enVar); ok {
		if strings.Contains(value, "#CMDSTART#") {
			return resolveCMD(os.ExpandEnv(value))
		}
	}
	return os.ExpandEnv(fallback)
}

/*
*	Run cmd embeded in environnement variable and run it
 */
func resolveCMD(cmdString string) string {
	inputString := strings.TrimSpace(cmdString)
	strCommand := getStringInBetween(inputString, "#CMDSTART#", "#CMDEND#")
	stringToReplace := "#CMDSTART#" + strCommand + "#CMDEND#"
	logf.VarDebug(strCommand, "strCommand")
	logf.VarDebug(stringToReplace, "stringToReplace")

	cmdArgs := strings.Split(strings.TrimSpace(strCommand), " ")
	logf.VarDebug(cmdArgs, "cmdArgs")

	var cmd *exec.Cmd
	if len(cmdArgs) == 1 {
		cmd = exec.Command(cmdArgs[0])
	} else {
		cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
	}
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Env = os.Environ()
	err := cmd.Run()
	if err != nil {
		logf.Err.Fatalf("%s\n", err)
	}

	/*
	*	Remove all unwanted characters
	 */
	cmdResult := strings.TrimFunc(stdout.String(), func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
	logf.VarDebug(cmdResult, "cmdResult")

	detectedCMDValue := strings.Replace(inputString, stringToReplace, cmdResult, -1)
	logf.VarDebug(detectedCMDValue, "detectedCMDValue")

	return detectedCMDValue
}

/*
*	Extract command from environnement variables
 */
func getStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return result
	}

	newS := str[s+len(start):]
	e := strings.Index(newS, end)
	if e == -1 {
		return result
	}
	result = newS[:e]
	return result
}
