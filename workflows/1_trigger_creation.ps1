$TRIGGER_SA = $args[0]
$PROJECT_ID = $args[1]
$WORKFLOW = $args[2]

# Create the Storage Bucket
gsutil mb -l asia-northeast1 gs://${PROJECT_ID}-bucket/

# Deploy Workflow
gcloud workflows deploy ${WORKFLOW} --source=${WORKFLOW}.yaml --location=asia-northeast1

# Create the trigger 
# Event is sent when a new object is created (or an existing object
# is overwritten, and a new generation of that object is created) 
# These flags are required:
# --event-filters="type=EVENT_FILTER_TYPE"
# --event-filters="bucket=BUCKET"
$DQUOTE = '"'
$BUCKET_FILTER = "bucket=${DQUOTE}${PROJECT_ID}-bucket${DQUOTE}"
$SA_FOR_TRIGGER = "${DQUOTE}${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com${DQUOTE}"

gcloud eventarc triggers create gcs-workflows-trigger `
--location=asia-northeast1 `
--destination-workflow=${WORKFLOW}  `
--destination-workflow-location=asia-northeast1 `
--event-filters="type=google.cloud.storage.object.v1.finalized"  `
--event-filters=${BUCKET_FILTER} `
--service-account=${SA_FOR_TRIGGER}

# List the triggers
gcloud eventarc triggers list --location=asia-northeast1
gcloud eventarc triggers describe gcs-workflows-trigger --location=asia-northeast1
