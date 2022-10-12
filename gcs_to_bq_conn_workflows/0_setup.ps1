$TRIGGER_SA = $args[0]
$PROJECT_ID = $args[1]



# Enable APIs
gcloud services enable eventarc.googleapis.com pubsub.googleapis.com `
workflows.googleapis.com workflowexecutions.googleapis.com

# gcloud alpha iam policies lint-condition

# Create the service account
gcloud iam service-accounts create ${TRIGGER_SA} 


# Grant Roles

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
--member=serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com `
--role="roles/eventarc.admin"

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/workflows.invoker"

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member "serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/eventarc.eventReceiver"


# Grant the pubsub.publisher role to the Cloud Storage service account
$SERVICE_ACCOUNT="$(gsutil kms serviceaccount -p ${PROJECT_ID})"

gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${SERVICE_ACCOUNT}" `
    --role="roles/pubsub.publisher"

