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
resource "google_os_config_patch_deployment" "patch" {
  patch_deployment_id = "patch-deploy"

  instance_filter {
    group_labels {
      labels = {
        env = "dev",
        app = "web"
      }
    }

    instance_name_prefixes = ["test-"]

    zones = ["us-central1-a", "us-central-1c"]
  }

  patch_config {
    mig_instances_allowed = true

    reboot_config = "ALWAYS"

    apt {
      type = "DIST"
      excludes = ["python"]
    }

    yum {
      security = true
      minimal = true
      excludes = ["bash"]
    }

    goo {
      enabled = true
    }

    zypper {
      categories = ["security"]
    }

    windows_update {
      classifications = ["CRITICAL", "SECURITY", "UPDATE"]
      excludes = ["5012170"]
    }

    pre_step {
      linux_exec_step_config {
        allowed_success_codes = [0,3]
        local_path = "/tmp/pre_patch_script.sh"
      }

      windows_exec_step_config {
        interpreter = "SHELL"
        allowed_success_codes = [0,2]
        local_path  = "C:\\Users\\user\\pre-patch-script.cmd"
      }
    }

    post_step {
      linux_exec_step_config {
        gcs_object {
          bucket = "my-patch-scripts"
          generation_number = "1523477886880" 
          object = "linux/post_patch_script"
        }
      }

      windows_exec_step_config {
        interpreter = "POWERSHELL"
        gcs_object {
          bucket = "my-patch-scripts"
          generation_number = "135920493447"
          object = "windows/post_patch_script.ps1"
        }
      }
    }
  }

  duration = "10s"

  recurring_schedule {
    time_zone {
      id = "America/New_York"
    }

    time_of_day {
      hours = 0
      minutes = 30
      seconds = 30
      nanos = 20
    }

    monthly {
      week_day_of_month {
        week_ordinal = -1
        day_of_week  = "TUESDAY"
      }
    }
  }

  rollout {
    mode = "ZONE_BY_ZONE"
    disruption_budget {
      fixed = 1
    }
  }
}
```
