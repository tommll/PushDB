package eval

import (
	"errors"
	"pushdb/internal/store"
)

func evalHGET(args []string, s *store.Store) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("wrong number of arguments for HGET, expected 2")
	}

	key := args[0]
	field := args[1]

	resp, err := s.GetValueFromHashMap(key, field)
	if err != nil {
		return nil, err
	}

	return resp.Serialize(), nil
}
