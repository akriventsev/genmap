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
