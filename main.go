package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Filei  string `short:"f" long:"file" env:"XMLFMT_FILEI" description:"The xml filepath if is - read from stdin" required:"true"`
	Prefix string `short:"p" long:"prefix" env:"XMLFMT_PREFIX" description:"Each element begins on a new line and this prefix"`
	Indent string `short:"i" long:"indent" env:"XMLFMT_INDENT" description:"Indent string for nested elements" default:"  "`
	Nested bool   `short:"n" long:"nested" env:"XMLFMT_NESTED" description:"Nested tags in comments"`
	Output string `short:"o" long:"output" description:"The output filepath"`
}

func main() {
	opts := &Options{}
	var parser = flags.NewParser(opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		return
	}

	var reader io.Reader
	if opts.Filei == "-" {
		reader = os.Stdin
	} else {
		file, err := os.Open(opts.Filei)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader = file
	}

	bb, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	out := xmlfmt.FormatXML(string(bb), opts.Prefix, opts.Indent, opts.Nested)
	out = strings.TrimSpace(out)

	var writer io.Writer
	if opts.Output == "" {
		writer = os.Stdout
	} else {
		err = os.MkdirAll(filepath.Dir(opts.Output), os.ModePerm)
		if err != nil {
			panic(err)
		}
		file, err := os.OpenFile(opts.Output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer = file
	}

	if _, err = writer.Write([]byte(out)); err != nil {
		panic(err)
	}
}
