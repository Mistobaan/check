package check

type validator struct{}

func New() int {
	return 0
}

func (v *validator) AddError(key string, err error) {

}

func (v *validator) Int(value int, name string, a ...func(int) error) {
	for _, f := range a {
		if err := f(value); err != nil {
			v.AddError(name, err)
		}
	}
}

func (v *validator) String(value string, name string, a ...func(string) error) {
	for _, f := range a {
		if err := f(value); err != nil {
			v.AddError(name, err)
		}
	}
}
