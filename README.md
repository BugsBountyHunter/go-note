# Go Note

A command-line note-taking app written in Go. You enter a title and content, it validates them, prints the note, and saves it as a JSON file. It's a small, focused exercise in **structs, methods, packages, struct tags and error handling** in Go.

![Go](https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

## Features

- Reads multi-word input from the console with `bufio.Reader`, handling both Unix and Windows line endings.
- Validates input in a constructor (`note.New`), which returns an error if the title or content is empty.
- Serializes to JSON with struct tags (`title`, `content`, `created_at`).
- Derives the file name from the title, for example `Learn Go` → `learn_go.json`.

## Project structure

```
.
├── main.go        # input handling and program flow
└── note/
    └── note.go    # Note struct, New(), Display(), Save()
```

## Usage

```bash
go run .
```

```
Note Title: Learn Go
Note Content: Structs, methods and interfaces

Your note title: Learn Go has following content:
...
Successfully save json file.
```

This produces `learn_go.json`:

```json
{
  "title": "Learn Go",
  "content": "Structs, methods and interfaces",
  "created_at": "2024-01-20T22:05:57Z"
}
```

## Possible next steps

- [ ] List and read saved notes
- [ ] Add a todo type that shares a `Saver` interface with notes
- [ ] Unit tests for `note.New` validation and file naming
