from selenium import webdriver

PROXY = 'http://localhost:8080'

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--proxy-server=%s' % PROXY)

driver = webdriver.Remote(
    command_executor='http://localhost:4444/wd/hub',
    desired_capabilities=options.to_capabilities(),
    options=options,
)

driver.get('https://www.google.com/')
print(driver.current_url)

driver.quit()
