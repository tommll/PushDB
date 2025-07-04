package object

type ObjType uint8

const (
	StringType ObjType = iota
	HashMapType
)

type ObjI interface {
	Serialize() []byte
}

type Obj struct {
	Type  ObjType
	Value interface{}
}

func (o *Obj) Serialize() []byte {
	switch o.Type {
	case StringType:
		if str, ok := o.Value.(string); ok {
			return []byte(str)
		}
		return nil
	}

	return nil
}
