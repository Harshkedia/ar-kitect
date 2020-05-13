# -*- coding: utf-8 -*-
import os

from flask import Flask, jsonify
from flask_cors import CORS


def create_app(config=None):
    app = Flask(__name__)

    app.config.update(dict(DEBUG=True))
    app.config.update(config or {})

    CORS(app)

    @app.route("/")
    def hello_world():
        return "Hello World"

    # TODO: create subprocess toconvert gltf
    # TODO: create usdz url 
    # NOTE: will the request fail because of large response time?
    # send dummy url as respone and process in own time?
    # invalidate url if process fails.. create err msg?

    return app


if __name__ == "__main__":
    port = int(os.environ.get("PORT", 8000))
    app = create_app()
    app.run(host="0.0.0.0", port=port)
