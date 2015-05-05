package check_test

import (
	"testing"

	"github.com/Mistobaan/check"
)

var ObjectValidator = check.New(
	check.AsInt("valueInt",
		check.IsRange(0, 100),
		check.IsInt,
	),
	check.AsString("valueString",
		check.StringNotEmptyf("custom message here"),
		check.IsUppercase,
	),
)

type ObjectToValidate struct {
	valueString string
	valueInt    int
}

func (o *ObjectToValidate) Validate(v *Validator) error {
	v.Int("valueInt", o.valueInt)
	v.String("valueString", o.valueString)
	v.MustBeTrue(o.valueInt != o.valueString, "some message")
	return v.Err()
}

// // func Validate() error {
// // 	func(obj *ObjectToValidate) {
// // 		return check.Apply(
// // 			check.IsRange(obj.valueInt, 0, 100),
// // 			check.IsInt(obj.valueInt, 0),
// // 			check.StringNotEmpty(obj.valueString),
// // 			check.StringIsUppercase(obj.valueString),
// // 		)
// // 	}
// // }

func TestValidator(t *testing.T) {
	v := check.New()
	o := &ObjectValidator{}
	o.Validate(v)

}
