/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package ldap

import (
	"context"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	ldapv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/ldap/v1alpha1"
	"kubeform.dev/provider-mongodbatlas-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// VerifyReconciler reconciles a Verify object
type VerifyReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk              schema.GroupVersionKind // GVK of the Resource
	Provider         *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource         *tfschema.Resource      // returns *schema.Resource
	TypeName         string                  // resource type
	WatchOnlyDefault bool
}

// +kubebuilder:rbac:groups=ldap.mongodbatlas.kubeform.com,resources=verifies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=ldap.mongodbatlas.kubeform.com,resources=verifies/status,verbs=get;update;patch

func (r *VerifyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("verify", req.NamespacedName)

	if r.WatchOnlyDefault && req.Namespace != v1.NamespaceDefault {
		log.Info("Only default namespace is supported for Kubeform Community, Please upgrade to Kubeform Enterprise to use any namespace.")
		return ctrl.Result{}, nil
	}
	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch Verify")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(ldapv1alpha1.GetEncoder(), ldapv1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *VerifyReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &ldapv1alpha1.Verify{}); err != nil {
			klog.Error(err, "unable to set up auditor", ldapv1alpha1.Verify{}.APIVersion, ldapv1alpha1.Verify{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&ldapv1alpha1.Verify{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		Owns(&v1.Secret{}).
		Complete(r)
}
