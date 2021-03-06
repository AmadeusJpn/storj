// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"
	"storj.io/storj/pkg/satellite"
	"storj.io/storj/pkg/satellite/satellitedb/dbx"
	"storj.io/storj/pkg/utils"
)

// implementation of Projects interface repository using spacemonkeygo/dbx orm
type projects struct {
	db *dbx.DB
}

// GetAll is a method for querying all projects from the database.
func (projects *projects) GetAll(ctx context.Context) ([]satellite.Project, error) {
	projectsDbx, err := projects.db.All_Project(ctx)
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(projectsDbx)
}

// GetByOwnerID is a method for querying projects from the database by ownerID.
func (projects *projects) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]satellite.Project, error) {
	projectsDbx, err := projects.db.All_Project_By_OwnerId(ctx, dbx.Project_OwnerId(ownerID[:]))
	if err != nil {
		return nil, err
	}

	return projectsFromDbxSlice(projectsDbx)
}

// Get is a method for querying project from the database by id.
func (projects *projects) Get(ctx context.Context, id uuid.UUID) (*satellite.Project, error) {
	project, err := projects.db.Get_Project_By_Id(ctx, dbx.Project_Id(id[:]))
	if err != nil {
		return nil, err
	}

	return projectFromDBX(project)
}

// Insert is a method for inserting project into the database.
func (projects *projects) Insert(ctx context.Context, project *satellite.Project) (*satellite.Project, error) {
	projectID, err := uuid.New()
	if err != nil {
		return nil, err
	}

	var ownerID dbx.Project_OwnerId_Field

	if project.OwnerID != nil {
		ownerID = dbx.Project_OwnerId(project.OwnerID[:])
	} else {
		ownerID = dbx.Project_OwnerId(nil)
	}

	createdProject, err := projects.db.Create_Project(ctx,
		dbx.Project_Id(projectID[:]),
		dbx.Project_Name(project.Name),
		dbx.Project_Description(project.Description),
		dbx.Project_TermsAccepted(project.TermsAccepted),
		dbx.Project_Create_Fields{
			OwnerId: ownerID,
		})

	if err != nil {
		return nil, err
	}

	return projectFromDBX(createdProject)
}

// Delete is a method for deleting project by Id from the database.
func (projects *projects) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := projects.db.Delete_Project_By_Id(ctx, dbx.Project_Id(id[:]))

	return err
}

// Update is a method for updating user entity
func (projects *projects) Update(ctx context.Context, project *satellite.Project) error {
	_, err := projects.db.Update_Project_By_Id(ctx,
		dbx.Project_Id(project.ID[:]),
		dbx.Project_Update_Fields{
			Name:          dbx.Project_Name(project.Name),
			Description:   dbx.Project_Description(project.Description),
			TermsAccepted: dbx.Project_TermsAccepted(project.TermsAccepted),
		})

	return err
}

// projectFromDBX is used for creating Project entity from autogenerated dbx.Project struct
func projectFromDBX(project *dbx.Project) (*satellite.Project, error) {
	if project == nil {
		return nil, errs.New("project parameter is nil")
	}

	id, err := bytesToUUID(project.Id)
	if err != nil {
		return nil, err
	}

	u := &satellite.Project{
		ID:            id,
		Name:          project.Name,
		Description:   project.Description,
		TermsAccepted: project.TermsAccepted,
		CreatedAt:     project.CreatedAt,
	}

	if project.OwnerId == nil {
		u.OwnerID = nil
	} else {
		ownerID, err := bytesToUUID(project.OwnerId)
		if err != nil {
			return nil, err
		}

		u.OwnerID = &ownerID
	}

	return u, nil
}

// projectsFromDbxSlice is used for creating []Project entities from autogenerated []*dbx.Project struct
func projectsFromDbxSlice(projectsDbx []*dbx.Project) ([]satellite.Project, error) {
	var projects []satellite.Project
	var errors []error

	// Generating []dbo from []dbx and collecting all errors
	for _, projectDbx := range projectsDbx {
		project, err := projectFromDBX(projectDbx)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		projects = append(projects, *project)
	}

	return projects, utils.CombineErrors(errors...)
}
