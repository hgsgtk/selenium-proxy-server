# selenium-proxy-server
Selenium Proxy Server to log request and response

## How WebDriver control Browser

As shown in the following documents [Selenium: Understanding the components](https://www.selenium.dev/documentation/en/webdriver/understanding_the_components/), There are generally three ways to control actual Browser.

1. WebDriver communicate driver **directly on same host system**
2. Remote communication through **Remote WebDriver** which is on the same system as the driver and the browser
3. Remote communication through **Selenium Server (Grid)**

This server is only relevant option 3 (using Selenium Server)

![](https://www.selenium.dev/documentation/images/remote_comms_server.png?width=400px)

## How does selenium-proxy-server act

This server acts as a proxy between Web Driver and SeleniumServer.

```plantuml
@startuml

control WebDriver
control SeleniumProxyServer #Red
control SeleniumServer
box "Host"
control Driver
control Browser
end box 

WebDriver -> SeleniumProxyServer
SeleniumProxyServer -> SeleniumServer
SeleniumServer -> Driver
Driver -> Browser
Browser -> Driver
Driver -> SeleniumServer
SeleniumServer -> SeleniumProxyServer
SeleniumProxyServer -> WebDriver
@enduml
```

## About Selenium Server

Selenium Grid is following...

> If you want to scale by distributing and running tests on several machines and manage multiple environments from a central point, making it easy to run the tests against a vast combination of browsers/OS, then you want to use Selenium Grid.
> 
> https://www.selenium.dev/

You can download at [Downloads](https://www.selenium.dev/downloads/) page.

Docker image for the Selenium Grid Server is here. https://github.com/SeleniumHQ/docker-selenium

## refs

- [10分で理解する Selenium](https://qiita.com/Chanmoro/items/9a3c86bb465c1cce738a)