#!/usr/bin/env python3
import subprocess

from . import spider_names


def run(query: str, spider: str = None):
	if spider:
		cmd = f'scrapy crawl {spider} -a query={query}'
		subprocess.run(cmd, shell=True)
	else:
		cmds = []
		for spider_name in spider_names:
			# Run a scrapy process for each spider available.
			cmds.append(f'scrapy crawl {spider_name} -a query={query}')
		# Wait for all of them to finish.
		cmds.append('wait')

		cmd = ' & '.join(cmds)
		subprocess.run(cmd, shell=True)
