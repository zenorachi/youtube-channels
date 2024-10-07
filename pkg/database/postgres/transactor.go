package postgres

type ITransactor interface {
	WithTransaction(work func() error) error
}
