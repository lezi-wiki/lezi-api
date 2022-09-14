package model

import (
	"fmt"
	"github.com/lezi-wiki/lezi-api/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lezi-wiki/lezi-api/pkg/conf"
	"github.com/lezi-wiki/lezi-api/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDatabase() (*gorm.DB, error) {
	var dialect gorm.Dialector

	if gin.Mode() == gin.TestMode {
		dialect = sqlite.Open("file::memory:?cache=shared")
	} else {
		switch conf.DataSourceConfig.Driver {
		case "sqlite3", "sqlite":
			dialect = sqlite.Open(util.RelativePath(conf.DataSourceConfig.File))
		case "mysql", "mariadb":
			dialect = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				conf.DataSourceConfig.Username,
				conf.DataSourceConfig.Password,
				conf.DataSourceConfig.Host,
				conf.DataSourceConfig.Port,
				conf.DataSourceConfig.Database,
			))
		case "postgres", "postgresql":
			dialect = postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
				conf.DataSourceConfig.Host,
				conf.DataSourceConfig.Port,
				conf.DataSourceConfig.Username,
				conf.DataSourceConfig.Database,
				conf.DataSourceConfig.Password,
				conf.DataSourceConfig.SSLMode,
				time.Local.String(),
			))
		case "mssql", "sqlserver":
			dialect = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
				conf.DataSourceConfig.Username,
				conf.DataSourceConfig.Password,
				conf.DataSourceConfig.Host,
				conf.DataSourceConfig.Port,
				conf.DataSourceConfig.Database,
			))
		default:
			log.Log().Panicf("不支持的数据库驱动 %s", conf.DataSourceConfig.Driver)
		}
	}

	logLevel := logger.Silent

	// Debug模式下，输出所有 SQL 日志
	if conf.SystemConfig.Debug {
		logLevel = logger.Info
	}

	database, err := gorm.Open(dialect, &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(new(log.GormLogger), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.DataSourceConfig.Prefix,
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return database, nil
}

func Init() {
	db, err := NewDatabase()

	if err != nil {
		log.Log().Panicf("数据库连接失败: %s", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Log().Panicf("数据库连接失败: %s", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	Client = NewDataClient(db)

	// 自动迁移
	migrate(db)
}
