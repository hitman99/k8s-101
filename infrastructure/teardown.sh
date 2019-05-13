#!/bin/bash

function teardown() {
	echo

    kubectl delete -f k8s/03-certificate.yaml
	kubectl delete -f k8s/02-issuer.yaml
	kubectl delete -f k8s/00-cert-manager.yaml

	rId=`doctl compute domain records list k8s.adomavicius.com -o json | jq  '.[] |.|select(.type=="A").id'`
	if [[ $rId != "" ]]; then
		doctl compute domain records delete -f k8s.adomavicius.com ${rId}
	fi

	kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml
	kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml

	echo "done"

}




echo "This will remove resources created in setup"
read -s -p "Do you understand what you're doing?[y/n]" answer
if [[ $answer == "y" ]]; then
	teardown
else
	echo
	echo "bye"
	exit 0
fi
