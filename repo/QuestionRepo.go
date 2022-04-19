/**
 * @Author: Pan
 * @Date: 2022/3/27 16:42
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

type QuestionRepo interface {
	CreateQuestion(q model.Question) (int, model.Question)
	DeleteQuestion(qid int64) bool
	GetQuestionByQid(qid int64) (int, model.Question)
	GetQuestionOneByTitle(str string) (int, model.Question)
	GetQuestionListByTitle(str string) (int, []model.Question)
	ModifyQuestion(q model.Question) (int, model.Question)
	GetQuestionListByUid(uid int64) (int, []model.Question)
	GetQuestionListByRandom(num int) (int, []model.Question)
}

type questionRepo struct{}

func NewQuestionRepo() QuestionRepo {
	return &questionRepo{}
}

func (qr questionRepo) CreateQuestion(q model.Question) (int, model.Question) {
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	_, err := questionEngine.Table("question").Insert(q)
	if err != nil {
		fmt.Println("Insert error: ", err)
		return -1, model.Question{}
	}

	errCode, q := qr.GetQuestionOneByTitle(q.Title)
	if errCode != 0 {
		fmt.Println("Get Obj After Insert error: ", err)
		return -1, model.Question{}
	}

	return 0, q
}

func (qr questionRepo) DeleteQuestion(qid int64) bool {
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	var question model.Question
	_, err := questionEngine.Table("question").ID(qid).Delete(&question)
	if err != nil {
		return false
	}
	return true
}

func (qr questionRepo) GetQuestionByQid(qid int64) (int, model.Question) {

	// 坑
	var question model.Question

	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	// &question会缓存，要新建一个对象
	_, err := questionEngine.Table("question").ID(qid).Get(&question)
	//fmt.Println("qid & question :", qid, question)
	if err != nil {
		fmt.Println("err: ", err)
		return -1, model.Question{}
	}
	if question.Id == 0 {
		fmt.Println("question.Id err: ", question.Id)
		return -1, model.Question{}
	}

	return 0, question
}

func (qr questionRepo) GetQuestionOneByTitle(str string) (int, model.Question) {
	var question model.Question
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	_, err := questionEngine.Table("question").Where("title = ?", str).Get(&question)
	if err != nil {
		return -1, model.Question{}
	}
	return 0, question
}

func (qr questionRepo) GetQuestionListByTitle(str string) (int, []model.Question) {
	var questionList []model.Question
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	err := questionEngine.Table("question").Where("title like ?", "%"+str+"%").Find(&questionList)
	if err != nil {
		return -1, nil
	}
	return 0, questionList
}

func (qr questionRepo) GetQuestionListByUid(uid int64) (int, []model.Question) {
	var questionList []model.Question
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	err := questionEngine.Table("question").Where("uid = ?", uid).Find(&questionList)
	if err != nil {
		return -1, nil
	}
	return 0, questionList
}

func (qr questionRepo) ModifyQuestion(q model.Question) (int, model.Question) {
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()

	_, err := questionEngine.Table("question").ID(q.Id).Update(q)
	if err != nil {
		fmt.Println("sql err: ", err)
		return -1, model.Question{}
	}

	return 0, q
}

func (qr questionRepo) GetQuestionListByRandom(num int) (int, []model.Question) {
	var questionList []model.Question
	var questionEngine = datasource.Init("Question")
	defer questionEngine.Close()
	sql := "select * from question where id >= ((select max(id) from question) - (select min(id) from question)) * rand() + (select min(id) from question)" + " limit 10"
	//fmt.Println("sql: ",sql)
	err := questionEngine.Table("question").SQL(sql).Find(&questionList)

	if err != nil {
		fmt.Println("sql err", err)
		return -1, nil
	}
	if len(questionList) == 0 {
		fmt.Println("No result yet: ", len(questionList))
		return -1, nil
	}

	return 0, questionList
}
