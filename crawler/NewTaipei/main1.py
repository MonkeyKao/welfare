from requests_html import HTMLSession
import validators

def scrape_data(city, url):
    results = []

    session = HTMLSession()
    r = session.get(url)
    about = r.html.find('.list a')
    for item in about:

        if "(另開新視窗)" in item.attrs["title"]:
            continue
        else:
            url = ("https://www.sw.ntpc.gov.tw/"+item.attrs["href"])
            r = session.get(url)
            about = r.html.find('.con a')
            
            for temp in about:
                if validators.url(temp.attrs['href']):
                    results.append({"category": [1], "city": city, "url": temp.attrs['href'], "title": temp.attrs['title']})
                else:
                    results.append({"category": [1], "city": city, "url": "https://www.sw.ntpc.gov.tw/"+temp.attrs['href'], "title": temp.attrs['title']})
    session.close()
    return results  # 返回包含所有结果的列表

def main(city,url):

    return scrape_data(city, url)  # 调用 scrape_data 函数获取数据

if __name__ == "__main__":
    main()
