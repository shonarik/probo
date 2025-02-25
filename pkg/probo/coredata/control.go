// Copyright (c) 2025 Probo Inc <hello@getprobo.com>.
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
// LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
// OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
// PERFORMANCE OF THIS SOFTWARE.

package coredata

import (
	"context"
	"fmt"
	"maps"
	"time"

	"github.com/getprobo/probo/pkg/gid"
	"github.com/getprobo/probo/pkg/page"
	"github.com/jackc/pgx/v5"

	"go.gearno.de/kit/pg"
)

type (
	Control struct {
		ID          gid.GID
		FrameworkID gid.GID
		Category    string
		Name        string
		Description string
		State       ControlState
		ContentRef  string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Controls []*Control
)

func (c Control) CursorKey() page.CursorKey {
	return page.NewCursorKey(c.ID, c.CreatedAt)
}

func (c *Control) scan(r pgx.Row) error {
	return r.Scan(
		&c.ID,
		&c.FrameworkID,
		&c.Category,
		&c.Name,
		&c.Description,
		&c.State,
		&c.ContentRef,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
}

func (v *Control) LoadByID(
	ctx context.Context,
	conn pg.Conn,
	scope *Scope,
	controlID gid.GID,
) error {
	q := `
WITH control_states AS (
    SELECT
        control_id,
        to_state,
        reason,
        RANK() OVER w
    FROM
        control_state_transitions
    WHERE
        control_id = @control_id
    WINDOW
        w AS (PARTITION BY control_id ORDER BY created_at DESC)
)
SELECT
    id,
    framework_id,
    category,
    name,
    description,
    cs.to_state AS state,
    content_ref,
    created_at,
    updated_at
FROM
    controls
INNER JOIN
    control_states cs ON cs.control_id = controls.id
WHERE
    %s
    AND id = @control_id
    AND cs.rank = 1
LIMIT 1;
`

	q = fmt.Sprintf(q, scope.SQLFragment())

	args := pgx.NamedArgs{"control_id": controlID}
	maps.Copy(args, scope.SQLArguments())

	r := conn.QueryRow(ctx, q, args)

	c2 := Control{}
	if err := c2.scan(r); err != nil {
		return err
	}

	*v = c2

	return nil
}

func (c Control) Insert(
	ctx context.Context,
	conn pg.Conn,
) error {
	q := `
INSERT INTO
    controls (
        id,
        control_id,
		category,
        name,
        description,
        content_ref,
        state,
        created_at,
        updated_at
    )
VALUES (
    @control_id,
    @framework_id,
    @name,
    @description,
    @content_ref,
    @state,
    @created_at,
    @updated_at
);
`

	args := pgx.NamedArgs{
		"control_id":   c.ID,
		"framework_id": c.FrameworkID,
		"name":         c.Name,
		"description":  c.Description,
		"content_ref":  c.ContentRef,
		"created_at":   c.CreatedAt,
		"updated_at":   c.UpdatedAt,
	}
	_, err := conn.Exec(ctx, q, args)
	return err
}

func (c *Controls) LoadByFrameworkID(
	ctx context.Context,
	conn pg.Conn,
	scope *Scope,
	frameworkID gid.GID,
	cursor *page.Cursor,
) error {
	q := `
WITH control_states AS (
    SELECT
        control_id,
        to_state,
        reason,
        RANK() OVER w
    FROM
        control_state_transitions
    WINDOW
        w AS (PARTITION BY control_id ORDER BY created_at DESC)
)
SELECT
    id,
    framework_id,
	category,
    name,
    description,
    cs.to_state AS state,
    content_ref,
    created_at,
    updated_at
FROM
    controls
INNER JOIN
    control_states cs ON cs.control_id = controls.id
WHERE
    %s
    AND framework_id = @framework_id
    AND cs.rank = 1
    AND %s
`
	q = fmt.Sprintf(q, scope.SQLFragment(), cursor.SQLFragment())

	args := pgx.NamedArgs{"framework_id": frameworkID}
	maps.Copy(args, scope.SQLArguments())
	maps.Copy(args, cursor.SQLArguments())

	r, err := conn.Query(ctx, q, args)
	if err != nil {
		return err
	}
	defer r.Close()

	controls := Controls{}
	for r.Next() {
		control := &Control{}
		if err := control.scan(r); err != nil {
			return err
		}

		controls = append(controls, control)
	}

	if err := r.Err(); err != nil {
		return err
	}

	*c = controls

	return nil
}
