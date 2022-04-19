/**
 * @Author: Pan
 * @Date: 2022/2/26 14:22
 */

package repo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type UserRepo interface {
	// GetUserByUsername 查询用户名
	GetUserByUsername(username string) (user model.User)
	// GetUserByEmail 查询邮箱
	GetUserByEmail(email string) (user model.User)
	// GetUserById 查询用户ID
	GetUserById(id int64) (user model.User)
	// GetUserByUsernameAndPassword 登录
	GetUserByUsernameAndPassword(username string, password string) (user model.User)
	// ModifyUser 修改
	ModifyUser(user model.User) (int, model.User)
	// CreateUser 新增
	CreateUser(user model.User) (int, model.User)
	// GetUserPasswordHashByUsername 获取数据库密码hash
	GetUserPasswordHashByUsername(username string) (pwdHash string)
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

// GetUserByUsername 查询
func (ur userRepo) GetUserByUsername(username string) (user model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	//fmt.Println("username: ", username)
	get, err := userRepoEngine.Table("User").Where("username = ?", username).Get(&user)
	if err != nil && get {
		fmt.Println("GetUserByUsername error: ", err)
		return model.User{}
	}
	return
}

func (ur userRepo) GetUserPasswordHashByUsername(username string) (pwdHash string) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	_, err := userRepoEngine.Table("User").Where("username = ?", username).Cols("password").Get(&pwdHash)
	if err != nil {
		return ""
	}
	return pwdHash
}

func (ur userRepo) GetUserByEmail(email string) (user model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	userRepoEngine.Table("User").Where("email = ?", email).Get(&user)
	//fmt.Println("GetUserByEmail: ", user)
	return
}

func (ur userRepo) GetUserById(id int64) (user model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	_, err := userRepoEngine.Table("User").ID(id).Get(&user)
	if err != nil {
		return model.User{}
	}
	return user
}

// GetUserByUsernameAndPassword 登录
func (ur userRepo) GetUserByUsernameAndPassword(username string, password string) (user model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	userRepoEngine.Table("User").Where("username = ?", username).And("password = ?", password).Get(&user)
	//fmt.Println(user)
	return
}

// ModifyUser 修改
func (ur userRepo) ModifyUser(user model.User) (int, model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	//user.Salt = util.CreateRandomString()
	errCode := 0
	_, err := userRepoEngine.Table("User").Where("username = ?", user.Username).Update(user)
	if err != nil {
		fmt.Println("Update error: ", err)
		errCode = -1
	}

	return errCode, user
}

func (ur userRepo) CreateUser(user model.User) (int, model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	errCode := 0
	_, err := userRepoEngine.Table("User").Insert(user)
	if err != nil {
		fmt.Println("Insert error: ", err)
		errCode = -1
	}
	userRepoEngine.Table("User").Where("email = ?", user.Email).Get(&user)
	//fmt.Println("233333333333:", user)
	return errCode, user
}
