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
	fpath := os.Getenv("HYPOLAS_LOGS_FILE")
	// Remove existing test file
	err := os.Remove(fpath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fpath)

	log := logg.NewLogger("")

	vars := InitHlckCustom{}
	v := vars.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	if v != "OK" {
		log.Err.Fatalln(v)
	}

	vars.ID = "MYID1"
	v = vars.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	if v != "OK" {
		log.Err.Fatalln(v)
	}

	// vars.ID = "MYID2"
	// v = vars.NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	// if v != "VAROK" {
	// 	log.Err.Fatalln(v)
	// }

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
