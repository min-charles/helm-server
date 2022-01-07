package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/min-charles/helm-server/internal"
	"github.com/min-charles/helm-server/pkg/schemas"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	helmclient "github.com/mittwald/go-helm-client"
)

// type Create struct {
// 	client.Client
// 	Log logr.Logger
// }

func UnInstallRelease(w http.ResponseWriter, r *http.Request) {
	req := schemas.ReleaseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Log.Error(err, "failed to decode request")
		return
	}
	///////////////////////////
	cfg, err := config.GetConfig()
	if err != nil {
		// Log.Error(err, "failed to get rest config")
		fmt.Println("failed to get rest config")
	}

	c, _ := client.New(cfg, client.Options{})

	sa, _ := internal.GetServiceAccount(c, types.NamespacedName{Name: "helm-server-test-sa", Namespace: "helm-ns"})
	var secretName string

	for _, sec := range sa.Secrets {
		secretName = sec.Name
	}

	testSecret, _ := internal.GetSecret(c, types.NamespacedName{Name: secretName, Namespace: "helm-ns"})
	token := testSecret.Data["token"]

	opt := &helmclient.Options{
		RepositoryCache:  "/tmp/.helmcache",
		RepositoryConfig: "/tmp/.helmrepo",
		Debug:            true,
		Linting:          true,
	}

	cfg.BearerToken = string(token)
	cfg.BearerTokenFile = ""

	helmClient, err := helmclient.NewClientFromRestConf(&helmclient.RestConfClientOptions{Options: opt, RestConfig: cfg})
	if err != nil {
		// Log.Error(err, "failed to create helm client")
		fmt.Println("failed to create helm client")
	}

	path := "/tmp/test"

	chartSpec := helmclient.ChartSpec{
		ReleaseName: req.Spec.ReleaseName,
		ChartName:   path + req.Spec.Path,
		Namespace:   req.Namespace,
		ValuesYaml:  req.Values,
		Version:     req.Spec.Version,
		UpgradeCRDs: true,
		Wait:        false,
	}

	if err := helmClient.UninstallRelease(&chartSpec); err != nil {
		// Log.Error(err, "failed to uninstall chart")
		fmt.Println(err, "failed to uninstall chart")
	}

	////////////////
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(""); err != nil {
		Log.Error(err, "failed to encode response")
	}
}
