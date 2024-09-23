

hosting docker port :  4338

# check if the port still free
sudo lsof -i:<port_number>
# If there is no output, it means the port is not in use and is free.

docker stop biatechauth1
docker rm biatechauth1
docker rmi 010309/biatechauth1:latest

docker-compose up -d

docker logs -f biatechauth1


https://cloudcalls.easipath.com/backend-easit/api/schedules/org/Org1


https://cloudcalls.easipath.com/backend-biatechauth1/api/logins-google/login?session_id=pmis0103069?redirect_uri=https://pmis.bizminder.co.za