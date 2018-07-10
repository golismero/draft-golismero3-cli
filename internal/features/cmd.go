package features

import (
	"fmt"
	"strings"
)

// CmdFeatures contains the basic test features
func CmdFeatures(ctx *Context) []Feat {
	return []Feat{
		Feat{`^no params$`, ctx.noParams},
		Feat{`^I call "([^"]*)" command$`, ctx.callCommand},
		Feat{`^tool show me the command help$`, ctx.toolShowMeTheCommandHelp},
	}
}

func (ctx *Context) noParams() error {
	return nil
}

func (ctx *Context) callCommand(cmd string) error {
	var (
		params []string
		stdin  []byte
		err    error
	)
	_, hasParams := ctx.m["params"]
	if hasParams {
		params, err = getStringSlice(ctx.m, "params")
		if err != nil {
			return err
		}
	}
	_, hasStdin := ctx.m["stdin"]
	if hasStdin {
		stdin, err = getByteSlice(ctx.m, "stdin")
		if err != nil {
			return err
		}
	}

	res, err := CallShell(cmd, params, stdin)
	if err != nil {
		return err
	}
	ctx.m["result"] = res
	return nil
}

func (ctx *Context) toolShowMeTheCommandHelp() error {
	result, err := getShellResult(ctx.m, "result")
	if err != nil {
		return err
	}
	resMsg := string(result.Stderr)
	if !strings.Contains(resMsg, "help") {
		return fmt.Errorf("Can't find the help")
	}
	return nil
}
