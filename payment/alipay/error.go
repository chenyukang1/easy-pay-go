package alipay

type IllegalApiParamError struct {
	Msg string
}

func (err *IllegalApiParamError) Error() string {
	return err.Msg
}
