# Example

## Set Up

- set up selenium (Python)

- start selenium server by Docker

```bash
$ docker run -d -p 4444:4444 -v /dev/shm:/dev/shm selenium/standalone-chrome:4.0.0-rc-1-prerelease-20210618
```

ref: https://github.com/SeleniumHQ/docker-selenium

When Selenium server starts, the following log is output.

```
$ docker logs 473fe763ab06
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

## Case: Go to google.com/

Try running the following python code.

```python
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
```

The selenium server outputs the following log.


```
07:02:08.679 INFO [LocalDistributor.newSession] - Session request received by the distributor:
 [Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}, platform: ANY, version: }, Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}}, Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}, platformName: any}]
Starting ChromeDriver 91.0.4472.101 (af52a90bf87030dd1523486a1cd3ae25c5d76c9b-refs/branch-heads/4472@{#1462}) on port 1929
Only local connections are allowed.
Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
ChromeDriver was started successfully.
[1624690928.726][SEVERE]: bind() failed: Cannot assign requested address (99)
07:02:09.258 INFO [ProtocolHandshake.createSession] - Detected dialect: W3C
07:02:09.287 INFO [LocalDistributor.newSession] - Session created by the distributor. Id: 2dcf9ba78fea5a08fee58700486e8f92, Caps: Capabilities {acceptInsecureCerts: false, browserName: chrome, browserVersion: 91.0.4472.114, chrome: {chromedriverVersion: 91.0.4472.101 (af52a90bf870..., userDataDir: /tmp/.com.google.Chrome.4h81bg}, goog:chromeOptions: {debuggerAddress: localhost:44475}, networkConnectionEnabled: false, pageLoadStrategy: normal, platformName: ANY, proxy: Proxy(), se:cdp: ws://172.17.0.2:4444/sessio..., se:cdpVersion: 91.0.4472.114, se:vnc: ws://172.17.0.2:4444/sessio..., se:vncEnabled: true, se:vncLocalAddress: ws://localhost:7900/websockify, setWindowRect: true, strictFileInteractability: false, timeouts: {implicit: 0, pageLoad: 300000, script: 30000}, unhandledPromptBehavior: dismiss and notify, webauthn:extension:largeBlob: true, webauthn:virtualAuthenticators: true}
07:02:10.230 INFO [LocalSessionMap.lambda$new$0] - Deleted session from local session map, Id: 2dcf9ba78fea5a08fee58700486e8f92
```


## Refs

- Auto complement on VSCode (https://qiita.com/4roro4/items/93f4851f1140e19753ce)


