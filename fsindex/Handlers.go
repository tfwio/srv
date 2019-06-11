package fsindex

import (
	"fmt"
	"os"
	"path/filepath"

	"tfw.io/Go/fsindex/util"
)

// CBPath is a simple callback;
// if you return true, then the caller function immediately returns.
type CBPath func(*Model, *PathEntry) bool // (*interface{}, error)
// CBFile is a simple callback
// if you return true, then the caller function immediately returns.
type CBFile func(*Model, *FileEntry) bool // (*interface{}, error)

// FileHandler is a simple callback.
type FileHandler func(*Model, *FileEntry) bool

// PathHandler is a simple callback.
type PathHandler func(*Model, *PathEntry) bool

// Handlers contains simple callbacks.
type Handlers struct {
	ChildPath PathHandler
	ChildFile FileHandler
}

// RefreshCB refreshes child directories and files.
// parameter `rootPathEntry`: root-path entry.
// parameter `counter (*int32)`: pointer to our indexing integer (counter).
// parameter `callback (RefreshAction)` is a method (if defined) which
//                                      can be used arbitrarily.
//
// Only difference here is that we're using CBPath and CBFile as opposed to a
// Handlers structure which contains callbacks for our `Refresh(…)`.
func (p *PathEntry) RefreshCB(rootPathEntry *Model, counter *(int32), cbPath *CBPath, cbFile *CBFile) {

	// create a reference node pointing to the tree-root
	var mRoot *Model

	// if the first parent element is root, we need to build some
	// reference memory (dictionary of ignore-paths).
	if p.IsRoot {
		mRoot = &Model{PathEntry: *p}
		// build absolute path list to ignore.
		for i := 0; i < len(p.IgnorePaths); i++ {
			p.IgnorePaths[i], _ = filepath.Abs(p.IgnorePaths[i])
		}
	} else {
		mRoot = rootPathEntry // assign mRoot
	}

	p.Index = *counter // Assign index
	*counter++

	if cbPath != nil {
		if (*cbPath)(mRoot, p) {
			return
		}
	}

	mPaths, mError := filepath.Glob(fmt.Sprintf("%s/*", p.FullPath))
	if mError != nil {
		fmt.Println("error in path:", mError)
		return
	}

	// FILE PATHS
	for _, mFullPath := range mPaths {

		fileinfo, err := os.Stat(mFullPath)
		if os.IsNotExist(err) {
			fmt.Println("Error reading file")
			return
		}

		if !fileinfo.IsDir() {

			for i := 0; i < len(mRoot.FileFilter); i++ {

				if mRoot.FileFilter[i].Match(mFullPath) {

					var child = FileEntry{
						Parent:    p,
						FullPath:  mFullPath,
						SHA1:      util.Sha1String(mFullPath),
						Name:      util.StripFileExtension(filepath.Base(mFullPath)),
						Extension: filepath.Ext(mFullPath),
					}
					child.Path = util.UnixSlash(util.Cat(mRoot.FauxPath, "/", child.RootedPath(mRoot)))
					p.Files = append(p.Files, child)
					if cbFile != nil {
						if (*cbFile)(mRoot, &child) {
							return
						}
						// println(fmt.Sprintf("  - %s", child.Base()))
					}

				}
			}
		}
	}

	// DIRECTORY PATHS
	for _, mFullPath := range mPaths {

		fileinfo, err := os.Stat(mFullPath)
		if os.IsNotExist(err) {
			fmt.Println("Error reading file")
			return
		}

		if fileinfo.IsDir() {

			var child = PathEntry{
				PathSpec: PathSpec{
					FileEntry: FileEntry{
						Parent:    p,
						FullPath:  mFullPath,
						SHA1:      util.Sha1String(mFullPath),
						Name:      util.StripFileExtension(filepath.Base(mFullPath)),
						Extension: filepath.Ext(mFullPath),
						Mod:       fileinfo.ModTime(),
					},
					IsRoot: false,
				},
			}
			child.Path = util.UnixSlash(util.Cat(mRoot.FauxPath, "/", child.RootedPath(mRoot)))

			if child.IsIgnore(mRoot) {

				fmt.Printf("- ignored: %s\n", child.FullPath)

			} else {

				child.RefreshCB(mRoot, counter, cbPath, cbFile)

				p.Paths = append(p.Paths, child)
			}

		}
	}
}

// var (
// 	xCounter   int32
// 	xpCounter  *int32
// 	localMedia = FileSpec{
// 		Name: "Media (images)",
// 		Extensions: []string{
// 			".bmp",
// 			".jpg",
// 			".svg",
// 			".png",
// 			".gif",
// 		},
// 	}
// 	localMarkdown = FileSpec{
// 		Name: "Markdown (hyper-text)",
// 		Extensions: []string{
// 			".md",
// 			".mmd",
// 		},
// 	}
// )

//func main() {
//	// flag.Parse()
//	// appRootPath := filepath.Dir(flag.Arg(0))
//	// root, _ := filepath.Abs(appRootPath)
//	// fmt.Println(fmt.Sprintf("root path: %s", root))
//	xCounter = 0
//	xpCounter = &xCounter
//	*xpCounter++
//	fmt.Printf("counter %d\n", *xpCounter)
//	*xpCounter++
//	fmt.Printf("counter %d\n", *xpCounter)
//	*xpCounter++
//	fmt.Printf("counter %d\n", *xpCounter)
//	startPath := "c:\\users\\tfwro\\.mmd"
//	rootPath := PathEntry{
//		PathSpec: PathSpec{
//			FileEntry: FileEntry{
//				Parent:   nil,
//				FullPath: startPath,
//			},
//			isRoot: true,
//		},
//		FileFilter: []FileSpec{localMedia, localMarkdown},
//		IgnorePaths: []string{
//			"c:\\users\\tfwro\\.mmd\\reveal.js",
//			"c:\\users\\tfwro\\.mmd\\.git",
//			"c:\\users\\tfwro\\.mmd\\.vscode",
//		},
//	}
//	xCounter = 0
//	rootPath.Refresh(nil, &xCounter)
//}
//
