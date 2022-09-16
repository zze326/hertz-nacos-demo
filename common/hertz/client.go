package hertz

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	hertzConfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"
	commonConfig "hertz-demo/common/config"
	"hertz-demo/common/util"
)

/**
 * @Author: zze
 * @Date: 2022/9/15 09:58
 * @Desc: Hertz 客户端
 */

type Client struct {
	*client.Client
}

func MustNewClient(nacosConfig *commonConfig.Nacos) *Client {
	c, err := NewClient(nacosConfig)
	if err != nil {
		hlog.Fatal(err)
	}
	return c
}

func NewClient(nacosConfig *commonConfig.Nacos) (*Client, error) {
	hertzClient, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	namingClient, err := nacosConfig.GetNamingClient()
	if err != nil {
		return nil, err
	}
	r := nacos.NewNacosResolver(namingClient)
	hertzClient.Use(sd.Discovery(r))
	return &Client{
		Client: hertzClient,
	}, nil
}

func (c *Client) Get(ctx context.Context, serviceName, path string, dest interface{}) error {
	return c.request(ctx, serviceName, path, consts.MethodGet, util.BlankStr, nil, dest)
}

func (c *Client) RequestWithJsonData(ctx context.Context, serviceName, path, method string, requestData, dest interface{}) error {
	return c.request(ctx, serviceName, path, method, "application/json", requestData, dest)
}

func (c *Client) request(ctx context.Context, serviceName, path, method, contentType string, requestData, dest interface{}) error {
	req := &protocol.Request{}
	res := &protocol.Response{}

	req.Header.SetMethod(method)
	if !util.IsBlank(contentType) {
		req.Header.SetContentTypeBytes([]byte(contentType))
	}

	req.SetRequestURI(fmt.Sprintf("http://%s%s", serviceName, path))
	req.SetOptions(hertzConfig.WithSD(true))
	req.SetHeader(HeaderXRequestID, ctx.Value(HeaderXRequestID).(string))

	if method != consts.MethodGet && method != consts.MethodDelete && requestData != nil {
		jsonByte, err := sonic.Marshal(requestData)
		if err != nil {
			return err
		}
		req.SetBody(jsonByte)
	}

	if err := c.Client.Do(ctx, req, res); err != nil {
		return err
	}

	if err := sonic.Unmarshal(res.BodyBytes(), dest); err != nil {
		return fmt.Errorf("unmarshal body to struct object error: %v", err)
	}

	return nil
}
