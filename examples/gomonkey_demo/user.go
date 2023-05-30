package gomonkey_demo

type MyUser struct {
	name string
}

// 返回用户的name
func (s *MyUser) GetUserName() string {
	return s.name
}
