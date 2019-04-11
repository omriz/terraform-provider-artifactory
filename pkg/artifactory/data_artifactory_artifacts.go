package artifactory

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceArtifactoryArtifacts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArtifactRead,

		Schema: map[string]*schema.Schema{
			"artifactory_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"artifact_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_token": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"download_location": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceArtifactRead(d *schema.ResourceData, meta interface{}) error {
	// Create the client,
	// Attempt to download to a random place.
	// Do the thing.
	d.Set("download_location", "aaa")
	return nil
}
