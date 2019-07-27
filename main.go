package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func usage() {
	msg := `

[usage]
キーバリュー形式で文字列情報を管理するコマンドです。
以下のサブコマンドが利用可能です。

list   ... 保存済みの内容を一覧表示します。
save   ... keyとvalueを渡して保存します。
get    ... keyを渡してvalueを表示します。
remove ... keyを渡してvalueを削除します。
help   ... ヘルプ情報（当内容と同じ）を表示します。

`
	println(msg)
}

// -------------------------------------------------------------------
// ここからメイン処理
// -------------------------------------------------------------------

var cmdStore = map[string]string{}

func main() {
	println("Start!")
	for {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {

			// アプリ終了判定
			if s.Text() == "end" {
				println("End!")
				os.Exit(-1)
			}

			// ヘルプ
			if s.Text() == "help" {
				usage()
			}

			if s.Text() == "" {
				usage()
				continue
			}

			// 以降は、引数ありコマンドの処理
			cmds := strings.Split(s.Text(), " ")

			// 保存
			if cmds[0] == "save" {
				if len(cmds) != 3 {
					usage()
					continue
				}
				cmdStore[cmds[1]] = cmds[2]
			}

			// 取得
			if cmds[0] == "get" {
				if len(cmds) != 2 {
					usage()
					continue
				}
				println(cmdStore[cmds[1]])
			}

			// 削除
			if cmds[0] == "remove" {
				if len(cmds) != 2 {
					usage()
					continue
				}
				delete(cmdStore, cmds[1])
			}

			// 一覧
			if cmds[0] == "list" {
				println(`"key","value"`)
				for k, v := range cmdStore {
					println("\"" + k + "\",\"" + v + "\"")
				}
			}

		}
		if s.Err() != nil {
			log.Fatal(s.Err())
		}
	}
}
