module github.com/min-charles/helm-server

go 1.16

require (
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-logr/logr v0.4.0
	github.com/gorilla/mux v1.8.0
	github.com/mittwald/go-helm-client v0.8.2
	helm.sh/helm/v3 v3.7.1 // indirect
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
	sigs.k8s.io/controller-runtime v0.10.3
)
