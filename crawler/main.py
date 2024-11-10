import importlib.util
import json
import os
import pandas as pd
import sys
import jieba
from fuzzywuzzy import fuzz
from transformers import pipeline

# 關鍵字與種類編號對應字典
category_mapping = {
    "家庭與育兒": 1, "教育": 2, "健康與退休": 3, "老人與退休": 4,
    "低收入戶與弱勢族群": 5, "殘疾與特殊需求": 6, "就業與創業": 7, "社會安全與基本生活支援": 8,
    "兒童及少年": 9, "其他特定族群": 10
}

# 設定模糊匹配的閾值
SIMILARITY_THRESHOLD = 80

def match_category(title):
    matched_categories = []
    words = jieba.lcut(title)
    for keyword, category_id in category_mapping.items():
        for word in words:
            if fuzz.partial_ratio(keyword, word) >= SIMILARITY_THRESHOLD:
                matched_categories.append(category_id)
                break
    return matched_categories

classifier = pipeline("zero-shot-classification", model="facebook/bart-large-mnli")
categories = ["家庭與育兒","教育","健康與退休","老人與退休","低收入戶與弱勢族群","殘疾與特殊需求","就業與創業","社會安全與基本生活支援","兒童及少年","其他特定族群"]

# 定義 JSON 檔案路徑
json_file_path = "../data.json"

if os.path.exists(json_file_path):
    with open(json_file_path, "r", encoding="utf-8") as f:
        old_data = json.load(f)
else:
    old_data = []

old_data_lookup = {
    item["city"]: {entry["title"]: entry for entry in item.get("output", [])}
    for item in old_data
}

next_id = max(
    (entry["id"] for item in old_data for entry in item.get("output", [])), default=0
) + 1
updated_data = []

# 載入 Excel 檔案
file_path = '../list.xlsx'
df = pd.read_excel(file_path)

new_crawled_data = []

# 迭代每一列，呼叫對應的爬蟲腳本並傳遞 city 和 url
for index, row in df.iterrows():
    city = str(row['city'])
    url = str(row['url'])
    script_path = str(row['name'])

    if not os.path.exists(script_path):
        print(f"{city} 異常: 找不到腳本 {script_path}")
        continue

    try:
        print(f"{city} 開始爬蟲")
        spec = importlib.util.spec_from_file_location("module.name", script_path)
        module = importlib.util.module_from_spec(spec)
        sys.modules["module.name"] = module
        spec.loader.exec_module(module)

        if hasattr(module, 'main'):
            result = module.main(city, url)
        else:
            print(f"Error: {script_path} does not含有 main() 函式。")
            result = []

        print(f"{city} 爬蟲成功，開始進行分類及編號")
        for item in result:
            item["id"] = next_id
            item['category'] = match_category(item['title'])
        print(f"{city} 已經分類完畢")

        new_crawled_data.append({
            'city': city,
            'output': result
        })

    except Exception as e:
        print(f"{city} 異常: {e}")
        new_crawled_data.append({
            'city': city,
            'output': []
        })

# print(new_crawled_data[0])

updated_output = []

# 比對新資料並處理 ID
for new_city_data in new_crawled_data:
    city = new_city_data["city"]
    updated_output = []  # 每次新的 city 處理時初始化

    if "output" not in new_city_data:
        print(f"{city} 異常: 缺少 output 資料")
        continue

    for new_item in new_city_data["output"]:
        # new_item每個爬蟲回來的資料
        if city in old_data_lookup and new_item["title"] in old_data_lookup[city]:
            # 如果舊的資料存在則沿用id
            new_item["id"] = old_data_lookup[city][new_item["title"]]["id"]
        else:
            # 不存在則新增id
            new_item["id"] = next_id
            next_id += 1
        updated_output.append(new_item)
    updated_data.append({"city": city, "output": updated_output})


# 將舊資料中未出現於新資料的項目加入 updated_data
for city, old_items in old_data_lookup.items():
    if city not in [entry["city"] for entry in updated_data]:
        updated_data.append({"city": city, "output": list(old_items.values())})
    else:
        # 將不在新 output 中的舊項目加入
        for entry in updated_data:
            if entry["city"] == city:
                current_titles = {item["title"] for item in entry["output"]}
                for old_title, old_item in old_items.items():
                    if old_title not in current_titles:
                        entry["output"].append(old_item)


# 將更新後的資料覆蓋寫回相同的 JSON 檔案
with open(json_file_path, "w", encoding="utf-8") as f:
    json.dump(updated_data, f, ensure_ascii=False, indent=4)

print(f"Results saved to {json_file_path}")