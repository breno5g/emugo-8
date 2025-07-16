# Emugo-8: A Chip-8 Emulator in Go

![Go Version](https://img.shields.io/badge/Go-1.20%2B-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)

Welcome to **Emugo-8**, a Chip-8 emulator developed entirely in Go, using the SDL2 library for graphics rendering and input handling. This project was created as a learning tool to deepen knowledge of computer architecture, low-level languages, and how emulators work.

Chip-8 is not a physical machine but rather an interpreted virtual machine that became popular in the 1970s. Emulating Chip-8 is often considered the "Hello, World!" of emulator development. For more technical details about the implementation, please refer to the [**Emugo-8 specification documentation**](https://breno5g.github.io/emugo-8).

## 🎯 Features

* **Full CPU Implementation:** All 35 Chip-8 opcodes have been implemented, allowing most classic ROMs to be executed.
* **Graphics Rendering:** Uses the **SDL2** library to render the 64x32 pixel display, with a customizable scale for better viewing on modern monitors.
* **Keyboard Controls:** Mapping of the original hexadecimal Chip-8 keypad to a modern QWERTY keyboard.
* **Timers and Audio:** Implementation of the delay (DT) and sound (ST) timers, with a simple "beep" sound when the sound timer is active.
* **Extensive Test Coverage:** The project has robust unit tests for the vast majority of opcodes, ensuring the reliability of the emulation.

## 🚀 How to Run

To compile and run Emugo-8, you will need **Go (version 1.20 or higher)** and the **SDL2 library** installed on your system. The project includes a `Makefile` to simplify common commands.

### Prerequisites

1.  **Install Go:**
    * Follow the instructions on the [official Go website](https://golang.org/doc/install).
2.  **Install the SDL2 library:**
    * **Ubuntu/Debian:**
        ```bash
        sudo apt-get install libsdl2-dev
        ```
    * **macOS (using Homebrew):**
        ```bash
        brew install sdl2
        ```
    * **Windows (using MSYS2/MinGW):**
        * Follow the [installation guide for the `go-sdl2` library](https://github.com/veandco/go-sdl2#installation).

### Makefile Commands

After cloning the repository and navigating to the root directory, you can use the following commands:

* **To run the emulator:**
    * By default, it will load the ROM defined in the code (`main.go`).
    ```bash
    make run
    ```
* **To run the tests:**
    ```bash
    make test
    ```
* **To compile the executable binary:**
    * This will create a file named `emugo-8` in the project root.
    ```bash
    make build
    ```
* **To clean the build files:**
    ```bash
    make clean
    ```

## ⌨️ Controls

The original Chip-8 keypad has 16 hexadecimal keys (0-F). Emugo-8 maps these keys to a modern QWERTY layout as follows:

| Original Keypad | QWERTY Keyboard |
| :-------------: | :-------------: |
|    `1 2 3 C`    |    `1 2 3 4`    |
|    `4 5 6 D`    |    `Q W E R`    |
|    `7 8 9 E`    |    `A S D F`    |
|    `A 0 B F`    |    `Z X C V`    |

## 💡 Future Improvements

This project was made for learning, and as such, there is always room to grow. Some of the planned future features include:

* **Interactive Debugger:** A tool to execute code step-by-step, view registers, and inspect memory in real-time.
* **Advanced Audio:** Replace the basic "beep" with higher-quality audio using the `SDL2_mixer` library.
* **UI Settings:** Add options to pause, reset, and control the emulation speed in real-time.
* **Save and Load States:** Implement the ability to save the current state of the emulation and load it later.

## 📜 License

This project is distributed under the MIT License. See the `LICENSE` file for more details.