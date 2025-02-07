package cache

import (
	"github.com/ljinf/template_project_v2/internal/model"
)

func SetUserInfoCache(rdb *redis.Client, user *model.User) error {
	return nil
}

func GetUserInfoCache(rdb *redis.Client, uid string) (*model.User, error) {
	return nil, nil
}
