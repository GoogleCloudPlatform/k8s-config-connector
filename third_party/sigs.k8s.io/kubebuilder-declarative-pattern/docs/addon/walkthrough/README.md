## Walkthrough: Creating a new Operator

This walkthrough is for creating an operator to run the [guestbook](https://github.com/kubernetes/examples/tree/master/guestbook) which is an example application for kubernetes.

### Basics

Install the following depenencies:

- [kubebuilder](https://book.kubebuilder.io/quick-start.html#installation) (tested with 3.1.0)
- [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) (tested with v3.8.7)
- docker
- kubectl
- golang (1.13 <= version < 1.17 for go modules)

Create a new directory and use kubebuilder to scaffold the operator:

```
export GO111MODULE=on
mkdir -p guestbook-operator/
cd guestbook-operator/
kubebuilder init --plugins go.kubebuilder.io/v3,declarative.go.kubebuilder.io/v1 --domain example.org --license apache2 --owner "TODO($USER): assign copyright" --repo example.org/guestbook-operator
```

### Adding our first CRD

```
# generate the API/controllers
kubebuilder create api --controller=true --group=addons --kind=Guestbook --make=false --namespaced=true --resource=true --version=v1alpha1
# remove the  test suites that are more checking that kubebuilder is working
find . -name "*_test.go" -delete
```


This creates API type definitions under `api/v1alpha1/`, and a basic
controller under `controllers/`

* Generate code: `make generate`

* You should now be able to `go run main.go` (or `make run`),
  though it will exit with an error from being unable to find the guestbook CRD.

### Adding a manifest

The addon operator pattern is based on declarative manifests; the framework is
able to load the manifests and apply them. Today we exec `kubectl apply`, but
when [server-side-apply](https://github.com/kubernetes/enhancements/issues/555)
is available we'll use that.

We suggest that even advanced operators use a manifest for their core objects.
It's always possible to manipulate the manifest before applying it (eg, adding labels,
changing namespaces, and tweaking flags)

Some other advantages:

* Working with manifests lets us release a new guestbook version without needing
  a new operator version
* The declarative manifest makes it easier for users to understand what is
  changing in each version
* It should result in less / simpler code

For now, we embed the manifests into the image, but we'll be evolving this, for example sourcing manifests from a bundle or over https.

Create a manifest under `channels/packages/<packagename>/<version>/manifest.yaml`

```bash
mkdir -p channels/packages/guestbook/0.1.0/
wget -O channels/packages/guestbook/0.1.0/manifest.yaml https://raw.githubusercontent.com/kubernetes/examples/master/guestbook/all-in-one/guestbook-all-in-one.yaml
```

We have a notion of "channels", which is a stream of updates.  We'll have
settings to automatically update or prompt-for-update when the channel updates.
Currently if you don't specify a channel in your CRD, you get the version
currently in the stable channel.

We need to define the default stable channel, so create `channels/stable`:

```bash
cat > channels/stable <<EOF
manifests:
- name: guestbook
  version: 0.1.0
EOF
```

### Using the framework in the controller

We replace the controller code `controllers/guestbook_controller.go`:

We are delegating most of the logic to `declarative.Reconciler`

```go
package controllers

/*
Copyright 2019 The Kubernetes Authors.

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

import (
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative"

	api "example.org/guestbook-operator/api/v1alpha1"
)

var _ reconcile.Reconciler = &GuestbookReconciler{}

// GuestbookReconciler reconciles a Guestbook object
type GuestbookReconciler struct {
	declarative.Reconciler
	client.Client
	Log logr.Logger
	Scheme *runtime.Scheme

	watchLabels declarative.LabelMaker
}

func (r *GuestbookReconciler) setupReconciler(mgr ctrl.Manager) error {
	labels := map[string]string{
		"example-app": "guestbook",
	}

	r.watchLabels = declarative.SourceLabel(mgr.GetScheme())

	return r.Reconciler.Init(mgr, &api.Guestbook{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(r.watchLabels),
		declarative.WithStatus(status.NewBasic(mgr.GetClient())),
		declarative.WithApplyPrune(),
		declarative.WithReconcileMetrics(0, nil),
	)
}

func (r *GuestbookReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := r.setupReconciler(mgr); err != nil {
		return err
	}

	c, err := controller.New("guestbook-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Guestbook
	err = c.Watch(&source.Kind{Type: &api.Guestbook{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to deployed objects
	_, err = declarative.WatchChildren(declarative.WatchChildrenOptions{Manager: mgr, Controller: c, Reconciler: r, LabelMaker: r.watchLabels})
	if err != nil {
		return err
	}

	return nil
}

// for WithApplyPrune
// +kubebuilder:rbac:groups=*,resources=*,verbs=list

// +kubebuilder:rbac:groups=addons.example.org,resources=guestbooks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=addons.example.org,resources=guestbooks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;delete;patch
// +kubebuilder:rbac:groups=apps;extensions,resources=deployments,verbs=get;list;watch;create;update;delete;patch
```

The important things to note here:

```go
	r.Reconciler.Init(mgr, &api.Guestbook{}, ...)
```

We bind the `api.Guestbook` type to the `guestbook` package in our `channels`
directory and pull in optional features of the declarative library.

Because api.Guestbook implements `addon.CommonObject` the
framework is then able to access CommonSpec and CommonStatus above, which
includes the version specifier.

### Misc

1. Add an import and init call to the top of the main() function in `main.go`:

	```go
	import (
		//..
		"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
	)
	func main() {
		// after: ctrl.SetLogger(zap.Logger(true))
		addon.Init()
	}
	```

### Testing it locally

We can register the Guestbook CRD and a Guestbook object, and then try running
the controller locally.

1) We need to generate and register the CRDs:

```bash
make install
```

You can verify that the CRD is regesterd successfully:

```
kubectl get crds guestbooks.addons.example.org
```

2) Create a guestbook CR:

Remove `spec.foo` key in `config/samples/addons_v1alpha1_guestbook.yaml` if exists.

```bash
kubectl apply -n kube-system -f config/samples/addons_v1alpha1_guestbook.yaml
```

You can verify the CR is created successfully:

```
kubectl get Guestbooks -n kube-system
```

3) You should now be able to run the controller using:

`make run`

You should see your operator apply the manifest.  You can then control-C and you
should see the deployment etc that the operator has created.

e.g. `kubectl get pods -n default` or
`kubectl get deploy -n default`.

## Running on-cluster

Previously we were running on your machine using your kubernetes credentials.
We want to run as a Pod on the cluster for real world operator. For that,
we'll need a Docker image and some manifests.

### Building the operator image

1. Modify the IMG value in the `Makefile` to reflect a docker registry that you
   can write to:

```make
# Image URL to use all building/pushing image targets
IMG ?= gcr.io/<my-cool-project>/guestbook-operator:latest
```

2. Create a patch to modify the memory limit for the operator:

```bash
cat << EOF > config/default/manager_resource_patch.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: controller-manager
  namespace: system
spec:
  template:
    spec:
      containers:
      - name: manager
        resources:
          limits:
            cpu: 100m
            memory: 150Mi
          requests:
            cpu: 100m
            memory: 20Mi
EOF
```

3. Reference the patch by adding `manager_resource_patch.yaml` to the `patches` section of `config/default/kustomization.yaml`:

```yaml
patchesStrategicMerge:
- manager_resource_patch.yaml
# ... existing patches
```

This is requried to run kubectl in the container.

4. Modify the `Dockerfile` to pull in kubectl, the manifests (in `channels/`),
   and run in a slim container:

```Dockerfile
# Build the manager binary
FROM golang:1.17 as builder

# Copy in the go src
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
COPY vendor/   vendor/

COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY channels/ channels/
RUN chmod -R a+rx channels/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Copy the operator and dependencies into a thin image
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /workspace/channels/ channels/

USER 65532:65532

ENTRYPOINT ["./manager"]
```

5. Modify the `Makefile` to run `go mod vendor` for building container image.

```make
...
# Add vendor prerequisites before test.
docker-build: vendor test
	docker build . -t ${IMG}

# Add vendor target.
vendor:
	go mod vendor
...
```

6. Verify everything worked by building and pushing the image:

```bash
make docker-build docker-push
```

### Generated RBAC Rules

We need a simple deployment to run our operator, and we want to run it under a
tightly-scoped RBAC role. To do that we use kubebuilder's RBAC role generation
based off of source annotations. In the future we may be able to generate RBAC
rules from the manfiest.

The RBAC rules are included in the `guestbook_controller.go` snippet you pasted above.

RBAC is the real pain-point here - we end up with a lot of permissions:
* The operator needs RBAC rules to see the CRDs.
* It needs permission to get / create / update the Deployments and other types
  that it is managing
* It needs permission to create the ClusterRoles / Roles that the guestbook
  needs
* Because of that, we also need permissions for all the permissions we are going
  to create.

The last one in particular can result in a non-trivial RBAC policy.  My approach:

* Start with minimal permissions (just watching addons.k8s.io guestbooks), and
  then add permissions iteratively
* If you're going to allow list, I tend to just allow get, list and watch -
  there's not a huge security reason to treat them separately as far as I can
  see
* Similarly I treat create and patch together
* No controller should be using update (because of version skew issues), so I
  tend to grant that one begrudgingly
* The RBAC policy in the manifest may scope down the permissions even more (for
  example scoping to resourceNames), in which case we can - and should - copy
  it.  That's what we did here for guestbook.

### Installing the operator in the cluster

```bash
# install the CRD, and start the operator
make deploy
```

You can troubleshoot the operator by inspecting the controller:

```bash
kubectl -n guestbook-operator-system get deploy
kubectl -n guestbook-operator-system logs <guestbook-operator-controller-manager-pod-name> manager
```

### Create a guestbook CR

```bash
kubectl apply -n kube-system -f config/samples/addons_v1alpha1_guestbook.yaml
```

You can verify the CR is created successfully:

```
kubectl get Guestbooks -n kube-system
```

You can verify that the operator has created the `kubernetes-guestbook`
deployment:

e.g. `kubectl get pods -n guestbook-operator-system` or
`kubectl get deployments -n guestbook-operator-system`.

## Manifest simplification: Automatic labels

Similar to how kustomize works, often you won't want labels hard-coded in the
manifest, but will use them to distinguish multiple instances.  Even if you're
writing something you expect to be a singleton instance, it can be tedious and
error-prone to specify labels on every object in the manifest.

Instead, the Reconciler can add labels to every object in the manifest:

```go
       labels := map[string]string{
               "example-app": "guestbook",
       }

       r := &ReconcileGuestbook{}
       r.Reconciler.Init(mgr, &api.Guestbook{}, "guestbook",
               declarative.WithObjectTransform(declarative.AddLabels(labels)),
			   ...
	   )
```

**NOTE**: operators.AddLabels does not [_yet_](https://github.com/kubernetes-sigs/kubebuilder-declarative-pattern/issues/21)
add selectors to Deployments/DaemonSets, nor to the templates.

## Manifest simplification: Automatic Namespace

The framework automatically creates objects in the same namespace as
the CR (by specifying the namespace to kubectl).  As such, we can remove the
namespaces from the manifest.

NOTE: We don't currently apply the namespace within objects.  For example, we
don't set the namespace on a RoleBinding subjects.namespace.  However, it seems
that most objects default to the same namespace - but presumably
ClusterRoleBinding will not.

NOTE: For non-namespaces objects (ClusterRole and ClusterRoleBinding), we often
need to name them with the namespace to support multiple instances.

### Manage an Application

The framework can manage an [application](https://github.com/kubernetes-sigs/application)
instance. The application contains human readable information in addition to deployment
status that can be surfaced in various user interfaces.

1. Fetch the Application CRD and place it with your operators CRD:

	```bash
	curl https://raw.githubusercontent.com/kubernetes-sigs/application/master/config/crds/app_v1beta1_application.yaml -o config/crd/app_v1beta1_application.yaml
	```

2. Add an instance of the Application CR in your manifest:

	```bash
	cat <<EOF >> channels/packages/guestbook/0.1.0/manifest.yaml
	# ------------------- Application ------------------- #
	apiVersion: app.k8s.io/v1beta1
	kind: Application
	metadata:
	  name: guestbook
	spec:
	  descriptor:
	    type: "guestbook"
	    description: "Guestbook is a simple, multi-tier web application using Kubernetes. This application consists of the following components: A single-instance Redis master to store guestbook entries, Multiple replicated Redis instances to serve reads, Multiple web frontend instances."
	    icons:
	    - src: "https://github.com/kubernetes/kubernetes/raw/master/logo/logo.png"
	      type: "image/png"
	    maintainers:
	    - name: Maintainer
	      email: maintainer@example.org
	    keywords:
	    - "addon"
	    - "guestbook"
	    links:
	    - description: Guide Document
	      url: "https://kubernetes.io/docs/tutorials/stateless-application/guestbook/"
	    - description: Source Code
	      url: "https://github.com/kubernetes/examples/tree/master/guestbook"
	EOF
	```

3. Add the two options for managing the Application to your controller:

	```go
		r.Reconciler.Init(mgr, &api.Guestbook{},
			...
			declarative.WithManagedApplication(r.watchLabels),
			declarative.WithObjectTransform(addon.TransformApplicationFromStatus),
			...
		)
	```

4. Rebuild the operator, reinstall the CRDs, and start the new operator. You can now see the Application:

	```bash
	kubectl -n kube-system get applications -oyaml
	```

### Next steps

* Read about [adding tests](tests.md)
* Remove cruft from the manifest yaml (Namespaces, Names, Labels)
* Explore avaliable [options](https://godoc.org/sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/declarative)
