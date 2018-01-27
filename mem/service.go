package mem

import (
	"errors"
	_ "fmt"
)

type member struct{}

//错误定义
var (
	ErrEmpty = errors.New("empty string")
)

func (member) Login(uid, pwd string) (bool, error) {}
func (member) Logout(uid string) (bool, error)     {}
func (member) Remark() error                       {}
