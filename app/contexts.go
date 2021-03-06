//************************************************************************//
// API "adder": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/tchssk/goa-example
// --design=github.com/tchssk/goa-example/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// AddOperandsContext provides the operands add action context.
type AddOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Left    int
	Right   int
}

// NewAddOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller add action.
func NewAddOperandsContext(ctx context.Context, service *goa.Service) (*AddOperandsContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := AddOperandsContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// SubOperandsContext provides the operands sub action context.
type SubOperandsContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Left    int
	Right   int
}

// NewSubOperandsContext parses the incoming request URL and body, performs validations and creates the
// context used by the operands controller sub action.
func NewSubOperandsContext(ctx context.Context, service *goa.Service) (*SubOperandsContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := SubOperandsContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramLeft := req.Params["left"]
	if len(paramLeft) > 0 {
		rawLeft := paramLeft[0]
		if left, err2 := strconv.Atoi(rawLeft); err2 == nil {
			rctx.Left = left
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("left", rawLeft, "integer"))
		}
	}
	paramRight := req.Params["right"]
	if len(paramRight) > 0 {
		rawRight := paramRight[0]
		if right, err2 := strconv.Atoi(rawRight); err2 == nil {
			rctx.Right = right
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("right", rawRight, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *SubOperandsContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
