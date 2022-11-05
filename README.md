# CODE1866EUR

## Carvel Demo

### Setup

```sh
brew tap vmware-tanzu/carvel
brew install kapp kbld ytt kctrl imgpkg vendir
```

### Part 1

```sh
MSG=terminal go run src/app.go

kubectl port-forward services/demo-deploy 8080

kapp deploy -a demo-app -f manifests/

kbld -f manifests/ -f config.yml

kbld -f manifests/ -f config.yml | kapp deploy -a demo-app -c -y -f-

ytt -f manifests/ -v msg=cli

ytt -f manifests/ -v msg=cli | kbld -f config.yml -f- | kapp deploy -a demo-app -c -y -f-

ytt -f manifests -f values.yml | kbld -f config.yml -f- | kapp deploy -a demo-app -c -y -f-

kubectl delte deploy demo-deploy
```

### Part 2

```sh
kubectl create sa demo-sa
kubectl create rolebinding demo-sa-admin --clusterrole=admin --serviceaccount=default:demo-sa
```

```sh
kctrl package init

kctrl dev -f package-resources.yml --local --kbld-build

kctrl package release -v 1.0.0
```

### Part 3

```sh
kctrl package repository release --chdir carvel-artifacts

cp carvel-artifacts/package-repository.yml gitops/packaging/repo.yml
```

> Kapp controller must be already installed on the cluster

```sh
git add .
git commit -m "ops"
```

```sh
kapp deploy -a gitops -f gitops/app/
```
