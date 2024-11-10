from requests_html import HTMLSession

def scrape_data(city, url):
    results = []

    session = HTMLSession()
    r = session.get(url)
    about = r.html.find('div.page>div>ul>li>a')
    for item in about:
        r = session.get("https://welfare.yunlin.gov.tw/" + item.attrs['href'])
        about1 = r.html.find("div.pink>ul>li>a")
        for temp1 in about1:
            results.append({"category": [1], "city": city, "url": ("https://welfare.yunlin.gov.tw/"+temp1.attrs['href']), "title": temp1.attrs['title']})
    session.close()
    return results

def main(city,url):
    return scrape_data(city, url)  # 調用 scrape_data 函數獲取資料

if __name__ == "__main__":
    main()
