package prototype

type Prototype interface {
	Clone() Prototype
}

type User struct {
	Name string
	Age  int
}

func (u *User) Clone() Prototype {
	copy := *u
	return &copy
}
