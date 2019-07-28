# octoprint-exporter
### OctoPrint Exporter for Prometheus written in Golang

## Metrics
Metrics are exposed under :9112/metrics and start with the prefix "octoprint" . The Metrics are a replicate from the Octoprint REST API

#### Metrics featured: 
- Temperature 
- PrintTime Stats about current Job 
- Print Progress

#### octoprint_status
- 0 == Unknown
- 1 == Operational
- 2 == Printing from SD
- 3 == Printing

## Building
Use the provided build.sh with './scripts/build.sh ./octo-export' or build yourself with 'go build ./octo-exporter -o $BINNAME'

## Deploying 
1. Ansible 
    - Run build.sh to create binarys or download them from the releases page
    - Change host and username in ansible-deploy.yaml (Ansible-Playbook)
    - Run ``` ansible-playbook ./ansible-deploy.yaml ```  (tested on CentOS and Raspbian)
    - Add scrape Job to your prometheus server and you're done
2. Docker
    - Clone the Repo and build the provided Dockerfile with ```docker build -t octoprint-exporter .```
    - Deploy your Container with ```docker run --net=host octoprint-exporter```
    - The Container is build with Alpine Linux and is just 17.5MB in size 
3. Do it yourself
    - Copy the binary and the example-config file to your remote server
    - Copy the systemd-service file to /etc/systemd/system/
    - Use ```systemctl``` to enable the Service

## Contributing 
- Pull-requests and bug reports wanted !

## Plattform 
Currently tested on:
- linux/amd64
- linux/arm

## Credits / Libarys used
https://mholt.github.io/json-to-go/     -- Used to create Go Structs out of JSON

## Author 
Copyright localleon(c) 2019
