package model

import (
	"fmt"
	"log"
)

func SetRecord(stateID int, userID int) {
	sqlStatement := `INSERT INTO records (state_id, user_id)
					VALUES ($1, $2)`
	_, err := DB.Exec(sqlStatement, stateID, userID)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRecords() ([]int, []int) {
	rows, err := DB.Query("SELECT state_id, user_id FROM records")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var stateIDArr []int
	var userIDArr []int

	for rows.Next() {
		var stateID int64
		var userID int64
		if err := rows.Scan(&stateID, &userID); err != nil {
			log.Fatal(err)
		}
		stateIDArr = append(stateIDArr, int(stateID))
		userIDArr = append(userIDArr, int(userID))
	}
	fmt.Println(stateIDArr, userIDArr)
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}
	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return stateIDArr, userIDArr
}
