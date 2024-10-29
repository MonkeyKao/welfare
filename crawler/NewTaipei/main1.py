import sys
import requests
from bs4 import BeautifulSoup
import json
import jieba
from fuzzywuzzy import fuzz

# 關鍵字與種類編號對應字典
category_mapping = {
    "育兒津貼": 1, "托育補助": 2, "孕婦補助": 3, "單親家庭津貼": 4,
    "學費減免": 5, "獎助學金": 6, "特殊教育支援": 7, "低收入戶子女教育補助": 8,
    "醫療補助": 9, "長照補助": 10, "精神健康補助": 11, "特殊疾病醫療補助": 12,
    "老人年金": 13, "長者健保補助": 14, "養老院或日照中心補助": 15, "老人生活津貼": 16,
    "低收入戶生活補助": 17, "無家可歸者支援": 18, "弱勢族群就業補助": 19, "住房補助": 20,
    "殘疾津貼": 21, "無障礙設施補助": 22, "身心障礙者就業補助": 23, "居家照顧津貼": 24,
    "就業津貼": 25, "創業貸款補助": 26, "技能培訓補助": 27, "職業重建服務": 28,
    "食物券或糧食補助": 29, "水電費補助": 30, "營養午餐補助": 31, "交通費補助": 32,
    "租/賃屋補助": 33, "弱勢兒童及少年補助": 34, "兒少保護通報": 35,
    "原住民福利": 36, "新住民福利": 37, "退伍軍人福利": 38, "性別平權相關補助": 39
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

def scrape_data(city, url):
    res = requests.get(url)
    res.encoding = 'gbk'
    res.encoding = 'utf-8'
    soup = BeautifulSoup(res.text, "html.parser").select(".con")[0].find_all("div")
    res.close()

    results = []  # 用于存储所有的 URL 和 title

    for div_list in soup:
        a = div_list.find_all("span")[1].find("a")
        detail_url = "https://www.sw.ntpc.gov.tw/" + a["href"]
        res = requests.get(detail_url)
        detail_soup = BeautifulSoup(res.content, "html.parser", from_encoding="utf-8").select(".con")
        res.close()

        for div_list_2 in detail_soup:
            a_list = div_list_2.find_all("a")
            for a in a_list:
                title = a["title"]
                href = a["href"]
                if "(開啟新視窗)" in title:
                    title = title.replace("(開啟新視窗)", "")
                    url = "https://www.sw.ntpc.gov.tw/" + href + "（在新窗口打开）"
                else:
                    url = "https://www.sw.ntpc.gov.tw/" + href
            
                # 匹配種類編號
                categories = match_category(title)
                
                results.append({"category": categories, "city": city, "url": url, "title": title})  # 将 URL 和 title 存储为字典并添加到列表中

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
