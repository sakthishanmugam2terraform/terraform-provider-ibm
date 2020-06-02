package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVPCRoutingTables_Basic(t *testing.T) {
	var resultroutetable string
	name := fmt.Sprintf("testroutetable-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRoutingTablesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPCRoutingTablesBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRoutingTablesExists("ibm_is_vpc_routing_tables.test_vpc_route_table", resultroutetable),
					resource.TestCheckResourceAttr("ibm_is_vpc_routing_tables.test_vpc_route_table", "name", name),
				),
			},
		},
	})
}

func TestAccIBMISVPCRoutingTablesImport(t *testing.T) {
	var resultroutetable string
	name := fmt.Sprintf("testroutetable-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPCRoutingTablesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPCRoutingTablesBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPCRoutingTablesExists("ibm_is_vpc_routing_tables.test_vpc_route_table", resultroutetable),
					resource.TestCheckResourceAttr("ibm_is_vpc_routing_tables.test_vpc_route_table", "name", name),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_is_vpc_routing_tables.test_vpc_route_table",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMISVPCRoutingTablesBasic(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		name = "default"
	}

	resource "ibm_is_vpc" "test_routing_table_vpc" {
		name = %s
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_is_vpc_routing_tables" "test_vpc_route_table" {
        name = "test-vpc-route-table"
        vpc_id = ibm_is_vpc.test_routing_table_vpc.id
    }
	`, name)
}

func testAccCheckIBMISVPCRoutingTablesDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpc_routing_tables" {
			continue
		}

		sess, err := testAccProvider.Meta().(ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts := rs.Primary.ID
		partslist := strings.Split(parts, "/")
		getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(partslist[0], partslist[1])
		_, _, err = sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMISVPCRoutingTablesDestroy: Error checking while route table (%s) has been destroyed?: %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccCheckIBMISVPCRoutingTablesExists(n string, result string) resource.TestCheckFunc {

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
		getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(partslist[0], partslist[1])
		r, _, err := sess.GetVpcRoutingTable(getVpcRoutingTableOptions)

		if err != nil {
			return fmt.Errorf("testAccCheckIBMPrivateDNSZoneExists: Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		result = *r.ID
		return nil
	}
}
