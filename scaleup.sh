#! /bin/sh
replica=`kubectl get pods | grep depl | wc -l`
kubectl scale -n default deployment camunda-deployment --replicas=$(($replica+1))



