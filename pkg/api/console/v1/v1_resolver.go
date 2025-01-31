package console_v1

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"

	"github.com/getprobo/probo/pkg/api/console/v1/schema"
	"github.com/getprobo/probo/pkg/api/console/v1/types"
	"github.com/getprobo/probo/pkg/probo/coredata"
	"github.com/getprobo/probo/pkg/probo/coredata/gid"
	"github.com/getprobo/probo/pkg/probo/coredata/page"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// StateTransisions is the resolver for the stateTransisions field.
func (r *controlResolver) StateTransisions(ctx context.Context, obj *types.Control, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.ControlStateTransitionConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListControlStateTransitions(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list control tasks: %w", err)
	}

	return types.NewControlStateTransitionConnection(page), nil
}

// Tasks is the resolver for the tasks field.
func (r *controlResolver) Tasks(ctx context.Context, obj *types.Control, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.TaskConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListControlTasks(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list control tasks: %w", err)
	}

	return types.NewTaskConnection(page), nil
}

// StateTransisions is the resolver for the stateTransisions field.
func (r *evidenceResolver) StateTransisions(ctx context.Context, obj *types.Evidence, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.EvidenceStateTransitionConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListEvidenceStateTransitions(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list evidence state transitions: %w", err)
	}

	return types.NewEvidenceStateTransitionConnection(page), nil
}

// Controls is the resolver for the controls field.
func (r *frameworkResolver) Controls(ctx context.Context, obj *types.Framework, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.ControlConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListFrameworkControls(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list framework controls: %w", err)
	}

	return types.NewControlConnection(page), nil
}

// Frameworks is the resolver for the frameworks field.
func (r *organizationResolver) Frameworks(ctx context.Context, obj *types.Organization, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.FrameworkConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListOrganizationFrameworks(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list organization frameworks: %w", err)
	}

	return types.NewFrameworkConnection(page), nil
}

// Vendors is the resolver for the vendors field.
func (r *organizationResolver) Vendors(ctx context.Context, obj *types.Organization, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.VendorConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListOrganizationVendors(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list organization vendors: %w", err)
	}

	return types.NewVendorConnection(page), nil
}

// Peoples is the resolver for the peoples field.
func (r *organizationResolver) Peoples(ctx context.Context, obj *types.Organization, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.PeopleConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListOrganizationPeoples(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list organization peoples: %w", err)
	}

	return types.NewPeopleConnection(page), nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id gid.GID) (types.Node, error) {
	switch id.EntityType() {
	case coredata.OrganizationEntityType:
		organization, err := r.svc.GetOrganization(ctx, id)
		if err != nil {
			return nil, err
		}

		return types.NewOrganization(organization), nil
	case coredata.PeopleEntityType:
		people, err := r.svc.GetPeople(ctx, id)
		if err != nil {
			return nil, err
		}

		return types.NewPeople(people), nil
	case coredata.VendorEntityType:
		vendor, err := r.svc.GetVendor(ctx, id)
		if err != nil {
			return nil, err
		}

		return types.NewVendor(vendor), nil
	default:
	}

	return nil, gqlerror.Errorf("node %q not found", id)
}

// StateTransisions is the resolver for the stateTransisions field.
func (r *taskResolver) StateTransisions(ctx context.Context, obj *types.Task, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.TaskStateTransitionConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListTaskStateTransitions(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list control tasks: %w", err)
	}

	return types.NewTaskStateTransitionConnection(page), nil
}

// Evidences is the resolver for the evidences field.
func (r *taskResolver) Evidences(ctx context.Context, obj *types.Task, first *int, after *page.CursorKey, last *int, before *page.CursorKey) (*types.EvidenceConnection, error) {
	cursor := types.NewCursor(first, after, last, before)

	page, err := r.svc.ListTaskEvidences(ctx, obj.ID, cursor)
	if err != nil {
		return nil, fmt.Errorf("cannot list organization frameworks: %w", err)
	}

	return types.NewEvidenceConnection(page), nil
}

// Control returns schema.ControlResolver implementation.
func (r *Resolver) Control() schema.ControlResolver { return &controlResolver{r} }

// Evidence returns schema.EvidenceResolver implementation.
func (r *Resolver) Evidence() schema.EvidenceResolver { return &evidenceResolver{r} }

// Framework returns schema.FrameworkResolver implementation.
func (r *Resolver) Framework() schema.FrameworkResolver { return &frameworkResolver{r} }

// Organization returns schema.OrganizationResolver implementation.
func (r *Resolver) Organization() schema.OrganizationResolver { return &organizationResolver{r} }

// Query returns schema.QueryResolver implementation.
func (r *Resolver) Query() schema.QueryResolver { return &queryResolver{r} }

// Task returns schema.TaskResolver implementation.
func (r *Resolver) Task() schema.TaskResolver { return &taskResolver{r} }

type controlResolver struct{ *Resolver }
type evidenceResolver struct{ *Resolver }
type frameworkResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type taskResolver struct{ *Resolver }
