package model

import "fmt"

type RequestArg struct {
	ArgOne string `json:"arg"`
}

func (this *RequestArg) String() string {
	return fmt.Sprintf("<arg: %s>", this.ArgOne)
}
