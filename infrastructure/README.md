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
    
    # Wait for LoadBalancer to assign external IP and configure DNS A record
    # ** This will only work if you have domain name configured in DO
    doctl compute domain records create \
    --record-type A --record-name '@' \
    --record-data `kubectl get svc -n ingress-nginx -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}'` \
    k8s.adomavicius.com
    
    # cert-manager operator
    kubectl apply -f k8s/00-cert-manager.yaml
    
    # Create new Token/Key in DigitalOcean UI, paste the command below, then copy token and execute command (MacOS only)
    kubectl create secret generic digitalocean-dns --from-literal access-token=`pbpaste` --namespace cert-manager
    
    
    # cert-manger config for demo app
    # ** This will only work if you have domain name configured in DigitalOcean **
    
    kubectl apply -f k8s/01-issuer.yaml
    kubectl apply -f k8s/02-certificate.yaml 
    ```
    ###or
    
    run the [setup.sh](./setup.sh) script. For script to work you will need to have
    `doctl` and `jq` installed

4.  Deploy the demo application
    ```
    kubectl apply -f k8s/03-demo-app.yaml
    ```
5.  Celebrate