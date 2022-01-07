package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/min-charles/helm-server/pkg/apis"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("helm-server")

func main() {

	log.Info("initializing server....")

	router := mux.NewRouter()

	// c, err := internal.Client(client.Options{})
	// if err != nil {
	// 	panic(err)
	// }

	// create := apis.Create{
	// 	Client: c,
	// 	Log:    logf.Log.WithName("Create"),
	// }
	router.HandleFunc("/helm", apis.InstallRelease).Methods("POST", "PUT")
	router.HandleFunc("/helm", apis.UnInstallRelease).Methods("DELETE")
	router.HandleFunc("/helm", apis.RollbackRelease).Methods("PATCH")

	http.Handle("/", router)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8081), nil); err != nil {
		log.Error(err, "failed to initialize a server")
	}

}
