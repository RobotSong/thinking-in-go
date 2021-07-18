package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, messageArgs ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(messagef, messageArgs),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

type LowLevelErr struct {
	error
}

func isGlobalExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{wrapError(err, err.Error())}
	}

	return info.Mode().Perm()&0100 == 0100, nil
}

type IntermediaErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGlobalExec(jobBinPath)
	if err != nil {
		return IntermediaErr{
			wrapError(
				err,
				"cannot run job %q: requisite binaries not available",
				id,
			),
		}
	} else if isExecutable == false {
		return wrapError(nil,
			"cannot run job %q: requisite binaries are not executable",
			id,
		)
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[LogId: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(IntermediaErr); ok {
			msg = err.Error()
		}
		handleError(1, err, msg)
	}
}
