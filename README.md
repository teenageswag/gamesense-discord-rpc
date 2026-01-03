<div align="center">

# 💫 Gamesense Discord RPC

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/platform-windows-lightgrey?style=for-the-badge)

**Stylish and customizable Rich Presence for your Discord profile.**
Show everyone you're with Gamesense.

[Download Release](https://github.com/teenageswag/gamesense-discord-rpc/releases) • [Report Bug](https://github.com/teenageswag/gamesense-discord-rpc/issues)

</div>

---

## 📖 Description
This is a lightweight Go application that integrates with Discord to set a custom "Gamesense" status.

## 🚀 How to Use

### For Users (Pre-built)
1. **Download** the latest `.exe` file from the [Releases](https://github.com/teenageswag/gamesense-discord-rpc/releases) tab.
2. **Run** the downloaded file.
3. **Enjoy** your new Discord status!

> [!IMPORTANT]
> If Discord was launched automatically with the system or started before the RPC, the status might not appear immediately. \
> In this case, **restart gamesense-discord-rpc**.

---

## 🛠️ Build and Configuration (For Developers)

If you want to customize icons, text, or the App ID, you need to build the project manually.

### Prerequisites
- [Go 1.20](https://go.dev/dl/) or newer.
- Git.

### Build Instructions

1. **Clone the repository:**
   ```bash
   git clone https://github.com/teenageswag/gamesense-discord-rpc.git
   ```

2. **Navigate to the source directory:**
   ```bash
   cd gamesense-discord-rpc/src
   ```

3. **⚙️ Configuration (Required):**
   Open [`main.go`](src/main.go) and find the initialization block:
   ```go
   var gs *rpc.RPCInfo = rpc.Initialize(
       "YOUR_APP_ID", // <-- Insert your App ID here
       "",
       "Get Good - Get Gamesense",
       "gs_logo640",  // Large image asset name
       "gamesense.pub",
       ...
   )
   ```
   Replace `YOUR_APP_ID` with your application ID from the [Discord Developer Portal](https://discord.com/developers/applications).

4. **Install dependencies:**
   ```bash
   go mod tidy
   ```

5. **Compile the project:**
   This command creates an optimized `.exe` file without a console window (hidden mode):
   ```bash
   go build -ldflags "-H=windowsgui" -o ../output/gamesense-rpc.exe main.go
   ```

6. **Run:**
   The executable will be located in the `output` folder.
   ```bash
   ..\output\gamesense-rpc.exe
   ```

---
