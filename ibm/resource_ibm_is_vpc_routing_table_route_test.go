package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVPCRoutingTableRoute_Basic(t *testing.T) {
	var resultrtroute string
	name := fmt.Sprintf("test-rt-route-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRoutingTableRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPCRoutingTableRouteBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRoutingTableRouteExists("ibm_is_vpc_routing_table_route.test_vpc_route_table_route", resultrtroute),
					resource.TestCheckResourceAttr("ibm_is_vpc_routing_table_route.test_vpc_route_table_route", "name", name),
				),
			},
		},
	})
}

func TestAccIBMISVPCRoutingTableRouteImport(t *testing.T) {
	var resultrtroute string
	name := fmt.Sprintf("test-rt-route-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRoutingTableRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPCRoutingTableRouteBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRoutingTableRouteExists("ibm_is_vpc_routing_table_route.test_vpc_route_table_route", resultrtroute),
					resource.TestCheckResourceAttr("ibm_is_vpc_routing_table_route.test_vpc_route_table_route", "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_vpc_routing_table_route.test_vpc_route_table_route",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMISVPCRoutingTableRouteBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
	}

	resource "ibm_is_vpc" "test_routing_table_vpc" {
		name = "test"-routing-table-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_is_vpc_routing_tables" "test_vpc_route_table" {
        name = "test-vpc-route-table"
        vpc_id = ibm_is_vpc.test_routing_table_vpc.id
	}
	
	resource "ibm_is_subnet" "test_subnet_route_table" {
		name = "test-subnet-route-table"
		vpc = ibm_is_vpc.test_routing_table_vpc.id
		zone = "us-south-1"
		ipv4_cidr_block = "10.0.0.0/24"
		routing_table_id = ibm_is_vpc_routing_tables.test_vpc_route_table.route_table_id
	}
	
	resource "ibm_is_vpc_routing_table_route" "test_vpc_route_table_route" {
		vpc_id = ibm_is_vpc.test_routing_table_vpc.id
		route_table_id = ibm_is_vpc_routing_tables.test_vpc_route_table.route_table_id
		zone = "us-south-1"
		route_name = %s
		next_hop = "10.0.0.2"
		action = "deliver"
		destination = "0.0.0.0/0"
	}
	`, name)
}

func testAccCheckIBMISVPCRoutingTableRouteDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_table_route" {
			continue
		}

		sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")
		getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(partslist[0], partslist[1], partslist[2])
		_, _, err = sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
		if err != nil {
			return err
		}
	}
	return nil
}

func testAccCheckIBMISVPCRoutingTableRouteExists(n string, result string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")
		getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(partslist[0], partslist[1], partslist[2])
		r, _, err := sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
		if err != nil {
			return err
		}

		result = *r.ID
		return nil
	}
}
