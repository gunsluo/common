package model

import "fmt"

// RequestArg example request
type RequestArg struct {
	ArgOne string `json:"arg"`
}

// String request to string
func (req *RequestArg) String() string {
	return fmt.Sprintf("<arg: %s>", req.ArgOne)
}
