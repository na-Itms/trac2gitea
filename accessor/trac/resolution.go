package trac

import "stevejefferson.co.uk/trac2gitea/log"

// GetResolutionNames retrieves all resolution names used in Trac tickets, passing each one to the provided "handler" function.
func (accessor *DefaultAccessor) GetResolutionNames(handlerFn func(resolution string)) {
	rows, err := accessor.db.Query(`SELECT DISTINCT resolution FROM ticket WHERE trim(resolution) != ''`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var resolution string
		if err := rows.Scan(&resolution); err != nil {
			log.Fatal(err)
		}

		handlerFn(resolution)
	}
}