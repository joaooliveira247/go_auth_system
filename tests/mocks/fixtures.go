package mocks

import (
	"github.com/go-faker/faker/v4"
)

func GenFakePassword() string {
	return faker.Password()
}
