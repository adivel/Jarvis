# Simple Terminal

A lightweight, cross-platform terminal emulator written in Go that provides basic shell functionality with built-in commands and external program execution.

## Features

- **Built-in Commands**: pwd, cd, ls, clear, history, echo, help
- **External Command Support**: Run any system command available in your PATH
- **Directory Navigation**: Change directories with relative/absolute paths
- **Command History**: Track and display previously executed commands
- **Clean Interface**: Minimal, user-friendly prompt and output

## Prerequisites

- Go 1.16 or higher installed on your system
- Terminal/Command Prompt access

## Installation & Usage

1. **Clone or download** the `terminal.go` file
2. **Navigate** to the directory containing the file
3. **Run** the terminal:
   ```bash
   go run terminal.go
   ```

Alternatively, you can build an executable:
```bash
go build -o terminal terminal.go
./terminal
```

## Built-in Commands

| Command | Description |
|---------|-------------|
| `help` | Show available commands |
| `exit` or `quit` | Exit the terminal |
| `pwd` | Print current working directory |
| `cd <directory>` | Change directory |
| `ls [directory]` | List files in current or specified directory |
| `clear` | Clear the terminal screen |
| `history` | Show command history |
| `echo <text>` | Display text |

## Examples

```bash
# Basic navigation
pwd                    # Show current directory
ls                     # List files
cd Documents           # Change to Documents folder
cd ..                  # Go up one directory
cd /home/user          # Absolute path

# Other commands
echo "Hello World"     # Display text
history               # Show previous commands
clear                 # Clear screen
exit                  # Quit terminal
```

## External Commands

The terminal can execute any system command available in your PATH:
```bash
# Examples (availability depends on your system)
date                  # Show current date/time
whoami               # Show current user
python --version     # Check Python version
git status           # Git commands
```

## Platform Support

- ✅ Linux
- ✅ macOS  
- ✅ Windows

## License

This project is open source and available under the MIT License.
