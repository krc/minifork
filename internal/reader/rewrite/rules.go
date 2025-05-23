// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rewrite // import "miniflux.app/v2/internal/reader/rewrite"

import (
	"net/url"
	"strings"
)

// List of predefined rewrite rules (alphabetically sorted)
// Available rules: "add_image_title", "add_youtube_video"
// domain => rule name
var predefinedRules = map[string]string{
	"abstrusegoose.com":      "add_image_title",
	"amazingsuperpowers.com": "add_image_title",
	"blog.cloudflare.com":    `add_image_title,remove("figure.kg-image-card figure.kg-image + img")`,
	"cowbirdsinlove.com":     "add_image_title",
	"drawingboardcomic.com":  "add_image_title",
	"exocomics.com":          "add_image_title",
	"framatube.org":          "nl2br,convert_text_link",
	"happletea.com":          "add_image_title",
	"ilpost.it":              `remove(".art_tag, #audioPlayerArticle, .author-container, .caption, .ilpostShare, .lastRecents, #mc_embed_signup, .outbrain_inread, p:has(.leggi-anche), .youtube-overlay")`,
	"imogenquest.net":        "add_image_title",
	"lukesurl.com":           "add_image_title",
	"medium.com":             "fix_medium_images",
	"mercworks.net":          "add_image_title",
	"monkeyuser.com":         "add_image_title",
	"mrlovenstein.com":       "add_image_title",
	"nedroid.com":            "add_image_title",
	"oglaf.com":              `replace("media.oglaf.com/story/tt(.+).gif"|"media.oglaf.com/comic/$1.jpg"),add_image_title`,
	"optipess.com":           "add_image_title",
	"peebleslab.com":         "add_image_title",
	"quantamagazine.org":     `add_youtube_video_from_id, remove("h6:not(.byline,.post__title__kicker), #comments, .next-post__content, .footer__section, figure .outer--content, script")`,
	"sentfromthemoon.com":    "add_image_title",
	"thedoghousediaries.com": "add_image_title",
	"theverge.com":           `add_dynamic_image, remove("div.duet--recirculation--related-list, .hidden")`,
	"treelobsters.com":       "add_image_title",
	"www.qwantz.com":         "add_image_title,add_mailto_subject",
	"xkcd.com":               "add_image_title",
	"youtube.com":            "add_youtube_video",
}

// GetRefererForURL returns the referer for the given URL if it exists, otherwise an empty string.
func GetRefererForURL(u string) string {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return ""
	}

	switch parsedUrl.Hostname() {
	case "appinn.com":
		return "https://appinn.com"
	case "bjp.org.cn":
		return "https://bjp.org.cn"
	case "cdnfile.sspai.com":
		return "https://sspai.com"
	case "f.video.weibocdn.com":
		return "https://weibo.com"
	case "i.pximg.net":
		return "https://www.pixiv.net"
	case "img.hellogithub.com":
		return "https://hellogithub.com"
	case "moyu.im":
		return "https://i.jandan.net"
	}

	switch {
	case strings.HasSuffix(parsedUrl.Hostname(), ".cdninstagram.com"):
		return "https://www.instagram.com"
	case strings.HasSuffix(parsedUrl.Hostname(), ".moyu.im"):
		return "https://i.jandan.net"
	case strings.HasSuffix(parsedUrl.Hostname(), ".sinaimg.cn"):
		return "https://weibo.com"
	}

	return ""
}
