package eval

type HashMap map[string]string

func (h HashMap) Get(k string) (*string, bool) {
	value, ok := h[k]
	if !ok {
		return nil, false
	}
	return &value, true
}

func (h HashMap) Set(k, v string) (*string, bool) {
	value, ok := h[k]
	if ok {
		oldValue := value
		h[k] = v
		return &oldValue, true
	}

	h[k] = v
	return nil, false
}
