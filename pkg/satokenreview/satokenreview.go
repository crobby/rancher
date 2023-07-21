package satokenreview

import (
	"errors"
	"net/http"

	authV1 "k8s.io/api/authorization/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/rancher/rancher/pkg/auth/util"
	"github.com/rancher/rancher/pkg/clustermanager"
	"github.com/rancher/rancher/pkg/types/config"
)

type saTokenReviewHandler struct {
	clusterManager *clustermanager.Manager
	k8sClient      kubernetes.Interface
	next           http.Handler
}

type RequestData struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}

func NewSATokenHandler(scaledContext *config.ScaledContext, clusterManager *clustermanager.Manager, handler http.Handler) http.Handler {
	return &saTokenReviewHandler{
		clusterManager: clusterManager,
		k8sClient:      scaledContext.K8sClient,
		next:           handler,
	}
}

func (h *saTokenReviewHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	var reqGroup []string
	if g, ok := req.Header["Impersonate-Group"]; ok {
		reqGroup = g
	}

	review := authV1.SubjectAccessReview{
		Spec: authV1.SubjectAccessReviewSpec{
			User:   req.Header.Get("Impersonate-User"),
			Groups: reqGroup,
			ResourceAttributes: &authV1.ResourceAttributes{
				Verb:     "get",
				Resource: "ranchermetrics",
				Group:    "management.cattle.io",
			},
		},
	}

	result, err := h.k8sClient.AuthorizationV1().SubjectAccessReviews().Create(req.Context(), &review, metav1.CreateOptions{})
	if err != nil {
		util.ReturnHTTPError(rw, req, 500, err.Error())
		return
	}

	if !result.Status.Allowed {
		util.ReturnHTTPError(rw, req, 401, "Unauthorized")
		return
	}

	//if id := mux.Vars(req)["clusterID"]; id != "" {
	//	h.getClusterObjectCount(id, rw, req)
	//	return
	//}

	h.next.ServeHTTP(rw, req)
}

// returnK8serror attempts to respond using the k8s error and its code/message
func returnK8serror(err error, rw http.ResponseWriter, req *http.Request) {
	var k8sError *k8serrors.StatusError
	if errors.As(err, &k8sError) {
		util.ReturnHTTPError(rw, req, int(k8sError.ErrStatus.Code), k8sError.Error())
		return
	}
	// Well, it wasn't a k8s error, give it back
	util.ReturnHTTPError(rw, req, 500, err.Error())
}
