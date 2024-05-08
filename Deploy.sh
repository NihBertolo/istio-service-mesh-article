# Execute apenas se já tiver o Kubernetes e Istio instalado em sua máquina
kubectl apply -f ./k8s/service_entry.yaml
kubectl apply -f ./k8s/destination_rule.yaml
kubectl apply -f ./k8s/virtual_service.yaml
kubectl apply -f ./k8s/deployment.yaml
kubectl apply -f ./k8s/service.yaml

printf "Ambiente configurado!"