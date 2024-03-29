package main

import (
	"os"
	"strings"

	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/tdiff"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	nocolor bool
	verbose bool
	version bool
	years   bool
}

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf("%s [OPTIONS] <date> <date>", os.Args[0])
	cli.BugEmail = "tdiff.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			hl.Sprintf("%d: Invalid option\n", InvalidOption),
			hl.Sprintf("%d: Missing option\n", MissingOption),
			hl.Sprintf("%d: Invalid argument\n", InvalidArgument),
			hl.Sprintf("%d: Missing argument\n", MissingArgument),
			hl.Sprintf("%d: Extra argument\n", ExtraArgument),
			hl.Sprintf("%d: Exception", Exception),
		},
		" ",
	)
	cli.Info = strings.Join(
		[]string{
			"Calculate the approximate difference in time between",
			"two dates. Dates should be formatted in RFC3339",
			"(YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS).",
		},
		" ",
	)
	cli.SeeAlso = []string{"date"}
	cli.Title = "TDiff"

	// Parse cli flags
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show stacktrace, if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Flag(
		&flags.years,
		"y",
		"yearly-training-mode",
		false,
		"Calculate time like a company does for yearly training.",
	)
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("tdiff version %s\n", tdiff.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() < 2 {
		cli.Usage(MissingArgument)
	} else if cli.NArg() > 2 {
		cli.Usage(ExtraArgument)
	}

	// Validate arguments match time format regex
	for _, arg := range cli.Args() {
		if !tdiff.RFC3339.MatchString(arg) {
			cli.Usage(InvalidArgument)
		}
	}
}
