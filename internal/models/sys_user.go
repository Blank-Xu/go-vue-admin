package models

type SysUser struct {
	ID                 int    `json:"id" xorm:"pk autoincr 'id'"`
	UserName           string `json:"user_name" xorm:"unique 'user_name'"`
	NickName           string `json:"nick_name" xorm:"nick_name"`
	Password           string `json:"-" xorm:"password"`
	Salt               string `json:"-" xorm:"salt"`
	State              int    `json:"state" xorm:"state"`
	Avatar             string `json:"avatar" xorm:"avatar"`
	Email              string `json:"email" xorm:"email"`
	Phone              string `json:"phone" xorm:"phone"`
	Remark             string `json:"remark" xorm:"remark"`
	ModifyPasswordTime int64  `json:"modify_password_time" xorm:"modify_password_time"`
	LastLoginTime      int64  `json:"last_login_time" xorm:"last_login_time"`
	LastLoginIP        string `json:"last_login_ip" xorm:"last_login_ip"`
	Updated            int64  `json:"updated" xorm:"updated"`
	Created            int64  `json:"created" xorm:"created"`
	Deleted            int64  `json:"deleted" xorm:"deleted"`
}
