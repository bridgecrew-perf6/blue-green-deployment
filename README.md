# Blue/Green Deployment K8s

## Prepara o pacote v1.0.0 e envia para Minikube

```bash
docker build -t api-version:1.0.0 v1/.
```

```bash
docker save api-version:1.0.0 | (eval $(minikube docker-env) && docker load)
```

## Configurando o deployment e service - v1.0.0

```bash
kubectl apply -f v1/api-version-deployment.yaml
```

```bash
kubectl apply -f api-version-service.yaml
```

## Resultado /api/version - v1.0.0

```bash
curl $(minikube service api-version --url)/api/version
```

```json
{"id":"1.0.0", "name":"Version One"}
```

## Prepara o pacote v2.0.0 e envia para Minikube

```bash
docker build -t api-version:2.0.0 v2/.
```

```bash
docker save api-version:2.0.0 | (eval $(minikube docker-env) && docker load)
```

## Configurando o deployment - v2.0.0

```bash
kubectl apply -f v2/api-version-deployment.yaml
```

## Alternado matchs version no service para v2.0.0

```bash
kubectl patch service api-version -p '{"spec":{"selector":{"version":"v2.0.0"}}}'
```

## Resultado /api/version - v2.0.0

```bash
curl $(minikube service api-version --url)/api/version
```

```json
{"id":"2.0.0","name":"Version Two", "describe": "Describe New Version"}
```

## Rollback para v1.0.0

```bash
kubectl patch service api-version -p '{"spec":{"selector":{"version":"v1.0.0"}}}'
```
