
# Sock Shop : Updating blueprint.yaml 

## Steps to do before using blueprint.yaml

- Set the values for the environmental variables ak, sk and projectid
```
export AK = xxxxxxxxxxxxxxxxxxx
export SK = xxxxxxxxxxxxxxxxxxxx
export PID = xxxxxxxxxxxxxxxxxxxx
```
- Run key.sh file
```
# This will update the AK/SK and Project ID in all the required fields
./key.sh
```
- ak, sk and projectid will be updated in blueprint.yaml file
```
# Tar the blueprint.yml to blueprint.zip and upload this file in Templates in ServiceStage
zip blueprint.zip blueprint.yaml
```
