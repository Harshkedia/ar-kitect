Build USDZ Server
```
docker build -t usdzserver .
```

Run USDZ Server
```
docker run --rm -d --mount type=bind,source=$(pwd)/sample-models,target=/usr/app/models -p 80:8090 --name usd1 usdzserver
```

Test the endpoint
```
curl -X POST 'https://ar.portfo.io/?mode=fbx' -F f.obj=@'C:/Users/29545/Downloads/testrhinofbx.fbx'
```