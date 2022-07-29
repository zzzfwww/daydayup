package test

type TStruct struct {
	num int
}

func NewTest() *TStruct {
	return &TStruct{
		num: 10,
	}
}
