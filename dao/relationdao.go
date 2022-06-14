package dao

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql users表
type Relation struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	//这里golang定义的Username变量和MYSQL表字段username一样，他们的名字可以不一样。
	Fuid     int64 `gorm:"column:fuid"`
	Tuid     int64 `gorm:"column:tuid"`
	Relation int64 `gorm:"column:relation"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u Relation) TableName() string {
	//绑定MYSQL表名为users
	return "relation"
}

type RelationDAO struct {
	db *gorm.DB
}

func NewRelationDao() *RelationDAO {
	//配置MySQL连接参数
	username := "root"     //账号
	password := "cheng"    //密码
	host := "127.0.0.1"    //数据库地址，可以是Ip或者域名
	port := 3306           //数据库端口
	Dbname := "mydatabase" //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	return &RelationDAO{db: db}
}

func (this *RelationDAO) Insert(u Relation) error {
	//插入一条用户数据
	if err := this.db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}

// select record according to (Fuid, Tuid)
func (this *RelationDAO) SelFuidTuid(fuid int64, tuid int64) *Relation {
	u := Relation{}
	//result := this.db.Where("fuid = ?", fuid).First(&u)

	result := this.db.Where("fuid = ? AND tuid = ?", fuid, tuid).First(&u)
	// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
		//panic("Record not found")
	}
	return &u
}

//update relation table
func (this *RelationDAO) UpRelation(fuid int64, tuid int64, relation int64) error {
	result := this.db.Model(&Relation{}).Where("fuid = ? AND tuid = ?", fuid, tuid).Update("relation", relation)
	if result.Error != nil {
		fmt.Println("update error ")
		return result.Error
	}
	return nil
}

func (this *RelationDAO) SelFollowLst(uid int64) []int64 {
	u := []int64{}

	u1 := []Relation{}
	this.db.Where("fuid = ? AND ( relation=? OR relation=?) ", uid, 1, 3).Find(&u1)
	for i := range u1 {
		tuid := u1[i].Tuid
		u = append(u, tuid)
	}

	u2 := []Relation{}
	this.db.Where("tuid = ? AND ( relation=? OR relation=?) ", uid, 2, 3).Find(&u2)
	for i := range u2 {
		fuid := u2[i].Fuid
		u = append(u, fuid)
	}

	return u
}

func (this *RelationDAO) SelFollowerLst(uid int64) []int64 {
	u := []int64{}

	u1 := []Relation{}
	this.db.Where("fuid = ? AND ( relation=? OR relation=?) ", uid, 2, 3).Find(&u1)
	for i := range u1 {
		tuid := u1[i].Tuid
		u = append(u, tuid)
	}

	u2 := []Relation{}
	this.db.Where("tuid = ? AND ( relation=? OR relation=?) ", uid, 1, 3).Find(&u2)
	for i := range u2 {
		fuid := u2[i].Fuid
		u = append(u, fuid)
	}

	return u
}

//follow list
//select * from relation where fuid=uid and ( relation=1 or relation=3 )
//select * from relation where tuid=uid and ( relation=2 or relation=3 )

//follow action
//query record (fuid, tuid) or (tuid, fuid)													--select from relation
//if (fuid, tuid) is not null
//	if record is (fuid, tuid, relation=0), revise relation=1								--update relation
// 	if record is (fuid, tuid, relation=1), return "Already follow" error
//	if record is (fuid, tuid, relation=2), revise relation=3, get them mutual follow
//	if record is (fuid, tuid, relation=3), return "Already follow" error
//if (tuid, fuid) is not null
//	if record is (tuid, fuid, relation=0), revise relation=2
//	if record is (tuid, fuid, relation=1), revise relation=3
//	if record is (tuid, fuid, relation=2), return "Aready follow" error
//	if record is (tuid, fuid, realtion=3), return "Aready follow" error
//if record (fuid, tuid) and (tuid, fuid) are all null
//	insert into relation table (fuid, tuid, relation=1)										--insert into relation

//unfollow action
//

//func (this *UserDAO) QueryUserByName(username string) *User {
//	u := User{}
//	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
//	result := this.db.Where("username = ?", username).First(&u)
//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//		return nil
//		//panic("Record not found")
//	}
//	return &u
//}
//
//func (this *UserDAO) QueryUserByID(id int) *User {
//	u := User{}
//	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
//	result := this.db.Where("id = ?", id).First(&u)
//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//		return nil
//		//panic("Record not found")
//	}
//	return &u
//}
//
//////定义一个用户，并初始化数据
////u := User{
////	Username:"tizi365",
////	Password:"123456",
////	CreateTime:time.Now().Unix(),
////}
//
//////查询并返回第一条数据
//////定义需要保存数据的struct变量
////u = User{}
//////自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
////result := db.Where("username = ?", "tizi365").First(&u)
////if errors.Is(result.Error, gorm.ErrRecordNotFound) {
////	fmt.Println("找不到记录")
////	return
////}
//////打印查询到的数据
////fmt.Println(u.Username,u.Password)
////
////	//更新
////	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
////	db.Model(&User{}).Where("username = ?", "tizi365").Update("password", "654321")
////
////	//删除
////	//自动生成Sql： DELETE FROM `users`  WHERE (username = 'tizi365')
////	db.Where("username = ?", "tizi365").Delete(&User{})
////}
