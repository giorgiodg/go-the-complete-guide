package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteContent(data interface{}) error
}
