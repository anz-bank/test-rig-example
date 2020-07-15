// Code generated by sysl DO NOT EDIT.
package clientapp

import (
	"time"

	"github.com/anz-bank/sysl-go/validator"
	"github.com/rickb777/date"
)

// Reference imports to suppress unused errors
var _ = time.Parse

// Reference imports to suppress unused errors
var _ = date.Parse

// StatusMsg ...
type StatusMsg struct {
	Status string `json:"status"`
}

// GetBarListRequest ...
type GetBarListRequest struct {
}

// GetEndpointRequest ...
type GetEndpointRequest struct {
	ID string
}

// GetFooListRequest ...
type GetFooListRequest struct {
}

// PostEndpointWithArgRequest ...
type PostEndpointWithArgRequest struct {
	ID string
}

// *StatusMsg validator
func (s *StatusMsg) Validate() error {
	return validator.Validate(s)
}

// Str ...
type Str string
