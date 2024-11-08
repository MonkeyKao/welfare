from requests_html import HTMLSession

def scrape_data(city, url):
    session = HTMLSession()
    r = session.get(url)

    about = r.html.find(".np ul li a")
    results = []
    for item in about:
        
        if "福利地圖" in item.attrs['title']:
            continue
        else:
            r = session.get(item.attrs['href'])
            about1 = r.html.find(".np ul li a")
            for temp1 in about1:
                r = session.get(temp1.attrs['href'])
                about2 = r.html.find(".list a")
                for temp2 in about2:
                    results.append({"category": [1], "city": city, "url": temp2.attrs['href'], "title": temp2.attrs['title']})

    session.close()
    return results

def main(city, url):
    return scrape_data(city, url)  # 調用 scrape_data 函數獲取數據

if __name__ == "__main__":
    main()
