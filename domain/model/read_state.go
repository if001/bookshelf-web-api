package model

type ReadState int8
const (
	notReadValue ReadState = 0
	readingValue ReadState = 1
	readValue ReadState = 2
)