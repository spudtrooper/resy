// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
)

type CalendarOption func(*calendarOptionImpl)

type CalendarOptions interface {
	StartDate() time.Time
	HasStartDate() bool
	EndDate() time.Time
	HasEndDate() bool
	NumSeats() int
	HasNumSeats() bool
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
	DebugPayload() bool
	HasDebugPayload() bool
	ToBaseOptions() []BaseOption
}

func CalendarStartDate(startDate time.Time) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_startDate = true
		opts.startDate = startDate
	}
}
func CalendarStartDateFlag(startDate *time.Time) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if startDate == nil {
			return
		}
		opts.has_startDate = true
		opts.startDate = *startDate
	}
}

func CalendarEndDate(endDate time.Time) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_endDate = true
		opts.endDate = endDate
	}
}
func CalendarEndDateFlag(endDate *time.Time) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if endDate == nil {
			return
		}
		opts.has_endDate = true
		opts.endDate = *endDate
	}
}

func CalendarNumSeats(numSeats int) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_numSeats = true
		opts.numSeats = numSeats
	}
}
func CalendarNumSeatsFlag(numSeats *int) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if numSeats == nil {
			return
		}
		opts.has_numSeats = true
		opts.numSeats = *numSeats
	}
}

func CalendarToken(token string) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func CalendarTokenFlag(token *string) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func CalendarDebugBody(debugBody bool) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func CalendarDebugBodyFlag(debugBody *bool) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

func CalendarDebugPayload(debugPayload bool) CalendarOption {
	return func(opts *calendarOptionImpl) {
		opts.has_debugPayload = true
		opts.debugPayload = debugPayload
	}
}
func CalendarDebugPayloadFlag(debugPayload *bool) CalendarOption {
	return func(opts *calendarOptionImpl) {
		if debugPayload == nil {
			return
		}
		opts.has_debugPayload = true
		opts.debugPayload = *debugPayload
	}
}

type calendarOptionImpl struct {
	startDate        time.Time
	has_startDate    bool
	endDate          time.Time
	has_endDate      bool
	numSeats         int
	has_numSeats     bool
	token            string
	has_token        bool
	debugBody        bool
	has_debugBody    bool
	debugPayload     bool
	has_debugPayload bool
}

func (c *calendarOptionImpl) StartDate() time.Time  { return c.startDate }
func (c *calendarOptionImpl) HasStartDate() bool    { return c.has_startDate }
func (c *calendarOptionImpl) EndDate() time.Time    { return c.endDate }
func (c *calendarOptionImpl) HasEndDate() bool      { return c.has_endDate }
func (c *calendarOptionImpl) NumSeats() int         { return or.Int(c.numSeats, 2) }
func (c *calendarOptionImpl) HasNumSeats() bool     { return c.has_numSeats }
func (c *calendarOptionImpl) Token() string         { return c.token }
func (c *calendarOptionImpl) HasToken() bool        { return c.has_token }
func (c *calendarOptionImpl) DebugBody() bool       { return c.debugBody }
func (c *calendarOptionImpl) HasDebugBody() bool    { return c.has_debugBody }
func (c *calendarOptionImpl) DebugPayload() bool    { return c.debugPayload }
func (c *calendarOptionImpl) HasDebugPayload() bool { return c.has_debugPayload }

type CalendarParams struct {
	VenueID      int       `json:"venue_id" required:"true"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	NumSeats     int       `json:"num_seats" default:"2"`
	Token        string    `json:"token"`
	DebugBody    bool      `json:"debug_body"`
	DebugPayload bool      `json:"debug_payload"`
}

func (o CalendarParams) Options() []CalendarOption {
	return []CalendarOption{
		CalendarStartDate(o.StartDate),
		CalendarEndDate(o.EndDate),
		CalendarNumSeats(o.NumSeats),
		CalendarToken(o.Token),
		CalendarDebugBody(o.DebugBody),
		CalendarDebugPayload(o.DebugPayload),
	}
}

// ToBaseOptions converts CalendarOption to an array of BaseOption
func (o *calendarOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
		BaseDebugPayload(o.DebugPayload()),
	}
}

func makeCalendarOptionImpl(opts ...CalendarOption) *calendarOptionImpl {
	res := &calendarOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCalendarOptions(opts ...CalendarOption) CalendarOptions {
	return makeCalendarOptionImpl(opts...)
}
