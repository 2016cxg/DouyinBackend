package dao

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//定义video模型，绑定video表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql video表
type Video struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	UserId        int64  `gorm:"column:userid"`
	Title         string `gorm:"column:title"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (v Video) TableName() string {
	//绑定MYSQL表名为users
	return "videos"
}

type VideoDAO struct {
	db *gorm.DB
}

func NewVideoDAO() *VideoDAO {
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

	return &VideoDAO{db: db}
}

func (this *VideoDAO) Insert(u *Video) error {
	//插入一条用户数据
	if err := this.db.Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return err
	}
	return nil
}

func (this *VideoDAO) QueryVideoByUID(id int64) []Video {
	u := []Video{}
	result := this.db.Where("userid = ?", id).Find(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
		//panic("Record not found")
	}
	return u
}

func (this *VideoDAO) QueryLast30Videos() []Video {
	u := []Video{}
	result := this.db.Limit(30).Order("id desc").Find(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return u
}
