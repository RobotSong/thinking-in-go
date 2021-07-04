package main

import "errors"

var Canceled = errors.New("context canceled")

var DeadlineExceeded = deadlineExceeded{}

type deadlineExceeded struct {
}

func (deadlineExceeded) Error() string {
	return "deadline exceeded"
}

func (deadlineExceeded) Timeout() bool {
	return true
}

func (deadlineExceeded) Temporary() bool {
	return true
}

// classifyError use type assertions to classify errors
func classifyError(err error) {
	if tmp, ok := err.(interface{ Temporary() bool }); ok {
		if tmp.Temporary() {
			// retry
		} else {
			// report
		}
	}
}
