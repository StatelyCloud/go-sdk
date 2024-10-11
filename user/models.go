package user

import (
	"github.com/StatelyCloud/go-sdk/dbmanagement"
	"github.com/StatelyCloud/go-sdk/stately"
)

// Organization is info about the org and the projects which it contains.
type Organization struct {
	OrganizationInfo *OrganizationInfo
	Projects         []*Project
}

// OrganizationInfo is info about the org.
type OrganizationInfo struct {
	ID   stately.OrganizationID
	Name string
}

// Project is info about the project as well as the stores which it contains.
type Project struct {
	ProjectInfo *ProjectInfo
	Stores      []*dbmanagement.StoreInfo
	Schemas     []*dbmanagement.SchemaInfo
}

// ProjectInfo is info about the project.
type ProjectInfo struct {
	ID          stately.ProjectID
	Name        string
	Description string
}

// UserInfo is information about the user.
//
//nolint:revive // This corresponds to an API type, ignore "stutter"
type UserInfo struct {
	OAuthSubject string
	UserID       stately.UserID
}

// WhoamiResponse is information about the user as well as the organizations which they are a member.
type WhoamiResponse struct {
	UserInfo      *UserInfo
	Organizations []*Organization
}
