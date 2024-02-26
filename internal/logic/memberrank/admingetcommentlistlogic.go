package memberrank

import (
	"context"
	"github.com/iot-synergy/synergy-member-api/internal/svc"
	"github.com/iot-synergy/synergy-member-api/internal/types"
	"github.com/iot-synergy/synergy-member-rpc/types/mms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetCommentListLogic {
	return &AdminGetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetCommentListLogic) AdminGetCommentList(req *types.CommentListReqVo) (resp *types.CommentListRespVo, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.MmsRpc.AdminGetCommentList(l.ctx, &mms.CommentListReq{
		IsReply:  &req.IsReply,
		PageNo:   &req.PageNo,
		PageSize: &req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	vos := make([]types.CommentRespVo, 0)
	for _, info := range list.Titles {
		reply := info.GetReply()
		respVos := make([]types.ReplyRespVo, 0)
		for _, replyInfo := range reply {
			respVos = append(respVos, types.ReplyRespVo{
				Id:         replyInfo.GetId(),
				CommentId:  replyInfo.GetCommentId(),
				Reply:      replyInfo.GetReply(),
				AdminId:    replyInfo.GetAdminId(),
				AdminName:  replyInfo.GetAdminName(),
				CreateTime: replyInfo.GetCreateTime(),
				UpdateTime: info.GetUpdateTime(),
			})
		}
		vos = append(vos, types.CommentRespVo{
			Id:          info.GetId(),
			Title:       info.GetTitle(),
			Content:     info.GetContent(),
			MemberId:    info.GetMemberId(),
			Create_time: info.GetCreateTime(),
			Update_time: info.GetUpdateTime(),
			Reply:       respVos,
		})
	}

	return &types.CommentListRespVo{CommentList: vos}, err
}
