package iface

type Line struct {
	Idx  int
	Text string
}

type DB interface {
	//	Search(str string) (error, []Line)

	//	RegExp(str string) (error, []Line)

	//	Insert(idx int, text string) error
	Update(idx int, text string) error

	Delete(idx int) error

	Length() int

	//	Unwrap(str string)
	//	Wrap() string
}
