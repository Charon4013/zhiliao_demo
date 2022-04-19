/**
 * @Author: Pan
 * @Date: 2022/3/27 16:39
 */

package model

import "time"

type Question struct {
	Id          int64
	Uid         int64     `xorm:"not null" json:"uid"`
	Title       string    `xorm:"unique not null" json:"title"`
	Description string    `xorm:"text" json:"description"`
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
	Deleted     time.Time `xorm:"deleted"`
}

type QuestionStringUid struct {
	Id          int64
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     time.Time
	Updated     time.Time
	Deleted     time.Time
}

type QuestionStringIdUid struct {
	Id          string
	Uid         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     time.Time
	Updated     time.Time
	Deleted     time.Time
}
