-include ../setup-env.mk
-include ../helm-cmd.mk
HELM_CMD ?= helm upgrade --install

deploy:
	${HELM_CMD} istio \
		deploy/helm/istio/ \
		--namespace aks-istio-system
.PHONY: deploy

undeploy:
	helm uninstall istio
.PHONY: undeploy
