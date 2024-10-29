import json
import sys
import time
import random
from bs4 import BeautifulSoup
import requests

def match_category(title):
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
    

    matched_categories = []
    for keyword, category_id in category_mapping.items():
        if keyword in title:
            matched_categories.append(category_id)
    return matched_categories

def scrape_sub_items(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
    }
    """抓取子項目內容"""
    try:
        res = requests.get(url,headers=headers)
        res.encoding = 'utf-8'
        soup = BeautifulSoup(res.text, "html.parser").find_all("a")

        sub_results = []
        for a_tag in soup:
            if 'href' in a_tag.attrs:
                link_url = a_tag["href"]
                if not link_url.startswith("http"):
                    link_url = f"https://social.chiayi.gov.tw/{link_url}"
                title = a_tag.get_text(strip=True)
                
                # 匹配子項目種類編號
                categories = match_category(title)
                
                sub_results.append({"url": link_url, "title": title, "category": categories})
        return sub_results
    except requests.exceptions.RequestException as e:
        print(f"RequestException in sub-items: {e}")
        return []

def scrape_data(city, url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"
    }

    try:
        # 模擬瀏覽器請求並添加隨機延遲
        time.sleep(random.uniform(2, 5))
        res = requests.get(url, headers=headers)
        res.encoding = 'utf-8'
        
        soup = BeautifulSoup(res.text, "html.parser").find_all("a")[21:48]

        results = []
        for a_tag in soup:
            if 'href' in a_tag.attrs:
                link_url = a_tag["href"]
                if not link_url.startswith("http"):
                    link_url = f"https://social.chiayi.gov.tw/{link_url}"
                title = a_tag.get_text(strip=True)
                
                # 抓取子項目
                sub_items = scrape_sub_items(link_url)

                # 匹配主項目種類編號
                categories = match_category(title)
                
                results.append({
                    "city": city,
                    "url": link_url,
                    "title": title,
                    "category": categories,
                    "sub_items": sub_items
                })
        
        return results
    
    except requests.exceptions.RequestException as e:
        print(f"RequestException in main: {e}")
        return []

def main():
    city = sys.argv[1]
    url = sys.argv[2]

    data = scrape_data(city, url)
    json_data = json.dumps(data, ensure_ascii=False, separators=(',', ':'))
    print(json_data)

if __name__ == "__main__":
    main()
