package geojson

import (
	"encoding/json"
	"errors"
	"net/url"
)

var (
	// ErrMultipleCRSs happens if both a Name & Link are specified on a
	// CRS struct
	ErrMultipleCRSs = errors.New("cannot specify multiple crss")
	// ErrInvalidCRS is for when an unknown CRS.Type is specified
	ErrInvalidCRS = errors.New("invalid crs specified")
)

type rawCRS struct {
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}

// CRSName is a Named CRS
type CRSName struct {
	// Name is a string identifying the coordinate system using OGC CRS URNs
	Name string `json:"name"`
}

// CRSLink is a link to a CRS
type CRSLink struct {
	// Href is a valid URI to a CRS on the Web
	Href string `json:"href"`
	// Type is the optional hint as to the format to be used for the CRS
	Type string `json:"type,omitempty"`
}

// CRS is the main struct that is used to generate a valid GeoJSON CRS object
//
// Only one of Name or Link may be specified
type CRS struct {
	rawCRS
	// Name is the CRSName of the CRS
	Name *CRSName
	// Link is the CRSLink of the CRS
	Link *CRSLink
}

func (c *CRS) setProperty() error {
	var d interface{}

	switch c.Type {
	case "name":
		c.Name = new(CRSName)
		d = c.Name
	case "link":
		c.Link = new(CRSLink)
		d = c.Link
	default:
		return ErrInvalidCRS
	}

	return json.Unmarshal(c.Properties, d)
}

// MarshalJSON will take a CRS and properly verify the struct and conditionally
// marshal a CRS name or link
func (c CRS) MarshalJSON() ([]byte, error) {
	type crs struct {
		Type       string      `json:"type"`
		Properties interface{} `json:"properties"`
	}

	var j crs
	i := 0

	if c.Name != nil {
		j = crs{Type: "name", Properties: c.Name}
		i++
	}
	if c.Link != nil {
		if _, err := url.Parse(c.Link.Href); err != nil {
			return nil, err
		}
		j = crs{Type: "link", Properties: c.Link}
		i++
	}

	// No crs property means don't include it in the JSON
	if i == 0 {
		return nil, nil
	}
	// There cannot be more than one crs specified
	if i >= 2 {
		return nil, ErrMultipleCRSs
	}

	return json.Marshal(j)

}

// UnmarshalJSON will take raw CRS JSON and properly fill out the CRS with either
// Name or Link set based on the parsed JSON
func (c *CRS) UnmarshalJSON(b []byte) error {
	var r rawCRS

	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	c.rawCRS = r

	return c.setProperty()
}
