# Instalación Kubernetes

Para esta sección se han tenido en cuenta diferentes *"distribuciones"* de Kubernetes (k8s) entre ellas [minikube](https://minikube.sigs.k8s.io/docs/), [Kubeadm](https://kubernetes.io/docs/reference/setup-tools/kubeadm/), para este curso se escogio la de [Microk8s](https://microk8s.io/), por ser  una de las mas simples de instalar, y por tener el respaldo de canonical. 

Esta sección es una guía de instalación de Microk8s adaptada para el sistema operativo Oracle Linux 8 

## Instalación de SNAP (manejador de paquetes)



```bash
# Agregar los paquetes necesarios 
sudo dnf install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

# Actualizar los paquetes del sistema operativo 
sudo dnf upgrade

# instalar el manejador de paquetes de snap (desarrollado por canonical)
sudo yum install snapd

# vincular el deamon de snapd al reboot del sistema operativo (enable)
# e iniciar el servicio 
sudo systemctl enable --now snapd.socket

# crear un enlace simbolico del comando al filesystem /snap 
sudo ln -s /var/lib/snapd/snap /snap

# Actualizar variables de entorno del sistema operativo 
# en caso de ver un error del tipo "operation too quickly" 
# ejecute nuevamente este comando ... 
sudo ldconfig
```




## Instalación de Mikrok8s

```bash
sudo snap install microk8s --classic
sudo usermod -a -G microk8s $USER
newgrp microk8s
```

### Validar instalación

```bash
microk8s status --wait-ready
```

### Habilitar autocompletado

```bash
source <(kubectl completion bash) 
echo "source <(kubectl completion bash)" >> ~/.bashrc
```




