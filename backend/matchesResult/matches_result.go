package matchesResult

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	db "pc3r/prisma"
	prismaDb "pc3r/prisma/db"
	"time"

	godotenv "github.com/joho/godotenv"
)

func fetchFromUrl(URL string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	
	if err != nil {
		return nil, err
	}
	
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func GetMatchesResultFromAPI() ([]MatchesResult, error) {
	// Retrieve my API base url from the .env file
	godotenv.Load(".env")
	BASE_URL := os.Getenv("API_BASE_URL")	

	URL := fmt.Sprintf("%s%s",BASE_URL, "match/results")
	res, err := fetchFromUrl(URL)
	if err != nil {
		fmt.Println("Error occured when retrieving data from API")
		return []MatchesResult{}, err
	}
	var results MatchesResultResponse
	err = json.NewDecoder(res.Body).Decode(&results)
	
	if err != nil {
		fmt.Println("Could not decode json")
		return []MatchesResult{}, err
	}
	return results.Data.Segments, nil

}

func CreateChatForMatchResult(name string, date time.Time) (*prismaDb.ChatModel, error) {
	prisma, ctx := db.GetPrisma()
	chat, err := prisma.Chat.CreateOne(
		prismaDb.Chat.Name.Set(name),
		prismaDb.Chat.Date.Set(date),
	).Exec(ctx)
	return chat, err
}

func PushMatchesResults(match MatchesResult, date time.Time) (*prismaDb.MatchResultModel, error) {
	prisma, ctx := db.GetPrisma()
	
	// avant de faire le push il faut vérifier que ce tuple n'existe pas dans la base de données // important
	result, err := prisma.MatchResult.FindMany(
		prismaDb.MatchResult.And(
			prismaDb.MatchResult.Team1.Equals(match.Team1), 
			prismaDb.MatchResult.Team2.Equals(match.Team2), 
			prismaDb.MatchResult.RoundInfo.Equals(match.Round_info),
			prismaDb.MatchResult.TournamentName.Equals(match.Tournament_name))).Exec(ctx)

	if len(result) != 0 {
		if !errors.Is(err, prismaDb.ErrNotFound) {
			fmt.Println("Match already existing in the database, No need to insert it to the database")
			return nil, nil
		}
	}

	chatName := match.Team1 + " VS " + match.Team2
	chat, err := CreateChatForMatchResult(chatName,date)
	if err != nil {
		fmt.Println("Error while creating a chat for a Match result")
	}

	createdMatch, err := prisma.MatchResult.CreateOne(
		prismaDb.MatchResult.Team1.Set(match.Team1),
		prismaDb.MatchResult.Team2.Set(match.Team2),
		prismaDb.MatchResult.Score1.Set(match.Score1),
		prismaDb.MatchResult.Score2.Set(match.Score2),
		prismaDb.MatchResult.Flag1.Set(match.Flag1),
		prismaDb.MatchResult.Flag2.Set(match.Flag2),
		prismaDb.MatchResult.TimeCompleted.Set(match.Time_completed),
		prismaDb.MatchResult.RoundInfo.Set(match.Round_info),
		prismaDb.MatchResult.TournamentName.Set(match.Tournament_name),
		prismaDb.MatchResult.MatchPage.Set(match.Match_page),
		prismaDb.MatchResult.TournamentIcon.Set(match.Tournament_icon),
		prismaDb.MatchResult.Chat.Link(
			prismaDb.Chat.ID.Equals(chat.ID),
		),
	).Exec(ctx)

	if err != nil {
		fmt.Println("Error while pushing a Match result")
	}

	return createdMatch, nil
}