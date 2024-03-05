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
	page := int64(req.Page)
	pageSize := int64(req.PageSize)
	list, err := l.svcCtx.MmsRpc.AdminGetCommentList(l.ctx, &mms.CommentListReq{
		IsReply:     &req.IsReply,
		Page:        &page,
		PageSize:    &pageSize,
		Title:       &req.Title,
		Content:     &req.Content,
		CommentTime: req.CommentTime,
	})

	if err != nil {
		return &types.CommentListRespVo{
			BaseMsgResp: types.BaseMsgResp{
				Code: -1,
				Msg:  err.Error(),
			},
		}, nil
	}

	vos := make([]types.CommentRespData, 0)
	for _, info := range list.List {
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
		vos = append(vos, types.CommentRespData{
			Id:          info.GetId(),
			Title:       info.GetTitle(),
			Content:     info.GetContent(),
			MemberId:    info.GetMemberId(),
			Create_time: info.GetCreateTime(),
			Update_time: info.GetUpdateTime(),
			IsReply:     info.GetIsReply(),
			Reply:       respVos,
		})
	}

	return &types.CommentListRespVo{
		BaseMsgResp: types.BaseMsgResp{
			Code: 0,
			Msg:  "成功",
		},
		Data: types.CommentListRespData{
			Data:  vos,
			Total: len(vos),
		},
	}, nil
}
