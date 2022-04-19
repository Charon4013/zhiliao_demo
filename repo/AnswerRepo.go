/**
 * @Author: Pan
 * @Date: 2022/4/8 18:38
 */

package repo

import (
	"fmt"
	"zhiliao_mvc_demo/datasource"
	"zhiliao_mvc_demo/model"
)

var answer model.Answer

type AnswerRepo interface {
	CreateAnswer(a model.Answer) (int, model.Answer)
	DeleteAnswer(aid int64) bool
	GetAnswerByAid(qid int64, aid int64) (int, model.Answer)
	GetAllAnswerByQid(qid int64) (int, []model.Answer)
	GetAnswerListByRandom(num int) (int, []model.Answer)
	GetUserAnswerByUid(uid int64) (int, []model.Answer)
	ModifyAnswer(a model.Answer) (int, model.Answer)
	GiveSupportByAid(aid int64) bool
}

type answerRepo struct{}

func NewAnswerRepo() AnswerRepo {
	return &answerRepo{}
}

func (ar answerRepo) CreateAnswer(a model.Answer) (int, model.Answer) {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	_, err := answerEngine.Table("answer").Insert(a)
	if err != nil {
		fmt.Println("Insert error: ", err)
		return -1, model.Answer{}
	}

	return 0, a
}

func (ar answerRepo) DeleteAnswer(aid int64) bool {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()
	var answer model.Answer
	_, err := answerEngine.Table("answer").ID(aid).Delete(&answer)
	fmt.Println("err: ", err)
	if err != nil {
		return false
	}
	return true
}

func (ar answerRepo) GetAnswerByAid(qid int64, aid int64) (int, model.Answer) {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	//fmt.Println("qid&aid: ", qid, aid)
	var answer model.Answer
	affected, err := answerEngine.Table("answer").Where("Id = ? AND qid = ?", aid, qid).Get(&answer)
	fmt.Println("affected: ", affected)
	fmt.Println("answer :", answer)
	if err != nil || affected == false {
		return -1, model.Answer{}
	}
	return 0, answer
}

func (ar answerRepo) GetAllAnswerByQid(qid int64) (int, []model.Answer) {
	var answerList []model.Answer

	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	err := answerEngine.Table("answer").Where("qid = ?", qid).Find(&answerList)
	if err != nil || len(answerList) == 0 {
		fmt.Println("sql err", err)
		return -1, nil
	}

	return 0, answerList
}

func (ar answerRepo) GetAnswerListByRandom(num int) (int, []model.Answer) {
	var answerList []model.Answer
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()
	sql := "select * from answer where id >= ((select max(id) from answer) - (select min(id) from answer)) * rand() + (select min(id) from answer)" + " limit 10"

	err := answerEngine.Table("answer").SQL(sql).Find(&answerList)

	if err != nil {
		fmt.Println("sql err", err)
		return -1, nil
	}
	if len(answerList) == 0 {
		fmt.Println("No result yet: ", len(answerList))
		return -1, nil
	}

	return 0, answerList
}

// GetUserAnswerByUid 获得用户所有回答
func (ar answerRepo) GetUserAnswerByUid(uid int64) (int, []model.Answer) {
	var answerList []model.Answer

	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	err := answerEngine.Table("answer").Where("uid = ?", uid).Find(&answerList)

	if err != nil || len(answerList) == 0 {
		fmt.Println("sql err or answerList = 0: ", err)
		return -1, nil
	}

	return 0, answerList
}

func (ar answerRepo) ModifyAnswer(a model.Answer) (int, model.Answer) {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	_, err := answerEngine.Table("answer").ID(a.Id).Update(a)
	if err != nil {
		return -1, model.Answer{}
	}

	return 0, a
}

func (ar answerRepo) GiveSupportByAid(aid int64) bool {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	answer := new(model.Answer)
	answer.Support = answer.Support + 1

	_, err := answerEngine.Table("answer").ID(aid).Cols("support").Update(&answer)
	if err != nil {
		fmt.Println("Give Support fail")
		return false
	}
	return true
}

func (ar answerRepo) CancelSupportByAid(aid int64) bool {
	var answerEngine = datasource.Init("Answer")
	defer answerEngine.Close()

	answer := new(model.Answer)
	answer.Support = answer.Support - 1

	_, err := answerEngine.Table("answer").ID(aid).Cols("support").Update(&answer)

	if err != nil {
		fmt.Println("Cancel Support fail")
		return false
	}
	return true
}
