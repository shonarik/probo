// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package types

import (
	"time"

	"github.com/getprobo/probo/pkg/probo/coredata/page"
)

type Node interface {
	IsNode()
	GetID() string
}

type Control struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Tasks       *TaskConnection `json:"tasks"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

func (Control) IsNode()            {}
func (this Control) GetID() string { return this.ID }

type ControlConnection struct {
	Edges    []*ControlEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type ControlEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Control       `json:"node"`
}

type Framework struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Controls    *ControlConnection `json:"controls"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func (Framework) IsNode()            {}
func (this Framework) GetID() string { return this.ID }

type FrameworkConnection struct {
	Edges    []*FrameworkEdge `json:"edges"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type FrameworkEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Framework     `json:"node"`
}

type Organization struct {
	ID         string               `json:"id"`
	Name       string               `json:"name"`
	Frameworks *FrameworkConnection `json:"frameworks"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
}

func (Organization) IsNode()            {}
func (this Organization) GetID() string { return this.ID }

type PageInfo struct {
	HasNextPage     bool            `json:"hasNextPage"`
	HasPreviousPage bool            `json:"hasPreviousPage"`
	StartCursor     *page.CursorKey `json:"startCursor,omitempty"`
	EndCursor       *page.CursorKey `json:"endCursor,omitempty"`
}

type Query struct {
}

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Task) IsNode()            {}
func (this Task) GetID() string { return this.ID }

type TaskConnection struct {
	Edges    []*TaskEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TaskEdge struct {
	Cursor page.CursorKey `json:"cursor"`
	Node   *Task          `json:"node"`
}
