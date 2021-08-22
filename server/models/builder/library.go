package builder

import (
	"../domain"
)

type libraryModifier func(account *domain.Library)
type LibraryBuilder struct {
	actions []libraryModifier
}

func NewLibraryBuilder() *LibraryBuilder {
	return &LibraryBuilder{}
}

func (b *LibraryBuilder) Build() domain.Library {
	library := domain.Library{}
	for _, action := range b.actions {
		action(&library)
	}

	return library
}

func (b *LibraryBuilder) Name(name string) *LibraryBuilder {
	b.actions = append(b.actions, func(library *domain.Library) {
		library.Name = name
	})
	return b
}

func (b *LibraryBuilder) Version(version string) *LibraryBuilder {
	b.actions = append(b.actions, func(library *domain.Library) {
		library.Version = version
	})
	return b
}

func (b *LibraryBuilder) Language(language uint) *LibraryBuilder {
	b.actions = append(b.actions, func(library *domain.Library) {
		library.LanguageID = language
	})
	return b
}
