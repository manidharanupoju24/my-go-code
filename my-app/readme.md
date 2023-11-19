Initialising a go module <br>
```go mod init my-app```<br>
Initialising kubebuilder <br>
```kubebebuilder init --domain <domain-name>``` <br>

----
```main.go``` is the main entry point for the operator, it contains the initialization code for the Kubernetes API server and the logic for starting and running the server. <br>
```Dockerfile``` is for building the k8s operator image. <br>
```Makefile```is the build automation tool that defines a set of rules for building and compiling the application. It helps automate repetitive tasks and ensure that the application is built consistently across the different environments. <br>
```PROJECT``` file contains metadata about the project, including the project name and version. <br>
```README``` provides overview of the project, including the instructions for building and running the app. <br>
```go.sum``` contains the checksum for the go modules used by the project. It helps ensure that the dependencies are downloaded securely and are not tampered with. <br>
```config``` folder contains the configuration files for the kubernetes API resources that will be defined by the project. It includes CRD files that defines the structre of the CRs, as well as RBAC. <br>
```hack``` folder contains the script and tools used during the build and deployment process. <br>

----

```kubebuilder create api --group <group-name> --version <version> --kind <kind-name>``` <br>
This command will be creating custom resources. <br>
For e.g. ```kubebuilder create api --group my-custom-app --version v1 --kind MyPythonApp``` <br>

This would create few folders. <br>

----

```api/v1``` contains the API definitions for the kubernetes custom resources that we will be defining. It includes the yaml files for the CRD that defines the structure of the custom resource as well as the types and code for the API that the project will provide. <br>
```bin``` contains the compiled binary for the application. It is typically generated during the build process and is used to start the K8s API server. <br>
```controllers``` folder contains the implemenation of the Kubernetes controllers for the custom resources defined in the api folder. <br>
The controllers are responsible for implementing the business logic for the custom resources, including handling events and updating the state of the resource, based on the changes in the Kubernetes API.<br> It includes go files for the controllers as well as any helper functions or libraries that are needed by the controllers. <br>

----
* Run ```make generate``` command after you make changes to the api/v1/mypythonapp_types.go file (in this particular scenario of operator) <br> 
* When you run this command, the kubebuilder framework reads the resource definition files and generates the corresponding go code, which includes the CRD, custom controller code and the utility code for interacting with the Kubernetes API. <br>
* Once the API is defined with spec and status field, CRD validation markers, run the ``` make manifests ``` command. <br>
* This is used to generate CRD manifest for deploying the operator and the custom resource to a kubernetes cluster. <br>
* Go to ```config/crd/bases``` and check the generated CRD. It has all the files what we defined in ```api/v1 folder``` <br>
* Whenever, we edit the types.go file in ```api/v1``` ,  we need to make sure to run ```make generate``` and ```make manifests``` commands, so that the generated CRD will be updated (in the ```config/crd/bases``` folder). <br>
* If you go to ```config/samples``` folder, you can see the sample CR file that is generated for us by kubebuilder. 
* Add all the fields and values which we defined in the previous types.go file in ```api/v1```
