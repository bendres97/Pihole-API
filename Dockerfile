FROM python:3.9-slim

RUN apt update && apt install ssh-tools -y

RUN pip install --upgrade pip

RUN mkdir /app
COPY requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
ADD . /app/

EXPOSE 5000
CMD [ "/bin/bash", "entrypoint.sh"]