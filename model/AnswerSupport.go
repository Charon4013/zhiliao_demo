/**
 * @Author: Pan
 * @Date: 2022/4/10 16:33
 */

package model

import "time"

type AnswerSupport struct {
	Id      int64
	Uid     int64     `xorm:"not null" json:"uid"`
	Aid     int64     `xorm:"not null" json:"aid"`
	Status  bool      `xorm:"default false" json:"status"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

type AnswerSupportStringUid struct {
	Id      int64
	Uid     string `json:"uid"`
	Aid     string `json:"aid"`
	Status  bool   `json:"status"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}
