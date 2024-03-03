package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/dmfed/tool/usecases"
)

var (
	cmdJP   bool
	cmdXP   bool
	cmdRun  bool
	input   string
	output  string
	threads int
)

var commit string

var credits string = fmt.Sprintf(`The Tool (c) 2023 Dmitry Fedotov
https://github.com/dmfed/tool

commit: %s
`, commit)

func main() {
	flag.BoolVar(&cmdJP, "jp", false, "json pretty")
	flag.BoolVar(&cmdXP, "xp", false, "xml pretty")
	flag.BoolVar(&cmdRun, "run", false, "run the following command")
	flag.StringVar(&input, "in", "", "input file (leave empty to read from stdin)")
	flag.StringVar(&output, "out", "", "output file (leave empty to write to stdout)")
	flag.IntVar(&threads, "t", 1, "how many threads to start")

	flag.Usage = func() {
		fmt.Println(credits)
		flag.PrintDefaults()
	}
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var err error

	switch {
	case cmdJP:
		err = usecases.JsonPretty(input, output)
	case cmdXP:
		err = usecases.XMLPretty(input, output)
	case cmdRun:
		err = usecases.RunThreaded(ctx, input, output, threads, flag.Args())
	default:
		flag.Usage()
	}

	if err != nil {
		fmt.Println(err)
	}
}
