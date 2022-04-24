package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type userInfo struct {
	Id   int    `gorm:"AUTO_INCRMENT"`
	Name string `gorm:"type:char(25); not null;comment:'name'"`
	Age  int    `gorm:"type:int(10); not null;comment:'age'"`
}

func main() {
	db, err := gorm.Open("mysql", "root:231231@(127.0.0.1:3306)/go_base?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println("***db open failed , error:", err)
		return
	}

	fmt.Println(db.AutoMigrate().HasTable(userInfo{}))

	db.AutoMigrate(&userInfo{})
	user01 := userInfo{
		Name: "lishengyi",
		Age:  18,
	}
	user02 := userInfo{
		Name: "zhangsi",
		Age:  14,
	}
	db.Create(&user01)
	db.Create(&user02)
	var res []userInfo
	res = make([]userInfo, 0)
	db.Find(&res)
	//for _, x := range res {
	//	fmt.Printf("%#v\n", x)
	//}
	tmp_u := userInfo{Id: 2}
	db.Model(&tmp_u).Update("age", "33")
	tmp_u_2 := userInfo{Name: "lishengyi"}
	db.Delete(&tmp_u_2, "age=?", "13")
	db.Select("name,age").Find(&res)
	fmt.Println(res)
}
