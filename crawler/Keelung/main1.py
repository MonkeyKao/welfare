import sys
import requests
from bs4 import BeautifulSoup
import json

def scrape_data(city, url):
    res = requests.get(url)
    res.encoding = 'utf-8'
    soup = BeautifulSoup(res.text, "html.parser")

    results = []  # 用於存儲所有的 URL 和 title

    # 查找所有的 <a> 標籤
    a_tags = soup.select(".np")[0].find_all('a')
    print(a_tags)
    for a_tag in a_tags:
        if a_tag.get("class") and "a" in a_tag["class"]:
            # 構造完整的 URL
            url = "https://www.klcg.gov.tw/tw/social/2748.html" + a_tag["href"]
            title = a_tag.get("title", a_tag.get_text(strip=True))
            results.append({"city": city, "url": url, "title": title})

    return results  # 返回包含所有結果的列表

def main():
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)  # 調用 scrape_data 函數獲取數據
    json_data = json.dumps(data, ensure_ascii=False, indent=4)  # 將數據轉換為 JSON 格式
    print(json_data)  # 輸出 JSON 數據

if __name__ == "__main__":
    main()
