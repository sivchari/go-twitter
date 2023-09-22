package gotwtr_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sivchari/gotwtr"
)

func Test_postTweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx    context.Context
		client *http.Client
		body   *gotwtr.PostTweetOption
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.PostTweetResponse
		wantErr bool
	}{
		{
			name: "201 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `
                        {
                            "data": {
                                "id": "1445880548472328192",
                                "text": "Hello world!"
                            }
                        }
                    `
					return &http.Response{
						StatusCode: http.StatusCreated,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				body: &gotwtr.PostTweetOption{
					Text: "Hello world!",
				},
			},
			want: &gotwtr.PostTweetResponse{
				PostTweetData: gotwtr.PostTweetData{
					ID:   "1445880548472328192",
					Text: "Hello world!",
				},
			},
			wantErr: false,
		},
		{
			name: "400 bad request payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := fmt.Sprintf(`
												{
														"errors": [{
																"parameters": { 
																		"text": ["%s"]
																},
																"message": "The Tweet contains an invalid URL."
														}],
														"title": "Invalid Request",
														"detail": "One or more parameters to your request was invalid.",
														"type": "https://api.twitter.com/2/problems/invalid-request"
												}
                    `, strings.Repeat("x", 281))
					return &http.Response{
						StatusCode: http.StatusBadRequest,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				body: &gotwtr.PostTweetOption{
					Text: strings.Repeat("x", 281),
				},
			},
			want: &gotwtr.PostTweetResponse{
				Errors: []*gotwtr.PostTweetResponseError{{
					Parameters: gotwtr.PostTweetResponseErrorParameter{
						Text: []string{strings.Repeat("x", 281)},
					},
					Message: "The Tweet contains an invalid URL.",
				},
				},
				Title:  "Invalid Request",
				Detail: "One or more parameters to your request was invalid.",
				Type:   "https://api.twitter.com/2/problems/invalid-request",
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.PostTweet(tt.args.ctx, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostTweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PostTweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}
		})
	}
}

func Test_deleteTweet(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx     context.Context
		client  *http.Client
		tweetID string
	}
	tests := []struct {
		name    string
		args    args
		want    *gotwtr.DeleteTweetResponse
		wantErr bool
	}{
		{
			name: "200 ok default payload",
			args: args{
				ctx: context.Background(),
				client: mockHTTPClient(func(request *http.Request) *http.Response {
					body := `
                        {
                            "data": {
                                "deleted": true
                            }
                        }
                    `
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
					}
				}),
				tweetID: "1445880548472328192",
			},
			want: &gotwtr.DeleteTweetResponse{
				Data: gotwtr.DeleteTweetData{
					Deleted: true,
				},
			},
			wantErr: false,
		},
		{
			name: "empty tweet id",
			args: args{
				ctx:     context.Background(),
				tweetID: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		i := i
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := gotwtr.New("test-key", gotwtr.WithHTTPClient(tt.args.client))
			got, err := c.DeleteTweet(tt.args.ctx, tt.args.tweetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTweet() index = %v error = %v, wantErr %v", i, err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("DeleteTweet() index = %v mismatch (-want +got):\n%s", i, diff)
				return
			}

		})
	}
}
