// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/getprobo/probo/pkg/probo/coredata/gid"
	"github.com/getprobo/probo/pkg/probo/coredata/page"
)

type Node interface {
	IsNode()
	GetID() gid.GID
}

type Control struct {
	ID               gid.GID                           `json:"id"`
	Name             string                            `json:"name"`
	Description      string                            `json:"description"`
	State            ControlState                      `json:"state"`
	StateTransisions *ControlStateTransitionConnection `json:"stateTransisions"`
	Tasks            *TaskConnection                   `json:"tasks"`
	CreatedAt        time.Time                         `json:"createdAt"`
	UpdatedAt        time.Time                         `json:"updatedAt"`
}

func (Control) IsNode()             {}
func (this Control) GetID() gid.GID { return this.ID }

type ControlConnection struct {
	Edges    []*ControlEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type ControlEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Control       `json:"node"`
}

type ControlStateTransition struct {
	ID        gid.GID       `json:"id"`
	FromState *ControlState `json:"fromState,omitempty"`
	ToState   ControlState  `json:"toState"`
	Reason    *string       `json:"reason,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type ControlStateTransitionConnection struct {
	Edges    []*ControlStateTransitionEdge `json:"edges"`
	PageInfo *PageInfo                     `json:"pageInfo"`
}

type ControlStateTransitionEdge struct {
	Cursor page.CursorKey          `json:"cursor"`
	Node   *ControlStateTransition `json:"node"`
}

type Evidence struct {
	ID        gid.GID   `json:"id"`
	FileURL   string    `json:"fileUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Evidence) IsNode()             {}
func (this Evidence) GetID() gid.GID { return this.ID }

type EvidenceConnection struct {
	Edges    []*EvidenceEdge `json:"edges"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type EvidenceEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Evidence      `json:"node"`
}

type Framework struct {
	ID          gid.GID            `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Controls    *ControlConnection `json:"controls"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func (Framework) IsNode()             {}
func (this Framework) GetID() gid.GID { return this.ID }

type FrameworkConnection struct {
	Edges    []*FrameworkEdge `json:"edges"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type FrameworkEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Framework     `json:"node"`
}

type Organization struct {
	ID         gid.GID              `json:"id"`
	Name       string               `json:"name"`
	Frameworks *FrameworkConnection `json:"frameworks"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
}

func (Organization) IsNode()             {}
func (this Organization) GetID() gid.GID { return this.ID }

type PageInfo struct {
	HasNextPage     bool            `json:"hasNextPage"`
	HasPreviousPage bool            `json:"hasPreviousPage"`
	StartCursor     *page.CursorKey `json:"startCursor,omitempty"`
	EndCursor       *page.CursorKey `json:"endCursor,omitempty"`
}

type Query struct {
}

type Task struct {
	ID          gid.GID             `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Evidences   *EvidenceConnection `json:"evidences"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
}

func (Task) IsNode()             {}
func (this Task) GetID() gid.GID { return this.ID }

type TaskConnection struct {
	Edges    []*TaskEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TaskEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Task          `json:"node"`
}

type ControlState string

const (
	ControlStateNotStarted    ControlState = "NOT_STARTED"
	ControlStateInProgress    ControlState = "IN_PROGRESS"
	ControlStateNotApplicable ControlState = "NOT_APPLICABLE"
	ControlStateImplemented   ControlState = "IMPLEMENTED"
)

var AllControlState = []ControlState{
	ControlStateNotStarted,
	ControlStateInProgress,
	ControlStateNotApplicable,
	ControlStateImplemented,
}

func (e ControlState) IsValid() bool {
	switch e {
	case ControlStateNotStarted, ControlStateInProgress, ControlStateNotApplicable, ControlStateImplemented:
		return true
	}
	return false
}

func (e ControlState) String() string {
	return string(e)
}

func (e *ControlState) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ControlState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ControlState", str)
	}
	return nil
}

func (e ControlState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
