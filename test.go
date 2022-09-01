package main

import (
	"math/rand"
	"strconv"
	"strings"
)

type Int interface {
	Add(other int) string
	Sub(other int) string
}

type MyInt struct {
	self int
}

func (m *MyInt) Add(other int) string {
	res := m.self + other
	l := rand.Intn(3)
	_list := []string{strconv.Itoa(m.self), strconv.Itoa(other), strconv.Itoa(res)}
	_list[l] = "( )"
	firstHalf := strings.Join(_list[:2], " + ")
	resStr := strings.Join([]string{firstHalf, _list[2]}, " = ")
	return resStr
}

func (m *MyInt) Sub(other int) string {
	res := 0
	resStr := ""
	l := rand.Intn(3)
	if m.self >= other {
		res = m.self - other
		_list := []string{strconv.Itoa(m.self), strconv.Itoa(other), strconv.Itoa(res)}
		_list[l] = "( )"
		firstHalf := strings.Join(_list[:2], " - ")
		resStr = strings.Join([]string{firstHalf, _list[2]}, " = ")
	} else {
		res = other - m.self
		_list := []string{strconv.Itoa(other), strconv.Itoa(m.self), strconv.Itoa(res)}
		_list[l] = "( )"
		firstHalf := strings.Join(_list[:2], " - ")
		resStr = strings.Join([]string{firstHalf, _list[2]}, " = ")
	}
	return resStr
}
