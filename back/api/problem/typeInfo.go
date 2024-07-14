package problem

type TypeInfo struct {
	Type_id   int
	Type_name string
}

func NewTypeInfo() *TypeInfo {
	return &TypeInfo{}
}
