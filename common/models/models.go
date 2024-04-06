package models

// Organization is info about the org and the projects which it contains.
type Organization struct {
	OrganizationInfo *OrganizationInfo
	Projects         []*Project
}

// OrganizationInfo is info about the org.
type OrganizationInfo struct {
	ID   uint64
	Name string
}

// Project is info about the project as well as the stores which it contains.
type Project struct {
	ProjectInfo *ProjectInfo
	Stores      []*StoreInfo
}

// ProjectInfo is info about the project.
type ProjectInfo struct {
	ID          uint64
	Name        string
	Description string
}

// StoreInfo is information about the store.
type StoreInfo struct {
	ID          uint64
	Name        string
	Description string
}

// UserInfo is information about the user.
type UserInfo struct {
	OAuthSubject string
	UserID       uint64
}

// WhoamiResponse is information about the user as well as the organizations which they are a member.
type WhoamiResponse struct {
	UserInfo      *UserInfo
	Organizations []*Organization
}
