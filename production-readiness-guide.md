# Config Connector Alpha Production Readiness Checklist

Config Connector Alpha is a limited-availability experiment to prepare it for the next release stage. Our focus with Alpha testing is to verify functionality and gather feedback from a limited set of customers. Participation in Config Connector Alpha is by invitation only and the binaries are not intended to be shared publicly. 

No SLAs are provided for Config Connector Alpha, and there are no technical support obligations. However, Config Connector Alpha is extensively tested and is suitable for use in test environments. 

We recommend that you understand and consider the following recommendations before proceeding with production deployment.

1. Create the intended number and types of GCP resources and verify that they are deployed successfully on your test cluster  before proceeding with production cluster deployment.
1. The default Config Connector behavior is that GCP resources created by Config Connector are deleted when Custom Objects, such as PubSubTopic or  Custom Resource Definitions installed as part of Config Connector are deleted. If you prefer that Config Connector does not delete the GCP resources under management, [enable the abandon deletion policy](https://github.com/GoogleCloudPlatform/k8s-config-connector#resource-abandonment).  
1. To prevent system error, it is recommended that for a given GCP Project, there is only a single namespace in a single Config-Connector enabled cluster that is mapped to the Project.
