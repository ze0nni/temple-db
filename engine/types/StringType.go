package types

type StringTypeDescription struct {
}

func (*StringTypeDescription) Config() []ConfigKey {
	return []ConfigKey{}
}

func (*StringTypeDescription) New(env CellTypeEnvironment, config []string) (CellType, error) {
	return &stringType{}, nil
}

type stringType struct {
}

func (*stringType) Validate(context CellTypeEnvironment, value string) error {
	return nil
}
