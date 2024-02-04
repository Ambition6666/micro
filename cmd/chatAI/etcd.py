import etcd3
import json
import socket

lease_timeout = 500
is_lease = False

def get_outside_ip():
    # 获取本机IP地址
    hostname = socket.gethostname()  # 获取本机主机名
    ip = socket.gethostbyname(hostname)  # 通过主机名获取本机IP地址
    return ip


def get_Instance_info(addr, tags):
    json_data = {
        "network": "tcp",
        "address": addr,
        "weight": 10,
        "tags": tags,
    }
    return json_data

class EtcdClient(etcd3.Etcd3Client):
    def get_values_by_key(self, key, **kwargs):
        values, _ = self.get(key, **kwargs)     
        if values is not None:
            try:
                endpoint = json.loads(values.decode('utf-8'))["address"]
            except:
                endpoint = None
                raise Exception()

        return endpoint

    def put_values_by_key(self, key, values, lease):
        if not isinstance(values, dict):
            raise Exception()
        self.put(key, json.dumps(values),lease)


class EtcdHandleServ():
    def __init__(self, service_port, etcd_ip, etcd_port, etcd_prefix):
        self.etcd_ip = etcd_ip
        self.etcd_port = etcd_port

        # service_ip = get_outside_ip()
        service_ip = '127.0.0.1' #在本机机器作实验使用
        self.endpoint = f'{service_ip}:{service_port}'
        self.etcd_prefix = f'{etcd_prefix}'

    def register_service(self):
        etcd_client = EtcdClient(host=self.etcd_ip, port=self.etcd_port)
        key_name = self.etcd_prefix
        lease = etcd_client.lease(lease_timeout)
        etcd_client.put_values_by_key(key_name, get_Instance_info(self.endpoint,{}), lease)


    def logout_service(self):
        etcd_client = EtcdClient(host=self.etcd_ip, port=self.etcd_port)
        key_name = self.etcd_prefix
        with etcd_client.lock(key_name):
            value = etcd_client.get_values_by_key(key_name)
            if self.endpoint == value:
                etcd_client.put_values_by_key(key_name, get_Instance_info(self.endpoint, {}), 0)


