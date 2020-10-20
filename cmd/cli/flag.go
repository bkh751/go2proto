package cli

import (
	"flag"
)

type ArrFlags []string

func (i *ArrFlags) String() string {
	return ""
}

func (i *ArrFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func ParseFlag() (filter *string, targetFolder *string, pkgFlags ArrFlags) {
	f := flag.String("filter", "", "Filter by struct names. Case insensitive.")
	t := flag.String("f", ".", "Protobuf output file path.")

	p := ArrFlags{}
	flag.Var(&pkgFlags, "p", `Fully qualified path of packages to analyse. Relative paths ("./example/in") are allowed.`)

	return f, t, p
}
