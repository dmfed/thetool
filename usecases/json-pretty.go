package usecases

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/dmfed/thetool/util"
)

func JsonPretty(in, out string) error {
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

	m := new(map[string]any)
	if err := json.NewDecoder(src).Decode(m); err != nil {
		return fmt.Errorf("err decoding json: %w", err)
	}

	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("err encoding output json: %w", err)
	}

	outbuf := bufio.NewWriter(dst)
	_, err = outbuf.Write(b)

	outbuf.Flush()

	return err
}
