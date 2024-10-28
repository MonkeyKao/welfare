import json
import sys
import time
import random
from bs4 import BeautifulSoup
import requests

def scrape_data(city, url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
    }

    try:
        # 模擬瀏覽器請求並添加隨機延遲
        time.sleep(random.uniform(2, 5))
        res = requests.get(url, headers=headers)
        res.encoding = 'gbk'
        res.encoding = 'utf-8'
        
        soup = BeautifulSoup(res.text, "html.parser").find_all("a")  # 取第 21 到 45 筆資料

        results = []
        for a_tag in soup:
            if 'href' in a_tag.attrs:
                link_url = a_tag["href"]
                if not link_url.startswith("http"):  # 處理相對 URL
                    link_url = f"https://welfare.yunlin.gov.tw/Home/WelfareList/63?Type=A"+a_tag["href"]
                title = a_tag.get_text(strip=True)
                results.append({"city":city,"url": link_url, "title": title})
        
        return results
    
    except requests.exceptions.RequestException as e:
        print(f"RequestException: {e}")
        return []

def main():
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)  # 調用 scrape_data 函數獲取資料
    json_data = json.dumps(data, ensure_ascii=False, separators=(',', ':'))  # 單行 JSON 格式輸出
    print(json_data)  # 輸出 JSON 資料

if __name__ == "__main__":
    main()
