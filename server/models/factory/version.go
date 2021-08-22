package factory

import (
	"time"

	"../domain"
)

type versionModifier func(version *domain.Version)
type VersionBuilder struct {
	actions []versionModifier
}

func NewVersionBuilder() *VersionBuilder {
	return &VersionBuilder{}
}

func (b *VersionBuilder) Build() domain.Version {
	version := domain.Version{}
	for _, action := range b.actions {
		action(&version)
	}

	return version
}

func (b *VersionBuilder) Name(name string) *VersionBuilder {
	b.actions = append(b.actions, func(version *domain.Version) {
		version.Name = name
	})
	return b
}

func (b *VersionBuilder) Code(code string) *VersionBuilder {
	b.actions = append(b.actions, func(version *domain.Version) {
		version.Code = code
	})
	return b
}

func (b *VersionBuilder) Date(date time.Time) *VersionBuilder {
	b.actions = append(b.actions, func(version *domain.Version) {
		version.Date = date
	})
	return b
}

func (b *VersionBuilder) Author(author uint) *VersionBuilder {
	b.actions = append(b.actions, func(version *domain.Version) {
		version.Author = author
	})
	return b
}

func (b *VersionBuilder) Implementation(impl string) *VersionBuilder {
	b.actions = append(b.actions, func(version *domain.Version) {
		version.Implementation = impl
	})
	return b
}
