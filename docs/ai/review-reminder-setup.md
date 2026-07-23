# PR Review Reminder Setup & WIF Integration Guide

This guide documents the setup for sending twice-daily PR review reminders to `kcc-eng@google.com` for the `k8s-config-connector-team` using GitHub Actions, GCP Workload Identity Federation (WIF), and the Gmail API.

---

## 1. Overview & Objectives

- **Target Audience**: `k8s-config-connector-team` (`acpana`, `anfernee`, `anhdle-sso`, `barney-s`, `gemmahou`, `maqiuyujoyce`).
- **Schedule**: Twice daily at 9:00 AM and 2:00 PM UTC (`0 9,14 * * *`).
- **Recipient**: `kcc-eng@google.com`
- **Authentication**: Zero static secrets in GitHub using GCP Workload Identity Federation (WIF) and Gmail API.
- **Target GCP Project**: `cnrm-eap`.

---

## 2. GCP Setup (`cnrm-eap` Project)

Run the following `gcloud` commands as a project admin:

```bash
# 1. Set active GCP project to cnrm-eap
gcloud config set project cnrm-eap

# 2. Enable necessary APIs
gcloud services enable iamcredentials.googleapis.com \
                       iam.googleapis.com \
                       gmail.googleapis.com

# 3. Create a dedicated GCP Service Account
gcloud iam service-accounts create kcc-review-reminder-sa \
  --display-name="KCC Review Reminder GitHub Action Bot"

# 4. Create the Workload Identity Pool
gcloud iam workload-identity-pools create "github-actions-pool" \
  --location="global" \
  --display-name="GitHub Actions Pool"

# 5. Create the Workload Identity Provider for GitHub OIDC
# Restricted specifically to GoogleCloudPlatform/k8s-config-connector repository
gcloud iam workload-identity-pools providers create-oidc "github-actions-provider" \
  --location="global" \
  --workload-identity-pool="github-actions-pool" \
  --display-name="GitHub Actions OIDC Provider" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
  --attribute-condition="attribute.repository == 'GoogleCloudPlatform/k8s-config-connector'" \
  --issuer-uri="https://token.actions.githubusercontent.com"

# 6. Fetch Project Number for cnrm-eap
PROJECT_NUMBER=$(gcloud projects describe cnrm-eap --format="value(projectNumber)")

# 7. Grant GitHub Actions permission to impersonate the Service Account
gcloud iam service-accounts add-iam-policy-binding "kcc-review-reminder-sa@cnrm-eap.iam.gserviceaccount.com" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/projects/${PROJECT_NUMBER}/locations/global/workloadIdentityPools/github-actions-pool/attribute.repository/GoogleCloudPlatform/k8s-config-connector"
```

---

## 3. Google Workspace Setup (Domain-Wide Delegation for Gmail API)

To allow the Service Account to send emails via Gmail API on behalf of the team identity:

1. Retrieve the Service Account's **Unique Client ID**:
   ```bash
   gcloud iam service-accounts describe kcc-review-reminder-sa@cnrm-eap.iam.gserviceaccount.com --format="value(uniqueId)"
   ```
2. Open the **Google Workspace Admin Console** (`admin.google.com`).
3. Navigate to **Security > Access and data control > API controls > Manage Domain Wide Delegation**.
4. Add a new API client using the **Unique Client ID** obtained above.
5. Grant the OAuth Scope:
   - `https://www.googleapis.com/auth/gmail.send`

---

## 4. GitHub Actions Workflow

The workflow is defined at [.github/workflows/review-reminder.yaml](file:///.github/workflows/review-reminder.yaml).
