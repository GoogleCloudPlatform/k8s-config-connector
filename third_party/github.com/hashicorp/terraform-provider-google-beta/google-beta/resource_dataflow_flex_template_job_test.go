package google

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"

	compute "google.golang.org/api/compute/v0.beta"
)

func TestAccDataflowFlexTemplateJob_basic(t *testing.T) {
	// This resource uses custom retry logic that cannot be sped up without
	// modifying the actual resource
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, "mytopic"),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobExists(t, "google_dataflow_flex_template_job.job"),
				),
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_streamUpdate(t *testing.T) {
	// This resource uses custom retry logic that cannot be sped up without
	// modifying the actual resource
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, "mytopic"),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobExists(t, "google_dataflow_flex_template_job.job"),
				),
			},
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, "mytopic2"),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobHasOption(t, "google_dataflow_flex_template_job.job", "topic", "projects/myproject/topics/mytopic2"),
				),
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_streamUpdateFail(t *testing.T) {
	// This resource uses custom retry logic that cannot be sped up without
	// modifying the actual resource
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, "mytopic"),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobExists(t, "google_dataflow_flex_template_job.job"),
				),
			},
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobHasOption(t, "google_dataflow_flex_template_job.job", "topic", "projects/myproject/topics/mytopic"),
				),
				ExpectError: regexp.MustCompile(`Error waiting for Job with job ID "[^"]+" to be updated: the job with ID "[^"]+" has terminated with state "JOB_STATE_FAILED" instead of expected state "JOB_STATE_RUNNING"`),
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withServiceAccount(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	accountId := "tf-test-dataflow-sa" + randStr
	zone := "us-central1-b"

	VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_serviceAccount(job, bucket, accountId, zone),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowJobExists(t, "google_dataflow_flex_template_job.job"),
					testAccDataflowFlexTemplateJobHasServiceAccount(t, "google_dataflow_flex_template_job.job", accountId, zone),
				),
			},
		},
	})
}

func testAccDataflowFlexTemplateJobHasServiceAccount(t *testing.T, res, expectedId, zone string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		instance, err := testAccDataflowFlexTemplateJobGetGeneratedInstance(t, s, res, zone)
		if err != nil {
			return fmt.Errorf("Error getting dataflow job instance: %s", err)
		}
		accounts := instance.ServiceAccounts
		if len(accounts) != 1 {
			return fmt.Errorf("Found multiple service accounts (%d) for dataflow job %q, expected 1", len(accounts), res)
		}
		actualId := strings.Split(accounts[0].Email, "@")[0]
		if expectedId != actualId {
			return fmt.Errorf("service account mismatch, expected account ID = %q, actual email = %q", expectedId, accounts[0].Email)
		}
		return nil
	}
}

func testAccDataflowFlexTemplateJobGetGeneratedInstance(t *testing.T, s *terraform.State, res, zone string) (*compute.Instance, error) {
	rs, ok := s.RootModule().Resources[res]
	if !ok {
		return nil, fmt.Errorf("resource %q not in state", res)
	}
	if rs.Primary.ID == "" {
		return nil, fmt.Errorf("resource %q does not have an ID set", res)
	}
	filter := fmt.Sprintf("labels.goog-dataflow-job-id = %s", rs.Primary.ID)

	config := GoogleProviderConfig(t)

	var instance *compute.Instance

	err := resource.Retry(1*time.Minute, func() *resource.RetryError {
		instances, rerr := config.NewComputeClient(config.UserAgent).Instances.
			List(config.Project, zone).
			Filter(filter).
			MaxResults(2).
			Do()
		if rerr != nil {
			return resource.NonRetryableError(rerr)
		}
		if len(instances.Items) == 0 {
			return resource.RetryableError(fmt.Errorf("no instance found for dataflow job %q", rs.Primary.ID))
		}
		if len(instances.Items) > 1 {
			return resource.NonRetryableError(fmt.Errorf("Wrong number of matching instances for dataflow job: %s, %d", rs.Primary.ID, len(instances.Items)))
		}
		instance = instances.Items[0]
		if instance == nil {
			return resource.NonRetryableError(fmt.Errorf("invalid instance"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// note: this config creates a job that doesn't actually do anything, but still runs
func testAccDataflowFlexTemplateJob_basic(job, bucket, topicName string) string {
	topicField := ""
	if topicName != "" {
		topicField = fmt.Sprintf("topic = \"projects/myproject/topics/%s\"", topicName)
	}
	return fmt.Sprintf(`
data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}

resource "google_storage_bucket" "bucket" {
  name = "%s"
  location = "US-CENTRAL1"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "schema" {
  name = "schema.json"
  bucket = google_storage_bucket.bucket.name
  content = "{}"
}

resource "google_dataflow_flex_template_job" "job" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    %s
  }
  labels = {
   "my_labels" = "value"
  }
}
`, bucket, job, topicField)
}

// note: this config creates a job that doesn't actually do anything, but still runs
func testAccDataflowFlexTemplateJob_serviceAccount(job, bucket, accountId, zone string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_service_account" "dataflow-sa" {
  account_id   = "%s"
  display_name = "DataFlow Service Account"
}

resource "google_project_iam_member" "dataflow-worker" {
  project = data.google_project.project.project_id
  role   = "roles/dataflow.worker"
  member = "serviceAccount:${google_service_account.dataflow-sa.email}"
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}

resource "google_storage_bucket" "bucket" {
  name = "%s"
  location = "US-CENTRAL1"
  force_destroy = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "schema" {
  name = "schema.json"
  bucket = google_storage_bucket.bucket.name
  content = "{}"
}

resource "google_dataflow_flex_template_job" "job" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = "projects/myproject/topics/mytopic"
    serviceAccount = google_service_account.dataflow-sa.email
    zone = "%s"
  }
  labels = {
   "my_labels" = "value"
  }
}
`, accountId, bucket, job, zone)
}

func testAccDataflowJobHasOption(t *testing.T, res, option, expectedValue string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[res]
		if !ok {
			return fmt.Errorf("resource %q not found in state", res)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		config := GoogleProviderConfig(t)

		job, err := config.NewDataflowClient(config.UserAgent).Projects.Jobs.Get(config.Project, rs.Primary.ID).View("JOB_VIEW_ALL").Do()
		if err != nil {
			return fmt.Errorf("dataflow job does not exist")
		}

		sdkPipelineOptions, err := tpgresource.ConvertToMap(job.Environment.SdkPipelineOptions)
		if err != nil {
			return fmt.Errorf("error from ConvertToMap: %s", err)
		}
		optionsMap := sdkPipelineOptions["options"].(map[string]interface{})

		if optionsMap[option] != expectedValue {
			return fmt.Errorf("Option %s do not match. Got %s while expecting %s", option, optionsMap[option], expectedValue)
		}

		return nil
	}
}
