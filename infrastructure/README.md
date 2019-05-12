# Infrastructure for this project

This is infrastructure setup for kubernetes-101 demo project. Follow steps below to get it running.

1.  Download [kubectl binary](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
2.  Create a Kubernetes cluster in the cloud.
    To create managed Kubernetes cluster in DigitalOcean cloud, follow the steps below.
    1.  create DigitalOcean account in https://digitalocean.com
    2.  [download](https://github.com/digitalocean/doctl) and initialize `doctl`
    3.  create cluster with 3 smallest nodes with `doctl`:
        ```
        doctl k8s cluster create kubernetes-101 --count 3 --size s-1vcpu-2gb --region fra1
        ``` 
    
3.  Once the cluster is up, provision ingress controller and cert-manager operator
    ```
    # NGINX ingress controller
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml
    
    # cert-manager operator
    kubectl apply -f https://raw.githubusercontent.com/jetstack/cert-manager/v0.7.2/deploy/manifests/00-crds.yaml
    kubectl apply -f https://raw.githubusercontent.com/jetstack/cert-manager/v0.7.2/deploy/manifests/01-namespace.yaml
    kubectl apply -f https://raw.githubusercontent.com/jetstack/cert-manager/v0.7.2/deploy/manifests/cert-manager.yaml
    
    # cert-manger config for demo app
    # ** This only works you have domain name configured in DigitalOcean, as this issuer is using DNS01 challenge
    
    kubectl apply -f k8s/01-issuer.yaml
    kubectl apply -f k8s/02-certificate.yaml 
    ```
    

4.  Deploy the demo application
    ```
    kubectl apply -f k8s/03-demo-app.yaml
    ```
5.  Celebrate