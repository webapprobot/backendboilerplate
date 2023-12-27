## Notes

### Required tools
- gluegunplus


### webappbot
1. Download and convert to json

```bash
webappbot d -j
```

Templates are in `template/`

### Renaming
Run in order 
```bash
webappbot r -m -n "github.com/webappbot/backendboilerplate
webappbot r -a -n "backendboilerplate"
```

Then manually rename /etc/app directory 

These are missing from the generated swagger docs:

### Date format
`YYYY-MM-DD`

### Social logins
- [LinkedIn](https://stateful.com/blog/linkedin-oauth)
- https://learn.microsoft.com/en-us/linkedin/shared/authentication/authorization-code-flow?tabs=HTTPS1


### Twitter
https://developer.twitter.com/en/portal/projects/1616075990358425600/apps/26559051/settings
1. https://developer.twitter.com/en/docs/apps/overview

https://developer.twitter.com/en/docs/authentication/oauth-2-0/user-access-token

https://twitter.com/i/oauth2/authorize?response_type=code&client_id=VDZnUWhmem1Zb1NMdDcteHc0Y2U6MTpjaQ&redirect_uri=http://ubuntu.cseco.co.ke:3000/twitter_login&scope=tweet.read%20users.read%20follows.read%20follows.write&state=state&code_challenge=challenge&code_challenge_method=plain



https://developer.twitter.com/en/portal/projects/1616075990358425600/apps/26559051/settings
https://developer.twitter.com/en/docs/authentication/guides/v2-authentication-mapping




### Google
https://accounts.google.com/o/oauth2/auth/identifier?client_id=717762328687-iludtf96g1hinl76e4lc1b9a82g457nn.apps.googleusercontent.com&scope=profile%20email&redirect_uri=https%3A%2F%2Fstackauth.com%2Fauth%2Foauth2%2Fgoogle&state=%7B%22sid%22%3A1%2C%22st%22%3A%2259%3A3%3A1b8%2C16%3A3bbc1287ec843cc5%2C10%3A1674150861%2C16%3A7fd46b5d454f25c1%2Cf3e39bfc2e3b843c929d75dfe14f0b3e19d38832d9d19984167a72b1cf8313e3%22%2C%22cid%22%3A%22717762328687-iludtf96g1hinl76e4lc1b9a82g457nn.apps.googleusercontent.com%22%2C%22k%22%3A%22Google%22%2C%22ses%22%3A%22a6279d0ab0204d7d922490ccbdb51dab%22%7D&response_type=code&service=lso&o2v=1&flowName=GeneralOAuthFlow

### Facebook
https://www.facebook.com/login.php?skip_api_login=1&api_key=145044622175352&kid_directed_site=0&app_id=145044622175352&signed_next=1&next=https%3A%2F%2Fwww.facebook.com%2Fv2.0%2Fdialog%2Foauth%3Fclient_id%3D145044622175352%26scope%3Demail%26redirect_uri%3Dhttps%253A%252F%252Fstackauth.com%252Fauth%252Foauth2%252Ffacebook%26state%3D%257B%2522sid%2522%253A1%252C%2522st%2522%253A%252259%253A3%253A1b8%252C16%253A55d6239e84e71082%252C10%253A1674150903%252C16%253Acabcaf6b7e488904%252C1283763d4c28036922efd2d1b9f0cb510563e919b321869163626bc494b7f2c1%2522%252C%2522cid%2522%253A%2522145044622175352%2522%252C%2522k%2522%253A%2522Facebook%2522%252C%2522ses%2522%253A%252234fcd7c53efc4035bf0c4ffc09330f7e%2522%257D%26ret%3Dlogin%26fbapp_pres%3D0%26logger_id%3Dae15c4bb-bba2-4ebd-bd66-5ff7386d8c47%26tp%3Dunspecified&cancel_url=https%3A%2F%2Fstackauth.com%2Fauth%2Foauth2%2Ffacebook%3Ferror%3Daccess_denied%26error_code%3D200%26error_description%3DPermissions%2Berror%26error_reason%3Duser_denied%26state%3D%257B%2522sid%2522%253A1%252C%2522st%2522%253A%252259%253A3%253A1b8%252C16%253A55d6239e84e71082%252C10%253A1674150903%252C16%253Acabcaf6b7e488904%252C1283763d4c28036922efd2d1b9f0cb510563e919b321869163626bc494b7f2c1%2522%252C%2522cid%2522%253A%2522145044622175352%2522%252C%2522k%2522%253A%2522Facebook%2522%252C%2522ses%2522%253A%252234fcd7c53efc4035bf0c4ffc09330f7e%2522%257D%23_%3D_&display=page&locale=en_GB&pl_dbl=0







https://www.facebook.com/v2.0/dialog/oauth?
client_id=145044622175352&scope=email&
redirect_uri=https%3a%2f%2fstackauth.com%2fauth%2foauth2%2ffacebook
&state=%7b%22sid%22%3a1%2c%22st%22%3a%2259%3a3%3a1b8%2c16%3a538b8aaeefb9cb51%2c10%3a1674151136%2c16%3a3ef4e37f5d52a334%2c5f23b40f4ba45b0664793c7ea971ab7caf2bc7b1c46c60657e8e702134c01f08%22%2c%22cid%22%3a%22145044622175352%22%2c%22k%22%3a%22Facebook%22%2c%22ses%22%3a%2240a13c7fb782421281cec6a4bd1aebfa%22%7d


https://www.facebook.com/v2.0/dialog/oauth?
client_id=8623091061099143&scope=email&
redirect_uri=http%3a%2f%2fubuntu.cseco.co.ke%2ffacebook_login
&state=%7b%22sid%22%3a1%2c%22st%22%3a%2259%3a3%3a1b8%2c16%3a538b8aaeefb9cb51%2c10%3a1674151136%2c16%3a3ef4e37f5d52a334%2c5f23b40f4ba45b0664793c7ea971ab7caf2bc7b1c46c60657e8e702134c01f08%22%2c%22cid%22%3a%22145044622175352%22%2c%22k%22%3a%22Facebook%22%2c%22ses%22%3a%2240a13c7fb782421281cec6a4bd1aebfa%22%7d




## Validations
### UniqueGroup
In `table`, there is only a single set of similar fields given under `UniqueGroup`

### EnsureExists
For `Table` entry



### Process
1. Create database
```bash
su - postgres
DROP USER "backendboilerplate";
CREATE USER "backendboilerplate" WITH PASSWORD 'backendboilerplate';
CREATE DATABASE ""backendboilerplate";
GRANT ALL PRIVILEGES ON DATABASE "backendboilerplate" TO "backendboilerplate";
```


Generate swagger docs
```bash
go install github.com/swaggo/swag/cmd/swag@latest


export GOROOT=/usr/lib/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


swag init --parseDependency --parseInternal
```


```
projectman@cseco.co.ke
```
