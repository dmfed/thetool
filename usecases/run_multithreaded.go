package usecases

import (
	"bufio"
	"context"
	"fmt"

	"github.com/dmfed/thetool/util"
)

func RunThreaded(ctx context.Context, in, out string, threads int, cmd []string) error {
	src, err := util.OpenInput(in)
	if err != nil {
		return fmt.Errorf("err opening input: %w", err)
	}
	defer src.Close()
	dst, err := util.OpenOutput(out)
	if err != nil {
		return fmt.Errorf("err opening output: %w", err)
	}
	defer dst.Close()
	outbuf := bufio.NewWriter(dst)

	reader := util.NewReader(src)
	runner := util.NewRunner(cmd)
	results := runner.Run(ctx, reader.Read(ctx), threads)

	var resultErr error
	for str := range results {
		if _, err := outbuf.WriteString(str); err != nil {
			resultErr = err
			break
		}
		// flush to immediately show the result
		outbuf.Flush()
	}

	if resultErr != nil {
		resultErr = fmt.Errorf("result write err: %w", resultErr)
	}

	if err := outbuf.Flush(); err != nil {
		resultErr = fmt.Errorf("%v, flush err: %w", resultErr, err)
	}

	if err := reader.Err(); err != nil {
		resultErr = fmt.Errorf("%v, reader err: %w", resultErr, err)
	}

	if err := runner.Err(); err != nil {
		resultErr = fmt.Errorf("%s, runner err: %w", resultErr, err)
	}

	return resultErr
}
