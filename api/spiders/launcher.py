#!/usr/bin/env python3
import subprocess
from urllib.parse import quote_plus


def run(spider: str, query: str):
    cmd = f'scrapy crawl {spider} -a query={quote_plus(query)}'
    subprocess.run(cmd.split(' '))
