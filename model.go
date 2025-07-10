package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

type Terminal struct {
	currentDir string
	history    []string
}

func NewTerminal() *Terminal {
	wd, _ := os.Getwd()
	return &Terminal{
		currentDir: wd,
		history:    make([]string, 0),
	}
}

func (t *Terminal) Run() {
	fmt.Println("🚀 Starting Simple Terminal v1.0...")
	fmt.Println("✅ Terminal initialized successfully!")
	fmt.Println("📁 Current directory:", t.currentDir)
	fmt.Println("💡 Type 'help' for available commands, 'exit' to quit")
	fmt.Println(strings.Repeat("-", 50))
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		// Display prompt
		fmt.Printf("%s $ ", t.getPrompt())
		
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		
		// Add to history
		t.history = append(t.history, input)
		
		// Parse and execute command
		if !t.executeCommand(input) {
			break
		}
	}
	
	fmt.Println("\nGoodbye!")
}

func (t *Terminal) getPrompt() string {
	// Show just the directory name, not full path
	return filepath.Base(t.currentDir)
}

func (t *Terminal) executeCommand(input string) bool {
	args := strings.Fields(input)
	if len(args) == 0 {
		return true
	}
	
	command := args[0]
	
	// Handle built-in commands
	switch command {
	case "exit", "quit":
		return false
	case "help":
		t.showHelp()
	case "pwd":
		fmt.Println(t.currentDir)
	case "cd":
		t.changeDirectory(args[1:])
	case "ls":
		t.listFiles(args[1:])
	case "clear":
		t.clearScreen()
	case "history":
		t.showHistory()
	case "echo":
		if len(args) > 1 {
			fmt.Println(strings.Join(args[1:], " "))
		}
	default:
		// Try to execute as external command
		t.executeExternal(args)
	}
	
	return true
}

func (t *Terminal) showHelp() {
	fmt.Println("Built-in commands:")
	fmt.Println("  help     - Show this help message")
	fmt.Println("  exit     - Exit the terminal")
	fmt.Println("  pwd      - Print working directory")
	fmt.Println("  cd <dir> - Change directory")
	fmt.Println("  ls       - List files in current directory")
	fmt.Println("  clear    - Clear the screen")
	fmt.Println("  history  - Show command history")
	fmt.Println("  echo     - Echo text")
	fmt.Println("\nYou can also run any system command available in your PATH.")
}

func (t *Terminal) changeDirectory(args []string) {
	var newDir string
	
	if len(args) == 0 {
		// No arguments, go to home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error getting home directory: %v\n", err)
			return
		}
		newDir = homeDir
	} else {
		newDir = args[0]
		
		// Handle relative paths
		if !filepath.IsAbs(newDir) {
			newDir = filepath.Join(t.currentDir, newDir)
		}
	}
	
	// Clean the path
	newDir = filepath.Clean(newDir)
	
	// Check if directory exists
	if info, err := os.Stat(newDir); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", newDir)
		return
	} else if !info.IsDir() {
		fmt.Printf("cd: %s: Not a directory\n", newDir)
		return
	}
	
	t.currentDir = newDir
}

func (t *Terminal) listFiles(args []string) {
	dir := t.currentDir
	if len(args) > 0 {
		dir = args[0]
		if !filepath.IsAbs(dir) {
			dir = filepath.Join(t.currentDir, dir)
		}
	}
	
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("ls: %v\n", err)
		return
	}
	
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s/\n", file.Name())
		} else {
			fmt.Printf("%s\n", file.Name())
		}
	}
}

func (t *Terminal) clearScreen() {
	// ANSI escape sequence to clear screen
	fmt.Print("\033[H\033[2J")
}

func (t *Terminal) showHistory() {
	for i, cmd := range t.history {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}

func (t *Terminal) executeExternal(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = t.currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	
	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		fmt.Printf("Error executing command: %v\n", err)
	}
}

func main() {
	terminal := NewTerminal()
	terminal.Run()
}
