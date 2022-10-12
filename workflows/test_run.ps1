$TRIGGER_SA = $args[0]
$PROJECT_ID = $args[1]

$WORKFLOW = "myEventWorkflow"

.\0_setup.ps1 $TRIGGER_SA $PROJECT_ID

.\1_create_workflow.ps1 $TRIGGER_SA $PROJECT_ID $WORKFLOW


gsutil cp test.csv gs://${PROJECT_ID}-bucket/test.csv

# wait 5 seconds
Start-Sleep -s 10

gcloud workflows executions list ${WORKFLOW} --limit=5 --location=asia-northeast1

echo "After you get the execution ID, run the following command to get the execution status:"
echo "gcloud workflows executions describe <execution_id> --location=asia-northeast1"