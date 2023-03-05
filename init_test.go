package hypolashlckhelpers

import (
	"bufio"
	"fmt"
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

	v := NewEnvVars("HYPOLAS_HEALTHCHECK_HTTP_URL", "")
	if v != "OK" {
		log.Err.Fatalln(v)
	} else {
		log.Info.Println(v)
	}

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
