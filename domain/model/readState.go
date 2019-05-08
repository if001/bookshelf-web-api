package model

type ReadState interface {
	Read
	Reading
	NotRead
}
func NextState(s ReadState) ReadState {
	switch s.(type) {
	case Read:
		return read{}
	case NotRead:
		return notRead{}
	case Reading:
		return reading{}
	default:
		return nil
	}
}
type Read interface {
}
type read struct {
}
type Reading interface {
}
type reading struct {
}
type NotRead interface {
}
type notRead struct {
}
