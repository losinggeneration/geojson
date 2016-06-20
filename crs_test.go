package geojson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func equalCRS(c1, c2 *CRS, t *testing.T) {
	if c1 == nil && c2 == nil {
		return
	} else if c1 != nil && c2 == nil {
		t.Errorf("expected CRS %#v but got nil", c1)
	} else if c1 == nil && c2 != nil {
		t.Errorf("expected CRS nil but got %#v", c2)
	} else if !reflect.DeepEqual(c1, c2) {
		t.Errorf("expected CRS %#v but got %#v", c1, c2)
	}
}

func TestCRSsetProperty(t *testing.T) {
	// Success for type name
	c := CRS{
		rawCRS: rawCRS{
			Type:       "name",
			Properties: json.RawMessage(`{"name": "urn:ogc:def:crs:OGC:1.3:CRS84"}`),
		},
	}
	if err := c.setProperty(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Success for type link
	c = CRS{
		rawCRS: rawCRS{
			Type:       "link",
			Properties: json.RawMessage(`{"href": "http://example.com/crs/42", "type": "proj4"}`),
		},
	}
	if err := c.setProperty(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	}

	// Fail on other types
	c = CRS{
		rawCRS: rawCRS{
			Type:       "crs",
			Properties: json.RawMessage(`{}`),
		},
	}
	if err := c.setProperty(); err != ErrInvalidCRS {
		t.Errorf("expected '%v' but got '%v'", ErrInvalidCRS, err)
	}
}

func TestCRSMarshalJSON(t *testing.T) {
	// Success on type name
	c := CRS{
		Name: &CRSName{
			Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
		},
	}
	expected := r.ReplaceAllString(`{"type":"name", "properties": {"name":"urn:ogc:def:crs:OGC:1.3:CRS84"}}`, "")
	if b, err := c.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Success on type link
	c = CRS{
		Link: &CRSLink{
			Href: "data.crs",
			Type: "ogcwkt",
		},
	}
	expected = r.ReplaceAllString(`{"type":"link", "properties": {"href": "data.crs", "type": "ogcwkt"}}`, "")
	if b, err := c.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if r.ReplaceAllString(string(b), "") != expected {
		t.Errorf("expected %q but got %q", expected, string(b))
	}

	// Should error when Link.Href is not parsable
	c = CRS{
		Link: &CRSLink{
			Href: `http://\data.crs`,
			Type: "ogcwkt",
		},
	}
	if b, err := c.MarshalJSON(); err == nil {
		t.Errorf("expected error but got '%v'", err)
	} else if b != nil {
		t.Errorf("expected nil but got %q", string(b))
	}

	// Should return nil when neither is set
	c = CRS{}
	if b, err := c.MarshalJSON(); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else if b != nil {
		t.Errorf("expected nil but got %q", string(b))
	}

	// Setting both Name & Link is an error
	c = CRS{
		Name: &CRSName{
			Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
		},
		Link: &CRSLink{
			Href: "date.crs",
			Type: "ogcwkt",
		},
	}
	if b, err := c.MarshalJSON(); err != ErrMultipleCRSs {
		t.Errorf("expected '%v' but got '%v'", ErrMultipleCRSs, err)
	} else if b != nil {
		t.Errorf("expected nil but got %q", string(b))
	}
}

func TestCRSUnmarshalJSON(t *testing.T) {
	equalCRS := func(c1, c2 CRS) {
		if c1.Type != c2.Type {
			t.Errorf("expected Type %q but got %q", c1.Type, c2.Type)
		}

		if c1.Name != nil && c2.Name == nil {
			t.Errorf("expected Name %#v but got nit", c1.Name)
		} else if c1.Name == nil && c2.Name != nil {
			t.Errorf("expected Name nil but got %#v", c2.Name)
		}

		if c1.Link != nil && c2.Link == nil {
			t.Errorf("expected Link %#v but got nil", c1.Link)
		} else if c1.Link == nil && c2.Link != nil {
			t.Errorf("expected Link nil but got %#v", c2.Link)
		}
	}

	// Success on type name
	expected := CRS{
		rawCRS: rawCRS{
			Type: "name",
		},
		Name: &CRSName{
			Name: "urn:ogc:def:crs:OGC:1.3:CRS84",
		},
	}
	c := CRS{}
	b := []byte(`{"type":"name", "properties": {"name":"urn:ogc:def:crs:OGC:1.3:CRS84"}}`)
	if err := c.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalCRS(expected, c)
	}

	// Success on type link
	expected = CRS{
		rawCRS: rawCRS{
			Type: "link",
		},
		Link: &CRSLink{
			Href: "data.crs",
			Type: "ogcwkt",
		},
	}
	c = CRS{}
	b = []byte(`{"type":"link", "properties": {"href": "data.crs", "type": "ogcwkt"}}`)
	if err := c.UnmarshalJSON(b); err != nil {
		t.Errorf("expected nil but got '%v'", err)
	} else {
		equalCRS(expected, c)
	}

	// Failure on invalid JSON
	c = CRS{}
	b = []byte(`{`)
	if err := c.UnmarshalJSON(b); err == nil {
		t.Error("expected error but got nil")
	}
}
