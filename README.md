# Micro Message

This application is for operating messages based on rest api with simple user authentication, registration and login. It is easy to deploy application on aws with only 2 steps.

## Getting Started

1. If you want to deploy application on aws, just need to run deploy script in terraform. 
2. If you want to test and build the project locally, you can just run go build.

Please refer steps below for details.

### Prerequisites

What things you need to install the software and how to install them

```
1. You need to install terraform locally and you can download from (https://www.terraform.io/downloads.html). Download corresponding package based on your local system.
2. You must have a ssh key pair for deployment. 
3. You must have a aws account associated with deployment. 
4. You must have golang and go dep package installed for build this project.
```

### Installing

1. Deploy application
* Have your aws access key and access secret key ready and use them to change the default value in <project-base>/terraform/userVariables.tf file. Also specify your ssh private key location for default value of variable 'privateKeyFilename' in the same file. (note you can also change the openshift username and password as you want.)

* Go to <project-base>/terraform and run deploy.sh script without passing any arguments. 

* The installation will take about 30 mins for deploying everything(including provisioning aws resources, installing openshift, configuring openshift, creating database pod, creating application pod and so on.)

* After installation and seeing the message "You have successfully deployed application on aws in openshift cluster.", then you are ready to go. You can check your master public dns in aws or you can just run: 
```
terraform output master-public_dns
```
* The openshift management console url is (https://<your master public dns>:8443), username and password are provided by you in userVariables.tf file.

* The application accessing url is (http://<your master public dns>). You can only create, view or delete message after you have register an account or login.

2. Build from the source
* Download project into your GOPATH by running
```
git clone git@github.com:xqianwang/micro-message.git
```
* Install necessary dependencies by running command in project base directory:
```
dep ensure -vendor-only
``` 
* If you can see vendor folder in project base already, then you don't need to run previous step.

## Running the tests

Go to project base folder and run command below:
```
go test -v ./...
```
## REST API for messages
Micro-Message application allows users to interact with application by REST API or UI. 
In terms of REST API:
* You can get all messages by running:
```
curl -X GET -H "Accept: application/json" http://<public-dns>/messages
```
* You can get query 1 message by running:
```
curl -X GET -H "Accept: application/json" http://<public-dns>/messages/view/<:messageid>
```
* You can delete a message by running:
```
curl -X DELETE -H "Accept: application/json" http://<public-dns>/messages/<:messageid>
```
* You can create a new message by running: 
```
curl -X POST -d "title=8&content=8" -H "Accept: application/json" http://<public-dns>/messages/create
```

Also you can absolutely interact with application through UI!


## Application infrastructure
Micro-Message's infrastructure has a few parts:
1) Backend database: Postgresql 9.6
2) REST API: golang
3) Template handler: golang template
4) Deployment and aws resource provisioning: terraform
5) Dockerization: docker
6) Container management platform: Openshift origin 3.9
7) UI: Bootstrap
8) Framework: [golang gin](https://github.com/gin-gonic/gin) [openshift-aws](https://github.com/dwmkerr/terraform-aws-openshift.git)

## Deployment

You can refer previous steps for deployment.


## Versioning

For the versions available, see the [tags on this repository](https://github.com/xqianwang/micro-message/tags). 

## Authors

* **Xueqian Wang** - *Initial work* - You can contact with me by email: quntas.wang@gmail.com

See also the list of [contributors](https://github.com/xqianwang/micro-message/graphs/contributors). There are 2 contributors, both of them are my accounts because I have 1 working account in my company computer and 1 private account in my private computer.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
