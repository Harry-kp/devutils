# 🚀 DevUtils: Your Go-To CLI Swiss Army Knife for Developers! 🧰

DevUtils is a blazingly fast and intuitive command-line tool designed to streamline your daily development workflow. ⚡️ Say goodbye to tedious tasks and hello to boosted productivity! 💪

## ✨ Features

DevUtils packs a punch with a rich set of features accessible through a sleek terminal-based user interface (TUI).

- **Base64 Encoder/Decoder:** Effortlessly encode and decode text to and from Base64. 🔐
  - **Encode:** Transform your text into Base64 with a single command.
  - **Decode:** Retrieve the original text from a Base64 string.
- **JSON Formatter:** Beautify and prettify your JSON data for enhanced readability. 💅
  - **Input:** Paste your messy JSON.
  - **Output:** Get perfectly formatted JSON, ready for inspection.
- **JWT Decoder:** Unveil the secrets hidden within your JSON Web Tokens (JWTs). 🕵️‍♀️
  - **Decode:** Extract the header and payload information from your JWTs.

## ⚡️ Installation

Getting DevUtils up and running is a breeze!

1. **Prerequisites:** Make sure you have Go installed on your system.
2. **Clone the Repository:**
   ```bash
   git clone https://github.com/harrykp/devutils.git
   ```
3. **Build the Tool:**
   ```bash
   cd devutils
   go build -o devutils
   ```

## 🕹️ Usage

Launch DevUtils directly from your terminal:

```bash
./devutils
```

You'll be greeted with an intuitive TUI.

![DevUtils TUI](devutils-tui.gif) _(Example GIF - Replace with an actual GIF of your TUI in action)_

### 🧭 Navigation

- **Arrow Keys (Up/Down):** Browse through the available tools.
- **Enter:** Select the tool you want to use.
- **Ctrl+C or q:** Exit the tool or return to the main menu.

### 🧰 Using the Tools

Each tool is designed for simplicity. Here's how to use them:

1. **Select the Tool:** Use the arrow keys and Enter to choose your desired tool.
2. **Input Your Data:** Enter the text, JSON, or JWT you want to process.
3. **Press Enter:** DevUtils will work its magic and display the results.
4. **Copy to Clipboard (Automatic):** The output is automatically copied to your clipboard for your convenience!

## 🧑‍💻 Development

Want to contribute or add new features? We welcome contributions!

### 🌱 Adding New Features

1. **Create a Handler:** Add a new handler function in `handlers.go`.
2. **Update the Toolbar:** Integrate your new function into the toolbar in `main.go`.
3. **Implement Logic (if needed):** Add any necessary utility functions in `utils.go`.

### 🔨 Building the Project

After making changes, rebuild the project:

```bash
go build -o devutils
```

## 📄 License

This project is licensed under the MIT License.

---
