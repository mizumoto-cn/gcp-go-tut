# GCP Cloud SDK Tutorial in Golang

This is a tutorial for using the GCP Cloud SDK in Golang. 

## App Engine Flexible

### Preparations

Enable Cloud Build API.

Then

```bash
gcloud app create --project=[YOUR_PROJECT_ID]

gcloud components install app-engine-go
```

### Deploy

Switch to the source directory containing the app.yaml file and run the following command:

```bash
gcloud app deploy
```

And the app will be deployed at `https://PROJECT_ID.REGION_ID.r.appspot.com`.

You can also access by typing the following command:

```bash
gcloud app browse
```

### Disable the app

To disable an App Engine application:

1. Go to the [Application settings page](https://console.cloud.google.com/appengine/settings?hl=zh-cn&_ga=2.241673547.83809591.1665311107-1681099199.1663672191&_gac=1.14170565.1665095474.Cj0KCQjw-fmZBhDtARIsAH6H8qizjiQbk4rQj-pUT-3lBzwjV_qBSQod20-YGpbY0_DP29k1nhvQpKcaAhvFEALw_wcB)
2. Click Disable application and then follow the prompts.

When you want your app to continue serving requests, enable the App Engine application by returning to the same Application settings page and then clicking Enable application.

## Pub/Sub to BigQuery