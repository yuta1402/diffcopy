package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb/v3"
	"github.com/yuta1402/diffcopy/pkg/diffcopy"
)

func main() {
	flag.Usage = func() {
		usageText := `Usage:
    diffcopy [options] <src> <dest>
Options:
`
		fmt.Fprint(flag.CommandLine.Output(), usageText)
		flag.PrintDefaults()
	}

	var (
		outDir string
	)

	flag.StringVar(&outDir, "output", "", "output directory ( default: <dest> )")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	args := flag.Args()

	var (
		srcDir  = args[0]
		destDir = args[1]
	)

	if outDir == "" {
		outDir = destDir
	}

	files, err := diffcopy.FindWaitingFiles(srcDir, destDir)
	if err != nil {
		return
	}

	bar := pb.StartNew(len(files))

	for _, f := range files {
		p, _ := filepath.Rel(srcDir, f)
		outPath := filepath.Join(outDir, p)

		if err := diffcopy.CopyFile(f, outPath); err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}

		bar.Increment()
	}

	bar.Finish()
}
