/**
 * @Author: Pan
 * @Date: 2022/4/13 15:55
 */

package model

import "time"

type CommentSupport struct {
	Id      int64
	Uid     int64     `xorm:"not null" json:"uid"`
	Cid     int64     `xorm:"not null" json:"cid"`
	Status  bool      `xorm:"default false" json:"status"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

type CommentSupportStringUid struct {
	Id      int64
	Uid     string `json:"uid"`
	Cid     string `json:"cid"`
	Status  bool   `json:"status"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
}
