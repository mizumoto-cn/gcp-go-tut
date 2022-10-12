$PROJECT_ID = args[0]
$WORKFLOW = args[1]

echo "Hello World" > random.txt
gsutil cp random.txt gs://${PROJECT_ID}-bucket/random.txt

gcloud workflows executions list ${WORKFLOW} --limit=2
