berlioz local push-run --quick --cluster berliozgo --service example
echo '==============================================================================='
echo '==============================================================================='
echo '==============================================================================='
read -p "Pausing to fetch logs..." -t 2
docker ps -a | grep berliozgo-example | head -n 1 | awk '{print $1}' | xargs docker logs