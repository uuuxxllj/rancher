package principals

import (
	"context"
	"net/url"

	"github.com/rancher/norman/types"
	"github.com/rancher/rancher/pkg/auth/requests"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	client "github.com/rancher/types/client/management/v3"
	"github.com/rancher/types/config"
)

func Schema(ctx context.Context, clusterRouter requests.ClusterRouter, management *config.ScaledContext, schemas *types.Schemas) error {
	p := newPrincipalsHandler(ctx, clusterRouter, management)
	schema := schemas.Schema(&managementSchema.Version, client.PrincipalType)
	schema.ActionHandler = p.actions
	schema.ListHandler = p.list
	schema.CollectionFormatter = collectionFormatter
	schema.Formatter = formatter
	return nil
}

func collectionFormatter(apiContext *types.APIContext, collection *types.GenericCollection) {
	collection.AddAction(apiContext, "search")
}

func formatter(request *types.APIContext, resource *types.RawResource) {
	schema := request.Schemas.Schema(&managementSchema.Version, client.PrincipalType)
	resource.Links = map[string]string{"self": request.URLBuilder.ResourceLinkByID(schema, url.PathEscape(resource.ID))}
}
