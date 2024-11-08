import importlib
import json
import os
import pandas as pd
import sys
import jieba
from fuzzywuzzy import fuzz

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
    
    # 使用 jieba 進行分詞
    words = jieba.lcut(title)
    
    for keyword, category_id in category_mapping.items():
        for word in words:
            # 使用模糊匹配，若相似度超過閾值則視為匹配
            if fuzz.partial_ratio(keyword, word) >= SIMILARITY_THRESHOLD:
                matched_categories.append(category_id)
                break  # 匹配到一個就跳出詞彙的迴圈，避免重複加入相同類別
    
    return matched_categories


# 載入 Excel 檔案
file_path = '../list.xlsx'
df = pd.read_excel(file_path)

results = []


# 迭代每一列，呼叫對應的爬蟲腳本並傳遞 city 和 url
for index, row in df.iterrows():
    city = str(row['city'])
    url = str(row['url'])
    script_path = str(row['name'])  # 這裡的 name 是腳本路徑
    
    try:
        print(city+"開始爬蟲")
        # 動態載入 Python 腳本
        spec = importlib.util.spec_from_file_location("module.name", script_path)
        module = importlib.util.module_from_spec(spec)
        sys.modules["module.name"] = module
        spec.loader.exec_module(module)

        # 假設每個腳本有一個 main 函數，並傳遞 city 和 url 參數
        if hasattr(module, 'main'):
            result = module.main(city, url)
        else:
            result = f"Error: {script_path} does not have a main() function."

        print(city+"爬蟲成功,開始進行分類")
        for item in result:
            item['category'] = match_category(item['title'])
        print(city+"已經分類完畢")
        # 儲存回傳結果
        results.append({
            'city': city,
            'output': result
        })
        
    except Exception as e:
        # 捕捉任何異常
        results.append({
            'city': city,
            'output': "異常"
        })
        print(city+"異常:"+str(e))

# 定義 JSON 檔案的路徑
json_file_path = os.path.join("../", 'data.json')

# 將結果儲存為 JSON 格式
with open(json_file_path, 'w', encoding='utf-8') as json_file:
    json.dump(results, json_file, ensure_ascii=False, indent=4)  # 保存為 JSON 文件

print(f"Results saved to {json_file_path}")