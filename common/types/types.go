package types

import dbCommon "github.com/StatelyCloud/go-sdk/db/common"

// Organization response.
type Organization struct {
	ID       uint64
	Name     string
	Projects []Project
}

// Project response.
type Project struct {
	ID          uint64
	Name        string
	Description string
	Stores      []dbCommon.Store
}
