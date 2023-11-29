// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package dataflow

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	"google.golang.org/api/googleapi"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	dataflow "google.golang.org/api/dataflow/v1b3"
)

// NOTE: resource_dataflow_flex_template currently does not support updating existing jobs.
// Changing any non-computed field will result in the job being deleted (according to its
// on_delete policy) and recreated with the updated parameters.

// ResourceDataflowFlexTemplateJob defines the schema for Dataflow FlexTemplate jobs.
func ResourceDataflowFlexTemplateJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataflowFlexTemplateJobCreate,
		Read:   resourceDataflowFlexTemplateJobRead,
		Update: resourceDataflowFlexTemplateJobUpdate,
		Delete: resourceDataflowFlexTemplateJobDelete,
		CustomizeDiff: customdiff.All(
			resourceDataflowFlexJobTypeCustomizeDiff,
		),
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{

			"container_spec_gcs_path": {
				Type:     schema.TypeString,
				Required: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: `The region in which the created job should run.`,
			},

			"transform_name_mapping": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.`,
			},

			"on_delete": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"cancel", "drain"}, false),
				Optional:     true,
				Default:      "cancel",
			},

			"labels": {
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: resourceDataflowJobLabelDiffSuppress,
				Description:      `User labels to be specified for the job. Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: Google-provided Dataflow templates often provide default labels that begin with goog-dataflow-provided. Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.`,
			},

			"parameters": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of this job, selected from the JobType enum.`,
			},

			"num_workers": {
				Type:     schema.TypeInt,
				Optional: true,
				// ForceNew applies to both stream and batch jobs
				ForceNew:    true,
				Description: `The initial number of Google Compute Engine instances for the job.`,
			},

			"max_workers": {
				Type:     schema.TypeInt,
				Optional: true,
				// ForceNew applies to both stream and batch jobs
				ForceNew:    true,
				Description: `The maximum number of Google Compute Engine instances to be made available to your pipeline during execution, from 1 to 1000.`,
			},

			"service_account_email": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Service Account email used to create the job.`,
			},

			"temp_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Cloud Storage path to use for temporary files. Must be a valid Cloud Storage URL, beginning with gs://.`,
			},

			"staging_location": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: `The Cloud Storage path to use for staging files. Must be a valid Cloud Storage URL, beginning with gs://.`,
			},

			"sdk_container_image": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Docker registry location of container image to use for the 'worker harness. Default is the container for the version of the SDK. Note this field is only valid for portable pipelines.`,
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The network to which VMs will be assigned. If it is not provided, "default" will be used.`,
			},

			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".`,
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The machine type to use for the job.`,
			},

			"kms_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`,
			},

			"ip_configuration": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  `The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".`,
				ValidateFunc: validation.StringInSlice([]string{"WORKER_IP_PUBLIC", "WORKER_IP_PRIVATE"}, false),
			},

			"additional_experiments": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: `List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"enable_streaming_engine": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `Indicates if the job should use the streaming engine feature.`,
			},

			"autoscaling_algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The algorithm to use for autoscaling`,
			},

			"launcher_machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The machine type to use for launching the job. The default is n1-standard-1.`,
			},

			"skip_wait_on_job_termination": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: `If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on. WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.`,
			},
		},
		UseJSONNumber: true,
	}
}

// resourceDataflowFlexTemplateJobCreate creates a Flex Template Job from TF code.
func resourceDataflowFlexTemplateJobCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}

	env, err := resourceDataflowFlexJobSetupEnv(d, config)
	if err != nil {
		return err
	}

	request := dataflow.LaunchFlexTemplateRequest{
		LaunchParameter: &dataflow.LaunchFlexTemplateParameter{
			ContainerSpecGcsPath: d.Get("container_spec_gcs_path").(string),
			JobName:              d.Get("name").(string),
			Parameters:           tpgresource.ExpandStringMap(d, "parameters"),
			Environment:          &env,
		},
	}

	response, err := config.NewDataflowClient(userAgent).Projects.Locations.FlexTemplates.Launch(project, region, &request).Do()
	if err != nil {
		return err
	}

	job := response.Job

	//adding wait time for setting all the parameters into state file
	err = waitForDataflowJobState(d, config, job.Id, userAgent, d.Timeout(schema.TimeoutUpdate), "JOB_STATE_RUNNING")
	if err != nil {
		return fmt.Errorf("Error waiting for job with job ID %q to be running: %s", job.Id, err)
	}

	d.SetId(job.Id)
	if err := d.Set("job_id", job.Id); err != nil {
		return fmt.Errorf("Error setting job_id: %s", err)
	}

	return resourceDataflowFlexTemplateJobRead(d, meta)
}

func resourceDataflowFlexJobSetupEnv(d *schema.ResourceData, config *transport_tpg.Config) (dataflow.FlexTemplateRuntimeEnvironment, error) {

	additionalExperiments := tpgresource.ConvertStringSet(d.Get("additional_experiments").(*schema.Set))

	env := dataflow.FlexTemplateRuntimeEnvironment{
		AdditionalUserLabels:  tpgresource.ExpandStringMap(d, "labels"),
		AutoscalingAlgorithm:  d.Get("autoscaling_algorithm").(string),
		NumWorkers:            int64(d.Get("num_workers").(int)),
		MaxWorkers:            int64(d.Get("max_workers").(int)),
		Network:               d.Get("network").(string),
		ServiceAccountEmail:   d.Get("service_account_email").(string),
		Subnetwork:            d.Get("subnetwork").(string),
		TempLocation:          d.Get("temp_location").(string),
		StagingLocation:       d.Get("staging_location").(string),
		MachineType:           d.Get("machine_type").(string),
		KmsKeyName:            d.Get("kms_key_name").(string),
		IpConfiguration:       d.Get("ip_configuration").(string),
		EnableStreamingEngine: d.Get("enable_streaming_engine").(bool),
		AdditionalExperiments: additionalExperiments,
		SdkContainerImage:     d.Get("sdk_container_image").(string),
		LauncherMachineType:   d.Get("launcher_machine_type").(string),
	}
	return env, nil
}

// resourceDataflowFlexTemplateJobRead reads a Flex Template Job resource.
func resourceDataflowFlexTemplateJobRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}

	jobId := d.Id()

	job, err := resourceDataflowJobGetJob(config, project, region, userAgent, jobId)
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("Dataflow job %s", jobId))
	}

	if err := d.Set("job_id", job.Id); err != nil {
		return fmt.Errorf("Error setting job_id: %s", err)
	}
	if err := d.Set("state", job.CurrentState); err != nil {
		return fmt.Errorf("Error setting state: %s", err)
	}
	if err := d.Set("name", job.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err := d.Set("type", job.Type); err != nil {
		return fmt.Errorf("Error setting type: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("labels", job.Labels); err != nil {
		return fmt.Errorf("Error setting labels: %s", err)
	}
	if err := d.Set("kms_key_name", job.Environment.ServiceKmsKeyName); err != nil {
		return fmt.Errorf("Error setting kms_key_name: %s", err)
	}
	if err := d.Set("service_account_email", job.Environment.ServiceAccountEmail); err != nil {
		return fmt.Errorf("Error setting service_account_email: %s", err)
	}

	sdkPipelineOptions, err := tpgresource.ConvertToMap(job.Environment.SdkPipelineOptions)
	if err != nil {
		return err
	}
	optionsMap := sdkPipelineOptions["options"].(map[string]interface{})

	if err := d.Set("temp_location", optionsMap["tempLocation"]); err != nil {
		return fmt.Errorf("Error setting temp_gcs_location: %s", err)
	}
	if err := d.Set("network", optionsMap["network"]); err != nil {
		return fmt.Errorf("Error setting network: %s", err)
	}
	if err := d.Set("num_workers", optionsMap["numWorkers"]); err != nil {
		return fmt.Errorf("Error setting num_workers: %s", err)
	}
	if err := d.Set("max_workers", optionsMap["maxWorkers"]); err != nil {
		return fmt.Errorf("Error setting max_workers: %s", err)
	}
	if err := d.Set("staging_location", optionsMap["stagingLocation"]); err != nil {
		return fmt.Errorf("Error setting staging_location: %s", err)
	}
	if err := d.Set("sdk_container_image", optionsMap["sdkContainerImage"]); err != nil {
		return fmt.Errorf("Error setting sdk_container_image: %s", err)
	}
	if err := d.Set("network", optionsMap["network"]); err != nil {
		return fmt.Errorf("Error setting network: %s", err)
	}
	if err := d.Set("subnetwork", optionsMap["subnetwork"]); err != nil {
		return fmt.Errorf("Error setting subnetwork: %s", err)
	}
	if err := d.Set("machine_type", optionsMap["workerMachineType"]); err != nil {
		return fmt.Errorf("Error setting machine_type: %s", err)
	}

	if ok := shouldStopDataflowJobDeleteQuery(job.CurrentState, d.Get("skip_wait_on_job_termination").(bool)); ok {
		log.Printf("[DEBUG] Removing resource '%s' because it is in state %s.\n", job.Name, job.CurrentState)
		d.SetId("")
		return nil
	}

	return nil
}

func waitForDataflowJobState(d *schema.ResourceData, config *transport_tpg.Config, jobID, userAgent string, timeout time.Duration, targetState string) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		project, err := tpgresource.GetProject(d, config)
		if err != nil {
			return resource.NonRetryableError(err)
		}

		region, err := tpgresource.GetRegion(d, config)
		if err != nil {
			return resource.NonRetryableError(err)
		}

		job, err := resourceDataflowJobGetJob(config, project, region, userAgent, jobID)
		if err != nil {
			if transport_tpg.IsRetryableError(err, nil, nil) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		state := job.CurrentState
		if state == targetState {
			log.Printf("[DEBUG] the job with ID %q has state %q.", jobID, state)
			return nil
		}
		_, terminating := DataflowTerminatingStatesMap[state]
		if terminating && targetState == "JOB_STATE_RUNNING" {
			return resource.NonRetryableError(fmt.Errorf("the job with ID %q is terminating with state %q and cannot reach expected state %q", jobID, state, targetState))
		}
		if _, terminated := DataflowTerminalStatesMap[state]; terminated {
			return resource.NonRetryableError(fmt.Errorf("the job with ID %q has terminated with state %q instead of expected state %q", jobID, state, targetState))
		} else {
			log.Printf("[DEBUG] the job with ID %q has state %q.", jobID, state)
			return resource.RetryableError(fmt.Errorf("the job with ID %q has state %q, waiting for %q", jobID, state, targetState))
		}
	})
}

// resourceDataflowFlexTemplateJobUpdate updates a Flex Template Job resource.
func resourceDataflowFlexTemplateJobUpdate(d *schema.ResourceData, meta interface{}) error {
	// Don't send an update request if only virtual fields have changes
	if resourceDataflowJobIsVirtualUpdate(d, ResourceDataflowFlexTemplateJob().Schema) {
		return nil
	}

	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}

	tnamemapping := tpgresource.ExpandStringMap(d, "transform_name_mapping")

	env, err := resourceDataflowFlexJobSetupEnv(d, config)
	if err != nil {
		return err
	}

	// wait until current job is running or terminated
	err = waitForDataflowJobState(d, config, d.Id(), userAgent, d.Timeout(schema.TimeoutUpdate), "JOB_STATE_RUNNING")
	if err != nil {
		return fmt.Errorf("Error waiting for job with job ID %q to be running: %s", d.Id(), err)
	}

	request := dataflow.LaunchFlexTemplateRequest{
		LaunchParameter: &dataflow.LaunchFlexTemplateParameter{

			ContainerSpecGcsPath:  d.Get("container_spec_gcs_path").(string),
			JobName:               d.Get("name").(string),
			Parameters:            tpgresource.ExpandStringMap(d, "parameters"),
			TransformNameMappings: tnamemapping,
			Environment:           &env,
			Update:                true,
		},
	}

	response, err := config.NewDataflowClient(userAgent).Projects.Locations.FlexTemplates.Launch(project, region, &request).Do()
	if err != nil {
		return err
	}

	// don't set id until new job is successfully running
	job := response.Job
	err = waitForDataflowJobState(d, config, job.Id, userAgent, d.Timeout(schema.TimeoutUpdate), "JOB_STATE_RUNNING")
	if err != nil {
		// the default behavior is to overwrite the resource's state with the state of the "new" job, even though we are returning an error here. this call to Partial prevents this behavior
		d.Partial(true)
		return fmt.Errorf("Error waiting for Job with job ID %q to be updated: %s", job.Id, err)
	}

	d.SetId(job.Id)
	if err := d.Set("job_id", job.Id); err != nil {
		return fmt.Errorf("Error setting job_id: %s", err)
	}

	return resourceDataflowFlexTemplateJobRead(d, meta)
}

func resourceDataflowFlexTemplateJobDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return err
	}

	region, err := tpgresource.GetRegion(d, config)
	if err != nil {
		return err
	}

	id := d.Id()

	requestedState, err := resourceDataflowJobMapRequestedState(d.Get("on_delete").(string))
	if err != nil {
		return err
	}

	// Retry updating the state while the job is not ready to be canceled/drained.
	err = resource.Retry(time.Minute*time.Duration(15), func() *resource.RetryError {
		// To terminate a dataflow job, we update the job with a requested
		// terminal state.
		job := &dataflow.Job{
			RequestedState: requestedState,
		}

		_, updateErr := resourceDataflowJobUpdateJob(config, project, region, userAgent, id, job)
		if updateErr != nil {
			gerr, isGoogleErr := updateErr.(*googleapi.Error)
			if !isGoogleErr {
				// If we have an error and it's not a google-specific error, we should go ahead and return.
				return resource.NonRetryableError(updateErr)
			}

			if strings.Contains(gerr.Message, "not yet ready for canceling") {
				// Retry cancelling job if it's not ready.
				// Sleep to avoid hitting update quota with repeated attempts.
				time.Sleep(5 * time.Second)
				return resource.RetryableError(updateErr)
			}

			if strings.Contains(gerr.Message, "Job has terminated") {
				// Job has already been terminated, skip.
				return nil
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	// Wait for state to reach terminal state (canceled/drained/done plus cancelling/draining if skipWait)
	skipWait := d.Get("skip_wait_on_job_termination").(bool)
	var ok bool
	ok = shouldStopDataflowJobDeleteQuery(d.Get("state").(string), skipWait)
	for !ok {
		log.Printf("[DEBUG] Waiting for job with job state %q to terminate...", d.Get("state").(string))
		time.Sleep(5 * time.Second)

		err = resourceDataflowFlexTemplateJobRead(d, meta)
		if err != nil {
			return fmt.Errorf("Error while reading job to see if it was properly terminated: %v", err)
		}
		ok = shouldStopDataflowJobDeleteQuery(d.Get("state").(string), skipWait)
	}

	// Only remove the job from state if it's actually successfully hit a final state.
	if ok = shouldStopDataflowJobDeleteQuery(d.Get("state").(string), skipWait); ok {
		log.Printf("[DEBUG] Removing dataflow job with final state %q", d.Get("state").(string))
		d.SetId("")
		return nil
	}
	return fmt.Errorf("Unable to cancel the dataflow job '%s' - final state was %q.", d.Id(), d.Get("state").(string))
}

func resourceDataflowFlexJobTypeCustomizeDiff(_ context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// All non-virtual fields are ForceNew for batch jobs
	if d.Get("type") == "JOB_TYPE_BATCH" {
		resourceSchema := ResourceDataflowFlexTemplateJob().Schema
		for field := range resourceSchema {
			if field == "on_delete" {
				continue
			}
			// Labels map will likely have suppressed changes, so we check each key instead of the parent field
			if field == "labels" {
				if err := resourceDataflowJobIterateMapForceNew(field, d); err != nil {
					return err
				}
			} else if d.HasChange(field) {
				if err := d.ForceNew(field); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
