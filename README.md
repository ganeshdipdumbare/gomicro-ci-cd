![Test, Deploy](https://github.com/ganeshdipdumbare/e2e_template_golang/workflows/Test,%20Deploy/badge.svg)
# e2e_template_golang
Simple project to demonstrate the end to end working code including CI/CD. 

## Introduction
This is a simple http server written in Go. The main purpose of the project
is to demonstrate the CI/CD capabilities using Github Actions.

## Description
When the code changes are pushed to master, following processes are happened-  
- Server test on linux platform  
- Build docker image for the server  
- Push docker image to Github Registry with tag ```latest```  
- Connect to DigitalOcean service with ```doctl```  
- Deploy the service with new docker image to k8s  
- Check if the changes are working by calling the endpoint-  
- ```http://159.65.104.97:30000/hello```

## Improvements
- Add config to fetch env vars from Github Secrets  
- Use the Github Secretes in k8s deployment

