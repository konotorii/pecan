package pecan

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

type ConfigStruct struct {
	token    string
	repo     string
	url      string
	owner    string
	interval int64
}

type CacheStruct struct {
	lastUpdate       time.Time
	isOutdated       bool
	refreshCache     bool
	loadCache        bool
	cacheReleaseList string
	latest           string
}

var cache *CacheStruct

var config *ConfigStruct

func (t *ConfigStruct) CacheInit(token string, repo string, url string, owner string, interval int64) {
	t.url = url
	t.token = token
	t.repo = repo
	t.owner = owner

	config.owner = owner
	config.repo = repo
	config.url = url
	config.token = token
	config.interval = interval

	if stringLength(owner) == 0 || stringLength(repo) == 0 {
		_ = errors.New("neither OWNER, nor REPO are defined")
	}

	if stringLength(token) > 0 && stringLength(url) == 0 {
		_ = errors.New("URL is not defined, which is mandatory for private repo mode")
	}
}

func cacheReleaseList(url string) string {
	headers := []CustomHeaders{{label: "Accept", value: "application/vnd.github.preview"}}

	if stringLength(config.token) > 0 {
		authHeader := CustomHeaders{label: "Authorization", value: fmt.Sprintf("Token %s", config.token)}
		headers = append(headers, authHeader)
	}

	res := getRequest(url, headers)

	if res.StatusCode != 200 {
		_ = errors.New(fmt.Sprintf("Tried to cache RELEASES, but failed fetching %u, status %s", url, res.Status))
	}

	buf := new(strings.Builder)

	_, err := io.Copy(buf, res.Body)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	content := buf.String()

	r := regexp.MustCompile(`/[^ ]*\.nupkg/gim`)

	matches := r.FindAllStringSubmatch(content, -1)

	for _, v := range matches {
		const nuPKG = strings.Replace(url, "RELEASES", v[1], 1)
		content = strings.Replace(content, v[1], nuPKG, -1)
	}

	return content
}

func refreshCache() bool {
	repo := fmt.Sprintf("%o/%r", config.owner, config.repo)
	url := fmt.Sprintf("https://api.github.com/repos/%r/releases?per_page_100", repo)
	headers := []CustomHeaders{{label: "Accept", value: "application/vnd.github.preview"}}

	if stringLength(config.token) > 0 {
		authHeader := CustomHeaders{label: "Authorization", value: fmt.Sprintf("Token %s", config.token)}
		headers = append(headers, authHeader)
	}

	res := getRequest(url, headers)

	if res.StatusCode != 200 {
		_ = errors.New(fmt.Sprintf("Github API responded with %s for url %u", res.Status, url))
	}

	data := res.Body

	if !data || len(data) == 0 {

	}

	cache.lastUpdate = time.Now()
}

func isOutdated() bool {
	lastUpdate := cache.lastUpdate

	if time.Now().UnixMilli()-lastUpdate.UnixMilli() > config.interval {
		return true
	}

	return false
}

func loadCache() {
	if cache.lastUpdate == 0 || isOutdated() {
		refreshCache()
	}

}
