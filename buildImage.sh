swagger generate server -f dsb-swagger.yaml -A k8svolumedsb
docker build -t ocopea/k8s-volume-dsb -f Dockerfile ../
