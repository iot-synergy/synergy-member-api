package publicmember

import (
	"context"

	"github.com/iot-synergy/synergy-common/config"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBySmsLogic {
	return &RegisterBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RegisterBySmsLogic) RegisterBySms(req *types.RegisterBySmsReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "sms" && l.svcCtx.Config.ProjectConf.RegisterVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for email validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}

	if captchaData == req.Captcha {
		_, err := l.svcCtx.MmsRpc.CreateMember(l.ctx,
			&mms.MemberInfo{
				Username: &req.Username,
				Password: &req.Password,
				Mobile:   &req.PhoneNumber,
				Nickname: &req.Username,
				Status:   pointy.GetPointer(uint32(1)),
				RankId:   pointy.GetPointer(l.svcCtx.Config.ProjectConf.DefaultRankId),
			})
		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		return &types.BaseMsgResp{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, "login.signupSuccessTitle"),
		}, nil
	}

	return nil, errorx.NewInvalidArgumentError("login.wrongCaptcha")
}
