import json
import scrapy


class SteamSpider(scrapy.Spider):
	name = "steam"
	allowed_domains = ["store.steampowered.com"]
	# Query term is provided at runtime.
	query_term = ""

	def start_requests(self):
		# URL encode whitespace chars
		adjusted_query = self.query_term.strip().replace(" ", "%20")
		yield scrapy.Request(url=f'https://store.steampowered.com/search/?term={adjusted_query}', callback=self.parse)

	def parse(self, response):
		# Capture first item in response list, if it exists.
		query_result = response.css('#search_result_container > div > a:nth-of-type(1)')
		if query_result:
			# Assemble response object.
			resp_obj = {'title': query_result.css('.title::text').get(),
			            'cost': query_result.css('.col.search_price_discount_combined::attr(data-price-final)').get()}

			# Convert dict to JSON.
			json_resp = json.dumps(resp_obj)
			print(f"Assembled JSON response: {json_resp}")
			return json_resp
		else:
			return None
