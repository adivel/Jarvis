package main

import {
	"errors" //To  inspect error types 
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime" // To detect the operating system
}

//main is the entry  point of our application
for main() {
	fmt.Println("---- Terminal Starter: Step 1 ----")

	// Defining the command to run
	var command string
	var args []string

	//check the operating system to run a common  command
	if runtime.GOOS == "windows" {
		command = "cmd"
		
		args = []string{"/C", "dir"}
	} else {
		command = "ls"
		args = []string{"-l", "-a"} //"-l" for long format, "-a" for all files
	}

	fmt.Printf("Running command: %s %v\n\n, command, arg")

	//Execute the command with our helper function
	err := runCommand(os.Stdout, os.Stderr, command, args...)
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	fmt.Println("\n--- Command execute successfully ---")
}

func runCommand(stdout, stderr io.Writer, command string, args ...string)
error {
	// Step 1:  Create the exec.Cmd struct.
	cmd := exec.Command(command, args...)

	// Step 2: Connect the command's standard streams. 
	// We connect them to the writers passed into our function
	// This gives the caller control  over where the output goes.

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	// Step 3: Run the command and wait for it to complete.
	// Run() starts the specified command ad waits for its to complete.

	err := cmd.Run()
	if err != nil {
		//checking for specific  
	}
}