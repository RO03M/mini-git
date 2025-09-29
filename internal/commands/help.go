package commands

import "fmt"

func Help() {
	fmt.Print("available commands:\n\n")
	fmt.Println("init\tCreate an empty MGit repository")
	fmt.Println("add\tStage files to the storage.\n\tusage: mgit add file1 file2 file3")
	fmt.Println("commit\tA commit is a reference in time of all the staged files. Create one.\n\tusage: mgit commit \"your commit message\"")
	fmt.Println("log\tShow commit logs")
	fmt.Println("checkout\tMove the head ref to a point in history")
	fmt.Println("switch\t\tGo to another branch")
	fmt.Println("trackable\tShow the files that MGit can see and track")
	fmt.Println("rev-parse\tResolve the hash commit from any mgit pointer (e.g branch, HEAD)")
	fmt.Println("rf\t\tShow files that have been removed and haven't been committed")
	fmt.Println("cat-file\tGet uncompressed content of a file in the of the object storage")
}
