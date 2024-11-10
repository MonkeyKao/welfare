from requests_html import HTMLSession
import validators

def scrape_data(city, url):
    results = []
    session = HTMLSession()
    r = session.get(url)

    about = r.html.find("div.content-list  > div.in >div.hd > div.in >div >span > a")
    for item in about:
        
        if "育兒專區" in item.attrs['title'] or "勞資及勞動條件業務" in item.attrs['title'] or "勞動力及青年發展業務" in item.attrs['title'] or "新住民福利" in item.attrs['title']:
            r = session.get("https://social.chiayi.gov.tw/"+item.attrs['href'])
            about1 = r.html.find("div.content-list  > div.in >div.hd > div.in >div >span > a")
            for temp1 in about1:
                if validators.url(temp1.attrs['href']):
                    results.append({"category": [1], "city": city, "url": temp1.attrs['href'], "title": temp1.attrs['title']})
                else:
                    results.append({"category": [1], "city": city, "url": "https://social.chiayi.gov.tw/"+temp1.attrs['href'], "title": temp1.attrs['title']})
        else:
            r = session.get("https://social.chiayi.gov.tw/"+item.attrs['href'])
            about1 = r.html.find("table a")
            for temp1 in about1:
                results.append({"category": [1], "city": city, "url": temp1.attrs['href'], "title": temp1.attrs['title']})
    session.close()
    return results

def main(city,url):

    return scrape_data(city, url)


if __name__ == "__main__":
    main()
