package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/dmfed/tool/util"
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

	dec := json.NewDecoder(src)
	enc := json.NewEncoder(dst)
	enc.SetIndent("", "  ")

	var raw json.RawMessage

	for dec.More() {
		if err = dec.Decode(&raw); err != nil {
			break
		}
		if err = enc.Encode(raw); err != nil {
			break
		}
	}

	return err
}
