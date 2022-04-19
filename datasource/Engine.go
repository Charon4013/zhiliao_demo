/**
 * @Author: Pan
 * @Date: 2022/2/25 18:39
 */

package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
	"zhiliao_mvc_demo/model"
)

func Init(tablename string) *xorm.Engine {
	// database password
	//engine, err := xorm.NewEngine("mysql", "root:***database password***@/zhiliao_demo")
	engine, err := xorm.NewEngine("mysql", "root:123456@/zhiliao_demo")
	if err != nil {
		fmt.Println("数据库初始化失败！")
		panic(err)
	}

	engine.SetConnMaxLifetime(time.Minute * 3)
	engine.SetMaxOpenConns(10)
	engine.SetMaxIdleConns(10)
	//engine.ShowSQL(true)
	engine.Logger().SetLevel(log.LOG_DEBUG)

	switch tablename {
	case "User":
		syncErr := engine.Sync2(new(model.User))
		if syncErr != nil {
			fmt.Println("User表初始化失败！")
			panic(syncErr)
		}
	case "UserInfo":
		syncErr := engine.Sync2(new(model.UserInfo))
		if syncErr != nil {
			fmt.Println("UserInfo表初始化失败！")
			panic(syncErr)
		}
	case "Question":
		syncErr := engine.Sync2(new(model.Question))
		if syncErr != nil {
			fmt.Println("Question表初始化失败！")
			panic(syncErr)
		}
	case "Answer":
		syncErr := engine.Sync2(new(model.Answer))
		if syncErr != nil {
			fmt.Println("Answer表初始化失败！")
			panic(syncErr)
		}
	case "AnswerSupport":
		syncErr := engine.Sync2(new(model.AnswerSupport))
		if syncErr != nil {
			fmt.Println("Support表初始化失败！")
			panic(syncErr)
		}
	case "Comment":
		syncErr := engine.Sync2(new(model.Comment))
		if syncErr != nil {
			fmt.Println("Comment表初始化失败！")
			panic(syncErr)
		}
	case "CommentSupport":
		syncErr := engine.Sync2(new(model.CommentSupport))
		if syncErr != nil {
			fmt.Println("CommentSupport表初始化失败！")
			panic(syncErr)
		}

	}

	return engine
}
