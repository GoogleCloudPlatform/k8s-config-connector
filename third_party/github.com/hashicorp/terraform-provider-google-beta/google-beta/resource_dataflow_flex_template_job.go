package google

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

	request := dataflow.LaunchFlexTemplateRequest{
		LaunchParameter: &dataflow.LaunchFlexTemplateParameter{
			ContainerSpecGcsPath: d.Get("container_spec_gcs_path").(string),
			JobName:              d.Get("name").(string),
			Parameters:           tpgresource.ExpandStringMap(d, "parameters"),
			Environment: &dataflow.FlexTemplateRuntimeEnvironment{
				AdditionalUserLabels: tpgresource.ExpandStringMap(d, "labels"),
			},
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

	if err := d.Set("state", job.CurrentState); err != nil {
		return fmt.Errorf("Error setting state: %s", err)
	}
	if err := d.Set("name", job.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("labels", job.Labels); err != nil {
		return fmt.Errorf("Error setting labels: %s", err)
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
			if transport_tpg.IsRetryableError(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}

		state := job.CurrentState
		if state == targetState {
			log.Printf("[DEBUG] the job with ID %q has state %q.", jobID, state)
			return nil
		}
		_, terminating := dataflowTerminatingStatesMap[state]
		if terminating && targetState == "JOB_STATE_RUNNING" {
			return resource.NonRetryableError(fmt.Errorf("the job with ID %q is terminating with state %q and cannot reach expected state %q", jobID, state, targetState))
		}
		if _, terminated := dataflowTerminalStatesMap[state]; terminated {
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

	// wait until current job is running or terminated
	err = waitForDataflowJobState(d, config, d.Id(), userAgent, d.Timeout(schema.TimeoutUpdate), "JOB_STATE_RUNNING")
	if err != nil {
		return fmt.Errorf("Error waiting for job with job ID %q to be running: %s", d.Id(), err)
	}

	request := dataflow.LaunchFlexTemplateRequest{
		LaunchParameter: &dataflow.LaunchFlexTemplateParameter{
			ContainerSpecGcsPath: d.Get("container_spec_gcs_path").(string),
			JobName:              d.Get("name").(string),
			Parameters:           tpgresource.ExpandStringMap(d, "parameters"),
			Environment: &dataflow.FlexTemplateRuntimeEnvironment{
				AdditionalUserLabels: tpgresource.ExpandStringMap(d, "labels"),
			},
			Update: true,
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
