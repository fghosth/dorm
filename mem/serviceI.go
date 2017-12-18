package mem

type Member interface {
	Login(uid, pwd string) (bool, error)
	Logout(uid string) (bool, error)
	Remark() error
}
