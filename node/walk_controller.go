package node

import (
	"meta"
)

type Request struct {
	Selection   *Selection
	Target      PathSlice
	Context     *Context
	Constraints *Constraints
}

type ActionRequest struct {
	Request
	Meta  *meta.Rpc
	Input *Selection
}

type ContainerRequest struct {
	Request
	New  bool
	Meta meta.MetaList
}

type ListRequest struct {
	Request
	New      bool
	StartRow int64
	Row      int64
	First    bool
	Meta     *meta.List
	Key      []*Value
}

type FieldRequest struct {
	Request
	Meta meta.HasDataType
}

type WalkController interface {
	ContainerIterator(sel *Selection, m meta.MetaList) (meta.MetaIterator, error)
	VisitList(r *ListRequest) (next *Selection, err error)
	VisitContainer(r *ContainerRequest) (child *Selection, err error)
	VisitAction(r *ActionRequest) (*Selection, error)
	VisitField(r *FieldRequest) (*Value, error)
}
