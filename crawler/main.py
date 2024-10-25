import json
import os
import pandas as pd
import subprocess

# 載入 Excel 檔案
file_path = 'list.xlsx'
df = pd.read_excel(file_path)

results = []

# 迭代每一列，呼叫對應的爬蟲腳本並傳遞 city 和 url
for index, row in df.iterrows():
    city = str(row['city'])
    url = str(row['url'])
    script_path = str(row['name'])  # 這裡的 name 是腳本路徑
    
    # 執行對應的爬蟲腳本，並將 city 和 url 作為參數傳遞
    result = subprocess.run(['python', script_path, city, url],capture_output=True, text=True)
    
    output = result.stdout.strip()
    results.append({"script": script_path, "output": output})  # 儲存結果

    print(f"Received from {script_path}: {output}")

    # 如果子腳本失敗（返回非0狀態碼），則記錄錯誤並繼續執行
    if result.returncode != 0:
        print(f"Error occurred in {script_path}. Continuing execution.")
        print(f"Error message: {result.stderr.strip()}")  # 打印錯誤信息

# 定義 JSON 檔案的路徑
json_file_path = os.path.join(os.path.dirname(__file__), 'data.json')

# 將結果儲存為 JSON 格式
with open(json_file_path, 'w', encoding='utf-8') as json_file:
    json.dump(results, json_file, ensure_ascii=False, indent=4)  # 保存為 JSON 文件

print(f"Results saved to {json_file_path}")