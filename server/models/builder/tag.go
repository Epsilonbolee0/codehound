package builder

import (
	"../domain"
)

type tagModifier func(tag *domain.Tag)
type TagBuilder struct {
	actions []tagModifier
}

func NewTagBuilder() *TagBuilder {
	return &TagBuilder{}
}

func (b *TagBuilder) Build() domain.Tag {
	tag := domain.Tag{}
	for _, action := range b.actions {
		action(&tag)
	}

	return tag
}

func (b *TagBuilder) Category(category string) *TagBuilder {
	b.actions = append(b.actions, func(tag *domain.Tag) {
		tag.Category = category
	})
	return b
}

func (b *TagBuilder) Content(content string) *TagBuilder {
	b.actions = append(b.actions, func(tag *domain.Tag) {
		tag.Content = content
	})
	return b
}
