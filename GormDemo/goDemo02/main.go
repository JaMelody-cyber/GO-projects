package main

//外键查询
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Students struct {
	Id   string `gorm:"type:char(10);not null;primary_key;"`
	Name string `gorm:"type:char(25);not null;"`
}
type Grades struct {
	StudentsId string    `gorm:"type:char(10);not null;primary_key;"`
	student    *Students `gorm:"ForeignKey:StudentsId;"`
	Lesson     string    `gorm:"not null;"`
	Grade      int       `gorm:"type:int(3);not null;"`
}

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:231231@(127.0.0.1:3306)/go_base?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("***db open failed , error:", err)
		return
	}
	//db.CreateTable(&Students{}, &Grades{})
}

func main() {
	var grades []Grades
	db.Table("grades").Joins("join students on grades.students_id = students.id").Find(&grades)
	for _, x := range grades {
		fmt.Println(x)
	}
}

func initialTB(db *gorm.DB) {
	db.AutoMigrate(&Students{})
	s1 := Students{Id: "1932978", Name: "lishengyi"}
	s2 := Students{Id: "1932932", Name: "zhangsan"}
	db.Create(s1)
	db.Create(s2)
	db.AutoMigrate(&Grades{})
	g1 := Grades{StudentsId: "1932978", student: &s1, Lesson: "Math", Grade: 96}
	g2 := Grades{StudentsId: "1932978", student: &s1, Lesson: "English", Grade: 100}
	g3 := Grades{StudentsId: "1932932", student: &s2, Lesson: "Math", Grade: 65}
	g4 := Grades{StudentsId: "1932932", student: &s2, Lesson: "History", Grade: 72}
	db.Create(g1)
	db.Create(g2)
	db.Create(g3)
	db.Create(g4)
}
