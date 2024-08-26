# Periodic tests

These tests run against real GCP resources, and are run periodically against test GCP accounts.

The results are not currently published publicly.

If you want to run the tests locally (they will incur GCP costs), you can do something like this:

```
GCP_PROJECT_ID=$(gcloud config get-value project)

PARENT_FOLDER_ID=$(gcloud projects describe ${GCP_PROJECT_ID} --format='value(parent.id)')
export PARENT_FOLDER_ID

BILLING_ACCOUNT_ID=$(gcloud billing projects describe ${GCP_PROJECT_ID} --format='value(billingAccountName)' | cut -f2 -d/)
export BILLING_ACCOUNT_ID

ARTIFACTS=`pwd`/artifactz
mkdir -p ${ARTIFACTS}
export ARTIFACTS

dev/ci/periodics/e2e-service-monitoring
```
