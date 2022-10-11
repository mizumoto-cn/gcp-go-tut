$DESTINATION_WORKFLOW = args[0]
$BUCKET_ID = args[1]
$TRIGGER_SA = args[2]
$PROJECT_ID = args[3]

# Create the trigger 
# Event is sent when a new object is created (or an existing object
# is overwritten, and a new generation of that object is created) 
# These flags are required:
# --event-filters="type=EVENT_FILTER_TYPE"
# --event-filters="bucket=BUCKET"
gcloud eventarc triggers create gcs-workflows-trigger `
--location=asia-northeast1 `
--destination-workflow=${DESTINATION_WORKFLOW}  `
--destination-workflow-location=asia-northeast1 `
--event-filters="type=google.cloud.storage.object.v1.finalized"  `
--event-filters="bucket=${BUCKET_ID}" `
--service-account=" ${TRIGGER_SA}@${PROJECT_ID}.iam.gserviceaccount.com"

# List the triggers
gcloud eventarc triggers list --location=asia-northeast1
