# Github Workflows - Kube RBAC Proxy
## Workflow: Build UBI8 (Minimal)
### What does this workflow do?
This workflow performs a go build for ```kube-rbac-proxy``` with GHA caching, and uses ```buildx``` to build and push multi-platform images based on Red Hat's ```ubi8-minimal``` base image. The ```kube-rbac-proxy``` upstream ```VERSION``` from the repoisitory and is tracked and used as our image tag for all releases. 
### Secrets 
##### REGISTRY_USER
* User for Registry access (required to push)
##### REGISTRY_PASSWORD
* Password required for Registry acesss (required to push)
##### (OPTIONAL) IMAGE_REPOISTIORY
* Image Repository will default to ```quay.io/owner/reponame``` but can be overwritten.
### Known Issues/Areas of Improvement 
There are some areas to promote workflow re-use:
* Make ```GO_VERSION``` dynamic - Currently is hardcoded value in the workflows files, this could be done with a lookup.
* Make ```PLATFORMS``` passable as workflow input - Currently platforms are harcoded within the build-workflow, this could be passed as input from other peroidic jobs like the image check to support re-use.

## Workflow: (Peroidic) Image Check for UBI
This workflow uses the specified ```Containerfile``` to detect a base-image, and uses Skopeo to query the image registry for ```vcs-ref``` label on parent and child images. If drift is detected, the ubi-build workflow is triggered passing required secrets as inputs.  
#### Secrets 
##### REGISTRY_USER
* User for Registry access (required to push)
##### REGISTRY_PASSWORD
* Password required for Registry acesss (required to push)
##### (OPTIONAL) IMAGE_REPOISTIORY
* Will default to ```quay.io/owner/reponame``` but can be overwritten.
### Known Issues/Areas of Improvement 
* Assumes UBI based image, but should work for RH images with ```vcs-ref```. Multi-stage or distroless images using ```...IMAGE as builder``` were not investigated.
* Multi-arch triggers - More testing is needed to verify mutli-arch triggers make sense, or to refactor build-ubi to support passing os-matrix platform parameters as inputs.

## WIP: Workflow: EXAMPLE (Periodic) Fork Sync
This workflow was setup to test PR updates with upstream action. This upstream action uses ```octokit``` under the hood, it may make sense to use public github actions for better support of fork-sync features, or other bots/apps to manage fork updates. For now this is moved to ```WIP_.example```