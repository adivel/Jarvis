#![cfg_attr(
    all(not(debug_assertions), target_os = "windows"),
    windows_subsystem = "windows"
)]

use tauri::api::process::Command;
use tauri::api::process::CommandEvent;

#[tauri::command]
async fn run_powershell(cmd: String) -> Result<String, String> {
    let (mut rx, mut child) = Command::new("powershell.exe")
        .args(["-Command", &cmd])
        .spawn()
        .map_err(|e| e.to_string())?;

    let mut output = String::new();

    while let Some(event) = rx.recv().await {
        if let CommandEvent::Stdout(line) = event {
            output.push_str(&line);
            output.push('\n');
        }
    }

    child.wait().await.map_err(|e| e.to_string())?;
    Ok(output)
}

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![run_powershell])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
