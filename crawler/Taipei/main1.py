import sys
import requests
from bs4 import BeautifulSoup
import json

def scrape_data(city, url):
    res = requests.get(url)
    res.encoding = 'utf-8'
    soup = BeautifulSoup(res.content, "html.parser")

    # 打印完整的 HTML 內容，幫助調試
    print(soup.prettify())

    results = []  # 用於存儲所有的 URL 和 title

    # 嘗試使用不同的選擇器
    a_tags = soup.find_all("a")
    if not a_tags:
        print("No elements found with class 'np'")
        return results

    print(a_tags)  # 打印所有 <a> 標籤，便於調試

    # for a_tag in a_tags:
    #     if a_tag.get("class") and "a" in a_tag["class"]:
    #         # 構造完整的 URL
    #         url = "https://dosw.gov.taipei/Content_List.aspx?n=A9C028834E9BDC90" + a_tag["href"]
    #         title = a_tag.get("title", a_tag.get_text(strip=True))
    #         results.append({"city": city, "url": url, "title": title})

    return results  # 返回包含所有結果的列表

def main():
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)  # 調用 scrape_data 函數獲取數據
    json_data = json.dumps(data, ensure_ascii=False, indent=4)  # 將數據轉換為 JSON 格式
    print(json_data)  # 輸出 JSON 數據

if __name__ == "__main__":
    main()
