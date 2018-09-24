package main

import (
	"os/exec"
)

type cmd struct {
	cmd    *exec.Cmd
	output string
}

func (c *cmd) runCommand() error {
	out, err := c.cmd.CombinedOutput()
	if err != nil {
		return err
	}
	c.output = string(out)
	return nil
}
