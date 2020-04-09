package clientapp_impl

import (
	"context"

	simple_app "github.com/anz-bank/test-rig-example/gen/simple_app"
)

type ServiceImpl struct {
	simple_app.ServiceInterface
}

func (s ServiceImpl) GetEndpoint(ctx context.Context, req *simple_app.GetEndpointRequest, client simple_app.GetEndpointClient) (*simple_app.Str, error) {
	return nil, nil
}

func (s ServiceImpl) PostEndpointWithArg(ctx context.Context, req *simple_app.PostEndpointWithArgRequest, client simple_app.PostEndpointWithArgClient) (*simple_app.StatusMsg, error) {
	return nil, nil
}
