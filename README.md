![logo](https://user-images.githubusercontent.com/28186014/62043602-10ea1b00-b201-11e9-8749-dd03bd9e0822.png)
# Octoprint-Exporter for Prometheus
**Simple GoLang-Application that exports Octoprint-Metrics into the Prometheus Format**

This projects tries to create a simple way to integrate your 3D-Printer into your existing Prometheus Monitoring System.It present's a standardized Prometheus Exporter to query. 

To quickly run this application use the docker container `localleon/octoprint-exporter` on hub.docker.com.  Adjust the `configs/config-example.yaml` file to your needs and run the container with it `docker run --net=host -v "$(pwd)"/config.yaml:/bin/config.yaml sam:latest`.

---

## Metrics
Metrics are exposed under **:9112/metrics** and start with the prefix "octoprint" . The Metrics are a replicate from the Octoprint REST API. 

**Most relevant Metrics**
- Temperature 
- PrintTime Stats about current Job 
- Print Progress

> The Status of the Printer is represented as Gauge

Prometheus Metric | OctoPrint Status
--- | --- 
0 | Unknown
1 | Operational
2 | Printing from SD
3 |Printing

## Contributing 
- Pull-requests and bug reports wanted !

## Plattform 
Currently tested on:
- linux/amd64
- linux/arm

## Building
Use the provided build.sh with './scripts/build.sh' or build yourself with 'go build . -o $BINNAME'

### Deploying binarys
You need to provide all binarys with your own config file. An example of the config file can be found in `configs/config-example.yaml`. Use the flag `--config` to specify the path to the file.

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

## Credits / Libarys used
- https://github.com/prometheus       <-- Prometheus (Open-Source Monitoring System)
- https://octoprint.org/              <-- The snappy web interface for your 3D printer
- https://mholt.github.io/json-to-go/ <-- Used to create Go Structs out of JSON

## Author 
Copyright localleon(c) 2019

This project is MIT Licensed 
