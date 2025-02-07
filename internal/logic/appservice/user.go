package appservice

import (
	"context"
	"github.com/ljinf/template_project_v2/api/reply"
	"github.com/ljinf/template_project_v2/internal/logic/domainservice"
	"github.com/ljinf/template_project_v2/pkg/util"
)

type UserAppService interface {
	GetUserProfile(ctx context.Context, uid string) *reply.UserInfoReply
}

type userAppService struct {
	userDomainSrv domainservice.UserDomainService
}

func (s *userAppService) GetUserProfile(ctx context.Context, uid string) *reply.UserInfoReply {
	userInfo := s.userDomainSrv.GetUserProfile(ctx, uid)
	if userInfo == nil || userInfo.ID == 0 {
		return nil
	}

	infoReply := new(reply.UserInfoReply)
	util.CopyProperties(infoReply, userInfo)
	return infoReply
}

func NewUserAppService(d domainservice.UserDomainService) UserAppService {
	return &userAppService{
		userDomainSrv: d,
	}
}
