<p align="center">
    <a href="http://kitura.io/">
        <img src="https://landscape.cncf.io/logos/ibm-cloud.svg" height="100" alt="IBM Cloud">
    </a>
</p>


<p align="center">
    <a href="https://cloud.ibm.com">
    <img src="https://img.shields.io/badge/IBM%20Cloud-powered-blue.svg" alt="IBM Cloud">
    </a>
    <img src="https://img.shields.io/badge/platform-go-lightgrey.svg?style=flat" alt="platform">
    <img src="https://img.shields.io/badge/license-Apache2-blue.svg?style=flat" alt="Apache 2">
</p>


# Create and deploy a basic Go Gin application

> We have similar applications available for [Node.js](https://github.com/IBM/nodejs-starter), [Java Spring](https://github.com/IBM/spring-starter), [Swift](https://github.com/IBM/swift-starter), [Python Flask](https://github.com/IBM/flask-starter), and [Java Liberty](https://github.com/IBM/java-liberty-starter).

In this sample web application, you will create a web application using Gin to serve web pages in Go, complete with standard best practices, including a health check and application metric monitoring.

This app contains an opinionated set of files for web serving:

- `public/index.html`
- `public/404.html`
- `public/500.html`

## Steps

You can [deploy this application to IBM Cloud](https://cloud.ibm.com/developer/appservice/create-app?navMode=starterkits) or [build it locally](#building-locally) by cloning this repo first.  Once your app is live, you can access the `/health` endpoint to build out your cloud native application.

### Deploying to IBM Cloud

<p align="center">
    <a href="https://cloud.ibm.com/developer/appservice/create-app?navMode=starterkits">
    <img src="https://cloud.ibm.com/devops/setup/deploy/button_x2.png" alt="Deploy to IBM Cloud">
    </a>
</p>

Use the button above to deploy this same application to IBM Cloud.  This option will create a deployment pipeline, complete with a hosted Git lab project and devops toolchain.  You will have the option of deploying to either CloudFoundry or a Kubernetes cluster. [IBM Cloud DevOps](https://www.ibm.com/cloud-computing/bluemix/devops) services provides toolchains as a set of tool integrations that support development, deployment, and operations tasks inside IBM Cloud. 


### Building Locally

To get started building this web application locally, you can either run the application natively or use the IBM Cloud Developer Tools for containerization and easy deployment to IBM Cloud.

All of your `dep` dependencies are stored inside of `Gopkg.toml`.

#### Native Application Development

- Install [Go](https://golang.org/dl/)
- Install [dep](https://github.com/golang/dep)

In order for Go applications to run locally, they must be placed in the correct file path. The application must exist in `$GOPATH/src/gostarter`

Once the Go toolchain has been installed, you can compile a Go project with:

```bash
go install
```

To run your application locally:

```bash
dep ensure
go run server.go
```

Your sources will be compiled to your `$GOPATH/bin` directory. Your application will be running at `http://localhost:8080`.

#### IBM Cloud Developer Tools

Install [IBM Cloud Developer Tools](https://cloud.ibm.com/docs/cli/index.html#overview) on your machine by using the following installation command: `curl -sL https://ibm.biz/idt-installer | bash`.

Your application will be compiled with Docker containers. To compile and run your app, run:

```bash
ibmcloud dev build
ibmcloud dev run
```

This will launch your application locally.  When you are ready to deploy to IBM Cloud on CloudFoundry or Kubernetes, run one of the commands below:

```bash
ibmcloud dev deploy -t buildpack
ibmcloud dev deploy -t container
```

You can build and debug your app locally with:

```bash
ibmcloud dev build --debug
ibmcloud dev debug
```

## Next Steps
* Learn more about augmenting your Go applications on IBM Cloud with the [Go Programming Guide](https://cloud.ibm.com/docs/go).
* Explore other [sample applications](https://cloud.ibm.com/developer/appservice/starter-kits) on IBM Cloud.

## License

This sample application is licensed under the Apache License, Version 2. Separate third-party code objects invoked within this code pattern are licensed by their respective providers pursuant to their own separate licenses. Contributions are subject to the [Developer Certificate of Origin, Version 1.1](https://developercertificate.org/) and the [Apache License, Version 2](https://www.apache.org/licenses/LICENSE-2.0.txt).

[Apache License FAQ](https://www.apache.org/foundation/license-faq.html#WhatDoesItMEAN)
