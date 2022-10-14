package api

import (
	"log"
	"time"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

type calendarInfoPayload struct {
	Scheduled []struct {
		Date      string `json:"date"`
		Inventory struct {
			Reservation string `json:"reservation"`
			Event       string `json:"event"`
			WalkIn      string `json:"walk-in"`
		} `json:"inventory"`
	} `json:"scheduled"`
	LastCalendarDay string `json:"last_calendar_day"`
}

type CalendarInfoInventory struct {
	Reservation string
	Event       string
	WalkIn      string
}

type CalendarInfoScheduled struct {
	Date      string
	Inventory CalendarInfoInventory
}

type CalendarInfo struct {
	Scheduled       []CalendarInfoScheduled
	LastCalendarDay string
}

func convertCalendarInfoPayload(c calendarInfoPayload) *CalendarInfo {
	var scheduled []CalendarInfoScheduled
	for _, s := range c.Scheduled {
		scheduled = append(scheduled, CalendarInfoScheduled{
			Date: s.Date,
			Inventory: CalendarInfoInventory{
				Reservation: s.Inventory.Reservation,
				Event:       s.Inventory.Event,
				WalkIn:      s.Inventory.WalkIn,
			},
		})
	}
	return &CalendarInfo{
		Scheduled:       scheduled,
		LastCalendarDay: c.LastCalendarDay,
	}
}

//go:generate genopts --function Calendar --extends Base --params --required "venueID int" startDate:time.Time endDate:time.Time numSeats:int:2
func (c *Client) Calendar(venueID int, optss ...CalendarOption) (*CalendarInfo, error) {
	opts := MakeCalendarOptions(optss...)

	startDateTime := or.Time(opts.StartDate(), time.Now())
	startDate := startDateTime.Format("2006-01-02")
	endDate := or.Time(opts.EndDate(), startDateTime.Add(time.Hour*24*365)).Format("2006-01-02")
	numSeats := opts.NumSeats()

	uri := request.MakeURL("https://api.resy.com/4/venue/calendar",
		request.Param{"start_date", startDate},
		request.Param{"end_date", endDate},
		request.Param{"venue_id", venueID},
		request.Param{"num_seats", numSeats},
	)
	headers := c.makeHeaders(false, opts.ToBaseOptions()...)

	var payload calendarInfoPayload

	if _, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %s", payload)
	}

	return convertCalendarInfoPayload(payload), nil
}
