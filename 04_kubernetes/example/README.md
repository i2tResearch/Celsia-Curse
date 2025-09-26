# Despliegue de una base de datos:

Definición de los secrets: 

Para postgresql: 

```bash
apiVersion: v1
kind: Secret
metadata:
  name: postgres-secret
  labels:
    app: postgres
type: Opaque
data:
  POSTGRES_DB: cG9zdGdyZXM= # Base 64 code 
  POSTGRES_USER: cG9zdGdyZXM= # Base 64 code 
  POSTGRES_PASSWORD: cXdlcnR5
```

Para pgp-admin

```bash
apiVersion: v1
kind: Secret
metadata:
  name: pgadmin-secret
  labels:
    app: postgres
    # meant that we can use arbitrary key pair values
    # https://kubernetes.io/docs/concepts/configuration/secret/
type: Opaque
data:
  PGADMIN_DEFAULT_EMAIL: YWRtaW5AYWRtaW4uY29t
  PGADMIN_DEFAULT_PASSWORD: cXdlcnR5
  PGADMIN_PORT: ODA=
```

Definición de los persistence volumes 

Para la base de datos postgres: 

```bash
#persistence volumen (PV) is a piece of storage that have idependent lifecycle from pods 
#thees preserve data throug restartin, rescheduling and even deleting pods
#PersistenceVolumeCalin is a request for storage by the user that can be fulfilled by teh PV
kind: PersistentVolume
#version of ApiServer on control panel node (/api/v1) check using kubectl api-versions
apiVersion: v1
metadata:
  name: postgres-volume
  labels:
    #it is aplugin suport many luke amazon EBS azure disk etc.  local = local storage devices mounted on nodes.
    type: local
    app: postgres
spec:
  storageClassName: manual
  capacity:
    storage: 5Gi
    #many pods on shcheduled on differents nodes can read and write
  accessModes:
    - ReadWriteMany
    #path on cluster's node
  hostPath:
    path: "/mnt/data/"
```

Para reclamar un espacio del volumen:

```bash
#it is a reques of resource (persistence volume) from a pod by example, teh pod claim by a storage throug PVC
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: postgres-claim
  labels:
    app: postgres
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 2Gi
```

Definir el config-map

Para inicializar las bases de datos: 

```bash
apiVersion: v1
kind: ConfigMap
metadata:
  name: postgres-init-script-configmap
data:
  initdb.sh: |-
   #!/bin/bash
   set -e
   psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER generalapp WITH PASSWORD 'qwerty';
    CREATE DATABASE billingapp_db;
    GRANT ALL PRIVILEGES ON DATABASE generalapp_db TO generalapp;
   EOSQL
```

Definición de los deployments: 

Para postgress

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres-deployment
  labels:
    app: postgres
spec:
#Pods number replicates
  replicas: 1
  #Define how the Deployment identify the pods that it could manage
  selector: 
    matchLabels:
     app: postgres
     #pod template specification
  template:
    metadata:
    #define teh labels for all pods
      labels:
        app: postgres       
    spec:
      containers:
        - name: postgres
          image: postgres:latest
          imagePullPolicy: IfNotPresent
          #open the port to allow send and receive traffic in teh container
          ports:
            - containerPort: 5432
            #read envars from secret file
          envFrom:
            - secretRef:
                name: postgres-secret
          volumeMounts:
          #This is the path in the container on which the mounting will take place.
            - mountPath: /var/lib/postgresql/data
              name: postgredb
            - mountPath: /docker-entrypoint-initdb.d
              name : init-script
      volumes:
        - name: postgredb
          persistentVolumeClaim:
            claimName: postgres-claim
        - name: init-script
          configMap:
             name: postgres-init-script-configmap
```

Para pgpadmin: 

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: pgadmin-deployment
spec:
  selector:
   matchLabels:
    app: pgadmin
  replicas: 1
  template:
    metadata:
      labels:
        app: pgadmin
    spec:
      containers:
        - name: pgadmin4
          image: dpage/pgadmin4        
          envFrom:
            - secretRef:
                name: pgadmin-secret
          ports:
            - containerPort: 80
              name: pgadminport
```

Definimos los services: 

Para postgres

```bash
kind: Service
apiVersion: v1
metadata:
  name: postgres-service
  labels:
    app: postgres
spec:  
  ports:
  - name: postgres
    port: 5432
    nodePort : 30432 
  #type: LoadBalancer
  type: NodePort
  selector:
   app: postgres
```

Para pgadmin

```bash
apiVersion: v1
kind: Service
metadata:
  name: pgadmin-service
  labels:
    app: pgadmin
spec:
  selector:
   app: pgadmin
  type: NodePort
  ports:
   - port: 80
     nodePort: 30200
```
