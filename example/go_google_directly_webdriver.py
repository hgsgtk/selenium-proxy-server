from selenium import webdriver

# https://www.proxynova.com/proxy-server-list/country-jp/
PROXY = '160.16.52.36:3128'

options = webdriver.ChromeOptions()
options.add_argument('--proxy-server=%s' % PROXY)
options.add_argument('--headless')

driver = webdriver.Chrome(options=options)

driver.get('https://www.google.com/')
print(driver.current_url)

driver.quit()
