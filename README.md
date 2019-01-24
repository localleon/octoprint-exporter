# octoprint-exporter
### OctoPrint Exporter for Prometheus written in Golang

## Metrics
Metrics are exposed under :9112/metrics and start with the prefix "octoprint" . The Metrics are a replicate from the octoprint API

#### octoprint_status
0 == Unknown
1 == Operational
2 == Printing from SD
3 == Printing

## Deploying 
1. Use Ansible 
    - Run build.sh to create binarys or download them from the releases page
    - Change host and username in ansible-deploy.yaml (Ansible-Playbook)
    - Run ``` ansible-playbook ./ansible-deploy.yaml ```  (tested on CentOS and Raspbian)
    - Add scrape Job to your prometheus server and you're done
2. Do it yourself
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