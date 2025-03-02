package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandelError(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
