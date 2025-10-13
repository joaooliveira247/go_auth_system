package mocks

import "github.com/jaswdr/faker"

func GenFakePassword() string {
	fake := faker.New()
	return fake.Internet().Password()
}
