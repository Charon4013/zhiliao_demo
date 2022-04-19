/**
 * @Author: Pan
 * @Date: 2022/3/18 19:53
 */

package model

import "time"

type UserInfo struct {
	Id          int64
	Uid         int64     `xorm:"unique not null" json:"uid"`
	Avatar      string    `json:"avatar"`
	Sex         string    `json:"sex"`
	Description string    `xorm:"varchar(100)" json:"description"`
	Birthday    string    `json:"birthday"`
	Location    string    `json:"location"`
	Created     time.Time `xorm:"created"`
	Updated     time.Time `xorm:"updated"`
	Deleted     time.Time `xorm:"deleted"`
}
