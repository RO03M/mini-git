package paths

import "fmt"

const RepoName string = ".mgit"
const HEAD string = RepoName + "/HEAD"

func GetBranchPath(name string) string {
	return fmt.Sprintf("%s/refs/heads/%s", RepoName, name)
}
