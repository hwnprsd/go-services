import aio_pika
import json


async def publish_summary_completion_message(connection: aio_pika.Connection, task_id: int, summary):
    message = json.dumps({
                    "MessagingBase": {
                        "TaskId": task_id,
                        "WORK_TYPE": "CALLBACK_API"
                    },
                    "Data": summary,
                    "Status": "SUMMARY_COMPLETE"
                    })
    channel: aio_pika.abc.AbstractChannel = await connection.channel()
    _ = await channel.default_exchange.publish(
        aio_pika.Message(
            body=str.encode(message),
        ),
        routing_key="api"
    ),
