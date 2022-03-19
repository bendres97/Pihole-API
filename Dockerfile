FROM python:3.9-slim

RUN apt update && apt install ssh-tools -y

RUN pip install --upgrade pip

RUN mkdir /app
RUN mkdir /data

COPY requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
ADD main.py /app/

EXPOSE 5000
CMD [ "python3", "main.py"]