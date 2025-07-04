package store

import (
	"errors"
	"pushdb/internal/object"
)

type Store struct {
	store            map[string]*object.Obj
	expire           map[*object.Obj]int64
	evictionStrategy EvictionStrategy
}

func (s *Store) GetValueFromHashMap(key, field string) (*object.Obj, error) {
	obj, exists := s.store[key]
	if !exists {
		return nil, nil // Key not found
	}
	if obj.Type != object.HashMapType {
		return nil, errors.New("value is not a HashMap")
	}

	hashMap, ok := obj.Value.(map[string]*object.Obj)
	if !ok {
		return nil, errors.New("value is not a HashMap")
	}

	value, ok := hashMap[field]
	if !ok {
		return nil, errors.New("field not found")
	}
	return value, nil
}
