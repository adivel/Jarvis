# app.py
from fastapi import FastAPI
import torch

app = FastAPI()

@app.get("/")
def home():
    return {"message": "AI Service is running!"}

@app.get("/predict")
def predict():
    x = torch.tensor([1.0, 2.0, 3.0])
    y = x * 2
    return {"input": x.tolist(), "output": y.tolist()}
