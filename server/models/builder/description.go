package builder

import (
	"../domain"
)

type descriptionModifier func(description *domain.Description)
type DescriptionBuilder struct {
	actions []descriptionModifier
}

func NewDescriptionBuilder() *DescriptionBuilder {
	return &DescriptionBuilder{}
}

func (b *DescriptionBuilder) Build() domain.Description {
	description := domain.Description{}
	for _, action := range b.actions {
		action(&description)
	}

	return description
}

func (b *DescriptionBuilder) Content(content string) *DescriptionBuilder {
	b.actions = append(b.actions, func(description *domain.Description) {
		description.Content = content
	})
	return b
}

func (b *DescriptionBuilder) Author(author uint) *DescriptionBuilder {
	b.actions = append(b.actions, func(description *domain.Description) {
		description.Author = author
	})
	return b
}

func (b *DescriptionBuilder) Implementation(impl string) *DescriptionBuilder {
	b.actions = append(b.actions, func(description *domain.Description) {
		description.Implementation = impl
	})
	return b
}
