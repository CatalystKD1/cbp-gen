# cbp-gen - Code Boilerplate Generator

A fast, native CLI tool for generating boilerplate code across multiple programming languages. Built with Go, `cbp-gen` runs on any operating system and requires no runtime. Just download and go.

> 🚧 **Current Status: Sprint 02 Complete**

---

## Supported Languages

| Flag | Language | Output |
|------|----------|--------|
| `-c` | C | `.c` source file |
| `-cpp` | C++ | `.cpp` source file |
| `-go` | Go | `.go` source file |
| `-html` | HTML | `.html` page |
| `-react` | React | Full project via `npm` |
| `-react-c` | React Component | `.jsx` / `.tsx` component file |

---

## Installation

Since `cbp-gen` is compiled to a **self-contained binary**, you do **not** need Go installed on your system to use it. Just download the binary for your platform and run it.

### Option 1 - Download a Pre-built Binary (Recommended)

Head to the [Releases](../../releases) page and download the binary for your operating system:

| OS | File |
|----|------|
| macOS (ARM / Apple Silicon) | `cbp-gen-darwin-arm64` |
| macOS (Intel) | `cbp-gen-darwin-amd64` |
| Linux (x86_64) | `cbp-gen-linux-amd64` |
| Windows | `cbp-gen-windows-amd64.exe` |

**macOS / Linux - make it executable and move it to your PATH:**

```bash
chmod +x cbp-gen-darwin-arm64        # (or your platform's file)
sudo mv cbp-gen-darwin-arm64 /usr/local/bin/cbp-gen
```

**Windows** - move the `.exe` to a folder that's on your system `PATH`, or run it directly from its location.

### Option 2 - Build from Source

> Requires [Go 1.21+](https://go.dev/dl/)

```bash
git clone https://github.com/your-username/cbp-gen.git
cd cbp-gen
go build -o cbp-gen .
```

Then move the built binary somewhere on your `PATH`:

```bash
# macOS / Linux
sudo mv cbp-gen /usr/local/bin/

# Windows (PowerShell - run as Administrator)
Move-Item cbp-gen.exe C:\Windows\System32\
```

---

## Usage

```
cbp-gen [language flag] [options]
```

### Flags

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-c` | bool | - | Generate a C file |
| `-cpp` | bool | - | Generate a C++ file |
| `-go` | bool | - | Generate a Go file |
| `-html` | bool | - | Generate an HTML file |
| `-react` | bool | - | Initialize a React project (runs `npm`) |
| `-react-c` | bool | - | Generate a React component (`.jsx`) |
| `-name` | string | `main` | Output file name |
| `-title` | string | `Hello World` | HTML page title |
| `-type` | string | `int` | Return type of the generated function |
| `-ts` | bool | - | Use TypeScript (`.tsx`) for React files |
| `-i` | string (multi) | - | Include / import files (repeatable) |
| `-pack` | string | `main` | Go package name |
| `-ns` | string (multi) | - | C++ namespaces (repeatable) |

### Examples

```bash
# Generate a C file named "parser"
cbp-gen -c -name parser

# Generate a C++ file with namespaces
cbp-gen -cpp -name engine -ns std -ns mylib

# Generate a Go file in a custom package
cbp-gen -go -name server -pack api

# Generate an HTML page with a custom title
cbp-gen -html -name index -title "My Portfolio"

# Initialize a new React (TypeScript) project
cbp-gen -react -name my-app -ts

# Generate a React component
cbp-gen -react-c -name Navbar -ts
```

---

## Project Roadmap

| Sprint | Status | Goals |
|--------|--------|-------|
| Sprint 01 | ✅ Complete | Language selection, file creation, OS detection, path handling |
| Sprint 02 | ✅ Complete | C template, CLI flags, user customization (name, type) |
| Sprint 03 | 🔲 Planned | Front-end UI, user input fields, function arguments |
| Sprint 04 | 🔲 Planned | UI polish, project website & download page |

---

## Project Goals

- Runs on **macOS, Linux, and Windows**
- Built with **Go** for speed and zero-dependency binaries
- Supports **user customization** (names, types, includes, namespaces)
- Scales easily to new languages over time
- Designed to eventually include a **React front-end**

---

## License

MIT - see [LICENSE](LICENSE) for details.
