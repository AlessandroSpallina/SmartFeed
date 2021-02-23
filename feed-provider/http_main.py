import uvicorn
from fastapi import FastAPI
from pydantic import BaseModel
from typing import List

import utils
import repository

app = FastAPI()


class ProducerMessage(BaseModel):
    tag: str
    args: List[str]
    data: str


@app.get("/ping")
def ping():
    return {"message": "pong", "status": "success"}


@app.post('/produce')
def produce(message: ProducerMessage):
    r = repository.InterestRepository()
    ret_value = r.create(message.tag, message.args, message.data)
    if ret_value:
        return {'message': message, 'status': 'success'}
    else:
        return {'message': '', 'status': 'failure'}


if __name__ == "__main__":
    config = utils.get_config_from_env()

    uvicorn.run(app, host="0.0.0.0", port=config['HTTP_SERVER_PORT'])
