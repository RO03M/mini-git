package commit

func (c Commit) TrackedFilesTree() []string {
	tree := c.Tree

	if tree == nil {
		return []string{}
	}

	tree.LoadBlobs()
	blobs := tree.Blobs
	var paths []string = make([]string, len(blobs))

	for i, blob := range blobs {
		paths[i] = blob.FilePath
	}

	return paths
}

func HeadTrackedFilesTree() []string {
	lastCommit := GetCommitFromHead()

	if lastCommit == nil {
		return []string{}
	}

	return lastCommit.TrackedFilesTree()
}
