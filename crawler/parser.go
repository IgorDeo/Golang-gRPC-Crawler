package crawler

import (
	"encoding/json"

	"github.com/buger/jsonparser"
	"github.com/go-resty/resty/v2"
)

type InstagramParser struct {
}

type InstagramProfileInfo struct {
	Username      string
	Followers     int32
	Following     int32
	FullName      string
	Biography     string
	IsVerified    bool
	ExternalUrl   string
	ProfilePicUrl string
}

func parseJsonBytes(jsonBytes []byte) (map[string]any, error) {
	var unstructuredJson map[string]any
	json.Unmarshal(jsonBytes, &unstructuredJson)
	return unstructuredJson, nil
}

func (s *InstagramParser) ParseProfileInfo(response resty.Response) (InstagramProfileInfo, error) {

	jsonBytes := response.Body()

	username, error := jsonparser.GetString(jsonBytes, "data", "user", "username")
	followers, error := jsonparser.GetFloat(jsonBytes, "data", "user", "edge_followed_by", "count")
	following, error := jsonparser.GetFloat(jsonBytes, "data", "user", "edge_follow", "count")
	fullName, error := jsonparser.GetString(jsonBytes, "data", "user", "full_name")
	biography, error := jsonparser.GetString(jsonBytes, "data", "user", "biography")
	isVerified, error := jsonparser.GetBoolean(jsonBytes, "data", "user", "is_verified")
	externalUrl, error := jsonparser.GetString(jsonBytes, "data", "user", "external_url")
	profilePicUrl, error := jsonparser.GetString(jsonBytes, "data", "user", "profile_pic_url")

	profileInfo := InstagramProfileInfo{
		Username:      username,
		Followers:     int32(followers),
		Following:     int32(following),
		FullName:      fullName,
		Biography:     biography,
		IsVerified:    isVerified,
		ExternalUrl:   externalUrl,
		ProfilePicUrl: profilePicUrl,
	}
	if error != nil {
		panic(error)
	}

	return profileInfo, nil
}
