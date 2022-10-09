// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
)

type SearchOption func(*searchOptionImpl)

type SearchOptions interface {
	Token() string
	HasToken() bool
	PartySize() int
	HasPartySize() bool
	Page() int
	HasPage() bool
	PerPage() int
	HasPerPage() bool
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	Radius() int
	HasRadius() bool
	Day() time.Time
	HasDay() bool
}

func SearchToken(token string) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func SearchTokenFlag(token *string) SearchOption {
	return func(opts *searchOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func SearchPartySize(partySize int) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_partySize = true
		opts.partySize = partySize
	}
}
func SearchPartySizeFlag(partySize *int) SearchOption {
	return func(opts *searchOptionImpl) {
		if partySize == nil {
			return
		}
		opts.has_partySize = true
		opts.partySize = *partySize
	}
}

func SearchPage(page int) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_page = true
		opts.page = page
	}
}
func SearchPageFlag(page *int) SearchOption {
	return func(opts *searchOptionImpl) {
		if page == nil {
			return
		}
		opts.has_page = true
		opts.page = *page
	}
}

func SearchPerPage(perPage int) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_perPage = true
		opts.perPage = perPage
	}
}
func SearchPerPageFlag(perPage *int) SearchOption {
	return func(opts *searchOptionImpl) {
		if perPage == nil {
			return
		}
		opts.has_perPage = true
		opts.perPage = *perPage
	}
}

func SearchLatitude(latitude float64) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}
}
func SearchLatitudeFlag(latitude *float64) SearchOption {
	return func(opts *searchOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}
}

func SearchLongitude(longitude float64) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}
}
func SearchLongitudeFlag(longitude *float64) SearchOption {
	return func(opts *searchOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}
}

func SearchRadius(radius int) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_radius = true
		opts.radius = radius
	}
}
func SearchRadiusFlag(radius *int) SearchOption {
	return func(opts *searchOptionImpl) {
		if radius == nil {
			return
		}
		opts.has_radius = true
		opts.radius = *radius
	}
}

func SearchDay(day time.Time) SearchOption {
	return func(opts *searchOptionImpl) {
		opts.has_day = true
		opts.day = day
	}
}
func SearchDayFlag(day *time.Time) SearchOption {
	return func(opts *searchOptionImpl) {
		if day == nil {
			return
		}
		opts.has_day = true
		opts.day = *day
	}
}

type searchOptionImpl struct {
	token         string
	has_token     bool
	partySize     int
	has_partySize bool
	page          int
	has_page      bool
	perPage       int
	has_perPage   bool
	latitude      float64
	has_latitude  bool
	longitude     float64
	has_longitude bool
	radius        int
	has_radius    bool
	day           time.Time
	has_day       bool
}

func (s *searchOptionImpl) Token() string      { return s.token }
func (s *searchOptionImpl) HasToken() bool     { return s.has_token }
func (s *searchOptionImpl) PartySize() int     { return or.Int(s.partySize, 2) }
func (s *searchOptionImpl) HasPartySize() bool { return s.has_partySize }
func (s *searchOptionImpl) Page() int          { return or.Int(s.page, 1) }
func (s *searchOptionImpl) HasPage() bool      { return s.has_page }
func (s *searchOptionImpl) PerPage() int       { return or.Int(s.perPage, 20) }
func (s *searchOptionImpl) HasPerPage() bool   { return s.has_perPage }
func (s *searchOptionImpl) Latitude() float64  { return or.Float64(s.latitude, 40.725562967812365) }
func (s *searchOptionImpl) HasLatitude() bool  { return s.has_latitude }
func (s *searchOptionImpl) Longitude() float64 { return or.Float64(s.longitude, -73.99434669171899) }
func (s *searchOptionImpl) HasLongitude() bool { return s.has_longitude }
func (s *searchOptionImpl) Radius() int        { return or.Int(s.radius, 35420) }
func (s *searchOptionImpl) HasRadius() bool    { return s.has_radius }
func (s *searchOptionImpl) Day() time.Time     { return s.day }
func (s *searchOptionImpl) HasDay() bool       { return s.has_day }

type SearchParams struct {
	Term      string    `json:"term" required:"true"`
	Token     string    `json:"token"`
	PartySize int       `json:"party_size" default:"2"`
	Page      int       `json:"page" default:"1"`
	PerPage   int       `json:"per_page" default:"20"`
	Latitude  float64   `json:"latitude" default:"40.725562967812365"`
	Longitude float64   `json:"longitude" default:"-73.99434669171899"`
	Radius    int       `json:"radius" default:"35420"`
	Day       time.Time `json:"day"`
}

func (o SearchParams) Options() []SearchOption {
	return []SearchOption{
		SearchToken(o.Token),
		SearchPartySize(o.PartySize),
		SearchPage(o.Page),
		SearchPerPage(o.PerPage),
		SearchLatitude(o.Latitude),
		SearchLongitude(o.Longitude),
		SearchRadius(o.Radius),
		SearchDay(o.Day),
	}
}

func makeSearchOptionImpl(opts ...SearchOption) *searchOptionImpl {
	res := &searchOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeSearchOptions(opts ...SearchOption) SearchOptions {
	return makeSearchOptionImpl(opts...)
}
