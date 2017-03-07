package main

import (
  "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceInstanceName() *schema.Resource {
  return &schema.Resource{
    Read:   dataSourceInstanceNameRead,

    Schema: map[string]*schema.Schema{
      "name": &schema.Schema{
        Type:     schema.TypeString,
        Optional: true,
        Computed: true,
      },
      "service": &schema.Schema{
        Type:     schema.TypeString,
        Required: true,
      },
    },
  }
}


func dataSourceInstanceNameRead(d *schema.ResourceData, m interface{}) error  {
  combinedName := d.Get("service").(string) + "-01.cantstopclapping.com"
  d.SetId(combinedName)
  d.Set("name", combinedName)
  return nil
}
