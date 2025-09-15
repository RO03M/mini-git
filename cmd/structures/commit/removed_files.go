package commit

import "mgit/cmd/paths"

// If the head has a file and the local dir doesn't, then it was deleted
func GetRemovedFiles() []string {
	localFileTree := paths.GetDirTree(".")
	headFileTree := HeadTrackedFilesTree()

	var localFileMap map[string]bool = make(map[string]bool)
	var removedFiles []string = []string{}

	for _, path := range localFileTree {
		localFileMap[path] = true
	}

	for _, path := range headFileTree {
		if _, found := localFileMap[path]; !found {
			removedFiles = append(removedFiles, path)
		}
	}

	return removedFiles
}
