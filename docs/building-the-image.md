# Building the Image

## Prerequisites

Before you can build the NGINX Kubernetes Gateway, make sure you have the following software installed on your machine:

- [git](https://git-scm.com/)
- [GNU Make](https://www.gnu.org/software/software.html)
- [Docker](https://www.docker.com/) v18.09+

## Steps

1. Clone the repo and change into the `nginx-kubernetes-gateway` directory:

   ```shell
   git clone https://github.com/nginxinc/nginx-kubernetes-gateway.git
   cd nginx-kubernetes-gateway
   ```

1. Build the image:

   ```makefile
   make PREFIX=myregistry.example.com/nginx-kubernetes-gateway container
   ```

   Set the `PREFIX` variable to the name of the registry you'd like to push the image to. By default, the image will be
   named `nginx-kubernetes-gateway:edge`.

1. Push the image to your container registry:

   ```shell
   docker push myregistry.example.com/nginx-kubernetes-gateway:edge
   ```

   Make sure to substitute `myregistry.example.com/nginx-kubernetes-gateway` with your registry.
