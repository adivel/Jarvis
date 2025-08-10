import { useState } from "react";
import { SayHello } from "../wailsjs/go/main/App"; // ✅ Import the actual Wails binding

function App() {
    const [msg, setMsg] = useState("");

    const greet = async () => {
        try {
            const res = await SayHello("Adiyanthy"); // ✅ Call Go method
            setMsg(res);
        } catch (err) {
            console.error("Error calling SayHello:", err);
            setMsg("Something went wrong!");
        }
    };

    return (
        <div style={{ padding: "20px", fontFamily: "sans-serif" }}>
            <h1>{msg || "Click Greet to start"}</h1>
            <button onClick={greet}>Greet</button>
        </div>
    );
}

export default App;
