package main

func NewHelpCommand() Command {
	return &helpCommand{}
}

type helpCommand struct {
}

func (c *helpCommand) Exec(args []string) {
	msg := `

[usage]
キーバリュー形式で文字列情報を管理するコマンドです。
以下のサブコマンドが利用可能です。

list   ... 保存済みの内容を一覧表示します。
save   ... keyとvalueを渡して保存します。
get    ... keyを渡してvalueを表示します。
remove ... keyを渡してvalueを削除します。
endCommand   ... ヘルプ情報（当内容と同じ）を表示します。

`
	println(msg)
}