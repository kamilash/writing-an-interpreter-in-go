package parser

import (
	"mycode/lexer"
	"testing"
)

func BenchmarkLetStatements(b *testing.B) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
		`

    for n := 0; n < b.N; n++ {
        l := lexer.New(input)
		p := New(l)
		p.ParseProgram()
	}
}
