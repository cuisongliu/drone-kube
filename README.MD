[![Build Status](https://cloud.drone.io/api/badges/cuisongliu/drone-kube/status.svg)](https://cloud.drone.io/cuisongliu/drone-kube)

# drone-kube

## config

> config admin kubeconfig, from env generator /root/.kube/config ,the user controller k8s cluster

### command generator 


```bash
 drone-kube config  --admin xxx  --admin-key xxx --ca xxx  --server xxx
```

### use docker  generator

```bash
 docker run -ti --network=host -e KUBE_SERVER=https://xxx:6443  -e KUBE_CA=xxx  -e KUBE_ADMIN=xxx -e KUBE_ADMIN_KEY=xxx cuisongliu/drone-kube bash 
 drone-kube config
```

| describe | env |
| :--- | :---  |  
| ~/.kube/config  server | KUBE_SERVER , PLUGIN_SERVER , PLUGIN_KUBE_SERVER|  
| ~/.kube/config certificate-authority-data | KUBE_CA, PLUGIN_CA, PLUGIN_KUBE_CA|  
| ~/.kube/config client-certificate-data | KUBE_ADMIN, PLUGIN_ADMIN, PLUGIN_KUBE_ADMIN |  
| ~/.kube/config client-key-data | KUBE_ADMIN_KEY, PLUGIN_ADMIN_KEY, PLUGIN_KUBE_ADMIN_KEY |  

### use drone generator 
1. use no prefix classpath
    ```yaml
    - name: deploy-font
      image: cuisongliu/drone-kube
      settings:
        server:
          from_secret: k8s-server
        ca:
          from_secret: k8s-ca
        admin:
          from_secret: k8s-admin
        admin_key:
          from_secret: k8s-admin-key
      commands:
        - drone-kube config  >> /dev/null
        - kubectl delete -f deploy/deploy.yaml || true
        - sleep 15
        - kubectl create -f deploy/deploy.yaml || true
        
    ```

2. use KUBE prefix classpath

    ```yaml
    - name: deploy-font
      image: cuisongliu/drone-kube
      settings:
        kube_server:
          from_secret: k8s-server
        kube_ca:
          from_secret: k8s-ca
        kube_admin:
          from_secret: k8s-admin
        kube_admin_key:
          from_secret: k8s-admin-key
      commands:
        - drone-kube config  
        - kubectl delete -f deploy/deploy.yaml || true
        - sleep 15
        - kubectl create -f deploy/deploy.yaml || true
    ```

## configToken

> config user kubeconfig, from env generator /root/.kube/config ,the user controller k8s cluster

### command generator 

```bash
 drone-kube configToken  --token xxx  --user xxx  --server xxx
```

### use docker  generator

```bash
 docker run -ti --network=host -e KUBE_SERVER=https://xxx:6443  -e KUBE_USER=xxx    -e KUBE_TOKEN=xxx cuisongliu/drone-kube bash 
 drone-kube configToken
```

| describe | env |
| :--- | :---  |  
| ~/.kube/config  server | KUBE_SERVER , PLUGIN_SERVER , PLUGIN_KUBE_SERVER|  
| ~/.kube/config  token | KUBE_TOKEN, PLUGIN_TOKEN, PLUGIN_KUBE_TOKEN |  
| ~/.kube/config  user | KUBE_USER, PLUGIN_USER, PLUGIN_KUBE_USER |  

### use drone generator 
1. use no prefix classpath
    ```yaml
    - name: deploy-font
      image: cuisongliu/drone-kube
      settings:
        server:
          from_secret: k8s-server
        user:
          from_secret: k8s-user
        token:
          from_secret: k8s-token
      commands:
        - drone-kube configToken  >> /dev/null
        - kubectl delete -f deploy/deploy.yaml || true
        - sleep 15
        - kubectl create -f deploy/deploy.yaml || true
        
    ```

2. use KUBE prefix classpath

    ```yaml
    - name: deploy-font
      image: cuisongliu/drone-kube
      settings:
        kube_server:
          from_secret: k8s-server
        kube_user:
          from_secret: k8s-user
        kube_token:
          from_secret: k8s-token
      commands:
        - drone-kube configToken  >> /dev/null
        - kubectl delete -f deploy/deploy.yaml || true
        - sleep 15
        - kubectl create -f deploy/deploy.yaml || true
    ```

## template

> deploy dir is template file,the file suffix is **yaml.tmpl** or **yml.tmpl** . need replace env. the files name auto rename to .yaml/.yml.
> the support env prefix: *TEMPLATE_*

### command template 


```bash
 drone-kube template  --deploy=xxx  
```

### use docker  template

```bash
 docker run -ti --network=host -e TEMPLATE_IMAGES1=alpine   cuisongliu/drone-kube bash 
 drone-kube template --deploy=xxx
```

### use drone template 

- drone script
```yaml
- name: deploy-font
  image: cuisongliu/drone-kube
  settings:
    server:
      from_secret: k8s-server
    ca:
      from_secret: k8s-ca
    admin:
      from_secret: k8s-admin
    admin_key:
      from_secret: k8s-admin-key
    template_tag1: alpine
    template_tag2: ${DRONE_TAG=drone-test}
  commands:
    - drone-kube config
    - drone-kube template  >> /dev/null
    - kubectl delete -f deploy/deploy.yaml || true
    - sleep 15
    - kubectl create -f deploy/deploy.yaml || true
    
```

- deploy script deploy/deploy.yaml.tmpl
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.TEMPLATE_TAG1}}
  labels:
    app: {{.TEMPLATE_TAG1}}
spec:
  replicas: 1
  template:
    metadata:
      name: {{.TEMPLATE_TAG1}}
      labels:
        app: {{.TEMPLATE_TAG1}}
    spec:
      containers:
        - name: {{.TEMPLATE_TAG1}}
          image: {{.TEMPLATE_TAG2}}
          imagePullPolicy: IfNotPresent
      restartPolicy: Always
  selector:
    matchLabels:
      app: {{.TEMPLATE_TAG1}}
```
