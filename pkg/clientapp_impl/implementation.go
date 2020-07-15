package clientapp_impl

import (
	"context"

	clientapp "github.com/anz-bank/test-rig-example/gen/clientapp"
)

type ServiceImpl struct {
	clientapp.ServiceInterface
}

func (s ServiceImpl) GetEndpoint(ctx context.Context, req *clientapp.GetEndpointRequest, client clientapp.GetEndpointClient) (*clientapp.Str, error) {
	return nil, nil
}

func (s ServiceImpl) PostEndpointWithArg(ctx context.Context, req *clientapp.PostEndpointWithArgRequest, client clientapp.PostEndpointWithArgClient) (*clientapp.StatusMsg, error) {
	return nil, nil
}
