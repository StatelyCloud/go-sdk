package user

import "github.com/StatelyCloud/go-sdk/common/types"

// Whoami response.
type Whoami struct {
	OAuthSubject  string
	UserID        uint64
	Organizations []types.Organization
}
