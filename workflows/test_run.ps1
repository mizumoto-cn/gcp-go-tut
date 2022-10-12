$PROJECT_ID = $args[0]
$WORKFLOW = $args[1]


gsutil cp test.csv gs://${PROJECT_ID}-bucket/test.csv

gcloud workflows executions list ${WORKFLOW} --limit=2 --location=asia-northeast1
