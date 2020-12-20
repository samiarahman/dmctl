# Quick Start

You need to install the pulumi kubernetes operator into a k8s cluster and set the credentials as secrets 
for stack state storage and aws resource management.

Follow [Create Pulumi Stacks using kubectl](https://github.com/pulumi/pulumi-kubernetes-operator/blob/master/docs/create-stacks-using-kubectl.md) as a reference

After the operator is installed and secrets created into your kubernetes cluster
try run this dmctl as following and see the progress by attaching to logs.

```
go build
./dmctl -f example/datamesh.yaml | kubectl apply -f -
./dmctl -f example/datamesh2.yaml | kubectl apply -f -

# following are how you can inspect the status of this provisioning process
kubectl logs pulumi-kubernetes-operator-* -f
kubectl get stack aws-datamesh-stack -o json | jq .status.outputs
kubectl get stack aws-datamesh-stack2 -o json | jq .status.outputs
```
