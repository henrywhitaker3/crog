package domain

type Action interface {
	Execute() Result

	GetName() string
	GetCommand() string
	GetCron() string
}

type Result interface {
	GetErr() error
	GetCode() int
	GetStdout() string
	GetAction() Action
}
