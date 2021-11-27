package gotwtr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func lookUpFollowing(ctx context.Context, c *client, id string, opt ...*FollowOption) (*FollowingResponse, error) {
	// check id
	if id == "" {
		return nil, errors.New("following by id: id parameter is required")
	}
	following := followingLookUp + "/" + id + "/following"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, following, nil)
	if err != nil {
		return nil, fmt.Errorf("following by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("following: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("following by id response: %w", err)
	}

	defer resp.Body.Close()

	var f FollowingResponse
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("following by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &f, &HTTPError{
			APIName: "following by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &f, nil
}

func lookUpFollowers(ctx context.Context, c *client, id string, opt ...*FollowOption) (*FollowersResponse, error) {
	// check id
	if id == "" {
		return nil, errors.New("followers by id: id parameter is required")
	}
	followers := followersLookUp + "/" + id + "/followers"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, followers, nil)
	if err != nil {
		return nil, fmt.Errorf("followers by id new request with ctx: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	var topt FollowOption
	switch len(opt) {
	case 0:
		// do nothing
	case 1:
		topt = *opt[0]
	default:
		return nil, errors.New("followers: only one option is allowed")
	}
	topt.addQuery(req)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("followers by id response: %w", err)
	}

	defer resp.Body.Close()

	var f FollowersResponse
	if err := json.NewDecoder(resp.Body).Decode(&f); err != nil {
		return nil, fmt.Errorf("followers by id decode: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return &f, &HTTPError{
			APIName: "followers by id",
			Status:  resp.Status,
			URL:     req.URL.String(),
		}
	}

	return &f, nil
}
