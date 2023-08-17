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
# VPC network
resource "google_compute_network" "default" {
    name                    = "tf-test-workstation-cluster%{random_suffix}"
    auto_create_subnetworks = true
}


data "google_project" "project" {

}

resource "google_dataplex_lake" "example_notebook" {
  name         = "tf-test-lake%{random_suffix}"
  location     = "us-central1"
  project = "my-project-name"
}


resource "google_dataplex_task" "example_notebook" {

    task_id      = "tf-test-task%{random_suffix}"
    location     = "us-central1"
    lake         = google_dataplex_lake.example_notebook.name
    trigger_spec  {
        type = "RECURRING"
        schedule = "1 * * * *"
    }
    
    execution_spec {
        service_account = "${data.google_project.project.number}-compute@developer.gserviceaccount.com"
        args = {
            TASK_ARGS  = "--output_location,gs://spark-job-jars-anrajitha/task-result, --output_format, json"
        }
    }
    notebook {
        notebook = "gs://terraform-test/test-notebook.ipynb"
        infrastructure_spec  {
            batch {
                executors_count = 2
                max_executors_count = 100
            }
            container_image {
                image = "test-image"
                java_jars = ["test-java-jars.jar"]
                python_packages = ["gs://bucket-name/my/path/to/lib.tar.gz"]
                properties = { "name": "wrench", "mass": "1.3kg", "count": "3" }
            }
            vpc_network  {
                    network_tags = ["test-network-tag"]
                    network = google_compute_network.default.id
                }
        }
        file_uris = ["gs://terraform-test/test.csv"]
        archive_uris = ["gs://terraform-test/test.csv"]

    }
    project = "my-project-name"
    
    
}
```
