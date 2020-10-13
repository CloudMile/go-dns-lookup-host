# NSlookup URL on GAE wiht GCP Internal DNS Server
dig with `169.254.169.254`

## Deploy
```
$ gcloud app deploy app.yaml
```

## Usage
```
$ curl https://<GAE_URL>/dns-check?url=<URL>
```

Example:
```
$ curl "https://<GCP_PROJECT_ID>.uc.r.appspot.com/dns-check?url=speech.googleapis.com"
URL: speech.googleapis.com
IP: 64.233.183.95	Times: 5
IP: 74.125.126.95	Times: 7
IP: 173.194.74.95	Times: 9
IP: 74.125.129.95	Times: 5
IP: 173.194.198.95	Times: 5
IP: 74.125.69.95	Times: 15
IP: 74.125.124.95	Times: 19
IP: 173.194.197.95	Times: 11
IP: 209.85.146.95	Times: 2
IP: 74.125.202.95	Times: 7
IP: 172.253.119.95	Times: 4
IP: 173.194.192.95	Times: 2
IP: 173.194.195.95	Times: 2
IP: 173.194.193.95	Times: 2
IP: 74.125.70.95	Times: 4
IP: 172.253.114.95	Times: 6
IP: 209.85.147.95	Times: 5
IP: 74.125.201.95	Times: 2
IP: 209.85.145.95	Times: 2
IP: 64.233.181.95	Times: 3
IP: 172.217.214.95	Times: 100
IP: 209.85.234.95	Times: 8
IP: 209.85.200.95	Times: 18
IP: 64.233.191.95	Times: 3
IP: 142.250.1.95	Times: 1
IP: 108.177.112.95	Times: 53
IP: 108.177.120.95	Times: 10
IP: 108.177.111.95	Times: 76
IP: 172.217.219.95	Times: 7
IP: 74.125.132.95	Times: 7
IP: 173.194.194.95	Times: 4
IP: 172.217.212.95	Times: 100
IP: 108.177.121.95	Times: 11
```
