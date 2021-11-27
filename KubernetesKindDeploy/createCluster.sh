#!/bin/sh

sudo ~/KinD/kind create cluster --name $1 --config config.yaml 

#loading all the images
for line in $(docker image ls | grep webshop-api-microservices | awk '{print $1 ":v1.0"}')
  do
    sudo ~/KinD/kind load docker-image $line --name $1
  done

#loading ingress controller (nginx)
sudo kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

sudo kubectl wait --namespace ingress-nginx   --for=condition=ready pod   --selector=app.kubernetes.io/component=controller   --timeout=300s

sudo kubectl apply -f service-host-list-config.yaml
sudo kubectl apply -f pv.yaml

#secrets and configs
for line in $(ls | grep -e secret -e config | grep -v ^config.yaml$  | awk '{print $1}')
  do
    sudo kubectl apply -f $line
  done



#statefuls -> databases and img service
for line in $(ls | grep -e db.yaml$ -e img-service | awk '{print $1}')
  do
    sudo kubectl apply -f $line
  done

#wait for dbs to be ready
for line in $(sudo kubectl get pods | grep db-ss | awk '{print $1}')
  do
    sudo kubectl wait --for=condition=ready pod/$line --timeout=600s
  done


#stateless services
for line in $(ls | grep service.yaml$ | grep -v img | awk '{print $1}')
  do
    sudo kubectl apply -f $line
  done


sudo kubectl apply -f ingress.yaml



