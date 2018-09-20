package errors

type CommonError struct {
	ErrCode		string
	ErrMsg		string
}

func (err CommonError) Error() string {
	return err.ErrMsg
}
