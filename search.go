package main

import (
	"flag"
	"log"
	"strings"

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
	uri := request.MakeURL("https://api.resy.com/3/venuesearch/search")
	cookie := [][2]string{}
	headers := map[string]string{
		"authority":             `api.resy.com`,
		"accept":                `application/json, text/plain, */*`,
		"accept-language":       `en-US,en;q=0.9`,
		"authorization":         `ResyAPI api_key="VbWk7s3L4KiK5fzlO7JD3Q5EYolJI7n5"`,
		"cache-control":         `no-cache`,
		"content-type":          `application/json`,
		"dnt":                   `1`,
		"origin":                `https://resy.com`,
		"pragma":                `no-cache`,
		"referer":               `https://resy.com/`,
		"sec-ch-ua":             `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":      `?0`,
		"sec-ch-ua-platform":    `"macOS"`,
		"sec-fetch-dest":        `empty`,
		"sec-fetch-mode":        `cors`,
		"sec-fetch-site":        `same-site`,
		"user-agent":            `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
		"x-origin":              `https://resy.com`,
		"x-resy-auth-token":     `eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NjkxNjQ4NzksInVpZCI6OTQyNDcyLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6NTAwMjEzfX0.AdAY4uFSFIo0sxNJ60bc0-n4zoOccXJI6j2TvmbZ6WtdBNQ6JjhJq7c1CnTEnMysyAEWDLpY-gud1U5Lbs3l-wA9AOUVGSWX9-0ujHsIkdYYKdaYZzqNqyDSTdvPnyTOMzcAENLXM6Jcqf8u73D57ewVgTpaQVD4UPQBoX_y340Ti7GX`,
		"x-resy-universal-auth": `eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE2NjkxNjQ4NzksInVpZCI6OTQyNDcyLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6NTAwMjEzfX0.AdAY4uFSFIo0sxNJ60bc0-n4zoOccXJI6j2TvmbZ6WtdBNQ6JjhJq7c1CnTEnMysyAEWDLpY-gud1U5Lbs3l-wA9AOUVGSWX9-0ujHsIkdYYKdaYZzqNqyDSTdvPnyTOMzcAENLXM6Jcqf8u73D57ewVgTpaQVD4UPQBoX_y340Ti7GX`,

		"cookie": request.CreateCookie(cookie),
	}
	body := `{"availability":true,"page":1,"per_page":20,"slot_filter":{"day":"2022-10-09","party_size":2},"types":["venue"],"geo":{"latitude":40.712941,"longitude":-74.006393,"radius":35420},"query":"mexican"}`

	{
		type Geo struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Radius    int     `json:"radius"`
		}
		type SlotFilter struct {
			Day       string `json:"day"`
			PartySize int    `json:"party_size"`
		}
		type Body struct {
			Availability bool       `json:"availability"`
			Page         int        `json:"page"`
			PerPage      int        `json:"per_page"`
			SlotFilter   SlotFilter `json:"slot_filter"`
			Types        []string   `json:"types"`
			Geo          Geo        `json:"geo"`
			Query        string     `json:"query"`
		}

		bodyObject := Body{
			Availability: true,
			Page:         1,
			PerPage:      20,
			SlotFilter:   SlotFilter{Day: "2022-10-09", PartySize: 2},
			Types:        []string{"venue"},
			Geo:          Geo{Latitude: 40.712941, Longitude: -74.006393, Radius: 35420},
			Query:        "mexican"}
		body = string(request.MustJSONMarshal(bodyObject))
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))

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
