package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"huahua_account_backend/account/internal/model"
	"huahua_account_backend/account/internal/svc"
	"huahua_account_backend/account/internal/types"
	"io"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const userIdKey = "userId"

func getUserID(ctx context.Context) string {
	v := ctx.Value(userIdKey)
	if v != nil {
		uid, _ := v.(string)
		return uid
	}
	return ""
}

func (l *LoginLogic) getJwtToken(userId string) string {
	claims := make(jwt.MapClaims)
	iat := time.Now().UnixMilli() / 1000

	claims["exp"] = iat + int64(l.svcCtx.Config.Auth.AccessExpire)
	claims["iat"] = iat
	claims[userIdKey] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	s, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		panic(err)
	}
	return s
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {

	openid, ok := l.GetOpenid(req.Code)
	if !ok {
		return nil, fmt.Errorf("code is invalid")
	}
	user, err := l.svcCtx.UserModel.FindByOpenid(context.Background(), openid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		user = &model.User{
			Openid:   openid,
			Username: req.Name,
			LastTime: time.Now(),
		}
		err := l.svcCtx.UserModel.Create(context.Background(), user)
		if err != nil {
			return nil, err
		}

		err = l.initUserInfo(context.Background(), user)
		if err != nil {
			return nil, err
		}

	}
	resp = new(types.LoginReply)
	resp.Name = user.Username
	resp.AvatarUrl = req.AvatarUrl
	resp.AccessToken = l.getJwtToken(user.ID)
	resp.AccessExpire = l.svcCtx.Config.Auth.AccessExpire
	return
}

func (l *LoginLogic) GetOpenid(code string) (string, bool) {
	var url = fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		l.svcCtx.Config.Wechat.AppID, l.svcCtx.Config.Wechat.AppSecret, code)

	res, err := http.Get(url)
	if err != nil {
		return "", false
	}
	body, _ := io.ReadAll(res.Body)
	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", false
	}
	if r["openid"] == nil {
		return "", false
	}
	return r["openid"].(string), true
}

func (l *LoginLogic) initUserInfo(ctx context.Context, user *model.User) error {
	defaultBook := &model.Book{
		Uid:  user.ID,
		Name: "支出账本",
	}
	defaultTag := &model.Tag{
		Uid:  user.ID,
		Name: "普通消费",
	}
	err := l.svcCtx.BookModel.Create(ctx, defaultBook)
	if err != nil {
		return err
	}

	err = l.svcCtx.TagModel.Create(ctx, defaultTag)
	if err != nil {
		return err
	}

	acTag := &model.TagRel{
		SID: defaultBook.ID,
		TID: defaultTag.ID,
	}

	err = l.svcCtx.TagModel.CreateTagRel(context.Background(), acTag)
	if err != nil {
		return err
	}

	return nil
}
