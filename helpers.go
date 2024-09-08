package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func proxyPrivateDownload(token string, url string, res gin.ResponseWriter) {
	headers := []CustomHeaders{{label: "Accept", value: "application/octet-stream"}}

	strings.Replace(url, "https://api.github.com/", fmt.Sprintf("https://%s@api.github.com/", token), 1)

	getRes := getRequest(url, headers)

	res.Header().Set("Location", getRes.Header.Get("Location"))
}

func stringLength(str string) int {
	return strings.Count(str, "")
}
