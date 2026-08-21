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
resource "google_data_loss_prevention_deidentify_template" "basic" {
	parent = "projects/my-project-name"
	description = "Description"
	display_name = "Displayname"

	deidentify_config {
		info_type_transformations {
			transformations {
				info_types {
					name = "FIRST_NAME"
				}

				primitive_transformation {
					replace_with_info_type_config = true
				}
			}

			transformations {
				info_types {
					name = "PHONE_NUMBER"
				}
				info_types {
					name = "AGE"
				}

				primitive_transformation {
					replace_config {
						new_value {
							integer_value = 9
						}
					}
				}
			}

			transformations {
				info_types {
					name = "EMAIL_ADDRESS"
				}
				info_types {
					name = "LAST_NAME"
				}

				primitive_transformation {
					character_mask_config {
						masking_character = "X"
						number_to_mask = 4
						reverse_order = true
						characters_to_ignore {
							common_characters_to_ignore = "PUNCTUATION"
						}
					}
				}
			}

			transformations {
				info_types {
					name = "DATE_OF_BIRTH"
				}

				primitive_transformation {
					replace_config {
						new_value {
							date_value {
								year  = 2020
								month = 1
								day   = 1
							}
						}
					}
				}
			}

      transformations {
        info_types {
          name = "CREDIT_CARD_NUMBER"
        }

        primitive_transformation {
          crypto_deterministic_config {
            context {
              name = "sometweak"
            }
            crypto_key {
              transient {
                name = "beep"
              }
            }
            surrogate_info_type {
              name = "abc"
            }
          }
        }
      }
		}
	}
}
```
