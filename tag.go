package main

import (
	"net/http"
	"strings"
)

type tagType string

const (
	router     tagType = "@router"
	middleware tagType = "@middleware"
	group      tagType = "@group"
	method     tagType = "@method"
	use        tagType = "@use"
)

func validComment(comment string) bool {
	if comment == string(router) ||
		comment == string(group) ||
		comment == string(method) ||
		comment == string(use) {
		return true
	}
	return false
}

type tag struct {
	// 注释 tag类型
	typ tagType
	// tag值 middleware无值
	value string
}

// newTag 实例化一个tag
func newTag(typ tagType, value string) *tag {
	if !validTag(typ, value) {
		return nil
	}
	return &tag{
		typ:   typ,
		value: value,
	}
}

// extractToTag 解析注释生成一个tag
func extractToTag(comment string) *tag {
	// middleware无值 特殊处理
	if comment == string(middleware) {
		return &tag{typ: middleware}
	}
	_tag := strings.Split(comment, ":")
	// 如果没有
	if len(_tag) == 0 {
		return nil
	}
	if !validComment(_tag[0]) {
		return nil
	}
	// 去除值的空格
	_tag[1] = strings.TrimSpace(_tag[1])
	// 生成tag
	return newTag(tagType(_tag[0]), _tag[1])
}

func validTag(typ tagType, value string) bool {
	if typ == method {
		if value == http.MethodGet ||
			value == http.MethodConnect ||
			value == http.MethodHead ||
			value == http.MethodPost ||
			value == http.MethodPut ||
			value == http.MethodPatch ||
			value == http.MethodDelete {
			return true
		}
	}
	return false
}
