// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"

	"github.com/spudtrooper/goutil/or"
)

type SearchOption struct {
	f func(*searchOptionImpl)
	s string
}

func (o SearchOption) String() string { return o.s }

type SearchOptions interface {
	Day() time.Time
	HasDay() bool
	DebugBody() bool
	HasDebugBody() bool
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	Page() int
	HasPage() bool
	PartySize() int
	HasPartySize() bool
	PerPage() int
	HasPerPage() bool
	Radius() int
	HasRadius() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func SearchDay(day time.Time) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_day = true
		opts.day = day
	}, fmt.Sprintf("api.SearchDay(time.Time %+v)}", day)}
}
func SearchDayFlag(day *time.Time) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if day == nil {
			return
		}
		opts.has_day = true
		opts.day = *day
	}, fmt.Sprintf("api.SearchDay(time.Time %+v)}", day)}
}

func SearchDebugBody(debugBody bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}, fmt.Sprintf("api.SearchDebugBody(bool %+v)}", debugBody)}
}
func SearchDebugBodyFlag(debugBody *bool) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}, fmt.Sprintf("api.SearchDebugBody(bool %+v)}", debugBody)}
}

func SearchLatitude(latitude float64) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.SearchLatitude(float64 %+v)}", latitude)}
}
func SearchLatitudeFlag(latitude *float64) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.SearchLatitude(float64 %+v)}", latitude)}
}

func SearchLongitude(longitude float64) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.SearchLongitude(float64 %+v)}", longitude)}
}
func SearchLongitudeFlag(longitude *float64) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.SearchLongitude(float64 %+v)}", longitude)}
}

func SearchPage(page int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_page = true
		opts.page = page
	}, fmt.Sprintf("api.SearchPage(int %+v)}", page)}
}
func SearchPageFlag(page *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if page == nil {
			return
		}
		opts.has_page = true
		opts.page = *page
	}, fmt.Sprintf("api.SearchPage(int %+v)}", page)}
}

func SearchPartySize(partySize int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_partySize = true
		opts.partySize = partySize
	}, fmt.Sprintf("api.SearchPartySize(int %+v)}", partySize)}
}
func SearchPartySizeFlag(partySize *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if partySize == nil {
			return
		}
		opts.has_partySize = true
		opts.partySize = *partySize
	}, fmt.Sprintf("api.SearchPartySize(int %+v)}", partySize)}
}

func SearchPerPage(perPage int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_perPage = true
		opts.perPage = perPage
	}, fmt.Sprintf("api.SearchPerPage(int %+v)}", perPage)}
}
func SearchPerPageFlag(perPage *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if perPage == nil {
			return
		}
		opts.has_perPage = true
		opts.perPage = *perPage
	}, fmt.Sprintf("api.SearchPerPage(int %+v)}", perPage)}
}

func SearchRadius(radius int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_radius = true
		opts.radius = radius
	}, fmt.Sprintf("api.SearchRadius(int %+v)}", radius)}
}
func SearchRadiusFlag(radius *int) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if radius == nil {
			return
		}
		opts.has_radius = true
		opts.radius = *radius
	}, fmt.Sprintf("api.SearchRadius(int %+v)}", radius)}
}

func SearchToken(token string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.SearchToken(string %+v)}", token)}
}
func SearchTokenFlag(token *string) SearchOption {
	return SearchOption{func(opts *searchOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.SearchToken(string %+v)}", token)}
}

type searchOptionImpl struct {
	day           time.Time
	has_day       bool
	debugBody     bool
	has_debugBody bool
	latitude      float64
	has_latitude  bool
	longitude     float64
	has_longitude bool
	page          int
	has_page      bool
	partySize     int
	has_partySize bool
	perPage       int
	has_perPage   bool
	radius        int
	has_radius    bool
	token         string
	has_token     bool
}

func (s *searchOptionImpl) Day() time.Time     { return s.day }
func (s *searchOptionImpl) HasDay() bool       { return s.has_day }
func (s *searchOptionImpl) DebugBody() bool    { return s.debugBody }
func (s *searchOptionImpl) HasDebugBody() bool { return s.has_debugBody }
func (s *searchOptionImpl) Latitude() float64  { return or.Float64(s.latitude, 40.712941) }
func (s *searchOptionImpl) HasLatitude() bool  { return s.has_latitude }
func (s *searchOptionImpl) Longitude() float64 { return or.Float64(s.longitude, -74.006393) }
func (s *searchOptionImpl) HasLongitude() bool { return s.has_longitude }
func (s *searchOptionImpl) Page() int          { return or.Int(s.page, 1) }
func (s *searchOptionImpl) HasPage() bool      { return s.has_page }
func (s *searchOptionImpl) PartySize() int     { return or.Int(s.partySize, 2) }
func (s *searchOptionImpl) HasPartySize() bool { return s.has_partySize }
func (s *searchOptionImpl) PerPage() int       { return or.Int(s.perPage, 20) }
func (s *searchOptionImpl) HasPerPage() bool   { return s.has_perPage }
func (s *searchOptionImpl) Radius() int        { return or.Int(s.radius, 35420) }
func (s *searchOptionImpl) HasRadius() bool    { return s.has_radius }
func (s *searchOptionImpl) Token() string      { return s.token }
func (s *searchOptionImpl) HasToken() bool     { return s.has_token }

type SearchParams struct {
	Day       time.Time `json:"day"`
	DebugBody bool      `json:"debug_body"`
	Latitude  float64   `json:"latitude" default:"40.712941"`
	Longitude float64   `json:"longitude" default:"-74.006393"`
	Page      int       `json:"page" default:"1"`
	PartySize int       `json:"party_size" default:"2"`
	PerPage   int       `json:"per_page" default:"20"`
	Radius    int       `json:"radius" default:"35420"`
	Term      string    `json:"term" required:"true"`
	Token     string    `json:"token"`
}

func (o SearchParams) Options() []SearchOption {
	return []SearchOption{
		SearchDay(o.Day),
		SearchDebugBody(o.DebugBody),
		SearchLatitude(o.Latitude),
		SearchLongitude(o.Longitude),
		SearchPage(o.Page),
		SearchPartySize(o.PartySize),
		SearchPerPage(o.PerPage),
		SearchRadius(o.Radius),
		SearchToken(o.Token),
	}
}

// ToBaseOptions converts SearchOption to an array of BaseOption
func (o *searchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
	}
}

func makeSearchOptionImpl(opts ...SearchOption) *searchOptionImpl {
	res := &searchOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeSearchOptions(opts ...SearchOption) SearchOptions {
	return makeSearchOptionImpl(opts...)
}
