package geojson

type GeoJSON struct {
	Type        string       `json:"type"`
	BoundingBox *BoundingBox `json:"bbox,omitempty"`
	CRS         *CRS         `json:"crs,omitempty"`
}
