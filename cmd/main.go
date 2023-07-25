package main

import (
    "os"
    "kbimg/cmd/internal"
)

// Go 程序的默认入口函数(主函数).
func main() {
    command := kbimg.ApplyUserConfiguration()
    if err := command.Execute(); err != nil {
        os.Exit(1)
    }
}
