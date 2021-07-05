from selenium import webdriver

# prepare: go run http_only_proxy.go
PROXY = '127.0.0.1:8080'

options = webdriver.ChromeOptions()
options.add_argument('--proxy-server=%s' % PROXY)
options.add_argument('--headless')

driver = webdriver.Chrome(options=options)

driver.get('http://example.com/')
print(driver.current_url)

driver.quit()
