package social

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gzipchrist/dont_at_me/pkg/style"
)

type Platform struct {
	Name                string
	URL                 string
	Match               string
	MatchMeansAvailable bool
}

var Platforms = []Platform{
	{"Instagram", "https://instagram.com/", "<title>Instagram</title>", true},
	{"TikTok", "https://us.tiktok.com/@", "Watch the latest video from .", true},
	{"GitHub", "https://github.com/", strconv.Itoa(http.StatusNotFound), true},
	{"Snapchat", "https://www.snapchat.com/add/", "content=\"Not_Found\"", true},
	{"Twitch", "https://www.twitch.tv/", "https://player.twitch.tv/?channel=", false},
	{"YouTube", "https://youtube.com/@", strconv.Itoa(http.StatusNotFound), true},
	{"Mastodon", "https://mastodon.social/@", "<title>The page you", true},
}

func (p Platform) String() string {
	return p.Name
}

func (p Platform) BaseUrl() string {
	return p.URL
}

func (p Platform) Spacer() int {
	return style.MaxCharWidth - len(p.Name)
}

type Status int

const (
	Unavailable Status = iota - 1
	Unknown
	Available
)

var StatusMessages = map[Status]string{
	Unavailable: "❌",
	Unknown:     "❓",
	Available:   "✅",
}

func (s Status) String() string {
	return StatusMessages[s]
}

func (p Platform) GetAvailability(username string) Status {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := p.BaseUrl() + username

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Unknown
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Unknown
	}

	defer res.Body.Close()

	if p.Match == strconv.Itoa(res.StatusCode) {
		if p.MatchMeansAvailable {
			return Available
		}
		return Unavailable
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Unknown
	}

	if strings.Contains(string(body), p.Match) {
		if p.MatchMeansAvailable {
			return Available
		}
		return Unavailable
	}

	return Unknown
}
