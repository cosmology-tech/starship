package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// runCommand will run the cmdStr as exec and return the results as a []byte
// Stderr is wrapped with error and returned as an error
func runCommand(cmdStr string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Env = os.Environ()
	var outb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &outb
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("err: %s, stderr: %s", err, outb.String())
	}

	return outb.Bytes(), nil
}
