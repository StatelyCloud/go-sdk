package user

import (
	"context"

	"connectrpc.com/connect"
	"github.com/planetscale/vtprotobuf/codec/grpc"

	"github.com/StatelyCloud/go-sdk/dbmanagement"
	pbuser "github.com/StatelyCloud/go-sdk/pb/user"
	"github.com/StatelyCloud/go-sdk/pb/user/userconnect"
	"github.com/StatelyCloud/go-sdk/sdkerror"
	"github.com/StatelyCloud/go-sdk/stately"
)

type clientImpl struct {
	client userconnect.UserServiceClient
}

// Client is a Stately user client that performs user management operations.
type Client interface {
	// Whoami returns information about the user who calls it (based on the auth
	// token). This includes information about what organizations the user belongs
	// to, what projects they have access to, what roles(?) they can use, etc. This
	// is meant to be called from the Web Console or CLI in order to populate some
	// basic information in the UI and allow calling other APIs like ListStores. It
	// probably shouldn't be included in the client SDK.
	Whoami(ctx context.Context) (*WhoamiResponse, error)

	// EnrollMachineUser enrolls the given oAuthSubject into the given organization and creates
	// a stately user with the given display name.
	// If called multiple times with different values the display name will not be updated
	// but the user will be added to the new organizations.
	// This is an admin only API.
	EnrollMachineUser(
		ctx context.Context,
		oAuthSubject string,
		displayName string,
		organizationID stately.OrganizationID,
	) error

	// CreateOrganization creates a new organization with the given name.
	CreateOrganization(ctx context.Context, name string, addCaller bool) (*OrganizationInfo, error)
}

// NewClient creates a new client with the given store and options.
func NewClient(appCtx context.Context, options ...*stately.Options) (Client, error) {
	opts := &stately.Options{}
	for _, o := range options {
		opts = opts.Merge(o)
	}
	opts, err := opts.ApplyDefaults(appCtx)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		client: userconnect.NewUserServiceClient(
			opts.HTTPClient(),
			opts.Endpoint,
			connect.WithCodec(grpc.Codec{}),
			connect.WithInterceptors(sdkerror.ConnectErrorInterceptor()),
		),
	}, nil
}

func (c *clientImpl) Whoami(appCtx context.Context) (*WhoamiResponse, error) {
	response, err := c.client.Whoami(appCtx, connect.NewRequest(&pbuser.WhoamiRequest{}))
	if err != nil {
		return nil, err
	}
	organizations := make([]*Organization, len(response.Msg.Organizations))
	for i, orgNode := range response.Msg.Organizations {
		organizations[i] = &Organization{
			OrganizationInfo: &OrganizationInfo{
				ID:   stately.OrganizationID(orgNode.Organization.OrganizationId),
				Name: orgNode.Organization.Name,
			},
		}
		projects := make([]*Project, len(orgNode.Projects))
		for j, projNode := range orgNode.Projects {
			projects[j] = &Project{
				ProjectInfo: &ProjectInfo{
					ID:          stately.ProjectID(projNode.Project.ProjectId),
					Name:        projNode.Project.Name,
					Description: projNode.Project.Description,
				},
				Stores:  make([]*dbmanagement.StoreInfo, len(projNode.Stores)),
				Schemas: make([]*dbmanagement.SchemaInfo, len(projNode.Schemas)),
			}

			for k, store := range projNode.Stores {
				projects[j].Stores[k] = &dbmanagement.StoreInfo{
					ID:            stately.StoreID(store.GetStore().GetStoreId()),
					Name:          store.GetStore().GetName(),
					Description:   store.GetStore().GetDescription(),
					DefaultRegion: store.GetStore().GetDefaultRegion(),
					SchemaID:      stately.SchemaID(store.GetStore().GetSchemaId()),
				}
			}
			for k, schema := range projNode.Schemas {
				projects[j].Schemas[k] = &dbmanagement.SchemaInfo{
					ID:          stately.SchemaID(schema.GetSchemaId()),
					Name:        schema.GetName(),
					Description: schema.GetDescription(),
				}
			}
		}
		organizations[i].Projects = projects
	}
	return &WhoamiResponse{
		UserInfo: &UserInfo{
			OAuthSubject: response.Msg.OauthSubject,
			UserID:       stately.UserID(response.Msg.UserId),
		},
		Organizations: organizations,
	}, nil
}

func (c *clientImpl) EnrollMachineUser(
	ctx context.Context,
	oAuthSubject string,
	displayName string,
	organizationID stately.OrganizationID,
) error {
	_, err := c.client.EnrollMachineUser(ctx, connect.NewRequest(&pbuser.EnrollMachineUserRequest{
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
) (*OrganizationInfo, error) {
	response, err := c.client.CreateOrganization(ctx, connect.NewRequest(&pbuser.CreateOrganizationRequest{
		Name:                name,
		DoNotAddCurrentUser: !addCaller,
	}))
	if err != nil {
		return nil, err
	}
	return &OrganizationInfo{
		ID:   stately.OrganizationID(response.Msg.GetOrganizationId()),
		Name: name,
	}, nil
}
