package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
	"github.com/stevenkitter/skills/common"
	pb "github.com/stevenkitter/skills/common/proto"
	"time"
)

type User struct {
	gorm.Model
	UserID       string `gorm:"unique_index"`
	Email        string `gorm:"unique"`
	Phone        string `gorm:"unique"`
	Password     string
	Address      string
	Lng          float64 `gorm:"type:DECIMAL(11,8)"`
	Lat          float64 `gorm:"type:DECIMAL(10,8)"`
	AreaCode     string
	Birthday     *time.Time
	MemberNumber string //show for everyone
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Updates(map[string]interface{}{
		"user_id":       xid.New().String(),
		"member_number": fmt.Sprintf("%d", time.Now().Unix()+int64(u.ID)),
	})
	return
}

//CreateUser For Sign up
func (db *Database) CreateUser(email, password string) *pb.Resp {
	user := &User{}
	err := db.Mariadb.Where(&User{Email: email}).First(user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return common.WrongServerResp(err.Error())
	}
	if user.ID != 0 {
		return common.WrongClientMsgResp("邮箱已注册，若忘记密码可找回", "")
	}
	newUser := &User{Email: email, Password: password}
	err = db.Mariadb.Create(newUser).Error
	if err != nil {
		return common.WrongServerResp(err.Error())
	}
	return common.OkMsgResp("新建成功,请验证邮箱", "")
}
