import React, { useEffect, useRef } from "react";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import "xterm/css/xterm.css";

// Wails bindings
import { StartTerminal, SendToTerminal } from "../wailsjs/go/main/App";
import { EventsOn } from "../wailsjs/runtime";

export default function App() {
  const terminalRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const term = new Terminal({
      rows: 20,
      cols: 80,
      cursorBlink: true,
    });
    const fitAddon = new FitAddon();
    term.loadAddon(fitAddon);
    term.open(terminalRef.current!);
    fitAddon.fit();

    // Start terminal in Go backend
    StartTerminal("t1", process.platform === "win32" ? "powershell.exe" : "bash");

    // Listen for output
    EventsOn("terminal-output-t1", (line: string) => {
      term.writeln(line);
    });

    // Send input from frontend → backend
    term.onData((data) => {
      SendToTerminal("t1", data);
    });

    return () => {
      term.dispose();
    };
  }, []);

  return (
    <div className="w-full h-screen bg-black text-white">
      <div ref={terminalRef} className="w-full h-full" />
    </div>
  );
}
