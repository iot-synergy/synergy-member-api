package svc

import (
	"github.com/iot-synergy/oms-core/rpc/coreclient"
	"github.com/iot-synergy/synergy-addx-proxy/synergyaddxproxyclient"
	"github.com/iot-synergy/synergy-common/utils/captcha"
	"github.com/iot-synergy/synergy-message-center/mcmsclient"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/iot-synergy/synergy-member-rpc/mmsclient"

	"github.com/iot-synergy/synergy-member-api/internal/config"
	i18n2 "github.com/iot-synergy/synergy-member-api/internal/i18n"
	"github.com/iot-synergy/synergy-member-api/internal/middleware"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	Casbin    *casbin.Enforcer
	Authority rest.Middleware
	Vip       rest.Middleware
	Trans     *i18n.Translator
	MmsRpc    mmsclient.Mms
	McmsRpc   mcmsclient.Mcms
	Captcha   *base64Captcha.Captcha
	CoreRpc   coreclient.Core
	Redis     redis.UniversalClient
	AddxProxy synergyaddxproxyclient.SynergyAddxProxy
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Vip:       middleware.NewVipMiddleware().Handle,
		Trans:     trans,
		Redis:     rds,
		Casbin:    cbn,
		MmsRpc:    mmsclient.NewMms(zrpc.NewClientIfEnable(c.MmsRpc)),
		Captcha:   captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		CoreRpc:   coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		McmsRpc:   mcmsclient.NewMcms(zrpc.NewClientIfEnable(c.McmsRpc)),
		AddxProxy: synergyaddxproxyclient.NewSynergyAddxProxy(zrpc.NewClientIfEnable(c.AddxProxyRpc)),
	}
}
