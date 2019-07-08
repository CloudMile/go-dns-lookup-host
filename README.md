# NSlookup with GAE

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
IP: 108.177.111.95	Times: 2000
IP: 172.217.212.95	Times: 2000
IP: 108.177.112.95	Times: 2000
IP: 2607:f8b0:4001:c15::5f	Times: 2000
IP: 172.217.214.95	Times: 2000
```
