// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"

	"github.com/spudtrooper/goutil/or"
)

type CalendarOption struct {
	f func(*calendarOptionImpl)
	s string
}

func (o CalendarOption) String() string { return o.s }

type CalendarOptions interface {
	DebugBody() bool
	HasDebugBody() bool
	EndDate() time.Time
	HasEndDate() bool
	NumSeats() int
	HasNumSeats() bool
	StartDate() time.Time
	HasStartDate() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func CalendarDebugBody(debugBody bool) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}, fmt.Sprintf("api.CalendarDebugBody(bool %+v)", debugBody)}
}
func CalendarDebugBodyFlag(debugBody *bool) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}, fmt.Sprintf("api.CalendarDebugBody(bool %+v)", debugBody)}
}

func CalendarEndDate(endDate time.Time) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		opts.has_endDate = true
		opts.endDate = endDate
	}, fmt.Sprintf("api.CalendarEndDate(time.Time %+v)", endDate)}
}
func CalendarEndDateFlag(endDate *time.Time) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		if endDate == nil {
			return
		}
		opts.has_endDate = true
		opts.endDate = *endDate
	}, fmt.Sprintf("api.CalendarEndDate(time.Time %+v)", endDate)}
}

func CalendarNumSeats(numSeats int) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		opts.has_numSeats = true
		opts.numSeats = numSeats
	}, fmt.Sprintf("api.CalendarNumSeats(int %+v)", numSeats)}
}
func CalendarNumSeatsFlag(numSeats *int) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		if numSeats == nil {
			return
		}
		opts.has_numSeats = true
		opts.numSeats = *numSeats
	}, fmt.Sprintf("api.CalendarNumSeats(int %+v)", numSeats)}
}

func CalendarStartDate(startDate time.Time) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		opts.has_startDate = true
		opts.startDate = startDate
	}, fmt.Sprintf("api.CalendarStartDate(time.Time %+v)", startDate)}
}
func CalendarStartDateFlag(startDate *time.Time) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		if startDate == nil {
			return
		}
		opts.has_startDate = true
		opts.startDate = *startDate
	}, fmt.Sprintf("api.CalendarStartDate(time.Time %+v)", startDate)}
}

func CalendarToken(token string) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.CalendarToken(string %+v)", token)}
}
func CalendarTokenFlag(token *string) CalendarOption {
	return CalendarOption{func(opts *calendarOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.CalendarToken(string %+v)", token)}
}

type calendarOptionImpl struct {
	debugBody     bool
	has_debugBody bool
	endDate       time.Time
	has_endDate   bool
	numSeats      int
	has_numSeats  bool
	startDate     time.Time
	has_startDate bool
	token         string
	has_token     bool
}

func (c *calendarOptionImpl) DebugBody() bool      { return c.debugBody }
func (c *calendarOptionImpl) HasDebugBody() bool   { return c.has_debugBody }
func (c *calendarOptionImpl) EndDate() time.Time   { return c.endDate }
func (c *calendarOptionImpl) HasEndDate() bool     { return c.has_endDate }
func (c *calendarOptionImpl) NumSeats() int        { return or.Int(c.numSeats, 2) }
func (c *calendarOptionImpl) HasNumSeats() bool    { return c.has_numSeats }
func (c *calendarOptionImpl) StartDate() time.Time { return c.startDate }
func (c *calendarOptionImpl) HasStartDate() bool   { return c.has_startDate }
func (c *calendarOptionImpl) Token() string        { return c.token }
func (c *calendarOptionImpl) HasToken() bool       { return c.has_token }

type CalendarParams struct {
	DebugBody bool      `json:"debug_body"`
	EndDate   time.Time `json:"end_date"`
	NumSeats  int       `json:"num_seats" default:"2"`
	StartDate time.Time `json:"start_date"`
	Token     string    `json:"token"`
	VenueID   int       `json:"venue_id" required:"true"`
}

func (o CalendarParams) Options() []CalendarOption {
	return []CalendarOption{
		CalendarDebugBody(o.DebugBody),
		CalendarEndDate(o.EndDate),
		CalendarNumSeats(o.NumSeats),
		CalendarStartDate(o.StartDate),
		CalendarToken(o.Token),
	}
}

// ToBaseOptions converts CalendarOption to an array of BaseOption
func (o *calendarOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseDebugBody(o.DebugBody()),
		BaseToken(o.Token()),
	}
}

func makeCalendarOptionImpl(opts ...CalendarOption) *calendarOptionImpl {
	res := &calendarOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeCalendarOptions(opts ...CalendarOption) CalendarOptions {
	return makeCalendarOptionImpl(opts...)
}
