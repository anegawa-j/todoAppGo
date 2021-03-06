// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "todolist": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=todoApp/design
// --out=$(GOPATH)\src\todoApp
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// TodolistHref returns the resource href.
func TodolistHref(todoID interface{}) string {
	paramtodoID := strings.TrimLeftFunc(fmt.Sprintf("%v", todoID), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/todolist/%v", paramtodoID)
}
