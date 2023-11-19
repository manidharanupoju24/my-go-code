* What is a k8s operator? <br>
Kubernetes operator is a software extension to Kubernetes that makes it easier to manage complex applications and infrastructure. <br>
    * Operator allows you to automate the management of applications on kubernetes cluster. 
* Why build a k8s operator? <br>
    * The main reason is to automate the management and deployment of complex applications. <br>
    * To manage the entire application lifecycle within Kubernetes, including updates, rollbacks and scaling. Also, managing the health of our application. <br>
    * To abstract the underlying infrastructure and provide a higher level of abstraction to manage applications. <br>
* When do you build a k8s operator? <br>
    * When you have a complex application that requires multiple components to be coordinated. <br>
    * When you want to automate the management of application within k8s. 
    * When you want to provide a higher level of abstraction for managing an application.
* How to build a k8s operator?
    * Kubernetes provides few built-in resources suchs as pods, replicasets, deployments etc. We can use these to create and manage our applications in a K8s cluster.
    * Custom resources on the other hand are the resources that we define ourselves. They allow us to extend Kubernetes with our own APIs and resources

* Difference between CRD and CR? 
    * CRDs play a role in defining the structure and validation rules for custom resources.
    * CRDs are used to define custom resources 
    * CRDs defines the blueprint for creating custom resources in a cluster. 
    * CRDs allow us to extend the K8S API with our own custom resources. 
    * (In simple terms, CRD is like a form that we fill out to create a ne object in our cluster)
    * CRD defines what fields that an object can have, what data types those fields should be, and any validation rules that should be applied. 
    * CRDs are managed at the cluster level

* Using a CRD, we can create a new object in our cluster, called CR. 

* Each CR is an instance of the CRS which are defined in the CRD. 

* CRs play a role in providing data for the custom resources. 

* CRs are used to store the data for those custom resources. 

* CRs are managed at the indvidual object level. (You can create, update or delete CRS independent of each other). 

* However, CRDs are the ones which create structure and validation rules and these are managed at cluster level by CRDs. 

* This would ensure consistency across all CRS of that particular type. 

* This is similar to Dockerfile (CRD here), Docker Image (CR here) and Docker container (Running instance of a CR)

---- 

Controller loop 

* Each controller has its own controller loop (pod, deployment, replicaset etc have their own)

* Custom Controller too have their own controller loops. 

----

Difference between Controller and Operator? 

Controller | Operator 
|---|---|
A component that runs one or more control loops | A custom controller that automates the deployment and management of complex apps |
Monitors the state of resources and makes changes to bring the actual state in line to the desired state | Extends the Kubernetes control loop to manage the lifecycle of a specific type of resource, such as database or a message queue

----

Tools available for creating operators : 
* kubebuilder
* operator-sdk
* metacontroller
* controller-runtime