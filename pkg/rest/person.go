package rest

import (
	"zupu/pkg/business"
)

func (rs *RS) GetPerson(ctx Context) {
	person, err := business.GetPerson()
	if err != nil {

	}
	ctx.SuccessWithResult(person)
}

func (rs *RS) SSH(ctx Context) {

}
