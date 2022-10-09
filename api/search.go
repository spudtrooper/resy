package api

import (
	"log"
	"strings"
	"time"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

type searchInfoPayload struct {
	Search struct {
		NbHits int `json:"nbHits"`
		Hits   []struct {
			// InventoryEvent  []interface{} `json:"inventory_event"`
			// HighlightResult struct {
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Name struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Cuisine []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// } `json:"_highlightResult,omitempty"`
			Keywords []string `json:"keywords"`
			// IsGns           int      `json:"is_gns"`
			Cuisine      []string `json:"cuisine"`
			Neighborhood string   `json:"neighborhood"`
			// InventoryTypeID int      `json:"inventory_type_id"`
			ID struct {
				Resy int `json:"resy"`
			} `json:"id"`
			Geoloc struct {
				Lng float64 `json:"lng"`
				Lat float64 `json:"lat"`
			} `json:"_geoloc"`
			Favorite       bool     `json:"favorite"`
			InventoryAny   []string `json:"inventory_any"`
			Images         []string `json:"images"`
			CurrencySymbol string   `json:"currency_symbol"`
			URLSlug        string   `json:"url_slug"`
			// Availability   struct {
			// 	Notifies     []interface{} `json:"notifies"`
			// 	ServiceTypes struct {
			// 		Num2 struct {
			// 		} `json:"2"`
			// 	} `json:"service_types"`
			// 	Status interface{}   `json:"status"`
			// 	Events []interface{} `json:"events"`
			// 	Slots  []struct {
			// 		DisplayConfig struct {
			// 			Color struct {
			// 				Font       interface{} `json:"font"`
			// 				Background interface{} `json:"background"`
			// 			} `json:"color"`
			// 		} `json:"display_config"`
			// 		IsGlobalDiningAccess bool `json:"is_global_dining_access"`
			// 		Config               struct {
			// 			Type  string `json:"type"`
			// 			ID    int    `json:"id"`
			// 			Token string `json:"token"`
			// 		} `json:"config"`
			// 		ReservationConfig struct {
			// 			Badge interface{} `json:"badge"`
			// 		} `json:"reservation_config"`
			// 		Shift struct {
			// 			Service struct {
			// 				Type struct {
			// 					ID int `json:"id"`
			// 				} `json:"type"`
			// 			} `json:"service"`
			// 			ID  int    `json:"id"`
			// 			Day string `json:"day"`
			// 		} `json:"shift"`
			// 		Exclusive struct {
			// 			ID int `json:"id"`
			// 		} `json:"exclusive"`
			// 		Date struct {
			// 			Start string `json:"start"`
			// 			End   string `json:"end"`
			// 		} `json:"date"`
			// 		Template struct {
			// 			ID int `json:"id"`
			// 		} `json:"template"`
			// 	} `json:"slots"`
			// } `json:"availability"`
			// Source struct {
			// 	TermsOfService interface{} `json:"terms_of_service"`
			// 	Name           interface{} `json:"name"`
			// 	PrivacyPolicy  interface{} `json:"privacy_policy"`
			// 	Logo           interface{} `json:"logo"`
			// } `json:"source"`
			// Contact struct {
			// 	PhoneNumber string `json:"phone_number"`
			// } `json:"contact"`
			// IsGlobalDiningAccess bool `json:"is_global_dining_access"`
			Location struct {
				Code string `json:"code"`
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"location"`
			EnableForApp int `json:"enable_for_app"`
			Rating       struct {
				Average float64 `json:"average"`
				Count   int     `json:"count"`
			} `json:"rating"`
			// ObjectID string `json:"objectID"`
			// Locality                     string      `json:"locality"`
			Name string `json:"name"`
			// CuisineDetail                interface{} `json:"cuisine_detail"`
			// RequiresReservationTransfers int         `json:"requires_reservation_transfers"`
			// IsRga                        int         `json:"is_rga"`
			// IsGdc                        int         `json:"is_gdc"`
			// Content                      []struct {
			// 	Body   string `json:"body"`
			// 	Name   string `json:"name"`
			// 	Locale struct {
			// 		Language string `json:"language"`
			// 	} `json:"locale"`
			// 	Display struct {
			// 		Type string `json:"type"`
			// 	} `json:"display"`
			// 	Attribution interface{} `json:"attribution"`
			// 	Icon        struct {
			// 		URL string `json:"url"`
			// 	} `json:"icon"`
			// 	Title interface{} `json:"title"`
			// } `json:"content"`
			Region string `json:"region"`
			// Collections []struct {
			// 	Image     string `json:"image"`
			// 	ID        string `json:"id"`
			// 	FileName  string `json:"file_name"`
			// 	Name      string `json:"name"`
			// 	TypeID    int    `json:"type_id"`
			// 	ShortName string `json:"short_name"`
			// } `json:"collections"`
			// PriceRangeID   int           `json:"price_range_id"`
			// ResySelect     int           `json:"resy_select"`
			// Collection     []string      `json:"collection"`
			// MenuHighlights []interface{} `json:"menu_highlights"`
			// Reopen         struct {
			// 	Date string `json:"date"`
			// } `json:"reopen"`
			Country string `json:"country"`
			// GdaConciergeBooking bool   `json:"gda_concierge_booking"`
			// HighlightResult0    struct {
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult1 struct {
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult2 struct {
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult3 struct {
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult4 struct {
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult5 struct {
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult6 struct {
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	CuisineDetail []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult7 struct {
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult8 struct {
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// } `json:"_highlightResult,omitempty"`
			// HighlightResult9 struct {
			// 	Neighborhood struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"neighborhood"`
			// 	Keywords []struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"keywords"`
			// 	CuisineDetail []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine_detail"`
			// 	Name struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"name"`
			// 	Locality struct {
			// 		Value        string        `json:"value"`
			// 		MatchLevel   string        `json:"matchLevel"`
			// 		MatchedWords []interface{} `json:"matchedWords"`
			// 	} `json:"locality"`
			// 	Location struct {
			// 		Name struct {
			// 			Value        string        `json:"value"`
			// 			MatchLevel   string        `json:"matchLevel"`
			// 			MatchedWords []interface{} `json:"matchedWords"`
			// 		} `json:"name"`
			// 	} `json:"location"`
			// 	Cuisine []struct {
			// 		Value            string   `json:"value"`
			// 		MatchLevel       string   `json:"matchLevel"`
			// 		FullyHighlighted bool     `json:"fullyHighlighted"`
			// 		MatchedWords     []string `json:"matchedWords"`
			// 	} `json:"cuisine"`
			// } `json:"_highlightResult,omitempty"`
		} `json:"hits"`
		// Page        int `json:"page"`
		// NbPages     int `json:"nbPages"`
		// HitsPerPage int `json:"hitsPerPage"`
	} `json:"search"`
	// Suggestions []interface{} `json:"suggestions"`
	// Meta        struct {
	// 	Page       int `json:"page"`
	// 	PerPage    int `json:"per_page"`
	// 	Total      int `json:"total"`
	// 	TotalPages int `json:"total_pages"`
	// } `json:"meta"`
}

type SearchInfoHit struct {
	ID             int
	Cuisines       []string
	Neighborhood   string
	Name           string
	Country        string
	Region         string
	RatingAverage  float64
	RatingCount    int
	Keywords       []string
	CurrencySymbol string
	URLSlug        string
	Location       SearchInfoHitLocation
}

type SearchInfoHitLocation struct {
	Code string
	ID   int
	Name string
}

type SearchInfo struct {
	Hits []SearchInfoHit
}

func convertSearchInfoPayload(p searchInfoPayload) *SearchInfo {
	var hits []SearchInfoHit
	for _, h := range p.Search.Hits {
		hits = append(hits, SearchInfoHit{
			ID:             h.ID.Resy,
			Cuisines:       h.Cuisine,
			Neighborhood:   h.Neighborhood,
			Name:           h.Name,
			Country:        h.Country,
			Region:         h.Region,
			RatingAverage:  h.Rating.Average,
			RatingCount:    h.Rating.Count,
			Keywords:       h.Keywords,
			CurrencySymbol: h.CurrencySymbol,
			URLSlug:        h.URLSlug,
			Location: SearchInfoHitLocation{
				Code: h.Location.Code,
				ID:   h.Location.ID,
				Name: h.Location.Name,
			},
		})
	}
	return &SearchInfo{
		Hits: hits,
	}
}

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

//go:generate genopts --function Search --params --required "term string" --extends Base partySize:int:2 page:int:1 perPage:int:20 latitude:float64:40.712941 longitude:float64:-74.006393 radius:int:35420 day:time.Time
func (c *Client) Search(term string, optss ...SearchOption) (*SearchInfo, error) {
	opts := MakeSearchOptions(optss...)

	d := or.Time(opts.Day(), time.Now())
	day := d.Format("2006-01-02")

	uri := request.MakeURL("https://api.resy.com/3/venuesearch/search")
	headers := c.makeHeaders(false, opts.ToBaseOptions()...)

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
		Page:         opts.Page(),
		PerPage:      opts.PerPage(),
		SlotFilter: SlotFilter{
			Day:       day,
			PartySize: opts.PartySize(),
		},
		Types: []string{"venue"},
		Geo: Geo{
			Latitude:  opts.Latitude(),
			Longitude: opts.Longitude(),
			Radius:    opts.Radius(),
		},
		Query: term,
	}
	body := string(request.MustJSONMarshal(bodyObject))

	if opts.DebugBody() {
		log.Printf("body: %s", body)
	}

	var payload searchInfoPayload
	if _, err := request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %s", payload)
	}

	return convertSearchInfoPayload(payload), nil
}
