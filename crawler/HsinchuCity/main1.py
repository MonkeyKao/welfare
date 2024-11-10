import math
from requests_html import HTMLSession

def scrape_data(city, url):
    results = []
    
    session = HTMLSession()
    r = session.get(url)
    num = r.html.find('.ap05_01>span',first=True).text[:3]
    page = math.ceil(int(num)/15)

    for i in range(page):
        temp = url + str (i+1)
        r = session.get(temp)
        about = r.html.find('#css_table a')
        for item in about:
            results.append({"category": [1], "city": city, "url": ("https://society.hccg.gov.tw/ch/"+item.attrs['href']), "title": item.attrs['title']})
    return results

def main(city,url):
    return scrape_data(city, url)  # 调用 scrape_data 函数获取数据

if __name__ == "__main__":
    main()
