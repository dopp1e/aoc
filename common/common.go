package common

import (
    "os"
    "log"
    "bufio"
)

func ReadFile(path string) []string {
    a := make([]string, 1)

    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s := scanner.Text()
        a = append(a, s)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return a
}
