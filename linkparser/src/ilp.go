package linkparser

type Ilp interface {
	LinkParse() (*[]link, error)
}
