import scrapy


class G2ASpider(scrapy.Spider):
    name = "g2a"

    def start_requests(self):
        base_url = 'https://www.g2a.com/en-us/search?query='

        query = getattr(self, 'query', None)
        if query is not None:
            url = f'{base_url}{query}'
            yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response, **kwargs):
        # Only get first result.
        for result in response.xpath('(//li[contains(@class, "pc-digital")])[1]'):
            yield {
                'name': result.css('a::text').get(),
                'price': result.css('span>span::text').get()
            }
