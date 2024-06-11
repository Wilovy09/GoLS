package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFolder(path string, showAll bool, maxDepth int) {
	readFolderRecursive(path, showAll, 0, maxDepth)
}

func readFolderRecursive(path string, showAll bool, depth, maxDepth int) {
	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error al leer la carpeta: %s : %v\n", path, err)
		return
	}

	for i, f := range files {
		var prefix string
		if i == len(files)-1 {
			prefix = "└──"
		} else {
			prefix = "├──"
		}
		if !showAll && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		var fileType string
		if f.IsDir() {
			fileType = "📁"
		} else {
			fileType = "📄"
		}
		fmt.Printf("%s%s%s\n", getIndentation(depth)+prefix, fileType, f.Name())
		if f.IsDir() {
			subdir := filepath.Join(path, f.Name())
			readFolderRecursive(subdir, showAll, depth+1, maxDepth)
		}
	}
}

func getIndentation(depth int) string {
	if depth == 0 {
		return ""
	}
	return strings.Repeat("│  ", depth-1) + "├──"
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [PATH]\n", os.Args[0])
		flag.PrintDefaults()
	}

	path := "."
	showAll := flag.Bool("a", false, "Show all files and folders, including those beginning with '.'.")
	maxDepth := flag.Int("depth", 0, "Maximum depth for folder traversal. Use -1 for unlimited depth.")

	flag.Parse()

	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	fmt.Printf("📁%s\n", path)
	ReadFolder(path, *showAll, *maxDepth)
}
