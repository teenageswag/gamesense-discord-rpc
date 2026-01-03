# Gamesense Discord-RPC

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-windows-lightgrey?style=for-the-badge)

## Overview

This application will allow you to display information about Gamesense in your Discord profile.

### Prerequisites

- [Go](https://go.dev/dl/) (1.20 or newer recommended)

## Build Instructions
1. Download the project
    ```powershell
    git clone https://github.com/teenageswag/gamesense-discord-rpc.git
    ```
    
2. Replace the string YOUR_APP_ID with your app's ID.
   
3. Navigate to the source directory
    ```powershell
    cd src
    ```
    
4. Install the necessary dependencies
    ```powershell
    go mod tidy
    ```

5. Create an executable file (to work in the background):
    ```powershell
    go build -ldflags -H=windowsgui -o ../output/gamesense-rpc.exe main.go
    ```

6. Run the created executable file:
    ```powershell
    .\output\gamesense-rpc.exe
    ```
