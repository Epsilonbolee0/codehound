package builder

import (
	"../domain"
)

type implementationModifier func(account *domain.Implementation)
type ImplementationBuilder struct {
	actions []implementationModifier
}

func NewImplementationBuilder() *ImplementationBuilder {
	return &ImplementationBuilder{}
}

func (b *ImplementationBuilder) Build() domain.Implementation {
	impl := domain.Implementation{}
	for _, action := range b.actions {
		action(&impl)
	}

	return impl
}

func (b *ImplementationBuilder) Name(name string) *ImplementationBuilder {
	b.actions = append(b.actions, func(impl *domain.Implementation) {
		impl.Name = name
	})
	return b
}

func (b *ImplementationBuilder) Language(languageID uint) *ImplementationBuilder {
	b.actions = append(b.actions, func(impl *domain.Implementation) {
		impl.LanguageID = languageID
	})
	return b
}
