package main

import (
	"gitlab.com/mjwhitta/cli"
	"gitlab.com/mjwhitta/log"
	"gitlab.com/mjwhitta/tdiff"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			log.ErrX(Exception, r.(error).Error())
		}
	}()

	validate()

	var e error
	var td *tdiff.TDiff

	if td, e = tdiff.New(cli.Arg(0), cli.Arg(1)); e != nil {
		panic(e)
	}

	if flags.years {
		log.Good(td.YearsShorthand())
	} else {
		log.Good(td.Years())
		log.Good(td.Months())
		log.Good(td.Days())
		log.Good(td.Hours())
		log.Good(td.Minutes())
		log.Good(td.Seconds())
		log.Good(td.Milliseconds())
		log.Good(td.Microseconds())
		log.Good(td.Nanoseconds())
	}
}
