# gridwalls3

A game.

# Prerequisites

You need a Kubernetes cluster. Digital ocean is nice.

## Cluster setup

Install helm:
```

```

### Install Nginx Ingress and Certmanager

The following is based on https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-with-cert-manager-on-digitalocean-kubernetes
However, it's outdated, so I had to do some changes to make it work.

TL;DR:

* Install Nginx ingress:

    ```sh
    cd digitalocean

    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.30.0/deploy/static/mandatory.yaml
    kubectl get svc --namespace ignress-nginx
    # - Now, wait for and note down the external IP of the service ingress-nginx.
    # - Add a DNS record for your hostname pointing to the external IP.
    # - Edit ingress-nginx.yaml with the DNS record (put it in service.beta.kubernetes.io/do-loadbalancer-hostname)
    kubectl apply -f ingress-nginx.yaml
    ```

    sources: 
    * https://github.com/kubernetes/ingress-nginx/blob/8f4d5f8b343612255fa311e2f9a0da216f61a1d8/docs/deploy/index.md
    * ingress-nginx.yaml is based on https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.30.0/deploy/static/provider/cloud-generic.yaml
    * https://github.com/digitalocean/digitalocean-cloud-controller-manager/blob/master/docs/controllers/services/examples/README.md#setup

(TODO: * Create a DNS A record at your domain provider: `*.k8s.mydomain.com -> public IP of your nginx ingress`)

* Install CertManager:
    
    ```sh
    kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.14.1/cert-manager.yaml
    ```

    source: https://cert-manager.io/docs/installation/kubernetes/

* Set up certificate issuers

    ```sh
    # kubectl apply -f staging_issuer.yaml
    kubectl apply -f prod_issuer.yaml
    ```


* Run a test application to verify that everything works:

    ```sh
    cd test kubectl apply --namespace=do-test -f .
    curl https://echo1.yngvark.com
    ```

#### Future improvements

* "!!! attention The default configuration watches Ingress object from all the namespaces. To change this behavior use the flag --watch-namespace to limit the scope to a particular namespace."
  source: https://github.com/kubernetes/ingress-nginx/blob/8f4d5f8b343612255fa311e2f9a0da216f61a1d8/docs/deploy/index.md
* If public IP of Digital Ocean's Load balancer changes, we must update the DNS A record in one.com.
  * Use ExternalDNS instead, which switching from one.com to something else, maby digital ocean works. Maybe linode?

## Wildnotes

DigitalOcean login:

```sh
doctl kubernetes cluster kubeconfig save use_your_cluster_name

```

* https://www.digitalocean.com/docs/kubernetes/how-to/connect-to-cluster/


### TODO


