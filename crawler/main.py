import importlib
import json
import os
import pandas as pd
import subprocess
import sys


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

        # 儲存回傳結果
        results.append({
            'city': city,
            'output': result
        })
        print(city+"成功")
    except Exception as e:
        # 捕捉任何異常
        results.append({
            'city': city,
            'output': "異常"
        })
        print(city+"異常")

# 定義 JSON 檔案的路徑
json_file_path = os.path.join("../", 'data.json')

# 將結果儲存為 JSON 格式
with open(json_file_path, 'w', encoding='utf-8') as json_file:
    json.dump(results, json_file, ensure_ascii=False, indent=4)  # 保存為 JSON 文件

print(f"Results saved to {json_file_path}")