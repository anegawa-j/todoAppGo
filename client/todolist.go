// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "todolist": todolist Resource Client
//
// Command:
// $ goagen
// --design=todoApp/design
// --out=$(GOPATH)\src\todoApp
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ShowTodolistPath computes a request path to the show action of todolist.
func ShowTodolistPath(todoID int) string {
	param0 := strconv.Itoa(todoID)

	return fmt.Sprintf("/todolist/%s", param0)
}

// Get todo by id
func (c *Client) ShowTodolist(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowTodolistRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowTodolistRequest create the request corresponding to the show action endpoint of the todolist resource.
func (c *Client) NewShowTodolistRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
