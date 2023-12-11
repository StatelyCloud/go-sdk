package user

import (
	"context"

	"google.golang.org/grpc"

	"github.com/StatelyCloud/go-sdk/common/auth"
	"github.com/StatelyCloud/go-sdk/common/identifiers"
	"github.com/StatelyCloud/go-sdk/common/net"
	"github.com/StatelyCloud/go-sdk/common/types"
	"github.com/StatelyCloud/go-sdk/db/common"
	pb "github.com/StatelyCloud/go-sdk/pb/user"
)

type clientImpl struct {
	client pb.UserClient
}

// Client is a Stately user client that performs user management operations.
type Client interface {
	// Whoami returns information about the user who calls it (based on the auth
	// token). This includes information about what organizations the user belongs
	// to, what projects they have access to, what roles(?) they can use, etc. This
	// is meant to be called from the Web Console or CLI in order to populate some
	// basic information in the UI and allow calling other APIs like ListStores. It
	// probably shouldn't be included in the client SDK.
	Whoami(ctx context.Context) (Whoami, error)

	// EnrollMachineUser enrolls the given oAuthSubject into the given organization and creates
	// a stately user with the given display name.
	// If called multiple times with different values the display name will not be updated
	// but the user will be added to the new organizations.
	// This is an admin only API.
	EnrollMachineUser(
		ctx context.Context,
		oAuthSubject string,
		displayName string,
		organizationID identifiers.OrganizationID,
	) error
}

// NewClientWithAuthProvider builds a new client that authenticates with the given AuthTokenProvider.
func NewClientWithAuthProvider(
	appCtx context.Context,
	provider auth.TokenProvider,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	conn, err := net.NewStatelyConnection(appCtx, provider, options...)
	if err != nil {
		return nil, err
	}
	return NewClientWithConn(conn), nil
}

// NewClientWithCreds builds a new client that authenticates with the given auth credentials.
func NewClientWithCreds(
	appCtx context.Context,
	clientID, clientSecret string,
	options ...net.StatelyConnectionOption,
) (Client, error) {
	return NewClientWithAuthProvider(appCtx, auth.NewAuth0TokenProvider(appCtx, clientID, clientSecret), options...)
}

// NewClientWithConn builds a new client with the provided connection to the stately API.
func NewClientWithConn(conn *grpc.ClientConn) Client {
	return &clientImpl{
		client: pb.NewUserClient(conn),
	}
}

func (c *clientImpl) Whoami(ctx context.Context) (Whoami, error) {
	response, err := c.client.Whoami(ctx, &pb.WhoamiRequest{})
	if err != nil {
		return Whoami{}, err
	}
	organizations := make([]types.Organization, len(response.Organizations))
	for i, orgNode := range response.Organizations {
		organizations[i] = types.Organization{
			ID:   orgNode.Organization.OrganizationId,
			Name: orgNode.Organization.Name,
		}
		projects := make([]types.Project, len(orgNode.Projects))
		for j, projNode := range orgNode.Projects {
			projects[j] = types.Project{
				ID:          projNode.Project.ProjectId,
				Name:        projNode.Project.Name,
				Description: projNode.Project.Description,
				Stores:      make([]common.Store, len(projNode.Stores)),
			}

			for k, store := range projNode.Stores {
				projects[j].Stores[k] = common.Store{
					ID:          store.GetStore().GetStoreId(),
					Name:        store.GetStore().GetName(),
					Description: store.GetStore().GetDescription(),
				}
			}
		}
		organizations[i].Projects = projects
	}
	return Whoami{
		OAuthSubject:  response.OauthSubject,
		UserID:        response.UserId,
		Organizations: organizations,
	}, nil
}

func (c *clientImpl) EnrollMachineUser(
	ctx context.Context,
	oAuthSubject string,
	displayName string,
	organizationID identifiers.OrganizationID,
) error {
	_, err := c.client.EnrollMachineUser(ctx, &pb.EnrollMachineUserRequest{
		OAuthSubject:   oAuthSubject,
		DisplayName:    displayName,
		OrganizationId: uint64(organizationID),
	})

	return err
}
