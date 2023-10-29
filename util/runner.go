package util

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

type Runner struct {
	cmd []string
	err error
}

func NewRunner(command []string) *Runner {
	return &Runner{cmd: command}
}

func (r *Runner) Run(ctx context.Context, in <-chan string, threads int) <-chan string {
	out := make(chan string)

	wg := new(sync.WaitGroup)
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			for line := range in {
				args := strings.Split(line, " ")
				cmd := append(copyCmd(r.cmd), args...)
				output, err := RunCmd(cmd)
				if err != nil {
					output = fmt.Sprintf("failed %s: %s\n, output: %s", strings.Join(cmd, " "), err.Error(), output)
				}

				select {
				case out <- output:
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (r *Runner) Err() error {
	return r.err
}

func copyCmd(c []string) []string {
	out := make([]string, len(c))
	copy(out, c)
	return out
}

func RunCmd(command []string) (string, error) {
	var cmd *exec.Cmd
	if len(command) == 1 {
		cmd = exec.Command(command[0])
	} else {
		cmd = exec.Command(command[0], command[1:]...)
	}
	out, err := cmd.CombinedOutput()
	return string(out), err
}
