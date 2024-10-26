import json
import sys

from bs4 import BeautifulSoup
import requests

def scrape_data(city, url): 
    headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
}
    res = requests.get(url,headers=headers)
    res.encoding = 'utf-8'

    try:
        soup = BeautifulSoup(res.content, "html.parser")
        print(soup)
    except UnicodeDecodeError as e:
        print(f"UnicodeDecodeError: {e}")
    
    
    
    return soup
    
    results = [] 
    
    for li in soup:
        a = li.find("span").find("a")
        url = "https://social.chiayi.gov.tw/cl.aspx?n=379" + a["href"]
        
        res = requests.get(url)
        res.close()
        
        soup = BeautifulSoup(res.text, "html.parser").find("tbody").find_all("tr");
        print(soup)
        
        for tr in soup:
            a= tr.find_all("td")[1].find("span").find("a")
            url = "https://social.chiayi.gov.tw/cl.aspx?n=379" + a["href"] + "&PageSize=100"
            title = a["title"]
            results.append({"url": url, "title": title})
    return results

def main():
    # 接收從主程式傳入的參數
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)  # 调用 scrape_data 函数获取数据
    json_data = json.dumps(data, ensure_ascii=False)  # 将数据转换为 JSON 格式
    print(json_data)  # 输出 JSON 数据

if __name__ == "__main__":
    main()
