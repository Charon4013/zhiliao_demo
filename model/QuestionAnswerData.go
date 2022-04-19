/**
 * @Author: Pan
 * @Date: 2022/4/9 15:59
 */

package model

type QuestionAnswerDataList struct {
	Question Question `json:"question"`
	Answer   []Answer `json:"answer"`
}

type QuestionAnswerData struct {
	Question Question `json:"question"`
	Answer   Answer   `json:"answer"`
}

type QuestionAnswerSupportCommentCount struct {
	Question     Question `json:"question"`
	Answer       Answer   `json:"answer"`
	SupportCount int64    `json:"support_count"`
	CommentCount int64    `json:"comment_count"`
}
