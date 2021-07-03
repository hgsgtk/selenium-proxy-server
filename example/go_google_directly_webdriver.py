from selenium import webdriver

PROXY = '127.0.0.1:8080'

options = webdriver.ChromeOptions()
options.add_argument('--proxy-server=%s' % PROXY)
options.add_argument('--headless')

driver = webdriver.Chrome(options=options)

driver.get('https://www.google.com/')
print(driver.current_url)

driver.quit()
