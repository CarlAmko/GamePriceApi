import scrapy


class HBSpider(scrapy.Spider):
    name = "hb"

    def start_requests(self):
        base_url = 'https://www.humblebundle.com/store/api/search?sort=alphabetical&filter=all&request=1&search='

        query = getattr(self, 'query', None)
        if query is not None:
            url = f'{base_url}{query}'
            yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response, **kwargs):
        results = response.json()['results']
        first_result = results[0]
        yield {
            'name': first_result['human_name'],
            'price': first_result['full_price']['amount']
        }
