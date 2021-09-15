package endpoint

import (
	"encoding/json"
	"net/url"
)

type Endpoint struct {
	Uri     string `json:"uri"`
	Module  string `json:"module"`
	Headers map[string]string
	Values  map[string][]string
}

type Container struct {
	Map map[string]*Endpoint
}

func NewContainer(list []*DBEndpoint, general []*DBFlowEndpoint) *Container {
	c := &Container{}
	c.Map = make(map[string]*Endpoint)
	for _, v := range list {
		c.Map[v.Module] = &Endpoint{
			Uri:     v.Uri,
			Module:  v.Module,
			Headers: getHeadersOrNothing(v.Headers),
			Values:  getUrlValuesOrNothing(v.Query),
		}
	}
	for _, v := range general {
		c.Map[v.Module] = &Endpoint{
			Uri:     v.Uri,
			Module:  v.Module,
			Headers: getHeadersOrNothing(v.Headers),
			Values:  getUrlValuesOrNothing(v.Query),
		}
	}
	return c
}

func getHeadersOrNothing(obj string) map[string]string {
	m := make(map[string]string)
	_ = json.Unmarshal([]byte(obj), &m)
	return m
}

func getUrlValuesOrNothing(obj string) url.Values {
	m := make(map[string]interface{})
	vals := make(url.Values)
	err := json.Unmarshal([]byte(obj), &m)
	if err != nil {
		return vals
	}
	for k, v := range m {
		switch tv := v.(type) {
		case string:
			vals[k] = []string{tv}
		case []interface{}:
			vals[k] = stringSlice(tv)
		}
	}
	return vals
}

func stringSlice(slice []interface{}) []string {
	var ret []string
	for _, v := range slice {
		if s, ok := v.(string); ok {
			ret = append(ret, s)
		}
	}
	return ret
}
