# NSlookup URL on GAE wiht GCP Internal DNS Server
dig with `169.254.169.254`

## Deplay
```
$ gcloud app deploy app.yaml
```

## Using
```
$ curl https://<GAE_URL>/dns-check?url=<URL>
```

ex:
```
$ curl "https://<GCP_PROJECT_ID>.appspot.com/dns-check?url=speech.googleapis.com"
URL: speech.googleapis.com
IP: 108.177.111.95	Times: 100
IP: 172.217.212.95	Times: 100
IP: 108.177.112.95	Times: 100
IP: 172.217.214.95	Times: 100
```
