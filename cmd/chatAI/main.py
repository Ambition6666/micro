import grpc
import time
from concurrent import futures
import aiservice_pb2, aiservice_pb2_grpc
from etcd import EtcdHandleServ
import threading
import signal
import chain
import sched
from manager import manager

# 定时任务对象
schedule = sched.scheduler(time.time, time.sleep)

# 给etcd键值对续约
def refresh(inc, schedule):
    manager.refreshLease()
    time.sleep(inc)
    schedule.enter(0, 0, refresh, (inc, schedule))

def schedul_job():
    inc = 10  # 定时任务,循环时间
    time.sleep(inc)
    schedule.enter(0, 0, refresh, (inc, schedule))  # 加入定时任务
    schedule.run()  # 执行

def something_schedule():
    t = threading.Thread(target=schedul_job)
    t.daemon = True
    t.start()


class AIService(aiservice_pb2_grpc.ChatServiceServicer):
    def Chat(self, request, context):
        chatting = chain.get_normal_chain("x")
        res = chatting.invoke(request.msg)
        return aiservice_pb2.ChatResponse(msg = res["response"])

def main(service_ip, service_port, etcd_ip, etcd_port, etcd_prefix):
    print('***service is starting...')
    grpc_server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    aiservice_pb2_grpc.add_ChatServiceServicer_to_server(AIService(), grpc_server)
    grpc_server.add_insecure_port(f'{service_ip}:{service_port}')
    
    grpc_server.start()

    etcd_handle_serv = EtcdHandleServ(service_port=service_port, etcd_ip=etcd_ip, etcd_port=etcd_port, etcd_prefix=etcd_prefix)
    etcd_handle_serv.register_service()

    event = threading.Event()
    def signal_handler(*args):
        etcd_handle_serv.logout_service()
        event.set()
    signal.signal(signal.SIGINT, signal_handler)
    signal.signal(signal.SIGTERM, signal_handler)

    print("***serveice started")
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        etcd_handle_serv.logout_service()
        grpc_server.stop(0)

if __name__ == '__main__':
    something_schedule()
    main("127.0.0.1", "8889", "192.168.40.134", "2379", "kitex/registry-etcd/chat/")