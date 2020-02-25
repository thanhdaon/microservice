package helper

type Auth interface {
	CreateToken(uint64) (string, error)
	HashPassword(string) ([]byte, error)
	VerifyPassword(string, string) error
}
