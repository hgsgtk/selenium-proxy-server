from selenium import webdriver


options = webdriver.ChromeOptions()
options.add_argument('--headless')

driver = webdriver.Remote(
    command_executor='http://localhost:4444/wd/hub',
    desired_capabilities=options.to_capabilities(),
    options=options,
)

driver.get('https://www.google.com/')
print(driver.current_url)

driver.quit()