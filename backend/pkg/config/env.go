package config

type Environment string

const (
	EnvDev Environment = "dev"
	EnvPrd Environment = "prd"
)

func (e Environment) String() string {
	return string(e)
}

func (e Environment) IsValid() bool {
	switch e {
	case EnvDev, EnvPrd:
		return true
	default:
		return false
	}
}

func (e Environment) IsProduction() bool {
	return e == EnvPrd
}
