#! /bin/bash

# Download the application
APP_LOCATION=gs://gcp-example-1290/gcp-example
gsutil cp "$APP_LOCATION" gcp-example
tar -xzf gcp-example

# Get templates
TEMPLATE_LOCATION=gs://gcp-example-1290/template
gsutil cp "$TEMPLATE_LOCATION" template

mkdir tpl

tar -xzf template -C /tpl

# Run the binary
service app start