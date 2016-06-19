package geojson

type Object struct {
	Type        string       `json:"type"`
	BoundingBox *BoundingBox `json:"bbox,omitempty"`
	CRS         *CRS         `json:"crs,omitempty"`
}
