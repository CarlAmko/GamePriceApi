# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


from itemadapter import ItemAdapter

from cache import cache


class CacheItemPipeline:
    def process_item(self, item, spider):
        adapter = ItemAdapter(item)
        cache.put(provider=spider.name, query=spider.query, result=adapter.asdict())
        print(item)
        return item
