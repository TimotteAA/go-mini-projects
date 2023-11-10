package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	bookEntity := NewBookEntity()
	
	// 定义命令行的参数
	flag.Parse();
	// 获取非标志参数（子命令）
    args := flag.Args()
    if len(args) < 1 {
        fmt.Println("No command provided")
        os.Exit(1)
    }

    // 根据子命令调用不同的函数
    switch args[0] {
    case "create":
        bookEntity.Create()
    // 可以添加更多的子命令
    default:
        fmt.Printf("Unknown command: %s\n", args[0])
        os.Exit(1)
    }
}