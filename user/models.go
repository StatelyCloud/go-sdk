package user

import (
	"github.com/StatelyCloud/go-sdk/client"
	"github.com/StatelyCloud/go-sdk/dbmanagement"
)

// Organization is info about the org and the projects which it contains.
type Organization struct {
	OrganizationInfo *OrganizationInfo
	Projects         []*Project
}

// OrganizationInfo is info about the org.
type OrganizationInfo struct {
	ID   client.OrganizationID
	Name string
}

// Project is info about the project as well as the stores which it contains.
type Project struct {
	ProjectInfo *ProjectInfo
	Stores      []*dbmanagement.StoreInfo
}

// ProjectInfo is info about the project.
type ProjectInfo struct {
	ID          client.ProjectID
	Name        string
	Description string
}

// UserInfo is information about the user.
//
//nolint:revive // This corresponds to an API type, ignore "stutter"
type UserInfo struct {
	OAuthSubject string
	UserID       client.UserID
}

// WhoamiResponse is information about the user as well as the organizations which they are a member.
type WhoamiResponse struct {
	UserInfo      *UserInfo
	Organizations []*Organization
}
