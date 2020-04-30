package helpers

import "strings"

type StringSlice []string

func (s StringSlice) ContainsAny(values ...string) bool {
	return len(s.Filter(values...)) != 0
}

func (s StringSlice) ContainsAll(values ...string) bool {
	return len(s.Filter(values...)) == len(values)
}

func (s StringSlice) Filter(values ...string) StringSlice {
	keys := make(map[string]struct{})
	for _, v := range values {
		keys[v] = struct{}{}
	}
	var result []string
	for _, v := range s {
		if _, ok := keys[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

func (s StringSlice) ToMap() map[string]struct{} {
	keys := make(map[string]struct{})
	for _, v := range s {
		keys[v] = struct{}{}
	}
	return keys
}

func (s StringSlice) Distinct() StringSlice {
	keys := s.ToMap()
	var result []string
	for k := range keys {
		result = append(result, k)
	}
	return result
}

func (s StringSlice) Except(values ...string) StringSlice {
	keys := make(map[string]struct{})
	for _, v := range values {
		keys[v] = struct{}{}
	}
	var result []string
	for _, v := range s {
		if _, ok := keys[v]; !ok {
			result = append(result, v)
		}
	}
	return result
}

func (s StringSlice) ToLower() StringSlice {
	var res StringSlice
	for _, v := range s {
		res = append(res, strings.ToLower(v))
	}
	return res
}
