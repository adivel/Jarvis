import React, { useState } from "react";

function App() {
  const [data, setData] = useState<any>(null);

  const callAI = async () => {
    try {
      const res = await fetch("http://127.0.0.1:8080/ai/predict");
      const json = await res.json();
      setData(json);
    } catch (err) {
      console.error("Error fetching AI data:", err);
    }
  };

  return (
    <div className="p-4">
      <h1>🚀 Jarvis Terminal</h1>
      <button
        onClick={callAI}
        className="bg-blue-500 text-white px-4 py-2 rounded"
      >
        Run AI Prediction
      </button>
      {data && (
        <pre className="mt-4 bg-gray-200 p-2 rounded">
          {JSON.stringify(data, null, 2)}
        </pre>
      )}
    </div>
  );
}

export default App;
