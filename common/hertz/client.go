package hertz

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/hertz-contrib/registry/nacos"
	commonConfig "hertz-demo/common/config"
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
	statusCode, body, err := c.Client.Get(ctx, nil, fmt.Sprintf("http://%s%s", serviceName, path), config.WithSD(true))
	if err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("response status code: %d", statusCode)
	}

	if err = sonic.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("unmarshal body to struct object error: %v", err)
	}

	return nil
}

func (c *Client) RequestWithJsonData(ctx context.Context, serviceName, path, method string, requestData, dest interface{}) error {
	req := &protocol.Request{}
	res := &protocol.Response{}

	req.Header.SetMethod(method)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI(fmt.Sprintf("http://%s%s", serviceName, path))
	req.SetOptions(config.WithSD(true))

	jsonByte, err := sonic.Marshal(requestData)
	if err != nil {
		return err
	}
	req.SetBody(jsonByte)
	if err := c.Client.Do(ctx, req, res); err != nil {
		return err
	}

	if err := sonic.Unmarshal(res.BodyBytes(), dest); err != nil {
		return fmt.Errorf("unmarshal body to struct object error: %v", err)
	}

	return nil
}
