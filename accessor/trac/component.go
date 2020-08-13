package trac

import "stevejefferson.co.uk/trac2gitea/log"

// GetComponentNames retrieves all Trac component names, passing each one to the provided "handler" function.
func (accessor *DefaultAccessor) GetComponentNames(handlerFn func(cmptName string)) {
	rows, err := accessor.db.Query(`SELECT name FROM component`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var cmptName string
		if err := rows.Scan(&cmptName); err != nil {
			log.Fatal(err)
		}

		handlerFn(cmptName)
	}
}
