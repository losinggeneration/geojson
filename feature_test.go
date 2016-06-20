package geojson

import (
	"reflect"
	"testing"
)

func equalFeatures(f1, f2 *Feature, t *testing.T) {
	if f1 == nil && f2 == nil {
		return
	} else if f1 != nil && f2 == nil {
		t.Errorf("expected Feature %#v but got nil", f1)
		return
	} else if f1 == nil && f2 != nil {
		t.Errorf("expected Feature nil but got %#v", f2)
		return
	}

	if f1.ID != f2.ID {
		t.Errorf("expected Feature.ID '%v' but got '%v'", f1.ID, f2.ID)
	}

	equalGeometries(f1.Geometry, f2.Geometry, t)
	if f1.Properties != nil && f2.Properties != nil {
		if !reflect.DeepEqual(f1.Properties, f2.Properties) {
			t.Errorf("expected Feature.Properties '%#v' but got '%#v'", f1.Properties, f2.Properties)
		}
	} else if f1.Properties != nil && f2.Properties == nil {
		t.Errorf("expected Feature.Properties%#v but got nil", f1)
	} else if f1.Properties == nil && f2.Properties != nil {
		t.Errorf("expected Feature.Properties nil but got %#v", f1)
	}

}

func equalFeatureCollections(f1, f2 *FeatureCollection, t *testing.T) {
	if f1 == nil && f2 == nil {
		return
	} else if f1 != nil && f2 == nil {
		t.Errorf("expected FeatureCollection %#v but got nil", f1)
	} else if f1 == nil && f2 != nil {
		t.Errorf("expected FeatureCollection nil but got %#v", f2)
	} else {
		equalObject(f1.Object, f2.Object, t)
		if len(f1.Features) != len(f2.Features) {
			t.Errorf("expected FeatureCollection.Feature length of %v but got %v", len(f1.Features), len(f2.Features))
		}
		for i := range f1.Features {
			equalFeatures(&f1.Features[i], &f2.Features[i], t)
		}
	}
}
