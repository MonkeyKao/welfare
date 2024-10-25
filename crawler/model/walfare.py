from dataclasses import dataclass
from datetime import date
import json

@dataclass
class Welfare:
    _id: int
    category: int
    city: int
    date: date
    priority: int
    title: str
    detail: str
    detailCondition: str
    detailDocument: str
    detailLink: str
    
    def to_json(self):
        # 将数据类实例转换为字典
        return json.dumps(self, default=str, ensure_ascii=False, indent=4)