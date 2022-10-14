package api

import "github.com/spudtrooper/goutil/or"

func (c *Client) makeHeaders(auth bool, optss ...BaseOption) map[string]string {
	opts := MakeBaseOptions(optss...)
	headers := map[string]string{
		"authority":          `api.resy.com`,
		"accept":             `application/json, text/plain, */*`,
		"accept-language":    `en-US,en;q=0.9`,
		"authorization":      `ResyAPI api_key="VbWk7s3L4KiK5fzlO7JD3Q5EYolJI7n5"`,
		"cache-control":      `no-cache`,
		"content-type":       `application/json`,
		"dnt":                `1`,
		"origin":             `https://resy.com`,
		"pragma":             `no-cache`,
		"referer":            `https://resy.com/`,
		"sec-ch-ua":          `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-site`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
		"x-origin":           `https://resy.com`,
	}
	if auth {
		token := or.String(opts.Token(), c.token)
		headers["x-resy-auth-token"] = token
		headers["x-resy-universal-auth"] = token
	}
	return headers
}
