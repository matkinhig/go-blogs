# HƯỚNG DẪN SỬ DỤNG
. facebook.com/matkinhig
. matkinhig@gmail.com
. matkinhig@outlook.com

## INSTALL PROJECT WITH DOCKER COMPOSE.

B1 : AutoMigrate data hoặc run script init.sql cùng folder <br/>

B2 : init go tool hoặc cài đặt golang <br/>

B3 : config docker-compose.yaml <br/>

B4 : run cript : docker-compose up <br/>

B5 : test service : docker ps -a <br/>
<br/>
NẾU MUỐN KẾT NỐI VỚI DataBase ĐANG CHẠY BỞI DOCKER : docker exec -it "namesOfDataBase" -l <br/>

mysql -u root -p ronglong01 <br/>

## INSTALL ELASTICSEARCH
1. Install container Elastic Search
2. Install container Kibana
3. run script : docker-compose up

## Install ES platinum 1 note link another node
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.6.2 
docker pull docker.elastic.co/kibana/kibana:7.6.2

# run Basic Auth Elasticsearch(user: 'elastic', pw 'secret') daemon
# in auto-remove mode, start takes 20+ seconds
docker run -d --rm -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e "transport.host=127.0.0.1" -e ELASTIC_PASSWORD=secret --name elastic docker.elastic.co/elasticsearch/elasticsearch:7.6.2 && sleep 20

# run Kibana daemon in auto-remove mode
# start takes 20+ seconds
docker run -d --rm --link elastic:elasticsearch-url -e "ELASTICSEARCH_URL=http://localhost:9200" -e ELASTICSEARCH_PASSWORD="secret"  -p 5601:5601 --name kibana docker.elastic.co/kibana/kibana:7.0.1 && sleep 20

# check connection to Elasticsearch (JSON is returned)
curl "http://localhost:9200/_count" -u 'elastic:secret'

# check connection to Kibana (HTML is returned)
curl http://localhost:5601 --location



