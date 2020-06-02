package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	rtID    = "route_table_id"
	rtVpcID = "vpc_id"
	rtName  = "name"
)

func resourceIBMISVPCRoutingTables() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPCRoutingTablesCreate,
		Read:     resourceIBMISVPCRoutingTablesRead,
		Update:   resourceIBMISVPCRoutingTablesUpdate,
		Delete:   resourceIBMISVPCRoutingTablesDelete,
		Exists:   resourceIBMISVPCRoutingTablesExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			rtID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			rtVpcID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			rtName: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIBMISVPCRoutingTablesCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	var (
		vpcID  string
		rtName string
	)
	vpcID = d.Get(rtVpcID).(string)
	rtName = d.Get(rtName).(string)

	createVpcRoutingTableOptions := sess.NewCreateVpcRoutingTableOptions(vpcID)
	createVpcRoutingTableOptions.SetName(rtName)
	response, _, err := sess.CreateVpcRoutingTable(createVpcRoutingTableOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", vpcID, *response.ID))

	return resourceIBMISVPCRoutingTablesRead(d, meta)
}

func resourceIBMISVPCRoutingTablesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(id_set[0], id_set[1])
	response, _, err := sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
	if err != nil {
		return err
	}

	d.Set("id", response.ID) // Check what to set in id & rtID
	d.Set(rtID, response.ID)
	d.Set(rtName, response.Name)

	return nil
}

func resourceIBMISVPCRoutingTablesUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(id_set[0], id_set[1])
	_, _, err = sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
	if err != nil {
		return err
	}

	if d.HasChange(rtName) {
		updateVpcRoutingTableOptions := sess.NewUpdateVpcRoutingTableOptions(id_set[0], id_set[1])
		name := d.Get(rtName).(string)
		updateVpcRoutingTableOptions.SetName(name)
		_, _, err := sess.UpdateVpcRoutingTable(updateVpcRoutingTableOptions)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVPCRoutingTablesRead(d, meta)
}

func resourceIBMISVPCRoutingTablesDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")

	deleteZoneOptions := sess.NewDeleteVpcRoutingTableOptions(id_set[0], id_set[1])
	response, err := sess.DeleteVpcRoutingTable(deleteZoneOptions)
	if err != nil && response.StatusCode != 404 {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISVPCRoutingTablesExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(id_set[0], id_set[1])
	_, response, err := sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
