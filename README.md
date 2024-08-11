# devutils

`devutils` is a command-line tool designed to help developers perform common tasks efficiently. The tool features a terminal-based user interface (TUI) that provides quick access to various utilities, such as Base64 encoding/decoding, JSON formatting, and JWT decoding.

## Features

- **Base64 Encoder**: Easily encode text into Base64 format.
- **Base64 Decoder**: Decode Base64 strings back to their original text.
- **JSON Formatter**: Format JSON strings to improve readability.
- **JWT Decoder**: Decode JSON Web Tokens (JWT) to view their header and payload.

## Installation

First, ensure you have Go installed on your machine. Then, clone the repository and build the tool:

```bash
git clone https://github.com/yourusername/devutils.git
cd devutils
go build -o devutils
```

## Usage

Run the tool directly from your terminal:

```bash
./devutils
```

### Navigating the Interface

- Use the **arrow keys** (`Up` and `Down`) to navigate through the list of available tools.
- Press **Enter** to select a tool.
- Use **Ctrl+C** or `q` to exit or go to previous menu at any time.

### Base64 Encoder

1. Select **Base64 Encoder** from the list.
2. Enter the text you want to encode.
3. The encoded string will be displayed and automatically copied to your clipboard.

### Base64 Decoder

1. Select **Base64 Decoder** from the list.
2. Enter the Base64 string you want to decode.
3. The decoded text will be displayed and automatically copied to your clipboard.

### JSON Formatter

1. Select **JSON Formatter** from the list.
2. Enter the JSON string you want to format.
3. The formatted JSON will be displayed and automatically copied to your clipboard.

### JWT Decoder

1. Select **JWT Decoder** from the list.
2. Enter the JWT string you want to decode.
3. The decoded header and payload will be displayed.

## Development

### Adding New Features

To add a new feature:

1. Create a new handler function in `handlers.go`.
2. Add the function to the toolbar in `main.go`.
3. Implement the logic in `utils.go` if necessary.

### Building the Project

After making changes, you can rebuild the project with:

```bash
go build -o devutils
```

## License

This project is licensed under the MIT License.
