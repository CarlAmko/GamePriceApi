import scrapy


class SteamSpider(scrapy.Spider):
    name = "steam"

    @staticmethod
    def convert_to_currency(price_str: str) -> float:
        return float(price_str) / 100

    def start_requests(self):
        base_url = 'https://store.steampowered.com/search/?term='

        query = getattr(self, 'query', None)
        if query is not None:
            url = f'{base_url}{query}'
            yield scrapy.Request(url=url, callback=self.parse)

    def parse(self, response, **kwargs):
        for result in response.css('#search_result_container > div > a:nth-of-type(1)'):
            yield {
                'name': result.css('.title::text').get(),
                'price': self.convert_to_currency(
                    result.css('.col.search_price_discount_combined').attrib['data-price-final']
                )
            }
