/*
Copyright 2024.

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
	"fmt"

	core1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	muhu3k8siov1 "github.com/muhyousri/ConfigMapSyncer/api/v1"
)

// SyncConfigReconciler reconciles a SyncConfig object
type SyncConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=muhu3.example.com,resources=syncconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=muhu3.example.com,resources=syncconfigs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=muhu3.example.com,resources=syncconfigs/finalizers,verbs=update
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;delete;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SyncConfig object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *SyncConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var syncConfig muhu3k8siov1.SyncConfig
	var testconfig core1.ConfigMap

	if err := r.Get(ctx, req.NamespacedName, &syncConfig); err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("Resource is deleted, Deleting ConfigMap.. \n ")
			fmt.Println(req.Name)
			fmt.Println(req.Namespace)
			cm_deleted := core1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      req.Name,
					Namespace: req.Namespace,
				},
			}

			r.Delete(ctx, &cm_deleted)
			fmt.Printf("Deleted! \n ")

		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	fmt.Println("Got a new SyncConfig")

	name := syncConfig.Name
	nameSpace := syncConfig.Namespace
	syncData := syncConfig.Spec.Data

	cmKey := client.ObjectKey{
		Name:      name,
		Namespace: nameSpace,
	}

	fmt.Println("cmKey:", cmKey)
	fmt.Println("syncData:", syncData)

	if err := r.Get(ctx, cmKey, &testconfig); err != nil {
		if errors.IsNotFound(err) {
			fmt.Println("ConfigMap doesn't exist, creating...")

			testconfig = core1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: nameSpace,
				},
				Data: syncData,
			}
			if err := r.Create(ctx, &testconfig); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil // Return after creating
		}

		return ctrl.Result{}, err
	}

	fmt.Println("ConfigMap exists, patching...")
	newConfig := testconfig.DeepCopy()
	newConfig.Data = syncData // Modify data in newConfig

	patch := client.MergeFrom(&testconfig) // Use testconfig as base for the patch
	if err := r.Patch(ctx, newConfig, patch); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SyncConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&muhu3k8siov1.SyncConfig{}).
		Named("syncconfig").
		Complete(r)
}
