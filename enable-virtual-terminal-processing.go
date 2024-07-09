package main

import (
    "fmt"
    "os"

    "golang.org/x/sys/windows"
)

const (
    ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
)

func enableVirtualTerminalProcessing() error {
    handle := windows.Handle(os.Stdout.Fd())
    var mode uint32

    // Get the current console mode
    if err := windows.GetConsoleMode(handle, &mode); err != nil {
        return err
    }

    // Enable the virtual terminal processing mode
    mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
    if err := windows.SetConsoleMode(handle, mode); err != nil {
        return err
    }

    return nil
}

func main() {
    // Enable virtual terminal processing
    if err := enableVirtualTerminalProcessing(); err != nil {
        fmt.Println("Error enabling virtual terminal processing:", err)
        return
    }

    // Print colored text
    fmt.Println("\033[31mThis is red text\033[0m")
    fmt.Println("\033[32mThis is green text\033[0m")
    fmt.Println("\033[33mThis is yellow text\033[0m")
    fmt.Println("\033[34mThis is blue text\033[0m")
}
