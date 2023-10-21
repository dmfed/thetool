package usecases

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"

	"github.com/dmfed/tool/util"
)

func XMLPretty(in, out string) error {
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

	dec := xml.NewDecoder(src)
	enc := xml.NewEncoder(dst)
	enc.Indent("", "  ")

	if err := decEncXml(dec, enc); err != nil {
		return err
	}

	err = enc.Flush()

	return err
}

func decEncXml(dec *xml.Decoder, enc *xml.Encoder) error {
	var (
		t   xml.Token
		err error
	)

	for {
		if t, err = dec.Token(); err != nil {
			break
		}

		if err = enc.EncodeToken(t); err != nil {
			break
		}
	}

	if err != nil && errors.Is(err, io.EOF) {
		// dec.Token() will return io.EOF, when
		// there is no data left
		err = nil
	}

	return err
}
