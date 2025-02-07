package repository

import (
	"context"
	"fmt"
	"github.com/ljinf/template_project_v2/pkg/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const ctxTxKey = "TxKey"

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		db:  db,
		rdb: rdb,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// MasterDB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewDB(conf *viper.Viper) *gorm.DB {
	db, err := gorm.Open(
		getDialector("mysql", conf.GetString("data.mysql.master.dsn")),
		&gorm.Config{Logger: log.NewGormLogger()})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(conf.GetInt("data.mysql.master.max_open"))
	sqlDb.SetMaxIdleConns(conf.GetInt("data.mysql.master.max_idle"))
	sqlDb.SetConnMaxLifetime(time.Duration(conf.GetInt("data.mysql.master.max_life_time")))
	if err = sqlDb.Ping(); err != nil {
		panic(err)
	}

	return db
}

func getDialector(t, dsn string) gorm.Dialector {
	//switch t { 项目数据库需要加载多数据源时去掉注释
	//case "postgres":
	//	return postgres.Open(dsn)
	//default:
	//	return mysql.Open(dsn)
	//}
	return mysql.Open(dsn)
}

type Cache interface {
	Redis() *redis.Client
}

func NewRedis(conf *viper.Viper) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.GetString("data.redis.addr"),
		Password:     conf.GetString("data.redis.password"),
		DB:           conf.GetInt("data.redis.db"),
		ReadTimeout:  time.Duration(conf.GetInt("data.redis.read_timeout")) * time.Millisecond,
		WriteTimeout: time.Duration(conf.GetInt("data.redis.write_timeout")) * time.Millisecond,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}

func (r *Repository) Redis() *redis.Client {
	return r.rdb
}
