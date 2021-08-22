package builder

import (
	"../domain"
)

type treeModifier func(account *domain.Tree)
type TreeBuilder struct {
	actions []treeModifier
}

func NewTreeBuilder() *TreeBuilder {
	return &TreeBuilder{}
}

func (b *TreeBuilder) Build() domain.Tree {
	tree := domain.Tree{}
	for _, action := range b.actions {
		action(&tree)
	}

	return tree
}

func (b *TreeBuilder) Name(name string) *TreeBuilder {
	b.actions = append(b.actions, func(tree *domain.Tree) {
		tree.Name = name
	})
	return b
}

func (b *TreeBuilder) Link(link string) *TreeBuilder {
	b.actions = append(b.actions, func(tree *domain.Tree) {
		tree.Link = link
	})
	return b
}

func (b *TreeBuilder) ParentName(parentName string) *TreeBuilder {
	b.actions = append(b.actions, func(tree *domain.Tree) {
		tree.ParentName = &parentName
	})
	return b
}
