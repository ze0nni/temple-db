package types

func CommonTypes() []StringTypeDescription {
	return []StringTypeDescription{
		StringTypeDescription{},
	}
}

// CellTypeEnvironment for overview environment
type CellTypeEnvironment interface {
	Type(name string, config []string) CellType
}

type ConfigKey struct {
	Key  string
	Type string
}

type CellTypeDescription interface {
	Config() []ConfigKey
	New(env CellTypeEnvironment, config []string) (CellType, error)
}

type CellType interface {
	Validate(
		env CellTypeEnvironment,
		value string,
	) error
}
