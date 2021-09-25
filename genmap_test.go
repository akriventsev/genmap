package genmap

import "testing"

func TestGenmap_Set(t *testing.T) {
	g := FromMap(map[string]interface{}{
		"payload": map[string]interface{}{},
	})
	g.Set("payload.level1.level2", "Karamba")

	if g.GetString("payload.level1.level2") != "Karamba" {
		t.FailNow()
	}
}

func TestGenmap_Set2(t *testing.T) {
	g := FromMap(map[string]interface{}{
		"payload": map[string]interface{}{},
	})
	g.Set("payload.level1.level2", float32(35.5))

	if g.GetFloat32("payload.level1.level2") != 35.5 {
		t.FailNow()
	}

	p := g.ToMap()
	if ((p["payload"].(map[string]interface{}))["level1"].(map[string]interface{}))["level2"].(float32) != 35.5 {
		t.FailNow()
	}
}

type TestType map[string]interface{}

func (t TestType) Map() map[string]interface{} {
	return t
}

func TestGenmap_Set3(t *testing.T) {
	g := FromMap(map[string]interface{}{
		"payload": TestType{
			"value": "Karamba",
		},
	})
	g.Set("payload.level1.level2", "Karamba2")

	if g.GetString("payload.value") != "Karamba" {
		t.FailNow()
	}
	if g.GetString("payload.level1.level2") != "Karamba2" {
		t.FailNow()
	}
}
