package service

import (
	"blog/app/pojo/req"
	"fmt"
)

type TestService interface {
	PrintInfo(req *req.TestPostRequest)
}

type Test struct {
}

func (t Test) PrintInfo(req *req.TestPostRequest) {
	fmt.Printf("test data: name=%s age=%d", req.Name, req.Age)
}
