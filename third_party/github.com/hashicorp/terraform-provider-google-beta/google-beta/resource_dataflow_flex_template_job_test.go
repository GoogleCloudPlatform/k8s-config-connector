// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
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

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job", false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_streamUpdate(t *testing.T) {
	// This resource uses custom retry logic that cannot be sped up without
	// modifying the actual resource
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr
	topic2 := "tf-test-topic-2" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job", false),
				),
			},
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, topic2),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job", true),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "transform_name_mapping", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_streamFailUpdate(t *testing.T) {
	// This resource uses custom retry logic that cannot be sped up without
	// modifying the actual resource
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_basic(job, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job", false),
				),
			},
			{
				Config: testAccDataflowFlexTemplateJob_basicfail(job, bucket),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobHasOption(t, "google_dataflow_flex_template_job.flex_job", "topic", "projects/myproject/topics/tf-test-topic"+randStr, true),
				),
				ExpectError: regexp.MustCompile(`Error waiting for Job with job ID "[^"]+" to be updated: the job with ID "[^"]+" has terminated with state "JOB_STATE_FAILED" instead of expected state "JOB_STATE_RUNNING"`),
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_FullUpdate(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_dataflowFlexTemplateJobFull(job, bucket, topic),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_fullupdate",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
			{
				Config: testAccDataflowFlexTemplateJob_dataflowFlexTemplateJobFullUpdate(job, bucket, topic),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_fullupdate",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withNetwork(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	network1 := "tf-test-dataflow-net" + randStr
	network2 := "tf-test-dataflow-net2" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_network(job, network1, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_network", false),
					testAccDataflowFlexTemplateJobHasNetwork(t, "google_dataflow_flex_template_job.flex_job_network", network1, false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_network",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
			{
				Config: testAccDataflowFlexTemplateJob_networkUpdate(job, network1, network2, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_network", true),
					testAccDataflowFlexTemplateJobHasNetwork(t, "google_dataflow_flex_template_job.flex_job_network", network2, true),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_network",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withSubNetwork(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	network := "tf-test-dataflow-net" + randStr
	subnetwork1 := "tf-test-dataflow-subnetwork" + randStr
	subnetwork2 := "tf-test-dataflow-subnetwork2" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_subnetwork(job, network, subnetwork1, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_subnetwork", false),
					testAccDataflowFlexTemplateJobHasSubNetwork(t, "google_dataflow_flex_template_job.flex_job_subnetwork", subnetwork1, false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_subnetwork",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
			{
				Config: testAccDataflowFlexTemplateJob_subnetworkUpdate(job, network, subnetwork1, subnetwork2, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_subnetwork", true),
					testAccDataflowFlexTemplateJobHasSubNetwork(t, "google_dataflow_flex_template_job.flex_job_subnetwork", subnetwork2, true),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_subnetwork",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withIpConfig(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr
	network := "tf-test-dataflow-net" + randStr
	subnetwork := "tf-test-dataflow-subnetwork" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_ipConfig(job, network, subnetwork, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_ipconfig", false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_ipconfig",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "ip_configuration", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withKmsKey(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	key_ring := "tf-test-dataflow-kms-ring-" + randStr
	crypto_key := "tf-test-dataflow-kms-key-" + randStr
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_kms(job, key_ring, crypto_key, bucket, topic),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_kms", false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_kms",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "container_spec_gcs_path"},
			},
		},
	})
}

func TestAccDataflowFlexTemplateJob_withAdditionalExperiments(t *testing.T) {
	// Dataflow responses include serialized java classes and bash commands
	// This makes body comparison infeasible
	acctest.SkipIfVcr(t)
	t.Parallel()

	randStr := acctest.RandString(t, 10)
	job := "tf-test-dataflow-job-" + randStr
	additionalExperiments := []string{"enable_stackdriver_agent_metrics", "use_runner_v2"}
	bucket := "tf-test-dataflow-bucket-" + randStr
	topic := "tf-test-topic" + randStr

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataflowJobDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataflowFlexTemplateJob_additionalExperiments(job, bucket, topic, additionalExperiments),
				Check: resource.ComposeTestCheckFunc(
					testAccDataflowFlexJobExists(t, "google_dataflow_flex_template_job.flex_job_experiments", false),
					testAccDataflowFlexTemplateJobHasAdditionalExperiments(t, "google_dataflow_flex_template_job.flex_job_experiments", additionalExperiments, false),
				),
			},
			{
				ResourceName:            "google_dataflow_flex_template_job.flex_job_experiments",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"on_delete", "parameters", "skip_wait_on_job_termination", "state", "additional_experiments", "container_spec_gcs_path"},
			},
		},
	})
}

func testAccDataflowFlexTemplateJobHasNetwork(t *testing.T, res, expected string, wait bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		instanceTmpl, err := testAccDataflowFlexTemplateGetGeneratedInstanceTemplate(t, s, res)
		if err != nil {
			return fmt.Errorf("Error getting dataflow job instance template: %s", err)
		}
		if len(instanceTmpl.Properties.NetworkInterfaces) == 0 {
			return fmt.Errorf("no network interfaces in template properties: %+v", instanceTmpl.Properties)
		}
		actual := instanceTmpl.Properties.NetworkInterfaces[0].Network
		if GetResourceNameFromSelfLink(actual) != GetResourceNameFromSelfLink(expected) {
			return fmt.Errorf("network mismatch: %s != %s", actual, expected)
		}
		return nil
	}
}

func testAccDataflowFlexTemplateJobHasSubNetwork(t *testing.T, res, expected string, wait bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		instanceTmpl, err := testAccDataflowFlexTemplateGetGeneratedInstanceTemplate(t, s, res)
		if err != nil {
			return fmt.Errorf("Error getting dataflow job instance template: %s", err)
		}
		if len(instanceTmpl.Properties.NetworkInterfaces) == 0 {
			return fmt.Errorf("no network interfaces in template properties: %+v", instanceTmpl.Properties)
		}
		actual := instanceTmpl.Properties.NetworkInterfaces[0].Subnetwork
		if GetResourceNameFromSelfLink(actual) != GetResourceNameFromSelfLink(expected) {
			return fmt.Errorf("subnetwork mismatch: %s != %s", actual, expected)
		}
		return nil
	}
}

func testAccDataflowFlexTemplateJobHasAdditionalExperiments(t *testing.T, res string, experiments []string, wait bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[res]
		if !ok {
			return fmt.Errorf("resource %q not found in state", res)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		config := acctest.GoogleProviderConfig(t)

		job, err := config.NewDataflowClient(config.UserAgent).Projects.Jobs.Get(config.Project, rs.Primary.ID).View("JOB_VIEW_ALL").Do()
		if err != nil {
			return fmt.Errorf("dataflow job does not exist")
		}

		for _, expectedExperiment := range experiments {
			var contains = false
			for _, actualExperiment := range job.Environment.Experiments {
				if actualExperiment == expectedExperiment {
					contains = true
				}
			}
			if contains != true {
				return fmt.Errorf("Expected experiment '%s' not found in experiments", expectedExperiment)
			}
		}

		return nil
	}
}

func testAccDataflowFlexTemplateGetGeneratedInstanceTemplate(t *testing.T, s *terraform.State, res string) (*compute.InstanceTemplate, error) {
	rs, ok := s.RootModule().Resources[res]
	if !ok {
		return nil, fmt.Errorf("resource %q not in state", res)
	}
	if rs.Primary.ID == "" {
		return nil, fmt.Errorf("resource %q does not have an ID set", res)
	}
	filter := fmt.Sprintf("properties.labels.dataflow_job_id = %s", rs.Primary.ID)

	config := acctest.GoogleProviderConfig(t)

	var instanceTemplate *compute.InstanceTemplate

	err := resource.Retry(1*time.Minute, func() *resource.RetryError {
		instanceTemplates, rerr := config.NewComputeClient(config.UserAgent).InstanceTemplates.
			List(config.Project).
			Filter(filter).
			MaxResults(2).
			Fields("items/properties").Do()
		if rerr != nil {
			return resource.NonRetryableError(rerr)
		}
		if len(instanceTemplates.Items) == 0 {
			return resource.RetryableError(fmt.Errorf("no instance template found for dataflow job %q", rs.Primary.ID))
		}
		if len(instanceTemplates.Items) > 1 {
			return resource.NonRetryableError(fmt.Errorf("Wrong number of matching instance templates for dataflow job: %s, %d", rs.Primary.ID, len(instanceTemplates.Items)))
		}
		instanceTemplate = instanceTemplates.Items[0]
		if instanceTemplate == nil || instanceTemplate.Properties == nil {
			return resource.NonRetryableError(fmt.Errorf("invalid instance template has no properties"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return instanceTemplate, nil
}

func testAccDataflowFlexTemplateJob_basic(job, bucket, topicName string) string {
	return fmt.Sprintf(`

resource "google_pubsub_topic" "example" {
  name = "%s"
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_dataflow_flex_template_job" "flex_job" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
}
`, topicName, bucket, job)
}

func testAccDataflowFlexTemplateJob_basicfail(job, bucket string) string {
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_dataflow_flex_template_job" "flex_job" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    sinkType= "BIGQUERY"
	outputTableSpec = "projectid:datasetid.tableid"
  }
  labels = {
   "my_labels" = "value"
  }
}
`, bucket, job)
}

func testAccDataflowFlexTemplateJob_dataflowFlexTemplateJobFull(job, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_service_account" "dataflow-sa" {
  count = 2
  account_id   = "dataflow-sa-${count.index}"
  display_name = "DataFlow Service Account"
}

resource "google_project_iam_member" "dataflow-worker" {
  count = 2
  project = data.google_project.project.project_id
  role   = "roles/dataflow.worker"
  member = "serviceAccount:${google_service_account.dataflow-sa[count.index].email}"
}

resource "google_project_iam_member" "dataflow-storage" {
  count = 2
  project = data.google_project.project.project_id
  role   = "roles/storage.admin"
  member = "serviceAccount:${google_service_account.dataflow-sa[count.index].email}"
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_dataflow_flex_template_job" "flex_job_fullupdate" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value-1"
  }
  service_account_email = google_service_account.dataflow-sa[0].email
  machine_type = "n1-standard-2"
}
`, topicName, bucket, job)
}

func testAccDataflowFlexTemplateJob_dataflowFlexTemplateJobFullUpdate(job, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_service_account" "dataflow-sa" {
	count = 2
	account_id   = "dataflow-sa-${count.index}"
	display_name = "DataFlow Service Account"
}

  resource "google_project_iam_member" "dataflow-worker" {
	count = 2
	project = data.google_project.project.project_id
	role   = "roles/dataflow.worker"
	member = "serviceAccount:${google_service_account.dataflow-sa[count.index].email}"
}

  resource "google_project_iam_member" "dataflow-storage" {
	count = 2
	project = data.google_project.project.project_id
	role   = "roles/storage.admin"
	member = "serviceAccount:${google_service_account.dataflow-sa[count.index].email}"
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}
resource "google_dataflow_flex_template_job" "flex_job_fullupdate" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value-update"
  }
  service_account_email = google_service_account.dataflow-sa[1].email
  machine_type = "n2-standard-2"
}
`, topicName, bucket, job)
}

func testAccDataflowFlexTemplateJob_network(job, network1, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_compute_network" "net1" {
  name                    = "%s"
  auto_create_subnetworks = true
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_dataflow_flex_template_job" "flex_job_network" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  network           = google_compute_network.net1.name

}
`, topicName, network1, bucket, job)
}

func testAccDataflowFlexTemplateJob_networkUpdate(job, network1, network2, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_compute_network" "net1" {
  name                    = "%s"
  auto_create_subnetworks = true
}

resource "google_compute_network" "net2" {
	name                    = "%s"
	auto_create_subnetworks = true
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_dataflow_flex_template_job" "flex_job_network" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  network           = google_compute_network.net2.name

}
`, topicName, network1, network2, bucket, job)
}

func testAccDataflowFlexTemplateJob_subnetwork(job, network, subnetwork1, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_compute_network" "net" {
  name                    = "%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet1" {
  name          = "%s"
  ip_cidr_range = "10.1.0.0/24"
  network       = google_compute_network.net.self_link
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}
resource "google_dataflow_flex_template_job" "flex_job_subnetwork" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  subnetwork        = google_compute_subnetwork.subnet1.self_link

}
`, topicName, network, subnetwork1, bucket, job)
}

func testAccDataflowFlexTemplateJob_subnetworkUpdate(job, network, subnetwork1, subnetwork2, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_compute_network" "net" {
  name                    = "%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet1" {
  name          = "%s"
  ip_cidr_range = "10.1.0.0/24"
  network       = google_compute_network.net.self_link
}

resource "google_compute_subnetwork" "subnet2" {
  name          = "%s"
  ip_cidr_range = "10.2.0.0/24"
  network       = google_compute_network.net.self_link
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}
resource "google_dataflow_flex_template_job" "flex_job_subnetwork" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  subnetwork        = google_compute_subnetwork.subnet2.self_link

}
`, topicName, network, subnetwork1, subnetwork2, bucket, job)
}

func testAccDataflowFlexTemplateJob_ipConfig(job, network, subnetwork, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
  name = "%s"
}

resource "google_compute_network" "net" {
  name          = "%s"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "subnet" {
  name          = "%s"
  ip_cidr_range = "10.1.0.0/24"
  network       = google_compute_network.net.self_link
  private_ip_google_access = true
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}
resource "google_dataflow_flex_template_job" "flex_job_ipconfig" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  ip_configuration = "WORKER_IP_PRIVATE"
  subnetwork        = google_compute_subnetwork.subnet.self_link

}
`, topicName, network, subnetwork, bucket, job)
}

func testAccDataflowFlexTemplateJob_kms(job, key_ring, crypto_key, bucket, topicName string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
	name = "%s"
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

resource "google_project_iam_member" "kms-project-dataflow-binding" {
  project = data.google_project.project.project_id
  role    = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member  = "serviceAccount:service-${data.google_project.project.number}@dataflow-service-producer-prod.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "kms-project-compute-binding" {
  project = data.google_project.project.project_id
  role    = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member  = "serviceAccount:service-${data.google_project.project.number}@compute-system.iam.gserviceaccount.com"
}

resource "google_kms_key_ring" "keyring" {
  name     = "%s"
  location = "global"
}

resource "google_kms_crypto_key" "crypto_key" {
  name            = "%s"
  key_ring        = google_kms_key_ring.keyring.id
  rotation_period = "100000s"
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}
resource "google_dataflow_flex_template_job" "flex_job_kms" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  kms_key_name		= google_kms_crypto_key.crypto_key.id

}
`, topicName, bucket, key_ring, crypto_key, job)
}

func testAccDataflowFlexTemplateJob_additionalExperiments(job, bucket, topicName string, experiments []string) string {
	return fmt.Sprintf(`
data "google_project" "project" {}

resource "google_pubsub_topic" "example" {
	name = "%s"
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
  content = <<EOF
{
	"eventId": "{{uuid()}}",
	"eventTimestamp": {{timestamp()}},
	"ipv4": "{{ipv4()}}",
	"ipv6": "{{ipv6()}}",
	"country": "{{country()}}",
	"username": "{{username()}}",
	"quest": "{{random("A Break In the Ice", "Ghosts of Perdition", "Survive the Low Road")}}",
	"score": {{integer(100, 10000)}},
	"completed": {{bool()}}
}
EOF
}

data "google_storage_bucket_object" "flex_template" {
  name   = "latest/flex/Streaming_Data_Generator"
  bucket = "dataflow-templates"
}
resource "google_dataflow_flex_template_job" "flex_job_experiments" {
  name = "%s"
  container_spec_gcs_path = "gs://${data.google_storage_bucket_object.flex_template.bucket}/${data.google_storage_bucket_object.flex_template.name}"
  on_delete = "cancel"
  parameters = {
    schemaLocation = "gs://${google_storage_bucket_object.schema.bucket}/schema.json"
    qps = "1"
    topic = google_pubsub_topic.example.id
  }
  labels = {
   "my_labels" = "value"
  }
  additional_experiments = ["%s"]

}
`, topicName, bucket, job, strings.Join(experiments, `", "`))
}

func testAccDataflowFlexJobHasOption(t *testing.T, res, option, expectedValue string, wait bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if wait {
			time.Sleep(300 * time.Second)
		}
		rs, ok := s.RootModule().Resources[res]
		if !ok {
			return fmt.Errorf("resource %q not found in state", res)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		config := acctest.GoogleProviderConfig(t)

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

func testAccDataflowFlexJobExists(t *testing.T, resource string, wait bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if wait {
			time.Sleep(300 * time.Second)
		}
		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("resource %q not in state", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := acctest.GoogleProviderConfig(t)
		_, err := config.NewDataflowClient(config.UserAgent).Projects.Locations.Jobs.Get(config.Project, config.Region, rs.Primary.ID).Do()
		if err != nil {
			return fmt.Errorf("could not confirm Dataflow Job %q exists: %v", rs.Primary.ID, err)
		}

		return nil
	}
}
