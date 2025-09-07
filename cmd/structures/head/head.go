package head

import (
	"errors"
	"log"
	"mgit/cmd/paths"
	"os"
	"strings"
)

func isBranchRef(text string) bool {
	return strings.HasPrefix(text, "ref:")
}

func GetHeadHash() (string, error) {
	head, err := os.ReadFile(paths.HEAD)

	if err != nil {
		return "", err
	}

	if isBranchRef(string(head)) {
		parts := strings.Split(string(head), "ref: ")

		if len(parts) >= 2 {
			path := parts[1]

			branch, err := os.ReadFile(paths.RepoName + "/" + path)
			return string(branch), err
		}

		return "", errors.New("invalid HEAD ref")
	}

	return string(head), err
}

func UpdateHead(hash string) {
	head, err := os.ReadFile(paths.HEAD)

	if err != nil {
		return
	}

	if isBranchRef(string(head)) {
		// essa lógica toda eu acho que deveria ficar no storage
		// então se for uma hash busca dentro de /objects
		// se for um caminho tenta achar ele com base no próprio caminho no ref
		parts := strings.Split(string(head), "ref: ")

		if len(parts) >= 2 {
			path := parts[1]

			err := os.WriteFile(paths.RepoName+"/"+path, []byte(hash), 0644)

			if err != nil {
				log.Fatalf("failed to update head ref: %v", err)
			}

			return
		}

		log.Fatal("couldn't update ref, invalid head ref")
	}

	err = os.WriteFile(paths.HEAD, []byte(hash), 0644)

	if err != nil {
		log.Fatal(err)
	}
}
