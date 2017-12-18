package lexer

import (
	"testing"
)

func TestMap(t *testing.T) {
	mysql := new(MysqlLexer)
	mysql.TableName()

}
