package object

type QueryKeyClass string

const (
	QueryKeyClassApplication QueryKeyClass = "1"
	QueryKeyClassBulk        QueryKeyClass = "2"
)

func (q QueryKeyClass) Value() string {
	return string(q)
}
