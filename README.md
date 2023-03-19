# Replace Workspace Environment Variable BitBucket 
In this program, we will replace Workspace variable of BitBucket using a simple Go script.

# Prepare
- Create [AppPassword](img/appPassword.png), then save/copy the secret.

# Use Case
- Currently, we will replace these Variables, with a new value:
  - `GCP_SA_DEV`
  - `GCP_SA_UAT`
  - `GCP_SA_PROD`
- NOTE: Just edit those variables with yours.

# How-to
- Define the Authentication Variable, by combining `username` of bitbucket with the `secret` that already created before. 
```
export BBAUTH="lukmanlabz:ATBBa6sFfvkxyz1235g7wSEED92D"
```
- Define also your `workspace` name
```
export WORKSPACE="company-workspace"
```
- Create three files with these following name:
    - `dev_sa.json` <- Fill with your data
    - `uat_sa.json` <- Fill with your data
    - `prod_sa.json` <- Fill with your data
