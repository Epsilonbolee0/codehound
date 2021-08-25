package builder

import (
	"../domain"
	"github.com/lib/pq"
)

type testModifier func(test *domain.Test)
type TestBuilder struct {
	actions []testModifier
}

func NewTestBuilder() *TestBuilder {
	return &TestBuilder{}
}

func (b *TestBuilder) Build() domain.Test {
	test := domain.Test{}
	for _, action := range b.actions {
		action(&test)
	}

	return test
}

func (b *TestBuilder) InArgs(len uint) *TestBuilder {
	b.actions = append(b.actions, func(test *domain.Test) {
		arr := make([]string, len)
		test.InArgs = pq.StringArray(arr)
	})
	return b
}

func (b *TestBuilder) OutArgs(len uint) *TestBuilder {
	b.actions = append(b.actions, func(test *domain.Test) {
		arr := make([]string, len)
		test.OutArgs = pq.StringArray(arr)
	})
	return b
}

func (b *TestBuilder) Description(descr string) *TestBuilder {
	b.actions = append(b.actions, func(test *domain.Test) {
		test.Description = descr
	})
	return b
}
func (b *TestBuilder) Author(author uint) *TestBuilder {
	b.actions = append(b.actions, func(test *domain.Test) {
		test.Author = author
	})
	return b
}

func (b *TestBuilder) Implementation(impl string) *TestBuilder {
	b.actions = append(b.actions, func(test *domain.Test) {
		test.Implementation = impl
	})
	return b
}
