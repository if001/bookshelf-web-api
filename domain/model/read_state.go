package model

type ReadState int8
const (
	NotReadValue ReadState = 0
	ReadingValue ReadState = 1
	ReadValue    ReadState = 2
)

