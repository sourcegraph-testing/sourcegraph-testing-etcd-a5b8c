FROM k8s.gcr.io/debian-base-arm64:latest@sha256:b62dad17f41f53b7caa291477264eedf5286e9639bff87551bb9b7f960d62a4d

ADD etcd /usr/local/bin/
ADD etcdctl /usr/local/bin/
ADD var/etcd /var/etcd
ADD var/lib/etcd /var/lib/etcd

EXPOSE 2379 2380

# Define default command.
CMD ["/usr/local/bin/etcd"]
