package lexer

type Lexer struct {
	input        string
	position     int  // 入力行における現在の位置
	readPosition int  // これから読み込む位置
	ch           byte // 現在検査中の位置
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
