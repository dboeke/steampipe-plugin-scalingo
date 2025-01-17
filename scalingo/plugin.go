package scalingo

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-scalingo",
		DefaultTransform: transform.FromGo().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"scalingo_addon":        tableScalingoAddon(),
			"scalingo_app":          tableScalingoApp(),
			"scalingo_app_event":    tableScalingoAppEvent(),
			"scalingo_collaborator": tableScalingoCollaborator(),
			"scalingo_deployment":   tableScalingoDeployment(),
			"scalingo_domain":       tableScalingoDomain(),
			"scalingo_key":          tableScalingoKey(),
			"scalingo_region":       tableScalingoRegion(),
			"scalingo_token":        tableScalingoToken(),
		},
	}
	return p
}
