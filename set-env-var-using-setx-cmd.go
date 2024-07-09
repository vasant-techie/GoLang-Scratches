package main

import (
    "fmt"
    "os/exec"
)

func setEnvVariable(key, value string) error {
    cmd := exec.Command("setx", key, value)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to set environment variable %s: %v, output: %s", key, err, output)
    }
    fmt.Printf("Successfully set %s to %s\n", key, value)
    return nil
}

func main() {
    err := setEnvVariable("AWS_ACCESS_KEY", "your-access-key")
    if err != nil {
        fmt.Println(err)
    }

    err = setEnvVariable("AWS_SECRET_KEY", "your-secret-key")
    if err != nil {
        fmt.Println(err)
    }

    err = setEnvVariable("AWS_SESSION_TOKEN", "your-session-token")
    if err != nil {
        fmt.Println(err)
    }
}
