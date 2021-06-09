package context

import "fmt"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (p *Response)Error()string  {
	return fmt.Sprintf("code: %d, message: %s", p.Code, p.Message)
}