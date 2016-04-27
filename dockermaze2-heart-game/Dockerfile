FROM python:2.7-alpine
RUN mkdir /heart
WORKDIR /heart
COPY requirements.txt ./
RUN pip install -r requirements.txt
COPY . ./
RUN python setup.py install
ENTRYPOINT /heart/beat.sh
