package domain

type Action interface {
	Execute() Result

	GetName() string
	GetCommand() string
	GetCron() string
	GetTries() int
}

type Result interface {
	GetErr() error
	GetCode() int
	GetStdout() string
	GetAction() Action
	GetTries() int
}
