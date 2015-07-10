package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unirita/cmptxt/comparer"
)

type arguments struct {
	freeOrderFlag  bool
	ignorePatterns []string
	baseFilePath   string
	targetFilePath string
}

func main() {
	args := fetchArgs()
	if args == nil {
		fmt.Println("Argument error.")
		flag.Usage()
	}

	var baseFile *os.File
	var targetFile *os.File
	var err error

	baseFile, err = os.Open(args.baseFilePath)
	if err != nil {
		fmt.Println("File open error: ", args.baseFilePath)
		fmt.Println(err)
		return
	}
	defer baseFile.Close()

	targetFile, err = os.Open(args.targetFilePath)
	if err != nil {
		fmt.Println("File open error: ", args.targetFilePath)
		fmt.Println(err)
		return
	}
	defer targetFile.Close()

	c := comparer.New()
	for _, ptn := range args.ignorePatterns {
		if err := c.AddIgnorePattern(ptn); err != nil {
			fmt.Println("Regular expression format error: ", ptn)
			fmt.Println(err)
			return
		}
	}

	var isSame bool
	if args.freeOrderFlag {
		isSame = c.CompareFreeOrder(baseFile, targetFile)
	} else {
		isSame = c.Compare(baseFile, targetFile)
	}

	if isSame {
		fmt.Println("Same.")
	} else {
		fmt.Println("Different.")
	}
}

func fetchArgs() *arguments {
	args := new(arguments)
	args.ignorePatterns = make([]string, 0)
	var hasIgnorePattern bool
	flag.Usage = showUsage
	flag.BoolVar(&args.freeOrderFlag, "f", false, "Compare with free line order.")
	flag.BoolVar(&hasIgnorePattern, "i", false, "Use ignore pattern.")
	flag.Parse()

	if flag.NArg() < 2 {
		return nil
	}
	otherArgs := flag.Args()
	if hasIgnorePattern {
		for i := 0; i < flag.NArg()-2; i++ {
			args.ignorePatterns = append(args.ignorePatterns, otherArgs[i])
		}
	}

	args.baseFilePath = otherArgs[flag.NArg()-2]
	args.targetFilePath = otherArgs[flag.NArg()-1]

	return args
}

func showUsage() {
	usage := `Usage:
    cmptxt [-f] [-i pattern1 pattern2 ...] file1 file2

Options:
    -f : Compare files in free order.
    -i : Use ignore patterns.
`
	fmt.Println(usage)
}
