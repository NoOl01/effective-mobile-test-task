package errs

import "errors"

var (
	ErrRecordAlreadyExist = errors.New("record already exist")
	ErrRecordsNotFound    = errors.New("records not found")
	ErrInvalidTimeFormat  = errors.New("end_date format is invalid. format should be '2006-01-02'")
)
