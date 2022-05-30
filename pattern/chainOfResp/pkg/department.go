package pkg

type Department interface {
	Execute(*Patient)
	SetNext(department Department)
}
