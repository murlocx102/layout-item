package main

import (
	"testing"
)

func Test_GetVip(t *testing.T) {
	data := RequestResult(t, addr+"/vip/exp/change/ratio", nil)

	t.Log(data)
}
