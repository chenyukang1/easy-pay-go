package payment

import "fmt"

type BusinessError struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func checkBusinessError(check func(resp *BaseResp) bool, resp *BaseResp) error {
	if !check(resp) {
		return &BusinessError{
			Code:    resp.Code,
			Msg:     resp.Msg,
			SubCode: resp.SubCode,
			SubMsg:  resp.SubMsg,
		}
	}
	return nil
}

func (err *BusinessError) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, err.Code, err.Msg, err.SubCode, err.SubMsg)
}
