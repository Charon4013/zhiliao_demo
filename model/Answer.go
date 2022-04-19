/**
 * @Author: Pan
 * @Date: 2022/4/8 18:34
 */

package model

import "time"

type Answer struct {
	Id      int64
	Qid     int64     `xorm:"not null" json:"qid"`
	Uid     int64     `xorm:"not null" json:"uid"`
	Content string    `xorm:"not null text" json:"content"`
	Support int64     `xorm:"default 0" json:"support"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

type AnswerStringUidAndQid struct {
	Id      int64
	Qid     string `json:"qid"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Support int64  `json:"support"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}

type AnswerStringIdUidQid struct {
	Id      string
	Qid     string `json:"qid"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Support int64  `json:"support"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}
