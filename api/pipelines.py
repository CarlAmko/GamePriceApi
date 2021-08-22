from itemadapter import ItemAdapter

from cache import cache


class CacheItemPipeline:
	def process_item(self, item, spider):
		adapter = ItemAdapter(item)
		cache.put(provider=spider.name, query=spider.query, result=adapter.asdict())
		print(item)
		return item
