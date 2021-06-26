from selenium import webdriver
from selenium.webdriver.common.by import By


options = webdriver.ChromeOptions()
options.add_argument('--headless')

driver = webdriver.Remote(
    command_executor='http://localhost:8080/wd/hub',
    desired_capabilities=options.to_capabilities(),
    options=options,
)

driver.get('https://www.selenium.dev')
print(driver.current_url)

navbar = driver.find_element_by_id('navbar')
print(navbar)

driver.quit()
