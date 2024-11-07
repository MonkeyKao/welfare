from requests_html import HTMLSession

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
                url = ""
                title = temp.attrs['title']
                if "(開啟新視窗)" in title:
                    url += temp.attrs['href']
                else:
                    url = ("https://www.sw.ntpc.gov.tw/"+item.attrs["href"])
                
                results.append({"category": [1], "city": city, "url": url, "title": title})
    session.close()
    return results  # 返回包含所有结果的列表

def main(city,url):

    return scrape_data(city, url)  # 调用 scrape_data 函数获取数据

if __name__ == "__main__":
    main()
