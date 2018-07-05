package model

import (
	"fmt"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 数据库初始化预配置
// 可以在Database结构体上添加更多的数据库实例配置，可以实现读写分离
type Database struct {
	Self *gorm.DB
	// Docker *gorm.DB
}

// 全局实例
var DB *Database

// 开启DB实例
func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	setupDB(db) // 设置DB实例的配置，线程池，长连接等等...
	return db
}

// 设置DB实例的配置，线程池，长连接等等...
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog")) // 设置数据库的日志记录设置
	db.DB().SetMaxOpenConns(2000)        // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(10)          // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// 使用本地普通的Mysql Cli访问
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

// 获取本机的DB实例
func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

// 使用Docker数据库进行访问
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

// 获取Docker的db实例
func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

// 全局DB初始化
func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		// Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}
