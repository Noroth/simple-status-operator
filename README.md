# kubebuilder tutorial
## 1. install the required tools

Follow the steps on the [documentation](https://book.kubebuilder.io/quick-start.html)

[Kustomize Binary Installation guide](https://kubernetes-sigs.github.io/kustomize/installation/binaries/)

Install with homebrew
```bash
brew install kustomize
```

```bash
os=$(go env GOOS)
arch=$(go env GOARCH)

# download kubebuilder and extract it to tmp
curl -L https://go.kubebuilder.io/dl/2.3.1/${os}/${arch} | tar -xz -C /tmp/

# move to a long-term location and put it on your path
# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
sudo mv /tmp/kubebuilder_2.3.1_${os}_${arch} /usr/local/kubebuilder
export PATH=$PATH:/usr/local/kubebuilder/bin
```


## 2. Initialize a new project
```bash
domain=noroth.io
kubebuilder init --domain ${domain}
```


## 3. Create a new API for show cases
```bash
# create a custom resource definition for a counter
kubebuilder create api --group app --version v1beta1 --kind Counter
```


## 4. Update the /api/v1beta1/counter_types.go
Add values for spec and status

If you want your status to be displayed with kubectl you need to add specific marker comments. 

See [Printer Columns](https://book.kubebuilder.io/reference/generating-crd.html#additional-printer-columns)


## 5. Install CRDs
```bash
# Make sure you have a connection to a kubernetes cluster or minikube available
# Run some commands
make generate; make manifests
# install the CRDs
make install

kubectl apply -f config/samples/{resource}
## run the application
make run
```

## 6. Deployment

See [Documentation](https://book.kubebuilder.io/cronjob-tutorial/running.html)