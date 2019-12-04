package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_DRIVER = "zbcg:fGg!WhQ8KWgKB@@tcp(rm-uf690v9wug7u7z877po.mysql.rds.aliyuncs.com:3306)/zbcg?charset=utf8"
)

var Db *gorm.DB

type  Customer struct {
	Id 				int64	`json:"id"`
	Nickname		string	`json:"nickname"`
	Password		string  `json:"password"`
	Clearpassword	string 	`json:"clearpassword"`
	Salesman		string	`json:"salesman"`
	Company			string	`json:"company"`
	Address 		string	`json:"address"`
	Phone			string 	`json:"phone"`
	Fax				string	`json:"fax"`
	Email			string	`json:"email"`
	PhoneMsg		string	`json:"phone_msg"`
	MainProducts	string	`json:"main_products"`
	Corporation		string	`json:"corporation"`
}

func init() {
	db, err := gorm.Open("mysql", DB_DRIVER)
	if err != nil {
		panic("MYSQL CONNECT FAIL:"+err.Error())
	}
	Db = db
}

//查询数据并返回
func GetData(page,size int) []Customer {
	var users []Customer
	field := "id,nickname,password,clearpassword,salesman,company,address,phone,fax,email,phone_msg,main_products,corporation";
	err := Db.Table("customer").Where("email != '' or phone != ''").Select(field).Limit(size).Offset((page - 1) * size).Find(&users).Error
	if err != nil {
		fmt.Println(err)
	}
	return users
}

