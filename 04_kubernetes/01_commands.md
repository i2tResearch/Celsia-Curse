# Comandos de resumen

## Interactuar con el comando kubectl

crear un pod partiendo d euna imagne en dockerhub

```bash
kubectl run <POD_NAME> --image=<IMAGE_NAME> --port=<LOCAL_PORT> <REMOTE_PORT>
```

Listar todos los pod del namespace por defecto

```bash
kubectl get pods
```

Obtener detalles del Pod

```bash
kubectl describe pod kbillingapp
```

Crear el servicio

```bash
kubectl expose pod <POD_NAME> --type=LoadBalancer --port <PORT_SERVICE> --target-port=<POD_PORT>
```

Obtener los logs de un servicio

```bash
kubectl logs <POD_NAME>
```

Eliminar un pod

```bash
kubectl delete service <POD_NAME>
```

Eliminar un servicio

```bash
kubectl delete service <SERVICE_NAME>
```

Consultar la version del apiserver

```bash
kubectl api-versions
```

Verificar namespace 

```bash
kubectl get namespace
```

Crear el namespace si no existe

```bash
kubectl create namespace <NAMESPACE_NAME>
```

verificar los pods del namespace monitoring

```bash
kubectl get all --namespace=<NAMESPACE_NAME>
```

Verificar las versiones del API que soporta `kubectl`

```bash
kubectl api-versions
```
