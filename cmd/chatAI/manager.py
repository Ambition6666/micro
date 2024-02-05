import etcd3


class Manager:
    def __init__(self) -> None:
        self.lease = None

    def setLease(self, lease: etcd3.Lease):
        self.lease = lease

    def refreshLease(self):
        if self.lease == None:
            return
        self.lease.refresh()

class ChatManger:
    def __init__(self) -> None:
        self.conns = dict()

    def addConn(self, id, ConversationChain):
        self.conns[id] = ConversationChain 

    def getConn(self, id: int):
        if id in self.conns.keys():
            return self.conns[id] 
        return None    

    def removeConn(self, id):
        del self.conns[id]

    
manager = Manager()
chatManger = ChatManger()