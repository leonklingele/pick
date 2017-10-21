package backends

type Client interface {
	Load() ([]byte, error)
	Save([]byte) error
	Backup() error
	SafeLocation() string
}

func New(config *Config) (Client, error) {
	c, err := clientByName(config.Type)
	if err != nil {
		return nil, err
	}
	return c.newClientFunc(config)
}

func NewWithType(t string, config *Config) (Client, error) {
	config.Type = t
	return New(config)
}
