resource "google_logging_folder_sink" "logginglogsink_sample_folder" {
  destination = "logginglogsinkdepfolder"
  filter      = "resource.type=\"bigquery_project\" AND logName:\"cloudaudit.googleapis.com\""
  folder      = "logginglogsink-dep-folder"
  name        = "logginglogsink-sample-folder"
}
# terraform import google_logging_folder_sink.logginglogsink_sample_folder #logginglogsink-dep-folder##logginglogsink-sample-folder
