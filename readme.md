consul
sudo docker run -d --name=cs -p 8500:8500 consul agent -server -bootstrap -ui -client 0.0.0.0 

http://127.0.0.1:8500/v1/agent/services

curl \
--request PUT \
--data @p1.json \
http://127.0.0.1:8500/v1/agent/service/register?replace-existing-checks=true



curl \
    --request PUT \
    http://127.0.0.1:8500/v1/agent/service/deregister/userservice1
