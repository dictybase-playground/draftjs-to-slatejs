# Draft.js to Slate.js converter

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictybase-playground/draftjs-to-slatejs/workflows/Continuous%20integration/badge.svg)
[![Issues](https://badgen.net/codeclimate/issues/dictybase-playground/draftjs-to-slatejs)](https://codeclimate.com/github/dictybase-playground/draftjs-to-slatejs/issues)
[![Maintainability](https://api.codeclimate.com/v1/badges/21ed283a6186cfa3d003/maintainability)](https://codeclimate.com/github/dictybase-playground/draftjs-to-slatejs/maintainability)  
![Last commit](https://badgen.net/github/last-commit/dictybase-playground/draftjs-to-slatejs/develop)
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,dictyBase/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9476993)
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,DSC/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9438930)

This is a hybrid Golang/Node.js script that converts Draft.js content to Slate.js 0.4x
compatible content.

The Golang portion is responsible for retrieving all Draft.js content from the content API (using gRPC),
executing the Node.js script to convert to HTML then to Slate.js content, then updating the content API
with the new data. During this process, all data is also uploaded to a user-specified Minio bucket.

There are two Helm charts that need to be installed in the following order.

1. [dictybase/convert-draftjs-content](./deployments/charts/convert-draftjs-content)
2. [dictybase/upload-slatejs-content](./deployments/charts/upload-slatejs-content)

Click the above links to see the documentation and values for each chart.

## Usage

```
NAME:
   draftjs-to-slate - cli for replacing draft.js data through grpc

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   get-draftjs-content  gets draft.js content and uploads their json to minio
   convert-content      runs node.js script to convert draft.js to slate.js content
   update-content       updates API with downloaded slate.js content
   help, h              Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-format value  format of the logging out, either of json or text. (default: "json")
   --log-level value   log level for the application (default: "error")
   --help, -h          show help
   --version, -v       print the version
```

## Subcommands

```
NAME:
   main get-draftjs-content - gets draft.js content and uploads their json to minio

USAGE:
   main get-draftjs-content [command options] [arguments...]

OPTIONS:
   --minio-host value         minio host [$MINIO_SERVICE_HOST]
   --minio-port value         minio port [$MINIO_SERVICE_PORT]
   --minio-access-key value   minio access key
   --minio-secret-key value   minio secret key
   --minio-bucket value       minio bucket (default: "content")
   --minio-location value     minio location (default: "us-east-1")
   --user-id value            user id to use for updating content
   --convert                  identifier to start conversion process
   --content-grpc-host value  content grpc host [$CONTENT_API_SERVICE_HOST]
   --content-grpc-port value  content grpc port [$CONTENT_API_SERVICE_PORT]
```

```
NAME:
   main convert-content - runs node.js script to convert draft.js to slate.js content

USAGE:
   main convert-content [command options] [arguments...]

OPTIONS:
   --minio-host value         minio host [$MINIO_SERVICE_HOST]
   --minio-port value         minio port [$MINIO_SERVICE_PORT]
   --minio-access-key value   minio access key
   --minio-secret-key value   minio secret key
   --minio-bucket value       minio bucket (default: "content")
   --minio-location value     minio location (default: "us-east-1")
   --user-id value            user id to use for updating content
   --convert                  identifier to start conversion process
   --content-grpc-host value  content grpc host [$CONTENT_API_SERVICE_HOST]
   --content-grpc-port value  content grpc port [$CONTENT_API_SERVICE_PORT]
```

```
NAME:
   main update-content - updates API with downloaded slate.js content

USAGE:
   main update-content [command options] [arguments...]

OPTIONS:
   --minio-host value         minio host [$MINIO_SERVICE_HOST]
   --minio-port value         minio port [$MINIO_SERVICE_PORT]
   --minio-access-key value   minio access key
   --minio-secret-key value   minio secret key
   --minio-bucket value       minio bucket (default: "content")
   --minio-location value     minio location (default: "us-east-1")
   --user-id value            user id to use for updating content
   --convert                  identifier to start conversion process
   --content-grpc-host value  content grpc host [$CONTENT_API_SERVICE_HOST]
   --content-grpc-port value  content grpc port [$CONTENT_API_SERVICE_PORT]
```
