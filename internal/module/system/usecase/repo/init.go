package repo

type InitRepo interface {
	IsInitialized(name string) (bool, error)
	SetInitialized(name, version, description string) error
}
