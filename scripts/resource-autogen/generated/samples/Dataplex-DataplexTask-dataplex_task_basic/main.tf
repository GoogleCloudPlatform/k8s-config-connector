/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
data "google_project" "project" {

}

resource "google_dataplex_lake" "example" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "my-project-name"
}


resource "google_dataplex_task" "example" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example.name
    
    description = "Test Task Basic"
    
    display_name = "task-basic"

    labels = { "count": "3" }

    trigger_spec  {
        type = "RECURRING"
        disabled = false
        max_retries = 3
        start_time = "2023-10-02T15:01:23Z"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        project = "my-project-name"
        max_job_execution_lifetime = "100s"
        kms_key = "234jn2kjn42k3n423"
    }
    
    spark {
        python_script_file = "gs://dataproc-examples/pyspark/hello-world/hello-world.py"

    }
    
    project = "my-project-name"
    
}
```
