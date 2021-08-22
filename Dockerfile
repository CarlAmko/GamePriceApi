FROM python:3.9-slim-bullseye

ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0

ADD requirements.txt requirements.txt
RUN pip install -r requirements.txt

ADD . .
RUN pip install -e .

ENTRYPOINT ["flask", "run"]