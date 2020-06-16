package model

type Fund struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Balance  string `json:"balance"`
	Username string `json:"username"`
}

//创建
func (this *Fund) Add() error {
	err := db.Create(this).Error
	return err
}

//检索
func (this Fund) GetFundById(id int) (fund Fund) {

	db.First(&fund, id)

	return
}

//通过username检索全部
func (this Fund) GetAllFundsByUsername(username string) (funds []Fund) {

	db.Where("username = ?", username).Find(&funds)

	return
}
