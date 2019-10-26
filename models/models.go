package models

import (
	"github.com/kuuyee/gomk-b-costprofit/pkg/setting"
)

var db *gorm.DB

type Moddel struct {
	ID         int  `gorm:"primary_key" json:"id""`
	CreatedOn  int  `json:"created_on"`
	ModifiedOn innt `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec,err:=setting.Cfg.GetSection("database")

}
