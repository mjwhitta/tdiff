package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/tdiff"
)

// Exit status
const (
	Good             int = 0
	InvalidOption    int = 1
	InvalidArgument  int = 2
	MissingArguments int = 3
	ExtraArguments   int = 4
	Exception        int = 5
)

// Flags
type cliFlags struct {
	nocolor bool
	verbose bool
	version bool
	years   bool
}

var flags cliFlags

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
			"1: Invalid option\n",
			"2: Invalid argument\n",
			"3: Missing arguments\n",
			"4: Extra arguments\n",
			"5: Exception",
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
		"Show show stacktrace if error.",
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
		cli.Usage(MissingArguments)
	} else if cli.NArg() > 2 {
		cli.Usage(ExtraArguments)
	}

	// Validate arguments match time format regex
	for _, arg := range cli.Args() {
		if !tdiff.RFC3339.MatchString(arg) {
			cli.Usage(InvalidArgument)
		}
	}
}
