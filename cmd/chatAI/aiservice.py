from langchain.memory import (
    ConversationBufferMemory, ConversationSummaryMemory
)
import aiservice_pb2, aiservice_pb2_grpc
import chain
from manager import chatManger

class AIService(aiservice_pb2_grpc.ChatServiceServicer):
    def Chat(self, request, context):
        if chatManger.getConn(request.userID) is None:
           chatManger.addConn(request.userID, chain.get_normal_chain("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"))
        chatting = chatManger.getConn(request.userID)
        res = chatting.invoke(request.msg)
        return aiservice_pb2.ChatResponse(msg = res["response"])