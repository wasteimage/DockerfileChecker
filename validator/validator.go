package validator

type Checker interface {
	Check([]byte) error
}

type Validator struct {
	dockerFileContent []byte

	checkers []Checker
}

func New(dockerFileContent []byte) *Validator {
	return &Validator{
		dockerFileContent: dockerFileContent,
	}
}

func (v *Validator) AddChecker(checker Checker) {
	v.checkers = append(v.checkers, checker)
}

func (v *Validator) Validate() error {
	for _, checker := range v.checkers {
		err := checker.Check(v.dockerFileContent)
		if err != nil {
			return err
		}
	}
	return nil
}
