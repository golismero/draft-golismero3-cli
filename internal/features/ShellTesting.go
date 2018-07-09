package features

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"
)

// ShellResult calling interface helper
type ShellResult struct {
	Stdout           []byte
	Stderr           []byte
	ExitCodeObtained int
}

// CallShell does exactly what name says
func CallShell(command string, params []string, input []byte) (ShellResult, error) {
	var outbuf, errbuf bytes.Buffer
	var cmd *exec.Cmd
	env := os.Environ()

	cmd = exec.Command(command, params...)
	cmd.Env = env

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	if len(input) > 0 {
		wr, err := cmd.StdinPipe()
		if err != nil {
			return ShellResult{}, err
		}
		err = cmd.Start()
		if err != nil {
			return ShellResult{}, err
		}
		_, err = wr.Write(input)
		if err != nil {
			return ShellResult{}, err
		}
		wr.Close()
	} else {
		err := cmd.Start()
		if err != nil {
			return ShellResult{}, err
		}
	}

	err := cmd.Wait()

	stdout := outbuf.Bytes()
	stderr := errbuf.Bytes()

	exitCode := 0

	if err != nil {
		exiterr, ok := err.(*exec.ExitError)
		if ok {
			status, ok := exiterr.Sys().(syscall.WaitStatus)
			if ok {
				exitCode = status.ExitStatus()
				if exitCode < 0 {
					return ShellResult{}, err
				}
			}
		} else {
			return ShellResult{}, err
		}
	}

	return ShellResult{
		Stdout:           stdout,
		Stderr:           stderr,
		ExitCodeObtained: exitCode,
	}, nil
}
