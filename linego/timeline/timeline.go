package timeline

import (
	"../service"
	"../channel"
	"../config"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
)

var client = &http.Client{}
var request = &http.Request{}

func LoginTimeline() {
	request.Header = make(http.Header)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", config.USER_AGENT)
	request.Header.Add("X-Line-Mid", service.MID)
	request.Header.Add("X-Line-Carrier", config.CARRIER)
	request.Header.Add("X-Line-Application", config.LINE_APPLICATION)
	request.Header.Add("X-Line-ChannelToken", channel.ChannelResult.ChannelAccessToken)
}

func TestResult() {
	u, _ := url.Parse(config.LINE_TIMELINE_MH+"/album/v3/albums.json?homeId=c6b3655da6da814cb30002f3012fc041d&type=g&sourceType=TALKROOM")
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

/* Timeline */

type Feed struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		FeedInfos []struct {
			Type   string      `json:"type"`
			ID     string      `json:"id"`
			Status string      `json:"status"`
			Score  interface{} `json:"score"`
		} `json:"feedInfos"`
		RequestTime int64 `json:"requestTime"`
		Feeds       []struct {
			FeedInfo struct {
				Type   string      `json:"type"`
				ID     string      `json:"id"`
				Status string      `json:"status"`
				Score  interface{} `json:"score"`
			} `json:"feedInfo"`
			Post struct {
				UserInfo struct {
					Mid        string `json:"mid"`
					Nickname   string `json:"nickname"`
					PictureURL string `json:"pictureUrl"`
					UserValid  bool   `json:"userValid"`
					WriterMid  string `json:"writerMid"`
				} `json:"userInfo"`
				PostInfo struct {
					AppSn        int      `json:"appSn"`
					HomeID       string   `json:"homeId"`
					PostID       string   `json:"postId"`
					Status       string   `json:"status"`
					LikeCount    int      `json:"likeCount"`
					CommentCount int      `json:"commentCount"`
					Liked        bool     `json:"liked"`
					TopLikes     []string `json:"topLikes"`
					URL          struct {
						Type      string `json:"type"`
						TargetURL string `json:"targetUrl"`
					} `json:"url"`
					ReadPermission struct {
						Type  string      `json:"type"`
						Gids  interface{} `json:"gids"`
						Count int         `json:"count"`
					} `json:"readPermission"`
					AllowShare            bool   `json:"allowShare"`
					AllowLikeShare        bool   `json:"allowLikeShare"`
					AllowComment          bool   `json:"allowComment"`
					AllowPreviewComment   bool   `json:"allowPreviewComment"`
					AllowPhotoComment     bool   `json:"allowPhotoComment"`
					AllowLike             bool   `json:"allowLike"`
					AllowRecall           bool   `json:"allowRecall"`
					AllowFriendRequest    bool   `json:"allowFriendRequest"`
					AllowCommentLike      bool   `json:"allowCommentLike"`
					AllowLikeProfiles     bool   `json:"allowLikeProfiles"`
					EnableCommentApproval bool   `json:"enableCommentApproval"`
					HasSharedToPost       bool   `json:"hasSharedToPost"`
					CommentLinkPermission string `json:"commentLinkPermission"`
					LikeLinkPermission    string `json:"likeLinkPermission"`
					OfficialHome          struct {
						HomeType    string `json:"homeType"`
						ApproveType string `json:"approveType"`
						HomeManager bool   `json:"homeManager"`
						HomeUse     bool   `json:"homeUse"`
						FriendCount int    `json:"friendCount"`
					} `json:"officialHome"`
					AllowEdit   bool  `json:"allowEdit"`
					CreatedTime int64 `json:"createdTime"`
					UpdatedTime int64 `json:"updatedTime"`
					SharedCount struct {
						ToTalk int `json:"toTalk"`
						ToPost int `json:"toPost"`
					} `json:"sharedCount"`
				} `json:"postInfo"`
				Contents struct {
					Text          string `json:"text"`
					ContentsStyle struct {
						TextStyle struct {
						} `json:"textStyle"`
						StickerStyle struct {
						} `json:"stickerStyle"`
						MediaStyle struct {
						} `json:"mediaStyle"`
					} `json:"contentsStyle"`
				} `json:"contents"`
				AdditionalContents struct {
					Title        string `json:"title"`
					Main         string `json:"main"`
					Sub          string `json:"sub"`
					Obsthumbnail struct {
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						PreferCdn    bool   `json:"preferCdn"`
						ObjectID     string `json:"objectId"`
						ServiceName  string `json:"serviceName"`
						ObsNamespace string `json:"obsNamespace"`
					} `json:"obsthumbnail"`
					Thumbnail struct {
						URL         string      `json:"url"`
						Width       interface{} `json:"width"`
						Height      interface{} `json:"height"`
						RequiredTid bool        `json:"requiredTid"`
					} `json:"thumbnail"`
					URL struct {
						Type      string `json:"type"`
						TargetURL string `json:"targetUrl"`
						MarketURL string `json:"marketUrl"`
					} `json:"url"`
				} `json:"additionalContents"`
				StatisticInfo struct {
					ContentType int `json:"contentType"`
				} `json:"statisticInfo"`
			} `json:"post"`
		} `json:"feeds"`
	} `json:"result"`
}
func GetFeed(postLimit string, commentLimit string, likeLimit string, order string) (*Feed, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v39/feed/list.json?postLimit="+postLimit+"&commentLimit="+commentLimit+"&likeLimit="+likeLimit+"&order="+order)
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var feed = new(Feed)
	err := json.Unmarshal(body, &feed)
	return feed, err
}

type HomeProfile struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		HomeInfo struct {
			HomeID       string `json:"homeId"`
			ObjectID     string `json:"objectId"`
			ObsNamespace string `json:"obsNamespace"`
			ServiceName  string `json:"serviceName"`
			PostCount    int    `json:"postCount"`
			UserInfo     struct {
				Mid          string `json:"mid"`
				Nickname     string `json:"nickname"`
				PictureURL   string `json:"pictureUrl"`
				VideoURLHash string `json:"videoUrlHash"`
				UserValid    bool   `json:"userValid"`
				WriterMid    string `json:"writerMid"`
			} `json:"userInfo"`
		} `json:"homeInfo"`
	} `json:"result"`
}
func GetHomeProfile(mid string, postLimit string, commentLimit string, likeLimit string) (*HomeProfile, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v39/post/list.json?homeId="+mid+"&postLimit="+postLimit+"&commentLimit="+commentLimit+"&likeLimit="+likeLimit+"&sourceType=LINE_PROFILE_COVER")
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var homeprofile = new(HomeProfile)
	err := json.Unmarshal(body, &homeprofile)
	return homeprofile, err
}

type ProfileDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		HasNewPost   bool          `json:"hasNewPost"`
		ExpireTime   int           `json:"expireTime"`
		RecentPhotos []interface{} `json:"recentPhotos"`
		UserMid      string        `json:"userMid"`
		ServiceName  string        `json:"serviceName"`
		ObsNamespace string        `json:"obsNamespace"`
		ObjectID     string        `json:"objectId"`
	} `json:"result"`
}
func GetProfileDetail(mid string) (*ProfileDetail, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v1/userpopup/getDetail.json?userMid="+mid)
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var profiledetail = new(ProfileDetail)
	err := json.Unmarshal(body, &profiledetail)
	return profiledetail, err
}

type Cover struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		HomeInfo struct {
			HomeID       string `json:"homeId"`
			ObjectID     string `json:"objectId"`
			ObsNamespace string `json:"obsNamespace"`
			ServiceName  string `json:"serviceName"`
			PostCount    int    `json:"postCount"`
		} `json:"homeInfo"`
		Feed struct {
			Post struct {
				UserInfo struct {
					Mid        string `json:"mid"`
					Nickname   string `json:"nickname"`
					PictureURL string `json:"pictureUrl"`
					UserValid  bool   `json:"userValid"`
					WriterMid  string `json:"writerMid"`
				} `json:"userInfo"`
				PostInfo struct {
					AppSn        int    `json:"appSn"`
					HomeID       string `json:"homeId"`
					PostID       string `json:"postId"`
					Status       string `json:"status"`
					LikeCount    int    `json:"likeCount"`
					CommentCount int    `json:"commentCount"`
					Liked        bool   `json:"liked"`
					URL          struct {
						Type      string `json:"type"`
						TargetURL string `json:"targetUrl"`
					} `json:"url"`
					ReadPermission struct {
						Type  string      `json:"type"`
						Gids  interface{} `json:"gids"`
						Count int         `json:"count"`
					} `json:"readPermission"`
					AllowShare            bool   `json:"allowShare"`
					AllowLikeShare        bool   `json:"allowLikeShare"`
					AllowComment          bool   `json:"allowComment"`
					AllowPreviewComment   bool   `json:"allowPreviewComment"`
					AllowPhotoComment     bool   `json:"allowPhotoComment"`
					AllowLike             bool   `json:"allowLike"`
					AllowRecall           bool   `json:"allowRecall"`
					AllowFriendRequest    bool   `json:"allowFriendRequest"`
					AllowCommentLike      bool   `json:"allowCommentLike"`
					AllowLikeProfiles     bool   `json:"allowLikeProfiles"`
					EnableCommentApproval bool   `json:"enableCommentApproval"`
					HasSharedToPost       bool   `json:"hasSharedToPost"`
					CommentLinkPermission string `json:"commentLinkPermission"`
					LikeLinkPermission    string `json:"likeLinkPermission"`
					AutoPostType          string `json:"autoPostType"`
					AllowEdit             bool   `json:"allowEdit"`
					CreatedTime           int64  `json:"createdTime"`
					UpdatedTime           int64  `json:"updatedTime"`
				} `json:"postInfo"`
				HeadLine struct {
					Body     string `json:"body"`
					BodyMeta []struct {
						Start int `json:"start"`
						End   int `json:"end"`
						User  struct {
							ActorID    string `json:"actorId"`
							UserValid  bool   `json:"userValid"`
							PictureURL string `json:"pictureUrl"`
							Nickname   string `json:"nickname"`
						} `json:"user"`
						URL struct {
							Type      string `json:"type"`
							TargetURL string `json:"targetUrl"`
						} `json:"url"`
						Bold bool `json:"bold"`
					} `json:"bodyMeta"`
				} `json:"headLine"`
				Contents struct {
				} `json:"contents"`
				AdditionalContents struct {
					Title     string `json:"title"`
					TitleMeta []struct {
						Start int `json:"start"`
						End   int `json:"end"`
						User  struct {
							ActorID    string `json:"actorId"`
							UserValid  bool   `json:"userValid"`
							PictureURL string `json:"pictureUrl"`
							Nickname   string `json:"nickname"`
						} `json:"user"`
						URL struct {
							Type      string `json:"type"`
							TargetURL string `json:"targetUrl"`
						} `json:"url"`
						Bold bool `json:"bold"`
					} `json:"titleMeta"`
					Main         string `json:"main"`
					Obsthumbnail struct {
						Width         int    `json:"width"`
						Height        int    `json:"height"`
						PreferCdn     bool   `json:"preferCdn"`
						ForbiddenSave bool   `json:"forbiddenSave"`
						ObjectID      string `json:"objectId"`
						ServiceName   string `json:"serviceName"`
						ObsNamespace  string `json:"obsNamespace"`
					} `json:"obsthumbnail"`
					Thumbnail struct {
						URL           interface{} `json:"url"`
						Width         interface{} `json:"width"`
						Height        interface{} `json:"height"`
						RequiredTid   bool        `json:"requiredTid"`
						ForbiddenSave bool        `json:"forbiddenSave"`
					} `json:"thumbnail"`
					URL struct {
						Type      string `json:"type"`
						TargetURL string `json:"targetUrl"`
						MarketURL string `json:"marketUrl"`
					} `json:"url"`
				} `json:"additionalContents"`
				StatisticInfo struct {
					ContentType int `json:"contentType"`
				} `json:"statisticInfo"`
			} `json:"post"`
		} `json:"feed"`
	} `json:"result"`
}
func UpdateCover(objId string) (*Cover, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v39/home/updateCover.json?coverImageId="+objId)
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var cover = new(Cover)
	err := json.Unmarshal(body, &cover)
	return cover, err
}

func GetProfileCoverId(mid string) (string, error) {
	profdet, err := GetProfileDetail(mid)
	return profdet.Result.ObjectID, err
}

func GetProfileCoverURL(mid string) (string, error) {
	profdet, err := GetProfileDetail(mid)
	oid := profdet.Result.ObjectID
	return config.LINE_OBS_DOMAIN+"/myhome/c/download.nhn?userid="+mid+"&oid="+oid, err
}

/* POST */

//will add CreatePost, CreateComment, DeleteComment, LikePost, and UnlikePost later, im lazy...

type PostToTalk struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		ReceiveMid string `json:"receiveMid"`
	} `json:"result"`
}
func SendPostToTalk(to string, postId string) (*PostToTalk, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v39/post/sendPostToTalk.json?receiveMid="+to+"&postId="+postId)
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var posttotalk = new(PostToTalk)
	err := json.Unmarshal(body, &posttotalk)
	return posttotalk, err
}

/* GROUP POST */

//will add CreateGroupAlbum, CreateGroupPost, and DeleteGroupAlbum later, im really lazy...

type GroupPost struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		HomeInfo struct {
			HomeID       string `json:"homeId"`
			ObjectID     string `json:"objectId"`
			ObsNamespace string `json:"obsNamespace"`
			ServiceName  string `json:"serviceName"`
			PostCount    int    `json:"postCount"`
			GroupHome    struct {
				GroupID    string `json:"groupId"`
				Name       string `json:"name"`
				PictureURL string `json:"pictureUrl"`
				GroupType  string `json:"groupType"`
			} `json:"groupHome"`
		} `json:"homeInfo"`
		Feeds []struct {
			FeedInfo struct {
				Type   string      `json:"type"`
				ID     string      `json:"id"`
				Status string      `json:"status"`
				Score  interface{} `json:"score"`
			} `json:"feedInfo"`
			Post struct {
				UserInfo struct {
					Mid        string `json:"mid"`
					Nickname   string `json:"nickname"`
					PictureURL string `json:"pictureUrl"`
					UserValid  bool   `json:"userValid"`
					WriterMid  string `json:"writerMid"`
				} `json:"userInfo"`
				PostInfo struct {
					AppSn        int    `json:"appSn"`
					HomeID       string `json:"homeId"`
					PostID       string `json:"postId"`
					Status       string `json:"status"`
					LikeCount    int    `json:"likeCount"`
					CommentCount int    `json:"commentCount"`
					Liked        bool   `json:"liked"`
					URL          struct {
						Type      string `json:"type"`
						TargetURL string `json:"targetUrl"`
					} `json:"url"`
					ReadPermission struct {
						Type  string        `json:"type"`
						Gids  []interface{} `json:"gids"`
						Count int           `json:"count"`
					} `json:"readPermission"`
					AllowShare            bool   `json:"allowShare"`
					AllowLikeShare        bool   `json:"allowLikeShare"`
					AllowComment          bool   `json:"allowComment"`
					AllowPreviewComment   bool   `json:"allowPreviewComment"`
					AllowPhotoComment     bool   `json:"allowPhotoComment"`
					AllowLike             bool   `json:"allowLike"`
					AllowRecall           bool   `json:"allowRecall"`
					AllowFriendRequest    bool   `json:"allowFriendRequest"`
					AllowCommentLike      bool   `json:"allowCommentLike"`
					AllowLikeProfiles     bool   `json:"allowLikeProfiles"`
					EnableCommentApproval bool   `json:"enableCommentApproval"`
					HasSharedToPost       bool   `json:"hasSharedToPost"`
					CommentLinkPermission string `json:"commentLinkPermission"`
					LikeLinkPermission    string `json:"likeLinkPermission"`
					GroupHome             struct {
						GroupID    string `json:"groupId"`
						Name       string `json:"name"`
						PictureURL string `json:"pictureUrl"`
						GroupType  string `json:"groupType"`
					} `json:"groupHome"`
					EditableContents []string `json:"editableContents"`
					AllowEdit        bool     `json:"allowEdit"`
					CreatedTime      int64    `json:"createdTime"`
					UpdatedTime      int64    `json:"updatedTime"`
				} `json:"postInfo"`
				Contents struct {
					Text          string `json:"text"`
					ContentsStyle struct {
						TextStyle struct {
							TextSizeMode    string `json:"textSizeMode"`
							BackgroundColor string `json:"backgroundColor"`
							TextAnimation   string `json:"textAnimation"`
						} `json:"textStyle"`
						StickerStyle struct {
						} `json:"stickerStyle"`
					} `json:"contentsStyle"`
				} `json:"contents"`
				StatisticInfo struct {
					ContentType int `json:"contentType"`
				} `json:"statisticInfo"`
			} `json:"post"`
		} `json:"feeds"`
	} `json:"result"`
}
func GetGroupPost(to string, commentLimit string, likeLimit string) (*GroupPost, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_API+"/v39/post/list.json?homeId="+to+"&commentLimit="+commentLimit+"&likeLimit="+likeLimit+"&sourceType=TALKROOM")
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var grouppost = new(GroupPost)
	err := json.Unmarshal(body, &grouppost)
	return grouppost, err
}

/* Group Album */

//by this point maybe you already know what im going to say...
//will add ChangeGroupAlbumName, AddImageToAlbum, and GetImageGroupAlbum later

type GroupAlbum struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		AlbumCountLimit string `json:"albumCountLimit"`
		Items           []struct {
			Created    string `json:"created"`
			ID         string `json:"id"`
			LastPosted string `json:"lastPosted"`
			ModifiedID string `json:"modifiedId"`
			NewFlag    bool   `json:"newFlag"`
			Owner      struct {
				Mid string `json:"mid"`
			} `json:"owner"`
			PhotoCount   string `json:"photoCount"`
			RecentPhotos []struct {
				Oid string `json:"oid"`
			} `json:"recentPhotos"`
			Status string `json:"status"`
			Title  string `json:"title"`
		} `json:"items"`
		PhotoCountLimit string `json:"photoCountLimit"`
	} `json:"result"`
}
func GetGroupAlbum(to string) (*GroupAlbum, error) {
	u, _ := url.Parse(config.LINE_TIMELINE_MH+"/album/v3/albums.json?homeId="+to+"&type=g&sourceType=TALKROOM")
	request.URL = u
	res, _ := client.Do(request)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var groupalbum = new(GroupAlbum)
	err := json.Unmarshal(body, &groupalbum)
	return groupalbum, err
}