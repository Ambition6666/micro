from langchain_openai import ChatOpenAI
from langchain.chains import ConversationChain
from langchain.memory import (
    ConversationBufferMemory, ConversationSummaryMemory
)


def get_normal_chain(openai_api_key) -> ConversationChain:
    llm = ChatOpenAI(openai_api_key = openai_api_key , temperature = 0.9, max_tokens = 256)
    memory = ConversationBufferMemory()
    chain = ConversationChain(
        llm = llm,
        memory =  memory,
    )
    return chain

def get_summary_chain(summary_api_key, conversation_api_key) -> ConversationChain:
    llm = ChatOpenAI(openai_api_key = conversation_api_key , temperature = 0.9, max_tokens = 256)
    memory = ConversationSummaryMemory(llm=ChatOpenAI(openai_api_key = summary_api_key))
    chain = ConversationChain(
        llm = llm,
        memory =  memory
    )
    return chain   