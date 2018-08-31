package pkgs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"reboot/pkg/dao"
	"reboot/pkg/dao/mysql"
	"sync"
)

const (
	DaoConnection = "dao.connection"
)

var (
	mysqlDao  dao.Storage
	mysqlOnce sync.Once
)

func init() {
	initDaoDefault()
}

func initDaoDefault() {
	viper.SetDefault(DaoConnection, "root:root@/reboot?charset=utf8&parseTime=true")
}

func GetDao() dao.Storage {
	mysqlOnce.Do(func() {
		mysqlDao = mysql.New(&mysql.Options{
			DbConnStr: viper.GetString(DaoConnection),
		})
		if mysqlDao == nil {
			panic("connect mysql failed")
		}
	})
	return mysqlDao

}
