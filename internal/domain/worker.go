package domain

type Worker interface {
	Start() error
	Stop() error
}
