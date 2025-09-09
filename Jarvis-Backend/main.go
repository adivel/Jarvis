// Filename: main.go
package main

import (
	"log"
	// Import our new internal package. The path is based on your module name.
	"Jarvis-Backend/internal/psrunner"
)

func main() {
	// Define the command and log file here, in the main application logic.
	psCommand := "Get-Process | Sort-Object CPU -Descending | Select-Object -First 10"
	logFile := "powershell_log.txt"

	// Call the public function from our psrunner package.
	err := psrunner.ExecuteAndLog(psCommand, logFile)
	if err != nil {
		// The main function decides how to handle errors.
		log.Fatalf("The psrunner module failed: %v", err)
	}

	log.Println("Main application finished successfully.")
}