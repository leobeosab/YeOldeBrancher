# Ye Olde Brancher
Huzzah fellow peasents! Welcome to this small simple project that will create branches in Bitbucket automatically based on a small JSON file.

The general idea is to create an [App Password](https://support.atlassian.com/bitbucket-cloud/docs/app-passwords/) with bitbucket and setup your credentials like so: 
```bash
export BITBUCKET_USERNAME=some_user
export BITBUCKET_PASSWORD=yourapppassword # please don't use your actual atlassian password
```

Create a simple JSON branch file example:
```json
[
  {
    "name": "testerton",
    "baseBranchName": "master",
    "workspace": "my_workspace",
    "repository": "testing"
  },
  {
    "name": "testerton2electricboogaloo",
    "baseBranchName": "develop",
    "workspace": "my_workspace",
    "repository": "testing"
  }
]
```

Then just run the executable:
```bash
yeoldebrancher -branch-list=path/to/json/file
```