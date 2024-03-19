package store

// A store holds the configuration

type ConfigStore interface {
	Content() (string, error)
	Update(string) error
}
