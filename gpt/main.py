#!/usr/bin/env python
import aio_pika
import asyncio
import os
import aio_pika.abc
import summarizer
import json
import requests



async def main(loop):
    print("Running GPT")
    amqp_url = os.getenv("AMQP_SERVER_URL", "")
    amqp_host, amqp_port = amqp_url.split(":")[0], (amqp_url.split(":")[1])
    print(f"Host {amqp_host} - Port {amqp_port}")

    # connection = pika.BlockingConnection(pika.ConnectionParameters(host=amqp_host, port=amqp_port))
    # connection = pika.BlockingConnection(pika.URLParameters(amqp_url))
    connection = await aio_pika.connect_robust(
        amqp_url, loop=loop
    )
    # print(f"Is Open? - {connection.is_open}")
    # channel = connection.channel()
    #
    # channel.queue_declare(queue='gptx', durable=True, auto_delete=False, exclusive=False)

    async with connection:
        queue_name = "gptx"
        # Creating channel
        channel: aio_pika.abc.AbstractChannel = await connection.channel()

        # Declaring queue
        queue: aio_pika.abc.AbstractQueue = await channel.declare_queue(
            queue_name,
            auto_delete=False,
            durable=True
        )

        async with queue.iterator() as queue_iter:
            # Cancel consuming after __aexit__
            async for message in queue_iter:
                async with message.process():
                    json_string = message.body.decode()
                    print(json_string)
                    data = json.loads(json_string)
                    scrape_data = data["ScrapeData"]
                    summary = summarizer.Summarize(scrape_data)
                    print(summary)
                    if queue.name in message.body.decode():
                        break



    #

    # while True:
    #     method_frame, header_frame, body = channel.basic_get(queue='gptx', auto_ack=False)
    #     if method_frame:
    #         print(body.decode())
    #         channel.basic_ack(method_frame.delivery_tag)
    #     else:
    #         print('No message in queue')
    #         time.sleep(1)

    # def callback(ch, method, properties, body):
    #     ch.basic_ack(delivery_tag=method.delivery_tag)
    #     print(" [x] Received %r" % body.decode())
    #
    # channel.basic_consume('gptx', callback, False)
    # print(' [*] Waiting for messages. To exit press CTRL+C')
    #
    # try:
    #     channel.start_consuming()
    # except: 
    #     print("BRUH")


if __name__ == "__main__":
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main(loop))
    loop.close()
