// Filename: internal/psrunner/psrunner.go
package psrunner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// ExecuteAndLog runs a given PowerShell command and appends its output to a log file.
// It is "exported" (public) because it starts with a capital letter.
func ExecuteAndLog(command, logFileName string) error {
	fmt.Printf("Executing PowerShell command: '%s'\n", command)

	cmd := exec.Command("powershell.exe", "-Command", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing PowerShell command: %v\n", err)
		log.Printf("PowerShell output (stderr):\n%s\n", string(output))
	}

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		// We return the error instead of fatally exiting, so the main app can decide what to do.
		return fmt.Errorf("failed to open log file %s: %w", logFileName, err)
	}
	defer file.Close()

	fileLogger := log.New(file, "", 0)

	fileLogger.Println("--------------------------------------------------")
	fileLogger.Printf("Log Entry: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fileLogger.Printf("Executed Command: %s\n", command)
	fileLogger.Println("Output:")
	fileLogger.Println(string(output))
	fileLogger.Println("--------------------------------------------------")

	fmt.Printf("Successfully saved logs to %s\n", logFileName)
	return nil // Return nil to indicate success
}