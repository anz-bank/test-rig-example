// Code generated by sysl DO NOT EDIT.
package dbfront

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

// Data ...
type Data struct {
	ID  int64  `json:"id"`
	Txt string `json:"txt"`
}

// GetEndpointRequest ...
type GetEndpointRequest struct {
	ID string
}

// PostEndpointWithArgRequest ...
type PostEndpointWithArgRequest struct {
	ID string
}

// *StatusMsg validator
func (s *StatusMsg) Validate() error {
	return validator.Validate(s)
}

// *Data validator
func (s *Data) Validate() error {
	return validator.Validate(s)
}
