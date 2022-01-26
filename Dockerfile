FROM python:3.9-slim

RUN apt update && apt install ssh-tools -y
RUN pip install --upgrade pip

RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN pip install -r requirements.txt

EXPOSE 5000
CMD [ "/bin/bash", "entrypoint.sh"]