$TRIGGER_SA = $args[0]
$PROJECT_ID = $args[1]

# Enable APIs
echo "Enabling APIs: eventarc, workflows, pubsub, workflows, workflows-executions"

gcloud services enable eventarc.googleapis.com pubsub.googleapis.com `
workflows.googleapis.com workflowexecutions.googleapis.com

# gcloud alpha iam policies lint-condition

# Create the service account
echo "Creating the service account:" ${TRIGGER_SA}

gcloud iam service-accounts create ${TRIGGER_SA} 


# Grant Roles
echo "Granting:roles/eventarc.admin"
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
--member=serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com `
--role="roles/eventarc.admin"

echo "Granting:roles/workflows.invoker"
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/workflows.invoker"

echo "Granting:roles/eventarc.eventReceiver"   
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member "serviceAccount:${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com" `
    --role="roles/eventarc.eventReceiver"

# Grant the pubsub.publisher role to the Cloud Storage service account
$SERVICE_ACCOUNT="$(gsutil kms serviceaccount -p ${PROJECT_ID})"
echo "Granting:roles/pubsub.publisher to" $SERVICE_ACCOUNT
gcloud projects add-iam-policy-binding ${PROJECT_ID} `
    --member="serviceAccount:${SERVICE_ACCOUNT}" `
    --role="roles/pubsub.publisher"

