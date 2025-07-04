package eval

type String struct {
	value string
}

func NewString(value string) *String {
	return &String{value: value}
}

func (s *String) Serialize() []byte {
	return []byte(s.value)
}
