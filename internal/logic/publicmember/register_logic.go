package publicmember

import (
	"context"

	"github.com/iot-synergy/synergy-common/config"
	"github.com/iot-synergy/synergy-common/utils/pointy"

	"github.com/iot-synergy/synergy-common/enum/errorcode"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "captcha" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	if ok := l.svcCtx.Captcha.Verify(config.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); ok {
		_, err := l.svcCtx.MmsRpc.CreateMember(l.ctx,
			&mms.MemberInfo{
				Username: &req.Username,
				Password: &req.Password,
				Email:    &req.Email,
				Nickname: &req.Username,
				Status:   pointy.GetPointer(uint32(1)),
				RankId:   pointy.GetPointer(l.svcCtx.Config.ProjectConf.DefaultRankId),
			})
		if err != nil {
			return nil, err
		}
		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, "login.signupSuccessTitle"),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}
}
