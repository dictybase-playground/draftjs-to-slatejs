# convert-draftjs-content

This [helm](https://github.com/kubernetes/helm) chart provides [kubernetes](http://kubernetes.io) manifests for running an [convert-draftjs-content](https://hub.docker.com/r/dictybase/draftjs-to-slatejs/) job.

# Managing the chart

## Install

```
helm install --name dev-release dictybase/convert-draftjs-content
```

For details, look [here](https://docs.helm.sh/using_helm/#helm-install-installing-a-package).

## Uninstall

```
helm uninstall dev-release
```

For details, look [here](https://docs.helm.sh/using_helm/#uninstall-a-release).

For upgrades and rollback, look [here](https://docs.helm.sh/using_helm/#helm-upgrade-and-helm-rollback-upgrading-a-release-and-recovering-on-failure).

## Configuration

The following tables lists the configurable parameters of the **convert-draftjs-content** chart and their default values.

| Parameter          | Description                     | Default                        |
| ------------------ | ------------------------------- | ------------------------------ |
| `image.repository` | convert-draftjs-content image   | `dictybase/draftjs-to-slatejs` |
| `image.tag`        | image tag                       | `latest`                       |
| `image.pullPolicy` | Image pull policy               | `IfNotPresent`                 |
| `minio.akey`       | Minio access key                | ``                             |
| `minio.skey`       | Minio secret key                | ``                             |
| `minio.bucket`     | Minio bucket to upload files to | `dsc-contents`                 |
| `minio.location`   | Minio region                    | `us-east-1`                    |
| `userId`           | Password for new ArangoDB user  | `9999`                         |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```bash
$ helm install --name my-release -f values.yaml convert-draftjs-content
```
