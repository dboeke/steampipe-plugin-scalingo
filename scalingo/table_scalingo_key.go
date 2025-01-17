package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableScalingoKey() *plugin.Table {
	return &plugin.Table{
		Name:        "scalingo_key",
		Description: "An SSH key associated to the account",
		List: &plugin.ListConfig{
			Hydrate: listKey,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "unique ID of the key"},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Given name to the key"},
			{Name: "content", Type: proto.ColumnType_STRING, Description: "Raw content of the SSH public key"},
		},
	}
}

func listKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	keys, err := client.KeysList()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		d.StreamListItem(ctx, key)
	}

	return nil, nil
}
