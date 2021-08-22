package builder

import (
	"../domain"
)

type languageModifier func(account *domain.Language)
type LanguageBuilder struct {
	actions []languageModifier
}

func NewLanguageBuilder() *LanguageBuilder {
	return &LanguageBuilder{}
}

func (b *LanguageBuilder) Build() domain.Language {
	language := domain.Language{}
	for _, action := range b.actions {
		action(&language)
	}

	return language
}

func (b *LanguageBuilder) Name(name string) *LanguageBuilder {
	b.actions = append(b.actions, func(language *domain.Language) {
		language.Name = name
	})
	return b
}

func (b *LanguageBuilder) Version(version string) *LanguageBuilder {
	b.actions = append(b.actions, func(language *domain.Language) {
		language.Version = version
	})
	return b
}
