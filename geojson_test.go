package geojson

import "regexp"

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile(`\s`)
}
