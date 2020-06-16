package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//创建
func (this *User) Create() error {
	err := db.Create(this).Error
	return err
}

//检索
func (this User) GetUserByUsername(username string) (user User) {

	db.First(&user, username)

	return
}
