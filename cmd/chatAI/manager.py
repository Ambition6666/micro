import etcd3
from langchain.chains import ConversationChain

class Manager:
    def __init__(self) -> None:
        self.conns = dict()
        self.lease = None

    def addConn(self, id: int, ConversationChain: ConversationChain) -> None:
        self.conns[id] = ConversationChain 

    def getConn(self, id) -> ConversationChain:
        return self.conns[id]     

    def removeConn(self, id) -> None:
        del self.conns[id]

    def setLease(self, lease: etcd3.Lease) -> None:
        self.lease = lease

    def refreshLease(self):
        if self.lease == None:
            return
        self.lease.refresh()
    
manager = Manager()