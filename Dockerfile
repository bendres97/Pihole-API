FROM python:3.9-slim

RUN mkdir /app
RUN mkdir /data

RUN useradd -m runner
RUN chown -R runner /app
RUN chown -R runner /data

WORKDIR /app

USER runner

RUN pip install --upgrade pip

COPY requirements.txt .
RUN pip install -r requirements.txt

ADD main.py /app/

EXPOSE 5000
CMD [ "python3", "main.py"]