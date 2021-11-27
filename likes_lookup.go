package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func usersLikingTweet(ctx context.Context, c *client, tweetID string, opt ...*UsersLikingTweetOption) (*UsersLikingTweetResponse, error) {
	if tweetID == "" {
		return nil, errors.New("users liking tweet: tweet id parameter is required")
	}
	ep := fmt.Sprintf(usersLikingTweetURL, tweetID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, fmt.Errorf("users liking tweet new request with ctx: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var uopt UsersLikingTweetOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		uopt = *opt[0]
	default:
		return nil, errors.New("users liking tweet: only one option is allowed")
	}
	uopt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("users liking tweet: %w", err)
	}
	defer resp.Body.Close()

	var ultr UsersLikingTweetResponse
	if err := json.NewDecoder(resp.Body).Decode(&ultr); err != nil {
		return nil, fmt.Errorf("users liking tweet: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &ultr, &HTTPError{
			APIName: "users liking tweet",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &ultr, nil
}
