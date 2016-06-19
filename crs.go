package geojson

import (
	"encoding/json"
	"errors"
	"net/url"
)

var (
	ErrMultipleCRSs = errors.New("cannot specify multiple crss")
	ErrInvalidCRS   = errors.New("invalid crs specified")
)

type rawCRS struct {
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}

type CRSName struct {
	Name string `json:"name"`
}

type CRSLink struct {
	Href string `json:"href"`
	Type string `json:"type"`
}

type CRS struct {
	rawCRS
	Name *CRSName
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

func (c CRS) MarshalJSON() ([]byte, error) {
	type crs struct {
		Type       string      `json:"type"`
		Properties interface{} `json:"properties"`
	}

	var j crs
	i := 0

	if c.Name != nil {
		j = crs{"name", c.Name}
		i++
	}
	if c.Link != nil {
		if _, err := url.Parse(c.Link.Href); err != nil {
			return nil, err
		}
		j = crs{"link", c.Link}
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

func (c *CRS) UnmarshalJSON(b []byte) error {
	var r rawCRS

	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}

	c.rawCRS = r

	return c.setProperty()
}
