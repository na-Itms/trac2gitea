package trac

import "log"

// GetSeverityNames retrieves all severity names used in Trac tickets, passing each one to the provided "handler" function.
func (accessor *Accessor) GetSeverityNames(handlerFn func(string)) {
	rows, err := accessor.db.Query(`SELECT DISTINCT COALESCE(severity,'') FROM ticket`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var severityName string
		if err := rows.Scan(&severityName); err != nil {
			log.Fatal(err)
		}

		handlerFn(severityName)
	}
}
