docker run -d --mount type=bind,source=$(pwd)/sample-models,target=/usr/app/models -p 80:8090 --name usd1 usdzserver