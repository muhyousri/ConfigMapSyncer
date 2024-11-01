# ConfigMapSyncer

A dummy controller that watches a custom resource called SyncConfig (syncconfigs.muhu3.example.com ),
Logic: When a new SyncConfig is created or updated, the controller creates or updates a ConfigMap in a target namespace with the specified key-value data.

Syncconfig resource Example:

```
apiVersion: muhu3.example.com/v1
kind: SyncConfig
metadata:
  namespace: default
  labels:
    app.kubernetes.io/name: configmapsyncer
    app.kubernetes.io/managed-by: kustomize
  name: syncconfig-sample-1
spec:
  targetNs: "default"
  data:
    month: "january"
    year: "2024"


```

**Install:**

```
git clone https://github.com/muhyousri/ConfigMapSyncer
cd ConfigMapSyncer
kubectl apply -k config/defauly
```
