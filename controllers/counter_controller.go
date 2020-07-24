/*


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

package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	appv1beta1 "github.com/noroth/simple-status-operator/api/v1beta1"
)

// CounterReconciler reconciles a Counter object
type CounterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// Reconcile will be always called when a Resource was added or deleted
// +kubebuilder:rbac:groups=app.noroth.io,resources=counters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=app.noroth.io,resources=counters/status,verbs=get;update;patch
func (r *CounterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("counter", req.NamespacedName)

	c := appv1beta1.Counter{}

	log.Info("In reconciling function for ", "resource", req.Name)

	if err := r.Get(ctx, req.NamespacedName, &c); err != nil {
		log.Info("Could not find resource")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if c.Spec.MinimumCount > c.Spec.MaximumCount {
		r.Delete(ctx, &c)
		return ctrl.Result{},
			client.IgnoreNotFound(
				errors.New(
					"The mininmum counter must be smaller than the maximum counter"))
	}

	// set the current value if the status was not set yet
	if c.Status.Current <= c.Spec.MinimumCount && c.Status.Description != appv1beta1.Active {
		c.Status.Current = c.Spec.MinimumCount
		c.Status.Description = appv1beta1.Active

		r.Status().Update(ctx, &c)
		return reconcile.Result{}, nil
	}

	if c.Status.Current < c.Spec.MaximumCount {
		time.Sleep(time.Second * c.Spec.CounterDelay)
		c.Status.Current++
		r.Status().Update(ctx, &c)

		return reconcile.Result{}, nil
	}

	c.Status.Description = appv1beta1.Disabled
	r.Status().Update(ctx, &c)

	return ctrl.Result{}, nil
}

// SetupWithManager does stuff
func (r *CounterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1beta1.Counter{}).
		Complete(r)
}
