FROM python:2.7-alpine
RUN mkdir /maze
WORKDIR /maze
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY . ./
RUN python setup.py install
ENTRYPOINT /maze/walk.sh
