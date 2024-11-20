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

// hasEvent checks for a specific event type in the transaction results.
// It first tries to fetch the events either from txMap["events"] or from txMap["logs"] and then checks for the event type.
func hasEvent(txResults map[string]interface{}, eventType string) error {
	var events []interface{}

	// Attempt to fetch events directly from txMap["events"]
	if directEvents, ok := txResults["events"].([]interface{}); ok && len(directEvents) > 0 {
		events = directEvents
	} else if logs, ok := txResults["logs"].([]interface{}); ok && len(logs) > 0 {
		// If no direct events, attempt to fetch events from logs
		for _, log := range logs {
			logMap, ok := log.(map[string]interface{})
			if !ok {
				continue
			}
			if logEvents, ok := logMap["events"].([]interface{}); ok && len(logEvents) > 0 {
				events = append(events, logEvents...)
			}
		}
	}

	// If events are found, check for the specific eventType
	if len(events) > 0 {
		for _, event := range events {
			eventMap, ok := event.(map[string]interface{})
			if !ok {
				continue
			}
			if eventMap["type"] == eventType {
				return nil
			}
		}
		return fmt.Errorf("event %s not found in transaction events", eventType)
	}

	// If no events were found in either place
	return fmt.Errorf("no events found in transaction")
}
