package routes

import (
	"github.com/jartek/goapi/models"
	"github.com/lunny/xorm"
)

func Init() {
	var orm *xorm.Engine
	orm = models.InitOrm()
}
