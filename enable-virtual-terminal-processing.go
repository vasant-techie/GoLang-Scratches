package main

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
)

const (
    ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
    STD_OUTPUT_HANDLE                  = -11 & 0xFFFFFFFF
)

func enableVirtualTerminalProcessing() error {
    handle := syscall.Handle(os.Stdout.Fd())
    var mode uint32

    // Get the current console mode
    r, _, e := syscall.Syscall(syscall.GetProcAddress(syscall.LoadLibrary("kernel32.dll"), "GetConsoleMode"), 2, uintptr(handle), uintptr(unsafe.Pointer(&mode)), 0)
    if r == 0 {
        return e
    }

    // Enable the virtual terminal processing mode
    mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
    r, _, e = syscall.Syscall(syscall.GetProcAddress(syscall.LoadLibrary("kernel32.dll"), "SetConsoleMode"), 2, uintptr(handle), uintptr(mode), 0)
    if r == 0 {
        return e
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
