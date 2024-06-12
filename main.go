package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func ReadFolder(showDetails bool, path string, showAll bool, maxDepth int, showEmojis bool, showOnlyFiles bool, showOnlyFolders bool) {
	readFolderRecursive(showDetails, path, showAll, 0, maxDepth, "", showEmojis, showOnlyFiles, showOnlyFolders)
}

func readFolderRecursive(showDetails bool, path string, showAll bool, depth, maxDepth int, prefix string, showEmojis bool, showOnlyFiles bool, showOnlyFolders bool) {
	if maxDepth >= 0 && depth > maxDepth {
		return
	}

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error al leer la carpeta: %s : %v\n", path, err)
		return
	}

	for i, f := range files {
		if !showAll && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		fileName := f.Name()
		filePath := filepath.Join(path, fileName)
		var details string
		if showDetails {
			details = getDetails(filePath)
		}

		if showOnlyFolders{
			if !f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			}
		} else if showOnlyFiles{
			if f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📄%s", fileName)
				}
			}
		} else {
			if f.IsDir() {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			} else if showEmojis {
				fileName = fmt.Sprintf("📄%s", fileName)
			}
		}

		if depth == 0 {
			if i == len(files)-1 {
				fmt.Printf("%s%s└──%s\n", details, prefix, fileName)
				if f.IsDir() {
					readFolderRecursive(showDetails, filePath, showAll, depth+1, maxDepth, prefix+"   ", showEmojis, showOnlyFiles, showOnlyFolders)
				}
			} else {
				fmt.Printf("%s%s├──%s\n", details, prefix, fileName)
				if f.IsDir() {
					readFolderRecursive(showDetails, filePath, showAll, depth+1, maxDepth, prefix+"│  ", showEmojis, showOnlyFiles, showOnlyFolders)
				}
			}
		} else {
			if i == len(files)-1 {
				fmt.Printf("%s%s└──%s\n", details, prefix, fileName)
				if f.IsDir() {
					readFolderRecursive(showDetails, filePath, showAll, depth+1, maxDepth, prefix+"   ", showEmojis, showOnlyFiles, showOnlyFolders)
				}
			} else {
				fmt.Printf("%s%s├──%s\n", details, prefix, fileName)
				if f.IsDir() {
					readFolderRecursive(showDetails, filePath, showAll, depth+1, maxDepth, prefix+"│  ", showEmojis, showOnlyFiles, showOnlyFolders)
				}
			}
		}
	}
}

func ListFiles(showDetails bool, path string, showAll bool, showEmojis bool, showOnlyFiles bool, showOnlyFolders bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error al leer la carpeta: %s : %v\n", path, err)
		return
	}

	var maxNameLen int
	for _, f := range files {
		if !showAll && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		fileName := f.Name()
		if showOnlyFolders{
			if !f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			}
		} else if showOnlyFiles{
			if f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📄%s", fileName)
				}
			}
		} else {
			if f.IsDir() {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			} else if showEmojis {
				fileName = fmt.Sprintf("📄%s", fileName)
			}
		}

		if len(fileName) > maxNameLen {
			maxNameLen = len(fileName)
		}
	}

	colWidth := maxNameLen + 2
	numCols := 80 / colWidth
	if numCols == 0 {
		numCols = 1
	}

	var details []string
	for _, f := range files {
		if !showAll && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		fileName := f.Name()
		filePath := filepath.Join(path, fileName)
		if showOnlyFolders{
			if !f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			}
		} else if showOnlyFiles{
			if f.IsDir(){
				continue
			} else {
				if showEmojis {
					fileName = fmt.Sprintf("📄%s", fileName)
				}
			}
		} else {
			if f.IsDir() {
				if showEmojis {
					fileName = fmt.Sprintf("📁%s", fileName)
				} else {
					fileName = fileName + "/"
				}
			} else if showEmojis {
				fileName = fmt.Sprintf("📄%s", fileName)
			}
		}

		if showDetails {
			details = append(details, fmt.Sprintf("%s %s", getDetails(filePath), fileName))
		} else {
			details = append(details, fileName)
		}
	}

	for i := 0; i < len(details); i += numCols {
		end := i + numCols
		if end > len(details) {
			end = len(details)
		}
		fmt.Println(strings.Join(details[i:end], strings.Repeat(" ", colWidth)))
	}
}


func getDetails(filePath string) string {
	info, err := os.Stat(filePath)
	if err != nil {
		return "Error obteniendo detalles"
	}

	sys := info.Sys().(*syscall.Stat_t)
	uid := sys.Uid
	u, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		return "Error obteniendo usuario"
	}

	perm := info.Mode().Perm()
	modTime := info.ModTime().Format(time.RFC3339)
	return fmt.Sprintf("%s %s %s", perm, u.Username, modTime)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [PATH]\n", os.Args[0])
		flag.PrintDefaults()
	}

	path := "."
	showAll := flag.Bool("a", false, "Show all files and folders, including those beginning with '.'.")
	maxDepth := flag.Int("depth", 0, "Maximum depth for folder traversal. Use -1 for unlimited depth.")
	showDetails := flag.Bool("details", false, "Show file details (permissions, creator, date).")
	showTree := flag.Bool("tree", false, "Show files and folders in a tree view.")
	showEmojis := flag.Bool("emoji", false, "Show emojis for files and folders.")
	showOnlyFiles := flag.Bool("f", false, "Show only files.")
	showOnlyFolders := flag.Bool("d", false, "Show only folders.")

	flag.Parse()

	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	if *showTree {
		if *showEmojis {
			fmt.Printf("📁%s\n", path)
		} else {
			fmt.Printf("%s/\n", path)
		}
		ReadFolder(*showDetails, path, *showAll, *maxDepth, *showEmojis, *showOnlyFiles, *showOnlyFolders)
	} else {
		ListFiles(*showDetails, path, *showAll, *showEmojis, *showOnlyFiles, *showOnlyFolders)
	}
}
