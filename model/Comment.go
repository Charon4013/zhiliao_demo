/**
 * @Author: Pan
 * @Date: 2022/4/12 12:54
 */

package model

import "time"

type Comment struct {
	Id      int64
	Aid     int64     `xorm:"not null" json:"aid"`
	Uid     int64     `xorm:"not null" json:"uid"`
	Content string    `xorm:"not null" json:"content"`
	Support int64     `xorm:"default 0" json:"support"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

type CommentStringAidAndUid struct {
	Id      int64
	Aid     string `json:"aid"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Support int64  `json:"support"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}
