package database

import (
	"fmt"
	matchesResult "pc3r/matchesResult"
	"time"
)

/*
In this function we are retrieving all the upcoming matches
and matches results and pushing the into our internal database.
*/
func PushData() {
	CURRENT_TIME := time.Now().Format("20060102")
	fmt.Println("Starting the process of pushing data into the database: ", CURRENT_TIME)

	matches, err := matchesResult.GetMatchesResultFromAPI()
	if err != nil {
		return
	}

	for _, matchResult := range matches {

		_, err :=  matchesResult.PushMatchesResults(matchResult, time.Now())

		if err != nil {
			fmt.Println("Error while trying to push all Matches Results")
			return 
		}
	}
	

}