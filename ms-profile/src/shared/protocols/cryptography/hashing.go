package cryptography

type Hashing interface {
	CreateHash(payload string) (string, error)
	CompareHash(payload string, hash string) bool
}
