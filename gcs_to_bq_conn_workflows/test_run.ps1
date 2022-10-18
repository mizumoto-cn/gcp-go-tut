$TRIGGER_SA = $args[0]
$PROJECT_ID = $args[1]

$WORKFLOW = "myEventWorkflow"

.\0_setup.ps1 $TRIGGER_SA $PROJECT_ID

.\1_trigger_creation.ps1 $TRIGGER_SA $PROJECT_ID $WORKFLOW
# For new projects, you may need to run 1_trigger_creation.ps1 twice.
# The first time, it will fail with the error:
# ERROR: (gcloud.workflows.deploy) FAILED_PRECONDITION: Workflows service agent does not exist
# ERROR: (gcloud.eventarc.triggers.describe) NOT_FOUND: Resource 'projects/bq-conn-jobload-id/locations/asia-northeast1/triggers/gcs-workflows-trigger' was not found

# .\1_trigger_creation.ps1 $TRIGGER_SA $PROJECT_ID $WORKFLOW

gsutil cp test.csv gs://${PROJECT_ID}-bucket/test.csv

# wait 5 seconds
Start-Sleep -s 10

gcloud workflows executions list ${WORKFLOW} --limit=5 --location=asia-northeast1

echo "After you get the execution ID, run the following command to get the execution status:"
echo "gcloud workflows executions describe <execution_id> --location=asia-northeast1"

Start-Sleep -s 10
bq query "select * from dataset_test.table_test limit 5"