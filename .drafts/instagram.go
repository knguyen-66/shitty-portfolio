package main

import (
	"fmt"
	"os"

	"github.com/Davincible/goinsta"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	insta := goinsta.New(os.Getenv("INSTAGRAM_USERNAME"), os.Getenv("INSTAGRAM_PASSWORD"))
	profile, err := insta.VisitProfile("knguyen.66")
	if err != nil {
		panic(err)
	}
	feed := profile.Feed
	fmt.Printf("%d posts fetched, more available = %v\n", len(feed.Items), feed.MoreAvailable)
}
