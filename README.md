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

## Setup local environment

Setup three components in local environment.

- Selenium Server
- Selenium Proxy Server
- Selenium Code (Python)

```
$ docker-compose up
```

When Selenium server starts, the following log is output.

```
2021-06-26 07:00:21,642 INFO Included extra file "/etc/supervisor/conf.d/selenium.conf" during parsing
2021-06-26 07:00:21,644 INFO supervisord started with pid 9
2021-06-26 07:00:22,653 INFO spawned: 'xvfb' with pid 11
2021-06-26 07:00:22,658 INFO spawned: 'vnc' with pid 12
2021-06-26 07:00:22,661 INFO spawned: 'novnc' with pid 13
2021-06-26 07:00:22,665 INFO spawned: 'selenium-standalone' with pid 14
2021-06-26 07:00:22,685 INFO success: xvfb entered RUNNING state, process has stayed up for > than 0 seconds (startsecs)
2021-06-26 07:00:22,685 INFO success: vnc entered RUNNING state, process has stayed up for > than 0 seconds (startsecs)
2021-06-26 07:00:22,685 INFO success: novnc entered RUNNING state, process has stayed up for > than 0 seconds (startsecs)
2021-06-26 07:00:22,685 INFO success: selenium-standalone entered RUNNING state, process has stayed up for > than 0 seconds (startsecs)
Selenium Grid Standalone configuration:
[network]
relax-checks = true

[node]
session-timeout = "300"
override-max-sessions = false
detect-drivers = false
max-sessions = 1

[[node.driver-configuration]]
name = "chrome"
stereotype = '{"browserName": "chrome", "browserVersion": "91.0", "platformName": "Linux"}'
max-sessions = 1

Starting Selenium Grid Standalone...
07:00:23.285 INFO [LoggingOptions.configureLogEncoding] - Using the system default encoding
07:00:23.291 INFO [OpenTelemetryTracer.createTracer] - Using OpenTelemetry for tracing
07:00:24.296 INFO [NodeOptions.getSessionFactories] - Detected 4 available processors
07:00:24.369 INFO [NodeOptions.report] - Adding chrome for {"browserVersion": "91.0","browserName": "chrome","platformName": "Linux","se:vncEnabled": true} 1 times
07:00:24.392 INFO [Node.<init>] - Binding additional locator mechanisms: id, name
07:00:24.415 INFO [LocalDistributor.add] - Added node 218787d9-1433-40fc-a6c2-792e573c1a68 at http://172.17.0.2:4444.
07:00:24.419 INFO [LocalGridModel.setAvailability] - Switching node 218787d9-1433-40fc-a6c2-792e573c1a68 (uri: http://172.17.0.2:4444) from DOWN to UP
07:00:24.652 INFO [Standalone.execute] - Started Selenium Standalone 4.0.0-rc-1 (revision 23ece4f646): http://172.17.0.2:4444
```

## refs

- [10分で理解する Selenium](https://qiita.com/Chanmoro/items/9a3c86bb465c1cce738a)
