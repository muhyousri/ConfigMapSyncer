# ConfigMapSyncer

**ConfigMap Syncer**
Goal: Create a controller that watches a custom resource called SyncConfig (e.g., syncconfigs.example.com), which defines the desired state of a ConfigMap.
Logic: When a new SyncConfig is created or updated, the controller creates or updates a ConfigMap in a target namespace with the specified key-value data.
Learning Focus: CRD and ConfigMap interaction, reconciliation loops, and basic object creation
