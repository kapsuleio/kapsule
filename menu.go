package main

import (
    "os"
    "bufio"
    "bytes"
    "io"
    "fmt"
)

// Read a whole file and brake it to an array of lines
func readFile(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )

    // Open the file
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    // Read the file content
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

func displayMenu() {
    
    // Load menu file
    lines, err := readFile("menu.txt")

    // Display error if menu file is missing
    if err != nil {
        fmt.Println("Error: %s\n", err)
        return
    }

    // Print menu options
    for _, line := range lines {
        fmt.Println(line)
    }
}
