# Development Quickstart

This guide will assist you in setting up your development environment for NGINX Kubernetes Gateway, covering the steps
to build, install, and execute tasks necessary for submitting pull requests. By following this guide, you'll have a
fully prepared development environment that allows you to contribute to the project effectively.

## Setup Your Development Environment

Follow these steps to set up your development environment.

1. Install:
    - [Go](https://golang.org/doc/install)
    - [Docker](https://docs.docker.com/get-docker/) v18.09+
    - [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
    - [Helm](https://helm.sh/docs/intro/quickstart/#install-helm)
    - [git](https://git-scm.com/)
    - [GNU Make](https://www.gnu.org/software/software.html)
    - [yq](https://github.com/mikefarah/yq/#install)
    - [fieldalignment](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/fieldalignment):

      ```shell
      go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
      ```

    - [pre-commit](https://pre-commit.com/#install):

     ```shell
     brew install pre-commit
     ```

     and then run

     ```shell
     pre-commit install
     ```

     in the project root directory to install the git hooks.

2. [Fork the project repository](https://github.com/nginxinc/nginx-kubernetes-gateway/fork)
3. Clone your repository, and install the project dependencies:

   ```shell
   git clone https://github.com/<YOUR-USERNAME>/nginx-kubernetes-gateway.git
   cd nginx-kubernetes-gateway
   ```

   ```makefile
   make deps
   ```

## Build the Binary and Image

### Build the Binary

To build the binary, run the make build command from the project's root directory:

```makefile
make build
```

This command will build the binary and output it to the `/build/.out` directory.

### Build the Image

To build an NGINX Kubernetes Gateway container image from source run the following make command:

```makefile
make TAG=$(whoami) container
```

This will build the docker image `nginx-kubernetes-gateway:<your-user>`.

## Deploy on Kind

1. Create a `kind` cluster:

   ```makefile
   make create-kind-cluster
   ```

2. Load the previously built image onto your `kind` cluster:

   ```shell
   kind load docker-image nginx-kubernetes-gateway:$(whoami)
   ```

3. Install Gateway API Resources

   ```shell
   kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v0.7.1/standard-install.yaml
   ```

4. Install NKG using your custom image and expose NKG with a NodePort Service:

   - To install with Helm (where your release name is `my-release`):

      ```shell
      helm install my-release ./deploy/helm-chart --create-namespace --wait --set service.type=NodePort --set nginxGateway.image.repository=nginx-kubernetes-gateway --set nginxGateway.image.tag=$(whoami) --set nginxGateway.image.pullPolicy=Never -n nginx-gateway
      ```

      > For more information on helm configuration options see the Helm [README](/deploy/helm-chart/README.md).

   - To install with manifests:

      ```shell
      make generate-manifests HELM_TEMPLATE_COMMON_ARGS="--set nginxGateway.image.repository=nginx-kubernetes-gateway --set nginxGateway.image.tag=$(whoami) --set nginxGateway.image.pullPolicy=Never"
      kubectl apply -f deploy/manifests/nginx-gateway.yaml
      kubectl apply -f deploy/manifests/nodeport.yaml
      ```

### Run Examples

To make sure NKG is running properly, try out the [examples](/examples).

## Run the Unit Tests

To run all the unit tests, run the make unit-test command from the project's root directory:

```makefile
make unit-test
```

For more details on testing, see the [testing](/docs/developer/testing.md) documentation.

## Gateway API Conformance Testing

To run Gateway API conformance tests, please follow the instructions on [this](/conformance/README.md) page.

## Run the Linter

To lint the code, run the following make command from the project's root directory:

```makefile
make lint
```

> **Note**
> fieldalignment errors can be fixed by running: `fieldalignment -fix <path-to-package>`

## Run the Helm Linter

Run the following make command from the project's root directory to lint the Helm Chart code:

```shell
make lint-helm
```

## Run go generate

To ensure all the generated code is up to date, run the following make command from the project's root directory:

```shell
make generate
```

## Update Generated Manifests

To update the NJS ConfigMap yaml, run the following make command from the project's root directory:

```shell
make generate-njs-yaml
```

Additionally, the [NJS ConfigMap Helm template](/deploy/helm-chart/templates/njs-modules.yaml) will need to be updated.
This is currently a manual process - ensure the content in the `data` field matches that in the
[NJS ConfigMap manifest](/deploy/manifests/njs-modules.yaml) `data` field.

Finally, to update all other generated manifests, run the following make command from the project's root directory:

```shell
make generate-manifests
```
