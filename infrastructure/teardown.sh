#!/bin/bash

DOMAIN=${1}

function teardown() {
	echo

    kubectl delete -f k8s/03-certificate.yaml
	kubectl delete -f k8s/02-issuer.yaml
	kubectl delete -f k8s/00-cert-manager.yaml

	rId=`doctl compute domain records list ${DOMAIN} -o json | jq  '.[] |.|select(.type=="A").id'`
	if [[ $rId != "" ]]; then
		doctl compute domain records delete -f ${DOMAIN} ${rId}
	fi

	kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml
	kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml

	echo "done"

}



if [[ $DOMAIN == "" ]]; then
	echo
	echo "domain name missing as first argument"
	echo "Usage: $0 <domain name (example.com)>"
	echo
	exit 1
fi

echo "This will remove resources created in setup"
read -s -p "Do you understand what you're doing?[y/n]" answer
if [[ $answer == "y" ]]; then
	teardown
else
	echo
	echo "bye"
	exit 0
fi
