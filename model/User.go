/**
 * @Author: Pan
 * @Date: 2022/2/25 17:50
 */

package model

import "time"

type User struct {
	Id       int64
	Username string    `xorm:"unique not null" json:"username"`
	Email    string    `xorm:"unique not null" json:"email"`
	Password string    `xorm:"not null varchar(100)" json:"password"`
	Token    string    `json:"token"`
	Manage   bool      `xorm:"default false" json:"manage"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
	Deleted  time.Time `xorm:"deleted"`
}
