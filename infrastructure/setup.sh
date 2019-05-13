#!/bin/bash

function get_loadbalancer_ip() {
	echo `kubectl get svc -n ingress-nginx -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}'`
}

function setup() {
	echo
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml

	ip=""
	while [[ $ip == "" ]]; do
		sleep 5
		ip=`get_loadbalancer_ip`
	done

	doctl compute domain records create \
    --record-type A --record-name '@' \
    --record-data ${ip} \
    k8s.adomavicius.com

	read -s -p "Paste DO Token" dotoken

	kubectl apply -f k8s/00-cert-manager.yaml
	kubectl create secret generic digitalocean-dns --from-literal access-token=${dotoken} --namespace cert-manager

	sleep 20
	kubectl apply -f k8s/01-namespace.yaml
	kubectl apply -f k8s/02-issuer.yaml
    kubectl apply -f k8s/03-certificate.yaml

	echo "done"
}




echo "This will setup Kubernetes 101 demo in active cluster. You should have the cluster running already
This is similar to running all commands in readme.md"
read -s -p "Do you understand what you're doing?[y/n]" answer
if [[ $answer == "y" ]]; then
	setup
else
	echo
	echo "bye"
	exit 0
fi
