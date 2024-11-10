from requests_html import HTMLSession
import validators

def scrape_data(city, url):
    results = []

    session = HTMLSession()
    r = session.get(url)
    about = r.html.find('.content-list  a',)
    for item in about:
        r = session.get('https://social.hsinchu.gov.tw/'+item.attrs['href'])
        about2 = r.html.find("tbody a")
        for temp in about2:
            if validators.url(temp.attrs['href']):
                results.append({"category": [1], "city": city, "url": temp.attrs['href'], "title": temp.attrs['title']})
            else:
                results.append({"category": [1], "city": city, "url": "https://social.hsinchu.gov.tw/"+temp.attrs['href'], "title": temp.attrs['title']})

    return results

def main(city, url):
    return scrape_data(city, url)  # 调用 scrape_data 函数获取数据


if __name__ == "__main__":
    main()
