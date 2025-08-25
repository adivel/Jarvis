package main

import (
	"bufio"
	"context"
	"io"
	"os/exec"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	processes map[string]*exec.Cmd
	stdins    map[string]io.WriteCloser
	mu        sync.Mutex
	ctx       context.Context
}

func NewApp() *App {
	return &App{
		processes: make(map[string]*exec.Cmd),
		stdins:    make(map[string]io.WriteCloser),
	}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// StartTerminal launches a terminal session
func (a *App) StartTerminal(id string, shell string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	cmd := exec.Command(shell)
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	a.processes[id] = cmd
	a.stdins[id] = stdin

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			runtime.EventsEmit(a.ctx, "terminal-output-"+id, scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			runtime.EventsEmit(a.ctx, "terminal-output-"+id, scanner.Text())
		}
	}()

	go cmd.Run()
}

// SendToTerminal writes input to the running terminal
func (a *App) SendToTerminal(id string, input string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if stdin, ok := a.stdins[id]; ok {
		stdin.Write([]byte(input))
	}
}
