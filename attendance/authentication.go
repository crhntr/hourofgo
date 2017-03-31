package attendance

type Passer interface {
	Validator
	Difference() int
}

type Passphrase string

func (ps Passphrase) Valid() bool {
	return false
}

func (ps Passphrase) Difference() int {
	return 1
}

type Password string

func (ps Password) Valid() bool {
	return false
}

func (ps Password) Difference() int {
	return 1
}
