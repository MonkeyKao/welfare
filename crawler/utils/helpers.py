def create_result(city, url, title, detail=None,date=None):
    """構建結果的字典"""
    result = {
        "category": [1],
        "city": city,
        "url": url,
        "title": title
    }
    if detail:
        result["detail"] = detail
    if date:
        result["date"] = date
    return result