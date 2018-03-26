# Contiv - VPP ONS 2018 DEMO

This Kubernetes network plugin uses FD.io VPP to provide network connectivity
between PODs. Currently, only Kubernetes 1.9.X versions are supported. This deployment uses SFC controller to deploy and manage VNF(s).


## Quickstart
Get started with Contiv-VPP using the Vagrant environment:
* Use the [Contiv-VPP Vagrant Installation][1] instructions to start a 
  simulated Kubernetes cluster with a couple of hosts running in VirtualBox
  VMs. This is the easiest way to bring up a cluster for exploring the 
  capabilities and features of Contiv-VPP.
* For the ONS 2018 demo, please use only a production environment (VirtualBox/VMWare Fusion)
   
### Deploying VNF(s) and/or Pods
To get started first ssh into the master node and browse into the ons2018 folder: 
```
vagrant ssh k8s-master
cd /home/vagrant/vpp/k8s/demo/ons2018
```

To run the different scenarios, you need to apply labels to your cluster's nodes. To do so, run the setLabels script located in the ons2018 folder. 

There are 3 different scenarios that can be tested. For every scenario, you will first need to create two K8s config Maps, common for all the VNFs. 
```
cd /home/vagrant/vpp/k8s/demo/ons2018/vnf-examples
kubectl apply -f configMaps.yaml
```

#### Scenario 1: 
Deploy an L2PP service chain in a single node. The two VNFs should be deployed on the master node.
```
cd /home/vagrant/vpp/k8s/demo/ons2018/vnf-examples/scenario1
kubectl apply -f vnf1.yaml
kubectl apply -f vnf2.yaml
```

You can verify that both VNFs run on the same host by: 
```
kubectl get po -o wide
```

#### Scenario 2: 
Deploy an L2PP service chain in two nodes. The two VNFs should be deployed on master and worker nodes.
Make sure you delete any previously deployed pods from scenario1. 
To clean up from scenario1: 
```
kubectl delete po vnf1
kubectl delete po vnf2
```

Verify that no pods are running with: 
```
kubectl get po -o wide
```

Browse into scenario2 folder and run:
```
cd /home/vagrant/vpp/k8s/demo/ons2018/vnf-examples/scenario2
kubectl apply -f vnf1.yaml
kubectl apply -f vnf2.yaml
```

Verify that VNFs are running on different nodes with:
```
kubectl get po -o wide
```

#### Scenario 3: 
Deploy two L2PP service chains in two nodes. The four VNFs should be deployed on master and worker nodes.
Again, make sure you delete any previously deployed pods from scenario1 or scenario2. 
To clean up from scenario1/2: 
```
kubectl delete po vnf1
kubectl delete po vnf2
```

Verify that no pods are running with: 
```
kubectl get po -o wide
```

Browse into scenario3 folder and run:
```
cd /home/vagrant/vpp/k8s/demo/ons2018/vnf-examples/scenario2
kubectl apply -f vnf1.yaml
kubectl apply -f vnf2.yaml
kubectl apply -f vnf3.yaml
kubectl apply -f vnf4.yaml
```

Verify that VNFs are running on different nodes with:
```
kubectl get po -o wide
```

#### Any of the 3 scenarios can be combined with Pods that won't necessarily use memif interfaces. 
You can test a combined setup by deploying an nginx server and a simple web client pod. 

Run a nginx Pod with labels app=web and expose it at port 80:
```
kubectl run web --image=nginx --labels app=web --expose --port 80
```

Verify that the nginx pod is running and get its IP address: 
```
kubectl get po -o wide
```

Run a temporary Pod and make a request to web Service by using the IP address from the previous output:
```
$ kubectl run --rm -i -t --image=alpine test-$RANDOM -- sh
/ # wget -qO- http://ipaddress
```

[1]: ../../../vagrant/README.md


