package gotwtr

import (
	"net/http"
	"sync"
)

type TweetField string

/*
	Tweet field will only return
	if you've also included the expansions=referenced_tweets.id query parameter in your request.
*/

const (
	TweetFieldID                 TweetField = "id"
	TweetFieldText               TweetField = "text"
	TweetEditHistoryIDs          TweetField = "edit_history_tweet_ids"
	TweetFieldAttachments        TweetField = "attachments"
	TweetFieldAuthorID           TweetField = "author_id"
	TweetFieldContextAnnotations TweetField = "context_annotations"
	TweetFieldConversationID     TweetField = "conversation_id"
	TweetFieldCreatedAt          TweetField = "created_at"
	TweetFieldEditControls       TweetField = "edit_controls"
	TweetFieldEntities           TweetField = "entities"
	TweetFieldInReplyToUserID    TweetField = "in_reply_to_user_id"
	TweetFieldLanguage           TweetField = "lang"
	TweetFieldNonPublicMetrics   TweetField = "non_public_metrics"
	TweetFieldOrganicMetrics     TweetField = "organic_metrics"
	TweetFieldPossiblySensitve   TweetField = "possibly_sensitive"
	TweetFieldPromotedMetrics    TweetField = "promoted_metrics"
	TweetFieldPublicMetrics      TweetField = "public_metrics"
	TweetFieldReferencedTweets   TweetField = "referenced_tweets"
	TweetReplySettings           TweetField = "reply_settings"
	TweetFieldSource             TweetField = "source"
	TweetFieldWithHeld           TweetField = "withheld"
	TweetFieldGeo                TweetField = "geo"
	TweetFieldMaxResults         TweetField = "max_results"
)

type Tweet struct {
	ID                 string                    `json:"id"`
	Text               string                    `json:"text"`
	EditHistoryIDs     []string                  `json:"edit_history_tweet_ids"`
	Attachments        *TweetAttachment          `json:"attachments,omitempty"`
	AuthorID           string                    `json:"author_id,omitempty"`
	ContextAnnotations []*TweetContextAnnotation `json:"context_annotations,omitempty"`
	ConversationID     string                    `json:"conversation_id,omitempty"`
	CreatedAt          string                    `json:"created_at"`
	Entities           *TweetEntity              `json:"entities,omitempty"`
	Geo                *TweetGeo                 `json:"geo,omitempty"`
	InReplyToUserID    string                    `json:"in_reply_to_user_id,omitempty"`
	Lang               string                    `json:"lang,omitempty"`
	NonPublicMetrics   *TweetMetrics             `json:"non_public_metrics,omitempty"`
	OrganicMetrics     *TweetMetrics             `json:"organic_metrics,omitempty"`
	PossiblySensitive  bool                      `json:"possibly_sensitive,omitempty"`
	PromotedMetrics    *TweetMetrics             `json:"promoted_metrics,omitempty"`
	PublicMetrics      *TweetMetrics             `json:"public_metrics,omitempty"`
	ReferencedTweets   []*TweetReferencedTweet   `json:"referenced_tweets,omitempty"`
	ReplySettings      string                    `json:"reply_settings,omitempty"`
	Source             string                    `json:"source,omitempty"`
	Withheld           *TweetWithheld            `json:"withheld,omitempty"`
}

type TweetAttachment struct {
	PollIDs   []string `json:"poll_ids"`
	MediaKeys []string `json:"media_keys"`
}

type TweetContextAnnotation struct {
	Domain *TweetContextObj `json:"domain"`
	Entity *TweetContextObj `json:"entity"`
}

type TweetContextObj struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetEntity struct {
	Annotations []*TweetAnnotation `json:"annotations"`
	Cashtags    []*TweetCashtag    `json:"cashtags"`
	Hashtags    []*TweetHashtag    `json:"hashtags"`
	Mentions    []*TweetMention    `json:"mentions"`
	URLs        []*TweetURL        `json:"urls"`
}

type TweetAnnotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type TweetCashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetHashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	UserName string `json:"user_name"`
}

type TweetURL struct {
	Start       int           `json:"start"`
	End         int           `json:"end"`
	URL         string        `json:"url"`
	ExpandedURL string        `json:"expanded_url"`
	DisplayURL  string        `json:"display_url"`
	Images      []*TweetImage `json:"images"`
	Status      int           `json:"status"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	UnwoundURL  string        `json:"unwound_url"`
}

type TweetImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type TweetGeo struct {
	Coordinates *TweetCoordinates `json:"coordinates"`
	PlaceID     string            `json:"place_id"`
}

type TweetCoordinates struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type TweetMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	LikeCount         int `json:"like_count"`
	ReplyCount        int `json:"reply_count"`
	RetweetCount      int `json:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
	QuoteCount        int `json:"quote_count"`
}

type TweetReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TweetWithheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
}

type TweetsResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type TweetResponse struct {
	Tweet    *Tweet              `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type TweetIncludes struct {
	Media  []*Media
	Places []*Place
	Polls  []*Poll
	Tweets []*Tweet
	Users  []*User
}

type UserTweetTimelineResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *UserTimelineMeta   `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type UserMentionTimelineResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *UserTimelineMeta   `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type UserTimelineMeta struct {
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	NextToken   string `json:"next_token"`
}

type SearchTweetsResponse struct {
	Tweets   []*Tweet            `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Meta     *SearchTweetsMeta   `json:"meta"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type SearchTweetsMeta struct {
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
	NextToken   string `json:"next_token,omitempty"`
}

type TimeseriesCount struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	TweetCount int    `json:"tweet_count"`
}

type TweetCountMeta struct {
	TotalTweetCount int    `json:"total_tweet_count"`
	NextToken       string `json:"next_token,omitempty"`
}

type TweetCountsResponse struct {
	Counts []*TimeseriesCount `json:"data"`
	Meta   *TweetCountMeta    `json:"meta"`
}

type RetweetsResponse struct {
	Users    []*User             `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
	Meta     *RetweetsLookupMeta `json:"meta"`
	Title    string              `json:"title,omitempty"`
	Detail   string              `json:"detail,omitempty"`
	Type     string              `json:"type,omitempty"`
}

type RetweetsLookupMeta struct {
	ResultCount int `json:"result_count"`
}

type RetrieveStreamRulesResponse struct {
	Rules  []*FilteredRule          `json:"data"`
	Meta   *RetrieveStreamRulesMeta `json:"meta"`
	Errors []*APIResponseError      `json:"errors,omitempty"`
}

type FilteredRule struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type RetrieveStreamRulesMeta struct {
	Sent string
}

type AddOrDeleteJSONBody struct {
	Add    []*AddRule  `json:"add,omitempty"`
	Delete *DeleteRule `json:"delete,omitempty"`
}

type AddRule struct {
	Value string `json:"value"`
	Tag   string `json:"tag,omitempty"`
}

type DeleteRule struct {
	IDs []string `json:"ids"`
}

type AddOrDeleteRulesResponse struct {
	Rules  []*FilteredRule       `json:"data"`
	Meta   *AddOrDeleteRulesMeta `json:"meta"`
	Errors []*APIResponseError   `json:"errors,omitempty"`
}

type AddOrDeleteRulesMeta struct {
	Sent    string                  `json:"sent"`
	Summary *AddOrDeleteMetaSummary `json:"summary"`
}

type AddOrDeleteMetaSummary struct {
	Created    int `json:"created"`
	NotCreated int `json:"not_created"`
	Deleted    int `json:"deleted"`
	NotDeleted int `json:"not_deleted"`
	Valid      int `json:"valid"`
	Invalid    int `json:"invalid"`
}

type ConnectToStreamResponse struct {
	Tweet         *Tweet          `json:"data"`
	Includes      *TweetIncludes  `json:"includes,omitempty"`
	MatchingRules []*MatchingRule `json:"matching_rules"`
}

type MatchingRule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type ConnectToStream struct {
	client *http.Client
	errCh  chan<- error
	ch     chan<- ConnectToStreamResponse
	done   chan struct{}
	wg     *sync.WaitGroup
}

type PostRetweetResponse struct {
	Retweeted *Retweeted          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type UndoRetweetResponse struct {
	Retweeted *Retweeted          `json:"data"`
	Errors    []*APIResponseError `json:"errors,omitempty"`
}

type Retweeted struct {
	Retweeted bool `json:"retweeted"`
}

type TweetBody struct {
	TweetID string `json:"tweet_id"`
}

type VolumeStreamsResponse struct {
	Tweet    *Tweet              `json:"data"`
	Includes *TweetIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type VolumeStreams struct {
	client *http.Client
	errCh  chan<- error
	ch     chan<- VolumeStreamsResponse
	done   chan struct{}
	wg     *sync.WaitGroup
}

type LookUpUsersWhoLikedWithheld struct {
	Scope        string   `json:"scope"`
	CountryCodes []string `json:"country_codes"`
}

type LookUpUsersWhoLikedEntity struct {
	URL         *LookUpUsersWhoLikedURL         `json:"url"`
	Description *LookUpUsersWhoLikedDescription `json:"description"`
}

type LookUpUsersWhoLikedURL struct {
	URLs []*LookUpUsersWhoLikedURLContent `json:"urls"`
}

type LookUpUsersWhoLikedDescription struct {
	URLs     []*LookUpUsersWhoLikedURLContent `json:"urls"`
	HashTags []*LookUpUsersWhoLikedHashTag    `json:"hashtags"`
	Mentions []*LookUpUsersWhoLikedMention    `json:"mentions"`
}

type LookUpUsersWhoLikedURLContent struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}

type LookUpUsersWhoLikedHashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	HashTag string `json:"hashtag"`
}

type LookUpUsersWhoLikedMention struct {
	Start    int                           `json:"start"`
	End      int                           `json:"end"`
	UserName string                        `json:"username"`
	CashTags []*LookUpUsersWhoLikedCashTag `json:"cashtags"`
}

type LookUpUsersWhoLikedCashTag struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	CashTag string `json:"cashtag"`
}

type LookUpUsersWhoLikedPublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type LookUpUsersWhoLiked struct {
	ID              string                            `json:"id"`
	Name            string                            `json:"name"`
	UserName        string                            `json:"username"`
	CreatedAt       string                            `json:"created_at,omitempty"`
	Protected       bool                              `json:"protected,omitempty"`
	Withheld        *LookUpUsersWhoLikedWithheld      `json:"withheld,omitempty"`
	Location        string                            `json:"location,omitempty"`
	URL             string                            `json:"url,omitempty"`
	Description     string                            `json:"description,omitempty"`
	Verified        bool                              `json:"verified,omitempty"`
	Entities        *LookUpUsersWhoLikedEntity        `json:"entities,omitempty"`
	ProfileImageURL string                            `json:"profile_image_url,omitempty"`
	PublicMetrics   *LookUpUsersWhoLikedPublicMetrics `json:"public_metrics,omitempty"`
	PinnedTweetID   string                            `json:"pinned_tweet_id,omitempty"`
}

type UsersLikingTweetResponse struct {
	Users    []*LookUpUsersWhoLiked       `json:"data"`
	Includes *LookUpUsersWhoLikedIncludes `json:"includes,omitempty"`
	Meta     *LookUpUsersWhoLikedMeta     `json:"meta"`
	Errors   []*APIResponseError          `json:"errors,omitempty"`
}

type PostUsersLikingTweetResponse struct {
	Liked  *Liked              `json:"data"`
	Errors []*APIResponseError `json:"errors,omitempty"`
}

type UndoUsersLikingTweetResponse struct {
	Liked  *Liked              `json:"data"`
	Errors []*APIResponseError `json:"errors,omitempty"`
}

type Liked struct {
	Liked bool `json:"liked"`
}

type UsersLikingBody struct {
	TweetID string `json:"tweet_id"`
}

type LookUpUsersWhoLikedIncludes struct {
	Tweets []*Tweet `json:"tweets"`
}

type LookUpUsersWhoLikedMeta struct {
	ResultCount int `json:"result_count"`
}

type TweetsUserLiked struct {
	ID                 string                               `json:"id"`
	Text               string                               `json:"text"`
	CreatedAt          string                               `json:"created_at,omitempty"`
	AuthorID           string                               `json:"author_id,omitempty"`
	ConversationID     string                               `json:"conversation_id,omitempty"`
	InReplyToUserID    string                               `json:"in_reply_to_user_id,omitempty"`
	ReferencedTweets   []*TweetsUserLikedReferencedTweets   `json:"referenced_tweets,omitempty"`
	Attachments        *TweetsUserLikedAttachments          `json:"attachments,omitempty"`
	Geo                *TweetsUserLikedGeo                  `json:"geo,omitempty"`
	ContextAnnotations []*TweetsUserLikedContextAnnotations `json:"context_annotations,omitempty"`
	Entities           *TweetsUserLikedEntities             `json:"entities,omitempty"`
	Withheld           *TweetsUserLikedWithheld             `json:"withheld,omitempty"`
	PublicMetrics      *TweetsUserLikedPublicMetrics        `json:"public_metrics,omitempty"`
	NonPublicMetrics   *TweetsUserLikedNonPublicMetrics     `json:"non_public_metrics,omitempty"` // requires the use of OAuth 1.0a User Context authentication.
	OrganicMetrics     *TweetsUserLikedOrganicMetrics       `json:"organic_metrics,omitempty"`    // requires user context authentication.
	PromotedMetrics    *TweetsUserLikedPromotedMetrics      `json:"promoted_metrics,omitempty"`   // requires user context authentication.
	PossiblySensitive  bool                                 `json:"possibly_sensitive,omitempty"`
	Lang               string                               `json:"lang,omitempty"`
	ReplySettings      string                               `json:"reply_settings,omitempty"`
	Source             string                               `json:"source,omitempty"`
	EditHistoryIDs     []string                             `json:"edit_history_ids,omitempty"`
	EditControls       *TweetsUserLikedEditControls         `json:"edit_controls,omitempty"`
}

type TweetsUserLikedReferencedTweets struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TweetsUserLikedAttachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIDs   []string `json:"poll_ids"`
}

type TweetsUserLikedGeo struct {
	Coordinates *TweetsUserLikedGeoCoordinates `json:"coordinates"`
	PlaceID     string                         `json:"place_id"`
}

type TweetsUserLikedGeoCoordinates struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type TweetsUserLikedContextAnnotations struct {
	Domain *TweetsUserLikedContextAnnotationsDomain `json:"domain"`
	Entity *TweetsUserLikedContextAnnotationsEntity `json:"entity"`
}

type TweetsUserLikedContextAnnotationsDomain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetsUserLikedContextAnnotationsEntity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetsUserLikedEntities struct {
	Annotations []*TweetsUserLikedEntitiesAnnotation `json:"annotations"`
	URLs        []*TweetsUserLikedEntitiesURLContent `json:"urls"`
	HashTags    []*TweetsUserLikedEntitiesHashTag    `json:"hashtags"`
	Mentions    []*TweetsUserLikedEntitiesMention    `json:"mentions"`
	CashTags    []*TweetsUserLikedEntitiesCashTag    `json:"cashtags"`
}

type TweetsUserLikedEntitiesAnnotation struct {
	Start          int    `json:"start"`
	End            int    `json:"end"`
	Probability    int    `json:"probability"`
	Type           string `json:"type"`
	NormalizedText string `json:"normalized_text"`
}

type TweetsUserLikedEntitiesURLContent struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	UnwoundURL  string `json:"unwound_url"`
}

type TweetsUserLikedEntitiesHashTag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetsUserLikedEntitiesMention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	UserName string `json:"username"`
}

type TweetsUserLikedEntitiesCashTag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type TweetsUserLikedWithheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type TweetsUserLikedPublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type TweetsUserLikedNonPublicMetrics struct {
	ImpressionCount   int `json:"impression_count"`
	URLLinkClicks     int `json:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks"`
}

type TweetsUserLikedOrganicMetrics struct {
	ImpressionCount   int `json:"impression_count"`    // requires the use of OAuth 1.0a User Context authentication.
	URLLinkClicks     int `json:"url_link_clicks"`     // requires the use of OAuth 1.0a User Context authentication.
	UserProfileClicks int `json:"user_profile_clicks"` // requires the use of OAuth 1.0a User Context authentication.
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}

type TweetsUserLikedPromotedMetrics struct {
	ImpressionCount   int `json:"impression_count"`    // requires the use of OAuth 1.0a User Context authentication.
	URLLinkClicks     int `json:"url_link_clicks"`     // requires the use of OAuth 1.0a User Context authentication.
	UserProfileClicks int `json:"user_profile_clicks"` // requires the use of OAuth 1.0a User Context authentication.
	RetweetCount      int `json:"retweet_count"`
	ReplyCount        int `json:"reply_count"`
	LikeCount         int `json:"like_count"`
}

type TweetsUserLikedResponse struct {
	Tweets   []*TweetsUserLiked   `json:"data"`
	Includes *TweetIncludes       `json:"includes,omitempty"`
	Meta     *TweetsUserLikedMeta `json:"meta"`
	Errors   []*APIResponseError  `json:"errors,omitempty"`
}

type TweetsUserLikedMeta struct {
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type TweetReply struct {
	ExcludeReplyUserIDs []string `json:"exclude_reply_user_ids"`
	InReplyToTweetID    string   `json:"in_reply_to_tweet_id"`
}

type PostTweetResponse struct {
	PostTweetData PostTweetData             `json:"data"`
	Errors        []*PostTweetResponseError `json:"errors,omitempty"`
	Title         string                    `json:"title,omitempty"`
	Detail        string                    `json:"detail,omitempty"`
	Type          string                    `json:"type,omitempty"`
}

type PostTweetData struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type DeleteTweetResponse struct {
	Data DeleteTweetData `json:"data"`
}

type DeleteTweetData struct {
	Deleted bool `json:"deleted"`
}

type HideRepliesResponse struct {
	HideRepliesResponseData *HideRepliesResponseData `json:"data"`
	Errors                  []*APIResponseError      `json:"errors,omitempty"`
}

type HideRepliesResponseData struct {
	Hidden bool `json:"hidden"`
}

type LookupUserBookmarksResponse struct {
	Tweets   []*Tweet                 `json:"data"`
	Includes *TweetIncludes           `json:"includes,omitempty"`
	Meta     *LookupUserBookmarksMeta `json:"meta"`
	Errors   []*APIResponseError      `json:"errors,omitempty"`
	Title    string                   `json:"title,omitempty"`
	Detail   string                   `json:"detail,omitempty"`
	Type     string                   `json:"type,omitempty"`
}

type LookupUserBookmarksMeta struct {
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type BookmarkTweetBody struct {
	TweetID string `json:"tweet_id"`
}

type BookmarkTweetResponse struct {
	BookmarkTweetData *BookmarkTweetData  `json:"data"`
	Errors            []*APIResponseError `json:"errors,omitempty"`
	Title             string              `json:"title,omitempty"`
	Detail            string              `json:"detail,omitempty"`
	Type              string              `json:"type,omitempty"`
}

type BookmarkTweetData struct {
	Bookmarked bool `json:"bookmarked"`
}

type RemoveBookmarkOfTweetResponse struct {
	RemoveBookmarkOfTweetData *RemoveBookmarkOfTweetData `json:"data"`
	Errors                    []*APIResponseError        `json:"errors,omitempty"`
	Title                     string                     `json:"title,omitempty"`
	Detail                    string                     `json:"detail,omitempty"`
	Type                      string                     `json:"type,omitempty"`
}

type RemoveBookmarkOfTweetData struct {
	Bookmarks bool `json:"bookmarks"`
}

type TweetsUserLikedEditControls struct {
	EditRemaining  int    `json:"edit_remaining"`
	IsEditEligible bool   `json:"is_edit_eligible"`
	EditableUntil  string `json:"editable_until"`
}
