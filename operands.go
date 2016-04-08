package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/tchssk/goa-example/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("operands")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}

// Sub runs the add action.
func (c *OperandsController) Sub(ctx *app.SubOperandsContext) error {
	sum := ctx.Left - ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}
