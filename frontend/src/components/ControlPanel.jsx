import React, { useState } from 'react';
import { ExecuteCommand } from '../../wailsjs/go/main/Jarvis';

export default function ControlPanel() {
    const [command, setCommand] = useState("");
    const [response, setResponse] = useState("");

    const send = () => {
        ExecuteCommand(command).then(res => setResponse(res));
    };

    return (
        <div className="control-panel">
            <input
                value={command}
                onChange={e => setCommand(e.target.value)}
                placeholder="Enter command"
            />
            <button onClick={send}>Run</button>
            <p>{response}</p>
        </div>
    );
}
