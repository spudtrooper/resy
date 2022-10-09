package main

import (
	"flag"
	"log"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://api.resy.com/3/venue",
		request.Param{"url_slug", `orale-mexican-kitchen-jersey-city`},
		request.Param{"location", `jsc`},
	)
	cookie := [][2]string{}
	headers := map[string]string{
		"authority":          `api.resy.com`,
		"accept":             `application/json, text/plain, */*`,
		"accept-language":    `en-US,en;q=0.9`,
		"authorization":      `ResyAPI api_key="VbWk7s3L4KiK5fzlO7JD3Q5EYolJI7n5"`,
		"cache-control":      `no-cache`,
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
		// "x-resy-auth-token":     `eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NjkxNjQ4NzksInVpZCI6OTQyNDcyLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6NTAwMjEzfX0.AdAY4uFSFIo0sxNJ60bc0-n4zoOccXJI6j2TvmbZ6WtdBNQ6JjhJq7c1CnTEnMysyAEWDLpY-gud1U5Lbs3l-wA9AOUVGSWX9-0ujHsIkdYYKdaYZzqNqyDSTdvPnyTOMzcAENLXM6Jcqf8u73D57ewVgTpaQVD4UPQBoX_y340Ti7GX`,
		// "x-resy-universal-auth": `eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NjkxNjQ4NzksInVpZCI6OTQyNDcyLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6NTAwMjEzfX0.AdAY4uFSFIo0sxNJ60bc0-n4zoOccXJI6j2TvmbZ6WtdBNQ6JjhJq7c1CnTEnMysyAEWDLpY-gud1U5Lbs3l-wA9AOUVGSWX9-0ujHsIkdYYKdaYZzqNqyDSTdvPnyTOMzcAENLXM6Jcqf8u73D57ewVgTpaQVD4UPQBoX_y340Ti7GX`,

		"cookie": request.CreateCookie(cookie),
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))

	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", json.MustColorMarshal(payload))
	}
	check.Err(err)
}
