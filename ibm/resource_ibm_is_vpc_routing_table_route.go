package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const (
	rID          = "route_id"
	rDestination = "destination"
	rAction      = "action"
	rNextHop     = "next_hop"
	rName        = "name"
	rZone        = "zone"
)

func resourceIBMISVPCRoutingTableRoute() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISVPCRoutingTableRouteCreate,
		Read:     resourceIBMISVPCRoutingTableRouteRead,
		Update:   resourceIBMISVPCRoutingTableRouteUpdate,
		Delete:   resourceIBMISVPCRoutingTableRouteDelete,
		Exists:   resourceIBMISVPCRoutingTableRouteExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			rID: {
				Type:     schema.TypeString,
				Computed: true,
			},
			rName: {
				Type:     schema.TypeString,
				Required: true,
			},
			rDestination: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			rAction: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			rNextHop: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			rZone: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIBMISVPCRoutingTableRouteCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	var (
		vpcID       string
		tableID     string
		action      string
		nextHop     string
		routeName   string
		destination string
		zone        string
	)

	vpcID = d.Get(rtVpcID).(string)
	tableID = d.Get(rtID).(string)
	action = d.Get(rAction).(string)
	nextHop = d.Get(rNextHop).(string)
	routeName = d.Get(rName).(string)
	destination = d.Get(rDestination).(string)

	zone = d.Get(rZone).(string)
	z := &vpcv1.ZoneIdentityByName{
		Name: core.StringPtr(zone),
	}

	nh := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr(nextHop),
	}

	createVpcRoutingTableRouteOptions := sess.NewCreateVpcRoutingTableRouteOptions(vpcID, tableID, destination, z)
	createVpcRoutingTableRouteOptions.SetName(routeName)
	createVpcRoutingTableRouteOptions.SetZone(z)
	createVpcRoutingTableRouteOptions.SetDestination(destination)
	createVpcRoutingTableRouteOptions.SetNextHop(nh)
	createVpcRoutingTableRouteOptions.SetAction(action)

	response, _, err := sess.CreateVpcRoutingTableRoute(createVpcRoutingTableRouteOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", vpcID, tableID, *response.ID))
	d.Set(rID, *response.ID)
	return resourceIBMISVPCRoutingTableRouteRead(d, meta)
}

func resourceIBMISVPCRoutingTableRouteRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	response, _, err := sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil {
		return err
	}

	d.Set("id", *response.ID)
	d.Set(rID, *response.ID)
	d.Set(rName, *response.Name)
	d.Set(rDestination, *response.Destination)
	d.Set(rAction, *response.Action)
	//nh := response.NextHop.(map[string]interface{})
	//nh := *response.NextHop.(vpcv1.RouteNextHopPrototype)
	//d.Set(rNextHop, nh.Address)
	d.Set(rZone, *response.Zone)

	return nil
}

func resourceIBMISVPCRoutingTableRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	_, _, err = sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil {
		return err
	}

	if d.HasChange(rName) {
		updateVpcRoutingTableRouteOptions := sess.NewUpdateVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
		name := d.Get(rName).(string)
		updateVpcRoutingTableRouteOptions.SetName(name)
		_, _, err := sess.UpdateVpcRoutingTableRoute(updateVpcRoutingTableRouteOptions)
		if err != nil {
			return err
		}
	}

	return resourceIBMISVPCRoutingTableRouteRead(d, meta)
}

func resourceIBMISVPCRoutingTableRouteDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := vpcClient(meta)
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	deleteVpcRoutingTableRouteOptions := sess.NewDeleteVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	response, err := sess.DeleteVpcRoutingTableRoute(deleteVpcRoutingTableRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISVPCRoutingTableRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := vpcClient(meta)
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	_, response, err := sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
