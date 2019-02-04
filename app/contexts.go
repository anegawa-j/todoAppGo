// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "todolist": Application Contexts
//
// Command:
// $ goagen
// --design=todoApp/design
// --out=$(GOPATH)\src\todoApp
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ShowTodolistContext provides the todolist show action context.
type ShowTodolistContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	TodoID int
}

// NewShowTodolistContext parses the incoming request URL and body, performs validations and creates the
// context used by the todolist controller show action.
func NewShowTodolistContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowTodolistContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowTodolistContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramTodoID := req.Params["todoID"]
	if len(paramTodoID) > 0 {
		rawTodoID := paramTodoID[0]
		if todoID, err2 := strconv.Atoi(rawTodoID); err2 == nil {
			rctx.TodoID = todoID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("todoID", rawTodoID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowTodolistContext) OK(r *GoaExampleTodolist) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.todolist+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowTodolistContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}