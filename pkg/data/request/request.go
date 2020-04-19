package request

type Request interface {
	GetPath() string
	GetFilter() (Filter, error)
}
