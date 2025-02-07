package domainservice

import (
	"context"
	"github.com/ljinf/template_project_v2/internal/cache"
	"github.com/ljinf/template_project_v2/internal/logic/do"
	"github.com/ljinf/template_project_v2/internal/model"
	"github.com/ljinf/template_project_v2/internal/repository"
	"github.com/ljinf/template_project_v2/pkg/log"
	"github.com/ljinf/template_project_v2/pkg/util"
)

type UserDomainService interface {
	GetUserProfile(ctx context.Context, uid string) *do.UserBaseInfo
}

type userDomainService struct {
	dao repository.UserRepository
}

func (u *userDomainService) GetUserProfile(ctx context.Context, uid string) *do.UserBaseInfo {

	var (
		user *model.User
		err  error
	)

	user, _ = cache.GetUserInfoCache(u.dao.Redis(), uid)

	if user == nil {
		user, err = u.dao.SelectById(ctx, uid)
		if err != nil {
			log.Error(ctx, "GetUserProfileError", "err", err)
			return nil
		}

		if err = cache.SetUserInfoCache(u.dao.Redis(), user); err != nil {
			log.Error(ctx, "SetUserInfoCacheError", "err", err)
		}
	}

	userBaseInfo := new(do.UserBaseInfo)
	err = util.CopyProperties(userBaseInfo, user)
	if err != nil {
		log.Error(ctx, "GetUserProfileError", "err", err)
		return nil
	}
	return userBaseInfo
}

func NewUserDomainService(r repository.UserRepository) UserDomainService {
	return &userDomainService{
		dao: r,
	}
}
