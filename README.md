# PixelEditor

[![Fyne](https://img.shields.io/badge/Fyne-2.6.0-blue)](https://fyne.io/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A simple pixel art editor built with Go and the Fyne UI framework.

https://github.com/user-attachments/assets/84e97a69-7c5f-499e-b275-1a8572e5965d

## Features

- **Pixel Canvas**: 50x50 grid for creating pixel art
- **Color Selection**: RGB color picker
- **Brush Tools**: Basic brush functionality
- **Pan & Navigate**: Move around the canvas
- **Simple UI**: Clean interface with essential controls

## Installation

### Prerequisites

- Go 1.22 or higher
- Fyne dependencies (see [Fyne Getting Started](https://developer.fyne.io/started/))

### Option 1: Run Natively (Recommended)

```bash
# Clone the repository
git clone https://github.com/abdul756/Pixel-Editor.git
cd Pixel-Editor

# Run the helper script (will install dependencies if needed)
./run-native.sh
```

### Option 2: Using Docker

Make sure Docker is installed and running on your system.

```bash
# Clone the repository
git clone https://github.com/abdul756/Pixel-Editor.git
cd Pixel-Editor

# Run the helper script
./run.sh
```

## Usage

1. **Start the Application**: Launch PixelEditor
2. **Draw**: Click and drag on the canvas to draw pixels
3. **Change Colors**: Use the color picker to select colors
4. **Navigate**: Right-click and drag to pan around the canvas

## Project Structure

- **src/core**: Canvas rendering and pixel manipulation
- **src/ui**: User interface components
- **src/utils**: Utility functions and types

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Built with Go and Fyne
