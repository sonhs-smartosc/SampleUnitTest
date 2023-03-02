package User

import "mocktest/IUser"

type User struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "new")
}
