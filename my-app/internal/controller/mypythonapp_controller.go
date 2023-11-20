/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"errors"
	"reflect"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mycustomappv1 "my-app/api/v1"
)

// MyPythonAppReconciler reconciles a MyPythonApp object
type MyPythonAppReconciler struct {
	client.Client // this is an interface that provides methods to interact with the Kubernetes API server
	Scheme        *runtime.Scheme
	//Scheme is used to manage the versioned types that are used in Kubernetes object
	Log logr.Logger
}

//+kubebuilder:rbac:groups=my-custom-app.stickers.com,resources=mypythonapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=my-custom-app.stickers.com,resources=mypythonapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=my-custom-app.stickers.com,resources=mypythonapps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MyPythonApp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
// The below function is the main logic of the controller. It is called everytime an object (In this case, myPythonApp object) is created, updated or deleted.
// This function purpose is to compare the desired state of the myPythonApp object with the actual state of the Kubernetes cluster and then perform operations to make the 2 states consistent.
func (r *MyPythonAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// _ = log.FromContext(ctx)
	log := r.Log.WithValues("MyPythonApp", req.NamespacedName)

	// TODO(user): your logic here
	operator := &mycustomappv1.MyPythonApp{}
	err := r.Client.Get{ctx, req.NamespacedName, operator}
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Operator resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Operator")
		return ctrl.Result{}, err
	}

	found := &appsv1.Deployment{}                                                                    // observed deployment object associated with the MyPythonApp resource in k8s cluster
	err = r.Get(ctx, types.NamespaceName{Name: operator.Name, Namespace: operator.Namespace}, found) //observe
	// condition to check if the deployment resource associated with MyPythonApp resource exists in the kubernetes cluster.
	if err != nil && errors.IsNotFound(err) { //check difference in this if
		dep := r.deploymentForOperator(operator) //we should write this method
		log.Info()
		err = r.Create(ctx, dep) // take action
		if err != nil {
			log.Error(err, "Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Deployment")
		return ctrl.Result{}, err
	}
	deploy := r.deploymentForOperator(operator) //deploy variable is deployment template with values
	if !equality.Semantic.DeepDerivative(deploy.Spec.Template, found.Spec.Template) {
		found = deploy
		err := r.Update(ctx, found)
		if err != nil {
			log.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil

	} // first parameter if desired pod spec template and second one is observed pod spec template
	size := operator.Spec.Size
	if *found.Spec.Replicas != size {
		found.Spec.Replicas = &size
		err = r.Update(ctx, found)
		if err != nil {
			log.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil

	}

	foundService := &corev1.Service{} 
	// The foundService variable is declared as a pointer to corev1.Service object. Service object belongs to Core API group, so use core go package.
	err = r.Get(ctx, types.NamespacedName{Name: operator.Name, Namespace: operator.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		dep := r.ServiceForOperator(operator) // dep is service template with values 
		// ServiceForOperator -> service template
		// operator -> passing service spec field values to service template
		log.Info("Creating a new service", "Service.Namespace", dep.Namespace, "Service.Name", dep.Name)
		err := r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create a new service", "Service.Namespace", dep.Namespace, "Service.Name", dep.Name)
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get service")
		return ctrl.Result{}, err
	}

	podList != &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(found.Namespace),
		client.MatchingLabels{map[string]string{"app": found.Name, "labels", found.Name}}
	}
	if err = r.List(ctx, podList, listOpts...); err!= nil {
		log.Error(err, "Failed to list Pods", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)
		return ctrl.Result{}, err
	}

	podNames := getPodNames(podList.Items)

	if reflect.DeepEqual(podNames, operator.Status.PodList)
}

func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod : range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}
// SetupWithManager sets up the controller with the Manager.
// This function is used to setup the controller with the Kubernetes manager which is responsible for running the controller in the Kubernetes cluster.
// This function creates a new controller instance, registers it with the "MyPythonApp" resource types and returns an error if there was an issue registering the controller.
func (r *MyPythonAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mycustomappv1.MyPythonApp{}).
		Complete(r)
}

// This code in this file will provide the basic framework for the Kubernetes controller that can handle the myPythonApp CRD
