
# Sock Shop : Updating blueprint.yaml 

## Steps to do before using blueprint.yaml
- Clone this repository
```
git clone https://github.com/huawei-microservice-demo/sockshop-demo
cd sockshop-demo/scripts/blueprint
```

- Set the values for the environmental variables ak, sk, projectid and mesher release
```
export AK=xxxxxxxxxxxxxxxxxxx
export SK=xxxxxxxxxxxxxxxxxxxx
export PROJECTID=xxxxxxxxxxxxxxxxxxxx
export MESHER_RELEASE=xxx
```
- Run key.sh file
```
# This will update the AK/SK and Project ID in all the required fields
bash -x key.sh
./key.sh
```
- ak, sk, projectid and mesher release will be updated in blueprint.yaml file
```
# Compress the blueprint.yaml to blueprint.zip and upload this file in Templates in ServiceStage
zip blueprint.zip blueprint.yaml
```
