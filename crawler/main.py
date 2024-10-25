import pandas as pd
import subprocess

# 載入 Excel 檔案
file_path = 'list.xlsx'
df = pd.read_excel(file_path)

# 迭代每一列，呼叫對應的爬蟲腳本並傳遞 city 和 url
for index, row in df.iterrows():
    city = str(row['city'])
    url = str(row['url'])
    script_path = str(row['name'])  # 這裡的 name 是腳本路徑
    
    # 執行對應的爬蟲腳本，並將 city 和 url 作為參數傳遞
    subprocess.run(['python', script_path, city, url])
