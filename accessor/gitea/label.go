package gitea

import (
	"fmt"
	"log"
)

// AddLabel adds a label to Gitea.
func (accessor *Accessor) AddLabel(label string, color string) {
	_, err := accessor.db.Exec(`
		INSERT INTO label(repo_id,name,color)
			SELECT $1,$2, $3 WHERE
			NOT EXISTS ( SELECT * FROM label WHERE repo_id = $1 AND name = $2 )`,
		accessor.repoID, label, color)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added label", label)
}