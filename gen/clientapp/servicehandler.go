// Code generated by sysl DO NOT EDIT.
package clientapp

import (
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
	"github.com/anz-bank/test-rig-example/gen//dep1"
	"github.com/anz-bank/test-rig-example/gen//dep2"
)

// Handler interface for clientapp
type Handler interface {
	GetBarListHandler(w http.ResponseWriter, r *http.Request)
	GetEndpointHandler(w http.ResponseWriter, r *http.Request)
	GetFooListHandler(w http.ResponseWriter, r *http.Request)
	PostEndpointWithArgHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for clientapp API
type ServiceHandler struct {
	genCallback      core.RestGenCallback
	serviceInterface *ServiceInterface
	dep2dep2Service  dep2.Service
	dep1dep1Service  dep1.Service
}

// NewServiceHandler for clientapp
func NewServiceHandler(genCallback core.RestGenCallback, serviceInterface *ServiceInterface, dep2dep2Service dep2.Service, dep1dep1Service dep1.Service) *ServiceHandler {
	return &ServiceHandler{genCallback, serviceInterface, dep2dep2Service, dep1dep1Service}
}

// GetBarListHandler ...
func (s *ServiceHandler) GetBarListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetBarList == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetBarListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetBarListClient{
		GetBarbar: s.dep2dep2Service.GetBarbar,
	}

	str, err := s.serviceInterface.GetBarList(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str)
}

// GetEndpointHandler ...
func (s *ServiceHandler) GetEndpointHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetEndpoint == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetEndpointRequest

	req.ID = restlib.GetURLParam(r, "id")
	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetEndpointClient{}

	str, err := s.serviceInterface.GetEndpoint(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str)
}

// GetFooListHandler ...
func (s *ServiceHandler) GetFooListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetFooList == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetFooListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetFooListClient{
		GetFoofoo: s.dep1dep1Service.GetFoofoo,
	}

	str, err := s.serviceInterface.GetFooList(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, str)
}

// PostEndpointWithArgHandler ...
func (s *ServiceHandler) PostEndpointWithArgHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.PostEndpointWithArg == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req PostEndpointWithArgRequest

	req.ID = restlib.GetURLParam(r, "id")
	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := PostEndpointWithArgClient{}

	statusmsg, err := s.serviceInterface.PostEndpointWithArg(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, statusmsg)
}
