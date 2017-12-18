package mem
import(
	_ "fmt"
	"errors"
	)
type member struct{}
//错误定义
var (
ErrEmpty = errors.New("empty string")
)
func (member) Remark() error {}
func (member) Login(uid, pwd string) (bool, error) {}
func (member) Logout(uid string) (bool, error) {}
