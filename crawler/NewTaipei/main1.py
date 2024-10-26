import sys
import requests
from bs4 import BeautifulSoup
import json

def scrape_data(city, url):
    res = requests.get(url)
    soup = BeautifulSoup(res.text, "html.parser").select(".con")[0].find_all("div")
    res.close()

    results = []  # 用于存储所有的 URL 和 title

    for div_list in soup:
        a = div_list.find_all("span")[1].find("a")
        detail_url = "https://www.sw.ntpc.gov.tw/" + a["href"]
        res = requests.get(detail_url)
        detail_soup = BeautifulSoup(res.content, "html.parser",from_encoding="utf-8").select(".con")
        res.close()

        for div_list_2 in detail_soup:
            a_list = div_list_2.find_all("a")
            for a in a_list:
                url = "https://www.sw.ntpc.gov.tw/" + a["href"]
                title = a["title"]
                results.append({"city":city,"url": url, "title": title})  # 将 URL 和 title 存储为字典并添加到列表中

    return results  # 返回包含所有结果的列表

def main():
    # 接收從主程式傳入的參數
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)  # 调用 scrape_data 函数获取数据
    json_data = json.dumps(data, ensure_ascii=False)  # 将数据转换为 JSON 格式
    print(json_data)  # 输出 JSON 数据

if __name__ == "__main__":
    main()