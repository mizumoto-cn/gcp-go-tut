$TRIGGER_SA = args[0]
$PROJECT_ID = args[1]

# Enable APIs
gcloud services enable eventarc.googleapis.com `
workflows.googleapis.com workflowexecutions.googleapis.com

# Create the service account
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
--member=serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com `
--role="roles/eventarc.admin"


# Grant Roles
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/workflows.invoker"

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member "serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/eventarc.eventReceiver"


# Grant the pubsub.publisher role to the Cloud Storage service account
SERVICE_ACCOUNT="$(gsutil kms serviceaccount -p ${PROJECT_ID})"

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${SERVICE_ACCOUNT}" `
    --role="roles/pubsub.publisher"


