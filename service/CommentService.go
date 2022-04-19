/**
 * @Author: Pan
 * @Date: 2022/4/12 14:44
 */

package service

import (
	"zhiliao_mvc_demo/model"
	"zhiliao_mvc_demo/repo"
)

type CommentService interface {
	Create(c model.Comment) (result model.Result)
	Delete(c int64) (result model.Result)
	SearchAllAnswerCommentByAid(aid int64) (result model.Result)
	GetCommentCountByAid(aid int64) (result model.Result)
}

type commentService struct{}

func NewCommentService() CommentService {
	return &commentService{}
}

var commentRepo = repo.NewCommentRepo()

func (cs commentService) Create(c model.Comment) (result model.Result) {
	errCode, data := commentRepo.CreateComment(c)

	if errCode != 0 {
		result.Code = errCode
		result.Msg = "Create fail"
		return
	}
	result.Code = 0
	result.Msg = "Create success"
	result.Data = data
	return
}

func (cs commentService) Delete(c int64) (result model.Result) {
	errCode := commentRepo.DeleteComment(c)

	if errCode {
		result.Code = -1
		result.Msg = "Delete fail"
		return
	}
	result.Code = 0
	result.Msg = "Delete success"
	return
}

func (cs commentService) SearchAllAnswerCommentByAid(aid int64) (result model.Result) {
	errCode, data := commentRepo.GetAllCommentByAid(aid)

	if errCode != 0 || len(data) == 0 {
		result.Code = errCode
		result.Msg = "SearchAllAnswerCommentByAid fail or len = 0"
		return
	}

	//fmt.Println("service data:", data, len(data))
	result.Code = 0
	result.Msg = "SearchAllAnswerCommentByAid success"
	result.Data = data
	return
}

func (cs commentService) GetCommentCountByAid(aid int64) (result model.Result) {
	result.Data = commentRepo.CountCommentByAid(aid)
	return
}
