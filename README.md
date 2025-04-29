# map-me-gcp-cloudrunjob
An implementation of managed-environment specific for google cloud plattform to create cloudrunjob which is responsible for creating google cloud project, IAM, service accounts needed in-order for your team to get a fresh project  
> Disclaimer: This is WIP, use in your environment at your own risk.

# Note
There are two packages in this github repo:  

1. github.com/mattilsynet/map-me-gcp-cloudrunjob:v{m:m:p}  
2. github.com/mattilsynet/me-gcp-cloudrun-job-admin:{m:m:p}  

First is the provider you'll link to from your wadm  
Second is the wit package `wash wit fetch` will use to fetch dependencies in your /wit/deps folder as part of your wash build command  

# How to use the me-gcp provider in your application
Look at `local.wadm.yaml` on how the link configuration is expected to look  
Add wkg config, if you haven't installed the tool `wkg` then look under "Requires" under "Development"   
1. `wkg config --edit`  
2. Add:  
```
[namespace_registries]
mattilsynet = "ghcr.io"
```
# Development
If you'd like to test this provider together with a working component, then clone the project, install requirements underneath and follow steps in "Quick start"

## Requires
tinygo >= 0.33.0 (`go install github.com/tinygo-org/tinygo@latest`)  
wrpc (`cargo install wrpc`)  
wasm-tools (`cargo install wasm-tools`)  
wash-cli (`cargo install wash`)   
wkg (`cargo install wkg`)

## Quick start
1. Add wkg config as described under "How to use the me-gcp provider in your application"  
2. Add in `./main.go` the `var gcpadmin = \`\`` to contain a valid jwt with `act as` and `cloudrun admin` permissions in the google cloud project you're operating on as viewed in the `local.wadm.yaml` manifest   
3. Uncomment and comment lines further down in `./main.go` under function `handleNewTargetLink` in section `INFO: Local development`  
4. Modify `local.wadm.yaml` with your project and location for where your service account has access and you'd like to operate  
5. `go mod tidy`  
6. `wash build` in root (OBS, you might need to alternate between `go get` and `wash build` a two times locally) 
7. `cd component`  
8. `wash build`    
9. Terminal 1: `wash up`   
10. Terminal 2: In root of project: `wash app deploy local.wadm.yaml --replace`  
11. Run testfile `./component/me-gcp_test.go`, modify according to `map.update`, `map.delete` or `map.get`  
12. Goto your google cloud console and check for changes according to step 10 above  
