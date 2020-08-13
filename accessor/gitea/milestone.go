package gitea

import (
	"database/sql"
	"fmt"

	"stevejefferson.co.uk/trac2gitea/log"
)

// AddMilestone adds a milestone to Gitea.
func (accessor *DefaultAccessor) AddMilestone(name string, content string, closed bool, deadlineTimestamp int64, closedTimestamp int64) {
	_, err := accessor.db.Exec(`
		INSERT INTO
			milestone(repo_id,name,content,is_closed,deadline_unix,closed_date_unix)
			SELECT $1,$2,$3,$4,$5,$6 WHERE
				NOT EXISTS (SELECT * FROM milestone WHERE repo_id = $1 AND name = $2)`,
		accessor.repoID, name, content, closed, deadlineTimestamp, closedTimestamp)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Added milestone %s\n", name)
}

// GetMilestoneID gets the ID of a named milestone - returns -1 if no such milestone
func (accessor *DefaultAccessor) GetMilestoneID(name string) int64 {
	var milestoneID int64 = -1
	err := accessor.db.QueryRow(`
		SELECT id FROM milestone WHERE name = $1
		`, name).Scan(&milestoneID)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	return milestoneID
}

// GetMilestoneURL gets the URL for accessing a given milestone
func (accessor *DefaultAccessor) GetMilestoneURL(milestoneID int64) string {
	repoURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/milestone/%d", repoURL, milestoneID)
}
