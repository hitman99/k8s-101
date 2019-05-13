.PHONY: design
design:
	mv vendor _vendor
	goagen app --design github.com/hitman99/k8s-101/demo-app/design
	goagen swagger --design github.com/hitman99/k8s-101/demo-app/design
	goagen client --design github.com/hitman99/k8s-101/demo-app/design
	goagen controller --pkg controllers --design github.com/hitman99/k8s-101/demo-app/design -o controllers
	mv _vendor vendor

build:
	go build . -o demoapp