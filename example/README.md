# Example

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


## Case go to selenium.dev and select element

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
selenium-server_1  | 07:59:57.744 INFO [LocalDistributor.newSession] - Session request received by the distributor:
selenium-server_1  |  [Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}}, Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}, platformName: any}, Capabilities {browserName: chrome, goog:chromeOptions: {args: [--headless], extensions: []}, platform: ANY, version: }]
selenium-server_1  | Starting ChromeDriver 91.0.4472.101 (af52a90bf87030dd1523486a1cd3ae25c5d76c9b-refs/branch-heads/4472@{#1462}) on port 10878
selenium-server_1  | Only local connections are allowed.
selenium-server_1  | Please see https://chromedriver.chromium.org/security-considerations for suggestions on keeping ChromeDriver safe.
selenium-server_1  | ChromeDriver was started successfully.
selenium-server_1  | [1624694397.764][SEVERE]: bind() failed: Cannot assign requested address (99)
selenium-server_1  | 07:59:57.936 INFO [ProtocolHandshake.createSession] - Detected dialect: W3C
selenium-server_1  | 07:59:57.944 INFO [LocalDistributor.newSession] - Session created by the distributor. Id: 8f54a15dc7527d84db94b3024260f463, Caps: Capabilities {acceptInsecureCerts: false, browserName: chrome, browserVersion: 91.0.4472.114, chrome: {chromedriverVersion: 91.0.4472.101 (af52a90bf870..., userDataDir: /tmp/.com.google.Chrome.0llfVU}, goog:chromeOptions: {debuggerAddress: localhost:40383}, networkConnectionEnabled: false, pageLoadStrategy: normal, platformName: linux, proxy: Proxy(), se:cdp: ws://172.19.0.2:4444/sessio..., se:cdpVersion: 91.0.4472.114, se:vnc: ws://172.19.0.2:4444/sessio..., se:vncEnabled: true, se:vncLocalAddress: ws://localhost:7900/websockify, setWindowRect: true, strictFileInteractability: false, timeouts: {implicit: 0, pageLoad: 300000, script: 30000}, unhandledPromptBehavior: dismiss and notify, webauthn:extension:largeBlob: true, webauthn:virtualAuthenticators: true}
selenium-server_1  | 07:59:58.773 INFO [LocalSessionMap.lambda$new$0] - Deleted session from local session map, Id: 8f54a15dc7527d84db94b3024260f463
```

## Client Code Details

Create a new driver that will issue commands using the wire protocol by selenium/py code.

```python
driver = webdriver.Remote(
    command_executor='http://localhost:4444/wd/hub',
    desired_capabilities=options.to_capabilities(),
    options=options,
)
```

The options are as follows

```python
         - command_executor - Either a string representing URL of the remote server or a custom
             remote_connection.RemoteConnection object. Defaults to 'http://127.0.0.1:4444/wd/hub'.
         - desired_capabilities - A dictionary of capabilities to request when
             starting the browser session. Required parameter.
         - browser_profile - A selenium.webdriver.firefox.firefox_profile.FirefoxProfile object.
             Only used if Firefox is requested. Optional.
         - proxy - A selenium.webdriver.common.proxy.Proxy object. The browser session will
             be started with given proxy settings, if possible. Optional.
         - keep_alive - Whether to configure remote_connection.RemoteConnection to use
             HTTP keep-alive. Defaults to True.
         - file_detector - Pass custom file detector object during instantiation. If None,
             then default LocalFileDetector() will be used.
         - options - instance of a driver options.Options class
```

It is deprecated that pass proxy in constructur, instead, pass proxy in options.

When `webdriver.Remote` is callled, a new session starts.

```
        self.start_session(capabilities, browser_profile)
```

https://github.com/SeleniumHQ/selenium/blob/1e3cc6b5f650fbb1da43aa0e400316fd37a5304b/py/selenium/webdriver/remote/webdriver.py#L247



## API Request via Curl command

For example, try to request /wd/hub. It's invalid request.

```
$ curl -X POST \
       -H "Content-Type: application/json" \
       -d '{"desiredCapabilities":{"browser":"chrome"}}' \
       http://localhost:4444/wd/hub

{
  "value": {
    "error": "unknown command",
    "message": "Unable to find handler for (POST) \u002fwd\u002fhub",
    "stacktrace": ""
  }
}
```

