FROM python:3.8

RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

EXPOSE 5000
CMD [ "/bin/bash", "entrypoint.sh"]