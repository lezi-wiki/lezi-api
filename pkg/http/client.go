package http

import "github.com/idoubi/goz"

var client *goz.Request

func init() {
	client = goz.NewClient()
}
