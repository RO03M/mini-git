package commands

import "fmt"

func PrintHelp() {
	fmt.Print("available commands:\n\n")
	fmt.Println("init\tCreate an empty MGit repository")
	fmt.Println("add\tStage files to the storage.\n\tusage: mgit add file1 file2 file3")
	fmt.Println("commit\tA commit is a reference in time of all the staged files. Create one.\n\tusage: mgit commit \"your commit message\"")
	fmt.Println("log\tShow commit logs")
	fmt.Println("checkout\tMove the head ref to a point in history")
	fmt.Println("trackable\tShow the files that MGit can see and track")
	fmt.Println("tracked\tShow files that has (at least) an older copy of itself in the current HEAD")
	fmt.Println("untracked\tShow files that has no track in the current HEAD (files that are new to mgit)")
	fmt.Println("rf\tShow files that have been removed and haven't been committed")
	fmt.Println("cat-file\tGet uncompressed content of a file in the of the object storage")
}
