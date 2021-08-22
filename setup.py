from setuptools import setup

with open("README.md", 'r') as f:
	long_description = f.read()

setup(
	name='GamePriceApi',
	version='1.0',
	description='HTTP REST API for fetching game prices',
	license="MIT",
	long_description=long_description,
	author='Carl Amko',
	author_email='carl@carlamko.me',
	install_requires=[
		'Flask==2.0.1',
		'Scrapy==2.5.0',
		'redis==3.5.3'
	]
)
