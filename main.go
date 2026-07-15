package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("使い方: todo <コマンド> [引数]")
		fmt.Println("  add <タスク> タスクを追加")
		fmt.Println("  list タスクを一覧表示")
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("タスクを指定してください")
			os.Exit(1)
		}
		fmt.Printf("タスクを追加しました: %s", args[1])
	case "list":
		fmt.Println("タスク一覧 (未実装)")
	default:
		fmt.Printf("不明なコマンド: %s\n", args[0])
		os.Exit(1)
	}
}
