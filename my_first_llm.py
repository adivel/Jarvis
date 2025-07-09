# Make sure you have the requests library
pip install requests

# main.py
import sys
import requests
import json

def ask_ai(prompt):
    # The URL for the Ollama API, assuming it runs on localhost
    url = "http://localhost:11434/api/generate"

    payload = {
        "model": "llama3:8b", # Or whatever model you want
        "prompt": prompt,
        "stream": False # We want the full response at once
    }

    try:
        response = requests.post(url, json=payload)
        response.raise_for_status() # Raise an exception for bad status codes

        # Parse the JSON response
        data = response.json()
        print(data['response'])

    except requests.exceptions.RequestException as e:
        print(f"Error: Could not connect to Ollama API. Is Ollama running?", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    # Get the prompt from the command-line arguments
    if len(sys.argv) < 2:
        print("Usage: python main.py \"<your question>\"", file=sys.stderr)
        sys.exit(1)
    
    user_prompt = " ".join(sys.argv[1:])
    ask_ai(user_prompt)