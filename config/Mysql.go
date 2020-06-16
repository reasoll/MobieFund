package config

import _ "gopkg.in/yaml.v2"

type Mysql struct {
	Addr         string `yaml:"Addr"`
	Port         string `yaml:"Port"`
	UserNmae     string `yaml:"UserName"`
	Password     string `yaml:"Password"`
	DataBaseName string `yaml:"DataBaseName"`
}

func (mysql *Mysql) DefaultMySqlConfig() {
	mysql.Addr = "localhost"
	mysql.Port = "3306"
	mysql.DataBaseName = "fund"
	mysql.UserNmae = "root"
	mysql.Password = "1257983649"
}

func (mysql *Mysql) InitMysqlConfig() {
	mysql.DefaultMySqlConfig()
}
