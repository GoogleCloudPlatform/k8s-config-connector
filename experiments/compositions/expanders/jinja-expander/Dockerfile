FROM python:3.7-slim

WORKDIR .
COPY . /

RUN pip install --upgrade pip
RUN pip install jinja2-cli[yaml,toml,xml,hjson,json5]

ENTRYPOINT ["jinja2"]