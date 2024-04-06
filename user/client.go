package user

import (
	"context"

	"connectrpc.com/connect"
	"github.com/StatelyCloud/go-sdk/common/client"
	"github.com/StatelyCloud/go-sdk/common/identifiers"
	"github.com/StatelyCloud/go-sdk/common/models"
	pb "github.com/StatelyCloud/go-sdk/pb/user"
	"github.com/StatelyCloud/go-sdk/pb/user/userconnect"
	"github.com/planetscale/vtprotobuf/codec/grpc"
)

type clientImpl struct {
	client userconnect.UserClient
}

// Client is a Stately user client that performs user management operations.
type Client interface {
	// Whoami returns information about the user who calls it (based on the auth
	// token). This includes information about what organizations the user belongs
	// to, what projects they have access to, what roles(?) they can use, etc. This
	// is meant to be called from the Web Console or CLI in order to populate some
	// basic information in the UI and allow calling other APIs like ListStores. It
	// probably shouldn't be included in the client SDK.
	Whoami(ctx context.Context) (*models.WhoamiResponse, error)

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

	// CreateOrganization creates a new organization with the given name.
	CreateOrganization(ctx context.Context, name string, addCaller bool) (*models.OrganizationInfo, error)
}

// NewClient creates a new client with the given store and options.
func NewClient(appCtx context.Context, options *client.Options) (Client, error) {
	options, err := options.ApplyDefaults(appCtx)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		client: userconnect.NewUserClient(options.HTTPClient(), options.Endpoint, connect.WithCodec(grpc.Codec{})),
	}, nil
}

func (c *clientImpl) Whoami(appCtx context.Context) (*models.WhoamiResponse, error) {
	response, err := c.client.Whoami(appCtx, connect.NewRequest(&pb.WhoamiRequest{}))
	if err != nil {
		return nil, err
	}
	organizations := make([]*models.Organization, len(response.Msg.Organizations))
	for i, orgNode := range response.Msg.Organizations {
		organizations[i] = &models.Organization{
			OrganizationInfo: &models.OrganizationInfo{
				ID:   orgNode.Organization.OrganizationId,
				Name: orgNode.Organization.Name,
			},
		}
		projects := make([]*models.Project, len(orgNode.Projects))
		for j, projNode := range orgNode.Projects {
			projects[j] = &models.Project{
				ProjectInfo: &models.ProjectInfo{
					ID:          projNode.Project.ProjectId,
					Name:        projNode.Project.Name,
					Description: projNode.Project.Description,
				},
				Stores: make([]*models.StoreInfo, len(projNode.Stores)),
			}

			for k, store := range projNode.Stores {
				projects[j].Stores[k] = &models.StoreInfo{
					ID:          store.GetStore().GetStoreId(),
					Name:        store.GetStore().GetName(),
					Description: store.GetStore().GetDescription(),
				}
			}
		}
		organizations[i].Projects = projects
	}
	return &models.WhoamiResponse{
		UserInfo: &models.UserInfo{
			OAuthSubject: response.Msg.OauthSubject,
			UserID:       response.Msg.UserId,
		},
		Organizations: organizations,
	}, nil
}

func (c *clientImpl) EnrollMachineUser(
	ctx context.Context,
	oAuthSubject string,
	displayName string,
	organizationID identifiers.OrganizationID,
) error {
	_, err := c.client.EnrollMachineUser(ctx, connect.NewRequest(&pb.EnrollMachineUserRequest{
		OAuthSubject:   oAuthSubject,
		DisplayName:    displayName,
		OrganizationId: uint64(organizationID),
	}))

	return err
}

func (c *clientImpl) CreateOrganization(
	ctx context.Context,
	name string,
	addCaller bool,
) (*models.OrganizationInfo, error) {
	response, err := c.client.CreateOrganization(ctx, connect.NewRequest(&pb.CreateOrganizationRequest{
		Name:                name,
		DoNotAddCurrentUser: !addCaller,
	}))
	if err != nil {
		return nil, err
	}
	return &models.OrganizationInfo{
		ID:   response.Msg.GetOrganizationId(),
		Name: name,
	}, nil
}
