package genmap

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Genmap struct {
	payload map[string]interface{}
}

func FromMap(m map[string]interface{}) *Genmap {
	return &Genmap{payload: m}
}

func FromRawStruct(i interface{}) *Genmap {
	m := map[string]interface{}{}
	mapstructure.Decode(i, &m)
	return &Genmap{payload: m}
}

func (g Genmap) getField(path string) interface{} {
	levels := strings.Split(path, ".")
	a := g.payload
	for i, l := range levels {
		x := a[l]
		if i == len(levels)-1 {
			return x
		}
		switch v := x.(type) {
		case map[string]interface{}:
			a = v
		default:
			myMap, ok := x.(map[string]interface{})
			if ok {
				a = myMap
				continue
			}
			return nil
		}
	}
	return nil
}
func (g Genmap) GetInt(path string) int {
	x := g.getField(path)
	if x == nil {
		return 0
	}
	switch v := x.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	default:
		return 0
	}
}

func (g Genmap) GetUint32(path string) uint32 {
	x := g.getField(path)
	if x == nil {
		return 0
	}
	switch v := x.(type) {
	case uint32:
		return v
	case int32:
		return uint32(v)
	case int64:
		return uint32(v)
	case uint:
		return uint32(v)
	case int:
		return uint32(v)
	case uint64:
		return uint32(v)
	default:
		return 0
	}
}

func (g Genmap) GetString(path string) string {
	x := g.getField(path)
	if x == nil {
		return ""
	}
	switch v := x.(type) {
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (g Genmap) GetFloat32(path string) float32 {
	x := g.getField(path)
	if x == nil {
		return 0
	}
	switch v := x.(type) {
	case float32:
		return v
	case float64:
		return float32(v)
	case int:
		return float32(v)
	case int32:
		return float32(v)
	case int64:
		return float32(v)
	case uint:
		return float32(v)
	case uint32:
		return float32(v)
	case uint64:
		return float32(v)
	default:
		return 0
	}
}

func (g Genmap) GetFloat64(path string) float64 {
	x := g.getField(path)
	if x == nil {
		return 0
	}
	switch v := x.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	default:
		return 0
	}
}

func (g *Genmap) Set(path string, value interface{}) {
	levels := strings.Split(path, ".")
	a := g.payload
	for i, l := range levels {
		x := a[l]
		switch x.(type) {
		case map[string]interface{}:
			if i == len(levels)-1 {
				a[l] = value
			}
			a = a[l].(map[string]interface{})
		case nil:
			if i == len(levels)-1 {
				a[l] = value
				return
			}
			a[l] = map[string]interface{}{}
			a = a[l].(map[string]interface{})
		default:
			return
		}
	}
}

func (g *Genmap) ToMap() map[string]interface{} {
	return g.payload
}
