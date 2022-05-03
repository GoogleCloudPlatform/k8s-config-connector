package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIAMBetaWorkloadIdentityPool_full(t *testing.T) {
	t.Parallel()

	randomSuffix := randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_full(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPool_update(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMBetaWorkloadIdentityPool_minimal(t *testing.T) {
	t.Parallel()

	randomSuffix := randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIAMBetaWorkloadIdentityPoolDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAMBetaWorkloadIdentityPool_minimal(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccIAMBetaWorkloadIdentityPool_update(randomSuffix),
			},
			{
				ResourceName:      "google_iam_workload_identity_pool.my_pool",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIAMBetaWorkloadIdentityPool_full(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
  display_name              = "Name of pool"
  description               = "Identity pool for automated test"
  disabled                  = true
}
`, suffix)
}

func testAccIAMBetaWorkloadIdentityPool_minimal(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
}
`, suffix)
}

func testAccIAMBetaWorkloadIdentityPool_update(suffix string) string {
	return fmt.Sprintf(`
resource "google_iam_workload_identity_pool" "my_pool" {
  workload_identity_pool_id = "my-pool-%s"
  display_name              = "Updated name of pool"
  description               = "Updated description"
  disabled                  = false
}
`, suffix)
}
