package publicmember

import (
	"context"

	"github.com/iot-synergy/synergy-common/config"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordByEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByEmailLogic {
	return &ResetPasswordByEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ResetPasswordByEmailLogic) ResetPasswordByEmail(req *types.ResetPasswordByEmailReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.ResetVerify != "email" && l.svcCtx.Config.ProjectConf.ResetVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.Email).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for email validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		memberData, err := l.svcCtx.MmsRpc.GetMemberList(l.ctx, &mms.MemberListReq{
			Page:     1,
			PageSize: 1,
			Email:    &req.Email,
		})
		if err != nil {
			return nil, err
		}

		if memberData.Total == 0 {
			return nil, errorx.NewCodeInvalidArgumentError("login.userNotExist")
		}

		_, err = l.svcCtx.MmsRpc.UpdateMember(l.ctx,
			&mms.MemberInfo{
				Id:       memberData.Data[0].Id,
				Password: &req.Password,
			})
		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.Email).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		return &types.BaseMsgResp{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		}, nil
	}

	return nil, errorx.NewInvalidArgumentError("login.wrongCaptcha")
}
