import sys

def main():
    # 接收從主程式傳入的參數
    city = sys.argv[1]
    url = sys.argv[2]
    
    # 打印 city 和 url
    print(f"City: {city}, URL: {url}")

if __name__ == "__main__":
    main()