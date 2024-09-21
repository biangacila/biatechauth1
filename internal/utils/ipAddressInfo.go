package utils

import (
	"net"
	"regexp"
)

func ContainsIPAddress(s string) bool {
	// Regular expression to match potential IP addresses
	ipRegex := `\b(?:\d{1,3}\.){3}\d{1,3}\b`

	// Compile the regex
	re := regexp.MustCompile(ipRegex)

	// Find all matches
	matches := re.FindAllString(s, -1)

	// Check if any of the matches are valid IP addresses
	for _, match := range matches {
		if net.ParseIP(match) != nil {
			return true
		}
	}
	return false
}

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

var m = map[string]string{
	"amazon":          "Amazon",
	"apple":           "Apple",
	"auth0":           "Auth0",
	"azuread":         "Azure AD",
	"battlenet":       "Battle.net",
	"bitbucket":       "Bitbucket",
	"box":             "Box",
	"dailymotion":     "Dailymotion",
	"deezer":          "Deezer",
	"digitalocean":    "Digital Ocean",
	"discord":         "Discord",
	"dropbox":         "Dropbox",
	"eveonline":       "Eve Online",
	"facebook":        "Facebook",
	"fitbit":          "Fitbit",
	"gitea":           "Gitea",
	"github":          "Github",
	"gitlab":          "Gitlab",
	"google":          "Google",
	"gplus":           "Google Plus",
	"heroku":          "Heroku",
	"instagram":       "Instagram",
	"intercom":        "Intercom",
	"kakao":           "Kakao",
	"lastfm":          "Last FM",
	"line":            "LINE",
	"linkedin":        "LinkedIn",
	"mastodon":        "Mastodon",
	"meetup":          "Meetup.com",
	"microsoftonline": "Microsoft Online",
	"naver":           "Naver",
	"nextcloud":       "NextCloud",
	"okta":            "Okta",
	"onedrive":        "Onedrive",
	"openid-connect":  "OpenID Connect",
	"patreon":         "Patreon",
	"paypal":          "Paypal",
	"salesforce":      "Salesforce",
	"seatalk":         "SeaTalk",
	"shopify":         "Shopify",
	"slack":           "Slack",
	"soundcloud":      "SoundCloud",
	"spotify":         "Spotify",
	"steam":           "Steam",
	"strava":          "Strava",
	"stripe":          "Stripe",
	"tiktok":          "TikTok",
	"twitch":          "Twitch",
	"twitter":         "Twitter",
	"twitterv2":       "Twitter",
	"typetalk":        "Typetalk",
	"uber":            "Uber",
	"vk":              "VK",
	"wecom":           "WeCom",
	"wepay":           "Wepay",
	"xero":            "Xero",
	"yahoo":           "Yahoo",
	"yammer":          "Yammer",
	"yandex":          "Yandex",
	"zoom":            "Zoom",
}
