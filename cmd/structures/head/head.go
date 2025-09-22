package head

import (
	"log"
	"mgit/cmd/paths"
	"mgit/cmd/storage"
	"mgit/cmd/structures/branch"
	"os"
)

func isBranchRef(text string) bool {
	for _, char := range text {
		if char == '/' {
			return true
		}
	}

	return false
}

func GetHeadHash() (string, error) {
	head, err := os.ReadFile(paths.HEAD)

	if err != nil {
		return "", err
	}

	hash := storage.GetHashFromRef(string(head))

	return hash, nil
}

// Search for branch first, not found?
//
// Then we check if there is a commit with that ref, if not we panic
func UpdateHead(ref string) {
	branchObj := branch.FindBranch(ref)

	if branchObj != nil {
		err := os.WriteFile(paths.HEAD, []byte(branchObj.Ref()), 0644)

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	object := storage.GetObjectByHash(ref)

	if object != nil {
		err := os.WriteFile(paths.HEAD, []byte(ref), 0644)

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	log.Fatalf("invalid ref %v", ref)
	// head, err := os.ReadFile(paths.HEAD)

	// if err != nil {
	// 	return
	// }

	// if isBranchRef(string(head)) {
	// 	// essa lógica toda eu acho que deveria ficar no storage
	// 	// então se for uma hash busca dentro de /objects
	// 	// se for um caminho tenta achar ele com base no próprio caminho no ref
	// 	parts := strings.Split(string(head), "ref: ")

	// 	if len(parts) >= 2 {
	// 		path := parts[1]

	// 		err := os.WriteFile(paths.RepoName+"/"+path, []byte(ref), 0644)

	// 		if err != nil {
	// 			log.Fatalf("failed to update head ref: %v", err)
	// 		}

	// 		return
	// 	}

	// 	log.Fatal("couldn't update ref, invalid head ref")
	// }

	// err = os.WriteFile(paths.HEAD, []byte(ref), 0644)

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
