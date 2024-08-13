package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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

// extractTxHash uses a regular expression to extract the txhash from the output
func extractTxHash(output string) (string, error) {
	// Regular expression to match the txhash value in the JSON output
	re := regexp.MustCompile(`"txhash":"([0-9A-Fa-f]{64})"`)
	match := re.FindStringSubmatch(output)
	if len(match) < 2 {
		return "", fmt.Errorf("txhash not found in output")
	}
	return match[1], nil
}

// hasEvent from txResults txMap. Returns nil if the event is found, otherwise an error
func hasEvent(txResults map[string]interface{}, eventType string) error {
	// Check for the "transfer" event
	logs, ok := txResults["logs"].([]interface{})
	if !ok || len(logs) == 0 {
		return fmt.Errorf("no logs found in transaction")
	}

	for _, log := range logs {
		logMap, ok := log.(map[string]interface{})
		if !ok {
			continue
		}
		events, ok := logMap["events"].([]interface{})
		if !ok {
			continue
		}
		for _, event := range events {
			eventMap, ok := event.(map[string]interface{})
			if !ok {
				continue
			}
			if eventMap["type"] == eventType {
				return nil
			}
		}
	}

	return fmt.Errorf("event %s not found in transaction", eventType)
}
