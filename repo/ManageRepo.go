/**
 * @Author: Pan
 * @Date: 2022/4/14 21:05
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type ManageRepo interface {
	GetAllUser() (int, []model.User)
	DeleteUserByUid(uid int64) bool
	GetAllUserInfo() (int, []model.UserInfo)
	DeleteUserInfoById(uid int64) bool
	GetAllQuestion() (int, []model.Question)
	DeleteQuestionByQid(qid int64) bool
	GetAllAnswer() (int, []model.Answer)
	DeleteAnswerByAid(aid int64) bool
	GetAllComment() (int, []model.Comment)
	DeleteCommentByCid(cid int64) bool
	CheckUserManageIdentityBy(uid int64) bool
}

type manageRepo struct{}

func NewManageRepo() ManageRepo {
	return &manageRepo{}
}

func (mr manageRepo) GetAllUser() (int, []model.User) {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	var userList []model.User
	err := userRepoEngine.Table("user").Omit("password", "token", "manage").Find(&userList)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, nil
	}
	return 0, userList
}

func (mr manageRepo) DeleteUserByUid(uid int64) bool {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	var user model.User
	affected, err := userRepoEngine.Table("user").ID(uid).Delete(&user)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func (mr manageRepo) GetAllUserInfo() (int, []model.UserInfo) {
	var userInfoEngine = datasource.Init("UserInfo")
	defer userInfoEngine.Close()

	var userInfoList []model.UserInfo
	err := userInfoEngine.Table("user_info").Find(&userInfoList)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, nil
	}
	return 0, userInfoList
}

func (mr manageRepo) DeleteUserInfoById(uid int64) bool {
	var userInfoEngine = datasource.Init("UserInfo")
	defer userInfoEngine.Close()

	var userInfo model.UserInfo
	affected, err := userInfoEngine.Table("user_info").ID(uid).Delete(&userInfo)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func (mr manageRepo) GetAllQuestion() (int, []model.Question) {
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	var questionList []model.Question
	err := questionEngine.Table("question").Find(&questionList)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, nil
	}
	return 0, questionList
}

func (mr manageRepo) DeleteQuestionByQid(qid int64) bool {
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	var question model.Question
	affected, err := questionEngine.Table("question").ID(qid).Delete(&question)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func (mr manageRepo) GetAllAnswer() (int, []model.Answer) {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	var answerList []model.Answer
	err := answerEngine.Table("answer").Find(&answerList)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, nil
	}
	return 0, answerList
}

func (mr manageRepo) DeleteAnswerByAid(aid int64) bool {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	var answer model.Answer
	affected, err := answerEngine.Table("answer").ID(aid).Delete(&answer)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func (mr manageRepo) GetAllComment() (int, []model.Comment) {
	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	var commentList []model.Comment
	err := commentEngine.Table("comment").Find(&commentList)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, nil
	}
	return 0, commentList
}

func (mr manageRepo) DeleteCommentByCid(cid int64) bool {
	var commentEngine = datasource.Init("Comment")
	defer commentEngine.Close()

	var comment model.Comment
	affected, err := commentEngine.Table("comment").ID(cid).Delete(&comment)
	if err != nil || affected == 0 {
		return false
	}
	return true
}

func (mr manageRepo) CheckUserManageIdentityBy(uid int64) bool {
	var userRepoEngine = datasource.Init("UserRepo")
	defer userRepoEngine.Close()

	isManage, err := userRepoEngine.Table("user").ID(uid).Where("manage = ?", true).Exist()
	if err != nil {
		fmt.Println("sql err: ", err)
		return false
	}
	return isManage
}
