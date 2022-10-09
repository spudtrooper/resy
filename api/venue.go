package api

import (
	"log"

	"github.com/spudtrooper/goutil/request"
)

type venueInfoPayload struct {
	// Announcement interface{} `json:"announcement"`
	Reopen struct {
		Date string `json:"date"`
	} `json:"reopen"`
	CustomAffiliations []interface{} `json:"custom_affiliations"`
	Collections        []struct {
		ID          int    `json:"id"`
		TypeID      int    `json:"type_id"`
		FileName    string `json:"file_name"`
		Image       string `json:"image"`
		Name        string `json:"name"`
		ShortName   string `json:"short_name"`
		Description string `json:"description"`
	} `json:"collections"`
	Config struct {
		AllowBypassPaymentMethod int `json:"allow_bypass_payment_method"`
		AllowMultipleResys       int `json:"allow_multiple_resys"`
		EnableDiscovery          int `json:"enable_discovery"`
		EnableInvite             int `json:"enable_invite"`
		EnableResypay            int `json:"enable_resypay"`
		HospitalityIncluded      int `json:"hospitality_included"`
	} `json:"config"`
	Contact struct {
		PhoneNumber string `json:"phone_number"`
		URL         string `json:"url"`
	} `json:"contact"`
	Content []struct {
		Attribution interface{} `json:"attribution"`
		Body        string      `json:"body"`
		Display     struct {
			Type string `json:"type"`
		} `json:"display"`
		Icon struct {
			URL string `json:"url"`
		} `json:"icon"`
		Locale struct {
			Language string `json:"language"`
		} `json:"locale"`
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"content"`
	ID struct {
		Foursquare string `json:"foursquare"`
		Google     string `json:"google"`
		Resy       int    `json:"resy"`
		FbPixel    string `json:"fb_pixel"`
	} `json:"id"`
	Images []string `json:"images"`
	// MapImages struct {
	// 	Two22X222 struct {
	// 		Num15 struct {
	// 			Ff2500 string `json:"FF2500"`
	// 		} `json:"15"`
	// 	} `json:"222x222"`
	// } `json:"map_images"`
	// ResponsiveImages struct {
	// 	Originals struct {
	// 		Eight392217F7Ee803D1Fec105D948Cea3E3Faf742F2 struct {
	// 			URL string `json:"url"`
	// 		} `json:"8392217f7ee803d1fec105d948cea3e3faf742f2"`
	// 		Da7492Feeef683Cc79814Ba6E8426D9D4059D2C6 struct {
	// 			URL string `json:"url"`
	// 		} `json:"da7492feeef683cc79814ba6e8426d9d4059d2c6"`
	// 		Two70365025Cc16877E9Fd33Aa8133A86C2380B84B struct {
	// 			URL string `json:"url"`
	// 		} `json:"270365025cc16877e9fd33aa8133a86c2380b84b"`
	// 		SixA2C21E1A04Cad6A286B5B1Ade3020992C41B66B struct {
	// 			URL string `json:"url"`
	// 		} `json:"6a2c21e1a04cad6a286b5b1ade3020992c41b66b"`
	// 		F44130519A5C79630E3446599827C9Dd8F8B969C struct {
	// 			URL string `json:"url"`
	// 		} `json:"f44130519a5c79630e3446599827c9dd8f8b969c"`
	// 		Be9Fd71E064F18746Fc2F70B3F01Fe44Eef49Cef struct {
	// 			URL string `json:"url"`
	// 		} `json:"be9fd71e064f18746fc2f70b3f01fe44eef49cef"`
	// 	} `json:"originals"`
	// 	Urls struct {
	// 		Eight392217F7Ee803D1Fec105D948Cea3E3Faf742F2 struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"8392217f7ee803d1fec105d948cea3e3faf742f2"`
	// 		Da7492Feeef683Cc79814Ba6E8426D9D4059D2C6 struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"da7492feeef683cc79814ba6e8426d9d4059d2c6"`
	// 		Two70365025Cc16877E9Fd33Aa8133A86C2380B84B struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"270365025cc16877e9fd33aa8133a86c2380b84b"`
	// 		SixA2C21E1A04Cad6A286B5B1Ade3020992C41B66B struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"6a2c21e1a04cad6a286b5b1ade3020992c41b66b"`
	// 		F44130519A5C79630E3446599827C9Dd8F8B969C struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"f44130519a5c79630e3446599827c9dd8f8b969c"`
	// 		Be9Fd71E064F18746Fc2F70B3F01Fe44Eef49Cef struct {
	// 			One1 struct {
	// 				Num200  string `json:"200"`
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"1:1"`
	// 			Four3 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"4:3"`
	// 			One69 struct {
	// 				Num400  string `json:"400"`
	// 				Num800  string `json:"800"`
	// 				Num1600 string `json:"1600"`
	// 			} `json:"16:9"`
	// 		} `json:"be9fd71e064f18746fc2f70b3f01fe44eef49cef"`
	// 	} `json:"urls"`
	// 	UrlsByResolution struct {
	// 		Eight392217F7Ee803D1Fec105D948Cea3E3Faf742F2 struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"8392217f7ee803d1fec105d948cea3e3faf742f2"`
	// 		Da7492Feeef683Cc79814Ba6E8426D9D4059D2C6 struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"da7492feeef683cc79814ba6e8426d9d4059d2c6"`
	// 		Two70365025Cc16877E9Fd33Aa8133A86C2380B84B struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"270365025cc16877e9fd33aa8133a86c2380b84b"`
	// 		SixA2C21E1A04Cad6A286B5B1Ade3020992C41B66B struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"6a2c21e1a04cad6a286b5b1ade3020992c41b66b"`
	// 		F44130519A5C79630E3446599827C9Dd8F8B969C struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"f44130519a5c79630e3446599827c9dd8f8b969c"`
	// 		Be9Fd71E064F18746Fc2F70B3F01Fe44Eef49Cef struct {
	// 			Num200 struct {
	// 				One1 string `json:"1:1"`
	// 			} `json:"200"`
	// 			Num400 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"400"`
	// 			Num800 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"800"`
	// 			Num1600 struct {
	// 				One1  string `json:"1:1"`
	// 				Four3 string `json:"4:3"`
	// 				One69 string `json:"16:9"`
	// 			} `json:"1600"`
	// 		} `json:"be9fd71e064f18746fc2f70b3f01fe44eef49cef"`
	// 	} `json:"urls_by_resolution"`
	// 	FileNames    []string `json:"file_names"`
	// 	AspectRatios struct {
	// 		One1 struct {
	// 			Num200  string `json:"200"`
	// 			Num400  string `json:"400"`
	// 			Num800  string `json:"800"`
	// 			Num1600 string `json:"1600"`
	// 		} `json:"1:1"`
	// 		Four3 struct {
	// 			Num400  string `json:"400"`
	// 			Num800  string `json:"800"`
	// 			Num1600 string `json:"1600"`
	// 		} `json:"4:3"`
	// 		One69 struct {
	// 			Num400  string `json:"400"`
	// 			Num800  string `json:"800"`
	// 			Num1600 string `json:"1600"`
	// 		} `json:"16:9"`
	// 	} `json:"aspect_ratios"`
	// } `json:"responsive_images"`
	Inventory struct {
		Type struct {
			ID int `json:"id"`
		} `json:"type"`
	} `json:"inventory"`
	Links struct {
		Deep string `json:"deep"`
		Web  string `json:"web"`
	} `json:"links"`
	Locale struct {
		Currency string `json:"currency"`
		TimeZone string `json:"time_zone"`
	} `json:"locale"`
	Location struct {
		Address1       string  `json:"address_1"`
		Address2       string  `json:"address_2"`
		Code           string  `json:"code"`
		Country        string  `json:"country"`
		CountryIso3166 string  `json:"country_iso3166"`
		CrossStreet1   string  `json:"cross_street_1"`
		CrossStreet2   string  `json:"cross_street_2"`
		ID             int     `json:"id"`
		Latitude       float64 `json:"latitude"`
		Locality       string  `json:"locality"`
		Longitude      float64 `json:"longitude"`
		Neighborhood   string  `json:"neighborhood"`
		PostalCode     string  `json:"postal_code"`
		Region         string  `json:"region"`
	} `json:"location"`
	CustomAffiliation struct {
		ID interface{} `json:"id"`
	} `json:"custom_affiliation"`
	VenueGroup struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Venues []int  `json:"venues"`
	} `json:"venue_group"`
	Metadata struct {
		Description string   `json:"description"`
		Keywords    []string `json:"keywords"`
	} `json:"metadata"`
	Name           string `json:"name"`
	PriceRangeID   int    `json:"price_range_id"`
	CurrencySymbol string `json:"currency_symbol"`
	Rater          []struct {
		Name  string  `json:"name"`
		Scale int     `json:"scale"`
		Score float64 `json:"score"`
		Total int     `json:"total"`
	} `json:"rater"`
	ResySelect                   int    `json:"resy_select"`
	IsGdc                        int    `json:"is_gdc"`
	IsGlobalDiningAccess         bool   `json:"is_global_dining_access"`
	IsGlobalDiningAccessOnly     bool   `json:"is_global_dining_access_only"`
	GdaConciergeBooking          bool   `json:"gda_concierge_booking"`
	IsRga                        bool   `json:"is_rga"`
	GdcPerk                      string `json:"gdc_perk"`
	RequiresReservationTransfers int    `json:"requires_reservation_transfers"`
	IsGns                        int    `json:"is_gns"`
	Social                       []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"social"`
	Ticket struct {
		Average    float32 `json:"average"`
		AverageStr string  `json:"average_str"`
	} `json:"ticket"`
	Type                string `json:"type"`
	URLSlug             string `json:"url_slug"`
	Favorite            bool   `json:"favorite"`
	UserLoginPreference struct {
		MobileFirst int `json:"mobile_first"`
	} `json:"user_login_preference"`
	Source struct {
		Name           interface{} `json:"name"`
		Logo           interface{} `json:"logo"`
		TermsOfService interface{} `json:"terms_of_service"`
		PrivacyPolicy  interface{} `json:"privacy_policy"`
	} `json:"source"`
	HideAllergyQuestion        bool   `json:"hide_allergy_question"`
	HideOccasionQuestion       bool   `json:"hide_occasion_question"`
	HideSpecialRequestQuestion bool   `json:"hide_special_request_question"`
	TaxIncluded                bool   `json:"tax_included"`
	TransactionProcessor       string `json:"transaction_processor"`
	SafetyCategories           []struct {
		ID       int    `json:"id"`
		Value    string `json:"value"`
		OrderNum int    `json:"order_num"`
		Items    []struct {
			ID               int    `json:"id"`
			SafetyCategoryID int    `json:"safety_category_id"`
			Value            string `json:"value"`
			OrderNum         int    `json:"order_num"`
			Active           bool   `json:"active"`
		} `json:"items"`
	} `json:"safety_categories"`
	SafetyInfoURL     interface{} `json:"safety_info_url"`
	MinPartySize      int         `json:"min_party_size"`
	MaxPartySize      int         `json:"max_party_size"`
	LargePartyMessage string      `json:"large_party_message"`
	IsActive          int         `json:"is_active"`
}

type VenueInfoContact struct {
	PhoneNumber string
	URL         string
}

type VenueInfoID struct {
	Foursquare string
	Google     string
	Resy       int
	FbPixel    string
}

type VenueInfoContent struct {
	Body           string `json:"body"`
	DisplayType    string
	IconURL        string
	LocaleLanguage string
	Name           string
	Title          string
}

type VenueRater struct {
	Name  string
	Scale int
	Score float64
	Total int
}

type VenueLocation struct {
	Address1       string
	Address2       string
	Code           string
	Country        string
	CountryIso3166 string
	CrossStreet1   string
	CrossStreet2   string
	ID             int
	Latitude       float64
	Locality       string
	Longitude      float64
	Neighborhood   string
	PostalCode     string
	Region         string
}

type VenueInfo struct {
	Contact        VenueInfoContact
	Content        []VenueInfoContent
	Name           string
	ID             VenueInfoID
	Images         []string
	DeepLinkURI    string
	DeepWebURI     string
	InventoryID    int
	CurrencySymbol string
	Location       VenueLocation
	Raters         []VenueRater
}

func convertvenueInfoPayload(p venueInfoPayload) *VenueInfo {
	var content []VenueInfoContent
	for _, c := range p.Content {
		content = append(content, VenueInfoContent{
			Body:           c.Body,
			DisplayType:    c.Display.Type,
			IconURL:        c.Icon.URL,
			LocaleLanguage: c.Locale.Language,
			Name:           c.Name,
			Title:          c.Title,
		})
	}
	var raters []VenueRater
	for _, r := range p.Rater {
		raters = append(raters, VenueRater{
			Name:  r.Name,
			Scale: r.Scale,
			Score: r.Score,
			Total: r.Total,
		})
	}
	return &VenueInfo{
		Contact: VenueInfoContact{
			PhoneNumber: p.Contact.PhoneNumber,
			URL:         p.Contact.URL,
		},
		Content: content,
		ID: VenueInfoID{
			Foursquare: p.ID.Foursquare,
			Google:     p.ID.Google,
			Resy:       p.ID.Resy,
			FbPixel:    p.ID.FbPixel,
		},
		Images:         p.Images,
		DeepLinkURI:    p.Links.Deep,
		DeepWebURI:     p.Links.Web,
		InventoryID:    p.Inventory.Type.ID,
		Name:           p.Name,
		CurrencySymbol: p.CurrencySymbol,
		Location: VenueLocation{
			Address1:       p.Location.Address1,
			Address2:       p.Location.Address2,
			Code:           p.Location.Code,
			Country:        p.Location.Country,
			CountryIso3166: p.Location.CountryIso3166,
			CrossStreet1:   p.Location.CrossStreet1,
			CrossStreet2:   p.Location.CrossStreet2,
			ID:             p.Location.ID,
			Latitude:       p.Location.Latitude,
			Locality:       p.Location.Locality,
			Longitude:      p.Location.Longitude,
			Neighborhood:   p.Location.Neighborhood,
			PostalCode:     p.Location.PostalCode,
			Region:         p.Location.Region,
		},
		Raters: raters,
	}
}

//go:generate genopts --function Venue --params --required "urlSlug string, location string" --extends Base
func (c *Client) Venue(urlSlug string, location string, optss ...VenueOption) (*VenueInfo, error) {
	opts := MakeVenueOptions(optss...)

	uri := request.MakeURL("https://api.resy.com/3/venue",
		request.Param{"url_slug", urlSlug},
		request.Param{"location", location},
	)
	headers := c.makeHeaders(false, opts.ToBaseOptions()...)

	var payload venueInfoPayload
	if _, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %s", payload)
	}

	res := convertvenueInfoPayload(payload)

	return res, nil
}
