package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"stair/easy-chat/apps/user/models"
	"stair/easy-chat/pkg/ctxdata"
	"stair/easy-chat/pkg/encrypt"
	"stair/easy-chat/pkg/wuid"
	"stair/easy-chat/pkg/xerr"
	"time"

	"stair/easy-chat/apps/user/rpc/internal/svc"
	"stair/easy-chat/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsRegister = xerr.New(xerr.SERVER_COMMON_ERROR, "手机号已经注册过")
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 1. 验证用户是否注册，根据手机号码验证
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewDBErr(), "find user by phone err %v,req %v", err, in.Phone)
	}

	if userEntity != nil {
		return nil, errors.WithStack(ErrPhoneIsRegister)
	}

	// 定义用户数据
	userEntity = &models.Users{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex: sql.NullInt64{
			Int64: int64(in.Sex),
			Valid: true,
		},
	}

	if len(in.Password) > 0 {
		genPassword, err1 := encrypt.GenPasswordHash([]byte(in.Password))
		if err1 != nil {
			return nil, errors.Wrapf(xerr.NewInternalErr(), "gen password hash err %v,req %v", err1, in.Password)
		}
		userEntity.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}

	_, err3 := l.svcCtx.UsersModel.Insert(l.ctx, userEntity)
	if err3 != nil {
		return nil, errors.Wrapf(xerr.NewDBErr(), "gen password hash err %v,req %v", err3, in.Password)
	}

	// 生成token
	now := time.Now().Unix()
	token, err2 := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire,
		userEntity.Id)
	if err2 != nil {
		return nil, errors.Wrapf(xerr.NewInternalErr(), "gen jwt token err %v,req %v", err2, in.Password)
	}

	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
