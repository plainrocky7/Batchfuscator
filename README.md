# Batchfuscator

**Batchfuscator** is a simple batch file obfuscator written in Go.  
It’s a personal learning project, so it’s not very polished or production-ready.  
The main goal is to practice Go programming and file manipulation.

---

## Features

- Reads an existing `.bat` file.
- Appends or prepends random `set` commands to the file.
- Generates simple random strings for obfuscation.
- Helps learn about:
  - File I/O in Go (`os.OpenFile`, `os.ReadFile`, `os.WriteString`)
  - Random string generation (`math/rand`)
  - Command-line arguments (`os.Args`)
  - Error handling in Go

---

## Usage

```bash
go run main.go <file.bat>
```
## License
This project is licensed under the MIT License. See Here For Details [MIT LICSENE](https://opensource.org/license/mit).
