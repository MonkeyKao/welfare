from requests_html import HTMLSession

def scrape_data(city, url):
    session = HTMLSession()
    r = session.get(url)
    about = r.html.find(".content-list a")
    results = []
    for item in about:

        r = session.get("https://sab.tycg.gov.tw/"+item.attrs["href"])
        about = r.html.find(".content-list a")
        if len(about) == 0:
            about = r.html.find("td a")
            for temp2 in about:
                if "https" in temp2.attrs['href']:
                    url = temp2.attrs['href']
                else:
                    url = "https://sab.tycg.gov.tw/" + temp2.attrs['href']
            
                results.append({"category": [1], "city": city, "url": url, "title": temp2.attrs['title']})
            
        for temp in about:
            r = session.get("https://sab.tycg.gov.tw/"+temp.attrs["href"])
            about = r.html.find("td a")
            
            for temp2 in about:
                if "https" in temp2.attrs['href']:
                    url = temp2.attrs['href']
                else:
                    url = "https://sab.tycg.gov.tw/" + temp2.attrs['href']
            
                results.append({"category": [1], "city": city, "url": url, "title": temp2.attrs['title']})
    session.close()
    return results

def main(city,url):
    return  scrape_data(city, url)  # 调用 scrape_data 函数获取数据

if __name__ == "__main__":
    main()
