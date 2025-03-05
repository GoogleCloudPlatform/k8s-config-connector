retries=30
command_to_retry="./bin/conductor runner --branch-repo=/usr/local/google/home/maqiuyu/go/src/github.com/maqiuyujoyce/5-k8s-config-connector --branch-conf=branches-3.yaml --logging-dir=./logs --command=4"

for i in $(seq 1 $retries); do
  $command_to_retry
  if [ $? -eq 0 ]; then
    echo "Command succeeded after $i attempt(s)."
    exit 0
  else
    echo "Command failed, retrying... ($i/$retries)"
    sleep 10 # Wait for 10 seconds before retrying
  fi
done

echo "Command failed after $retries attempts."
exit 1
