package main

import (
	"context"
	"fmt"

	"github.com/sivchari/gotwtr"
)

func main() {
	client := gotwtr.New("key")
	// get users who liked the tweet that id is "tweet_id"
	t, err := client.LikesLookUpUsers(context.Background(), "tweet_id")
	if err != nil {
		panic(err)
	}
	fmt.Println("------no option------")
	for _, user := range t.Users {
		fmt.Printf("id: %s, name: %s\n", user.ID, user.UserName)
	}

	// get users who liked the tweet that id is "tweet_id" with option
	t, err = client.LikesLookUpUsers(context.Background(), "tweet_id", &gotwtr.LikesLookUpByTweetOpts{
		Expansions:  []gotwtr.Expansion{gotwtr.ExpansionPinnedTweetID},
		UserFields:  []gotwtr.UserField{gotwtr.UserFieldCreatedAt},
		TweetFields: []gotwtr.TweetField{gotwtr.TweetFieldCreatedAt},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("------with option------")
	for _, user := range t.Users {
		fmt.Printf("id: %s, name: %s, created_at: %v\n", user.ID, user.UserName, user.CreatedAt)
	}
	for _, tweet := range t.Includes.Tweets {
		fmt.Printf("tweet_id: %s, created_at: %v\n", tweet.ID, tweet.CreatedAt)
	}
}