package hypolashlckhelpers

import (
	"bufio"
	"fmt"
	logg "github.com/hypolas/hypolaslogger"
	"os"
	"testing"
)

// Test if log is wheel write.
func TestLogs(t *testing.T) {
	os.Setenv("HYPOLAS_LOGS_FILE", "test/test.log")
	os.Setenv("HYPOLAS_HEALTHCHECK_HTTP_URL", "http://#CMDSTART# hostname #CMDEND#:8082/ping")

	fpath := os.Getenv("HYPOLAS_LOGS_FILE")
	// Remove existing test file
	os.Remove(fpath)

	log := logg.NewLogger("")

	vars := InitHlckCustom{}
	v := vars.InitEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "titi")
	log.Info.Println(v)

	// Test variable
	var tesvar os.Process
	log.VarDebug(tesvar, "tesvar")

	readFile, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
