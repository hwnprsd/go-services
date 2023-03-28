from llama_index import GPTSimpleVectorIndex,  LLMPredictor
from llama_index import download_loader
from langchain.chat_models import ChatOpenAI

llm_predictor = LLMPredictor(llm=ChatOpenAI(model_name="gpt-3.5-turbo", max_tokens=512, client=""))
StringIterableReader = download_loader("StringIterableReader")
loader = StringIterableReader()


def Summarize(text):
    documents = loader.load_data(texts=[f"""
                                        You are acting as a summarizer bot, who's job is to summarize the given blog content. The content itself is scraped from a website and is given below. You need to summarize the article within 500 words and 3 paragraphs. You can only respond in Markdown format with appropriate highlighting for keywords.

                                        {text}

                                        """])
    index = GPTSimpleVectorIndex(documents, llm_predictor=llm_predictor)
    # index = GPTListIndex(documents)
    #
    #
    # response = index.query("What is the given context?")
    response = index.query("Generate a Complete Valid and Correct JSON text for 'blog_title', 'author_details', 'blog_categories' and 'content' (which is a 500 word text summary of the blog in 3 paragraphs, formatted in Markdown). The text has to be a valid json")
    return response
