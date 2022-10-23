# CODE1866EUR

## Carvel Demo

### Setup

brew tap vmware-tanzu/carvel
brew install kapp kbld ytt kctrl imgpkg vendir

### Part 1

kapp deploy -a demo-app -c -f manifests/

kbld -f manifests/ -f config.yml

kbld -f manifests/ -f config.yml | kapp deploy -a demo-app -c -f-

ytt -f manifests/ -v msg=cli

ytt -f manifests/ -v msg=cli | kbld -f config.yml -f- | kapp deploy -a demo-app -c -f-

ytt -f manifests -f values.yml | kbld -f config.yml -f- | kapp deploy -a demo-app -c -f-

kapp delete -a demo-app

### Part 2

kubectl create sa demo-sa
kubectl create rolebinding demo-sa-admin --clusterrole=admin --serviceaccount=default:demo-sa

kctrl package init

kctrl dev -f package-resources.yml --local --kbld-build

kctrl package release -v 1.0.0

### Part 3

kctrl package repository release --chdir carvel-artifacts

cp carvel-artifacts/package-repository.yml gitops/packaging/repo.yml

> Kapp controller must be already installed on the cluster

git add .
git commit -m "ops"

kapp deploy -a gitops -f gitops/app/
