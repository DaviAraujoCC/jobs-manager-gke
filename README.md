# About

A project created to manage jobs/cronjobs in GKE.

## Requirements

* Go 1.18+
* Make

## Authentication

This application supports ~/kube/config authentication or via service account.

## Variables:

| Variable | Description |
| --- | --- |
| `HTTP_PORT` | Port to listen on |
| `LOG_LEVEL` | Log level |
| `SECRET_KEY` | Secret key to access API |


## Usage

Run `make run` to run the application locally, press Ctrl + C to exit.

* Before running the application locally, create a .env file with the variables.

### Docker 

To create the image and push it to docker hub:

```
$ make docker-build docker-push IMG={Image name} TARGET=release
```

Edit the manifest inside manifests folder and after that run:

```
$ kubectl apply -k manifests/
```

To undeploy:

```
$ kubectl delete -k manifests/
```


## TODO

- [ ] Create swagger documentation
