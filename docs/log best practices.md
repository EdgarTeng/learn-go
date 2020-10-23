# Log best practices



## 1. What is log

Logging is important when developing an application. It helps statistics, debugging and troubleshooting. You may consider what you need for these scenarios when you decide whether to write log and what data to write in log.

Good logs can help developer and maintenance person: 

1. Understand the operating status of the online system
2. Quickly and accurately locate online problems
3. Find system bottlenecks
4. Potential risks of early warning system
5. Mining the maximum value of products
6. ……

Bad logs lead to: 

1. Little or no knowing of the system state
2. Problem cannot be located, or it takes a lot of time and effort
3. Unable to find system bottlenecks, don’t know where to optimize
4. Unable to monitor and alarm errors for potential risks
5. No help for mining user behavior and enhancing product value
6. ……



## 2. Classification of logs

logs can be divided into diagnostic logs, statistical logs, and audit logs.

Diagnostic logs, typically:

- Request entry and exit
- External service call and return
- Resource operations: such as reading and writing files, etc.
- Fault-tolerant behavior: such as cloud hard disk copy repair operation
- Program exception: If the database cannot be connected
- Background operation: periodically delete threads
- Startup, shutdown, configuration loading

Statistics log:

- User access statistics: user IP, amount of data uploaded and downloaded, request time-consuming, etc.
- Billing log (such as recording the network resources used by the user or disk occupation, the format is stricter, and it is convenient for statistics)

Audit log:

- Management operations

For a simple system, all logs can be output to the same log file and differentiated by different keywords. For complex systems, it is necessary to output logs with different requirements to different log files. By adopting different log formats for different types of files (for example, billing logs can be directly output to Json format), which can be easily accessed.

## 3. What is recorded in the log

The good log should record **no more, no less**.

**No more:** means not to record useless information in the log. 

Common useless logs in practice are: 

1) Things that can be placed in one log are output in multiple logs; 

2) An exception that is expected to occur and can be handled normally, prints out a bunch of useless stacks;

3) "Temporary" logs added by developers for debugging convenience during development

**No less:** means that users of the log can get all the information they need from the log. 

In practice, there are often insufficient logs. For example: 

1) When the request is wrong, the log cannot be used to locate the problem, but the developer needs to temporarily increase the log and request the sender to resend the same request to locate the problem; 

2) Unable to determine whether the background tasks in the service are executed as expected; 

3) Unable to determine the status of the service's memory data structure; 

4) Unable to determine whether the service's exception handling logic (such as retry) is executed correctly; 

5) Unable to determine the configuration when the service starts Is it loaded correctly; 

6) etc.

When outputting the log, consider the user of the log. For example, if the log is mainly viewed by the operation and maintenance personnel of the system, it cannot be output:

```bash
[INFO] RequestID:b1946ac92492d2347c6235b4d2611184, ErrorCode:1426 
```

At least:

```bash
[INFO] RequestID:b1946ac92492d2347c6235b4d2611184, ErrorCode:1426, Message: callback request (to http://example.com/callback) failed due to socket timeout
```

In this way, the operation and maintenance personnel can clearly understand the cause of the problem at a glance, and no longer need to develop to view the specific error corresponding to the ErrorCode.

Sort out the logs that are usually missed:

1. System configuration parameters: The system usually reads the startup parameters during the startup process, and these parameters can be output to the log after the system is started to facilitate confirmation that the system is started according to the expected parameters;

2. Tasks that are executed regularly in the background: such as periodically updating the cached tasks, you can record the start time of the task, the end time of the task, how many cache configurations have been updated, etc., so that you can grasp the status of the tasks that are executed regularly;

3. Exception handling logic: For distributed storage systems, when the system fails to read data on one storage node, it needs to go to another data node to retry. You can record the failure of reading data, and using for confirms whether the disk of some nodes may be faulty. For another example, if the system needs to request an external resource, it can record the occasional failure of requesting this external resource and the success of retrying. Specifically:

   ```bash
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth request (to http://auth1.example.com/v2) timeout ... 1 try
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth request (to http://auth1.example.com/v2) timeout ... 2 try
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth request (to http://auth1.example.com/v2) success
   ```

   Better than

   ```bash
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth request (to http://auth1.example.com/v2) success
   ```

   Because the former allows us to predict that the service quality of the depended server is at risk, and may need to be expanded;

4. The key parameters and the key reasons for errors should be recorded in the log. E.g:

   ```bash
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth failed
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611185, content digest does not match
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611186, request ip not in whitelist
   ```

   Not as good as:

   ```bash
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611184, auth failed due to token expiration
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611185, content digest does not match, expect 7b3f050bfa060b86ba781151c563c953, actual f60645e7107917250a6408f2f302d048
   [INFO] RequestID:b1946ac92492d2347c6235b4d2611186, request ip(=202.17.34.1) not in whitelist
   ```

## 4. Log levels

- **Trace** - Only when I would be "tracing" the code and trying to find one **part** of a function specifically.

  

  Trace running detail status, use in test environment (=debug in python)

  ```python
  log.trace('call_function|function=on_action,elapsed=2s,params=["test", 100]')
  ```

- **Debug** - Information that is diagnostically helpful to people more than just developers (IT, sysadmins, etc.).

  

  Helpful info for diagnostic purpose, not only for developers, but to SA etc. too

  ```python
  log.debug('query_db|time_elapsed=2s,sql=xxx')
  ```

- **Info** - Generally useful information to log (service start/stop, configuration assumptions, etc). Info I want to always have available but usually don't care about under normal circumstances. This is my out-of-the-box config level.

  

  Record key events for future reference, which are not very important under normal situation

  ```python
  log.info('tcp_client_connect|remote_ip=127.0.0.1:12345')
  ```

  

- **Warn** - Anything that can potentially cause application oddities, but for which I am automatically recovering. (Such as switching from a primary to backup server, retrying an operation, missing secondary data, etc.)

  

  Anything that can potentially cause abnormal result, but are correctly handled by program

  ```python
  log.warning('error_password|uid=10000,password=xxx,ip=10.10.10.10')
  ```

- **Error** - Any error which is fatal to the **operation**, but not the service or application (can't open a required file, missing data, etc.). These errors will force user (administrator, or direct user) intervention. These are usually reserved (in my apps) for incorrect connection strings, missing services, etc.

  

  Error that only affects current/short-term requests but not to the program or service

  ```python
  try: 
    requests.post(url,data,timeout=10)
  except: 
    log.error('http_request_fail|url=https://xx.yy.zz,error=timeout')
  ```

- **Fatal** - Any error that is forcing a shutdown of the service or application to prevent data loss (or further data loss). I reserve these only for the most heinous errors and situations where there is guaranteed to have been data corruption or loss.

  Critical error that make program can not run normally

  ```python
   log.fatal('mysql_connect_fail|ip=10.10.10.10,db=test_db')
  ```

  

## 5. Continuously optimize the log

One thing is certain. A good log is like a good article. It cannot be written once, but needs to be continuously optimized in the actual operation and maintenance process, combined with the positioning of online problems. The most important point is that the team should pay attention to log optimization, and don't let the quality of the log continue to decrease (when the project becomes larger, the project code also has the same problem, the more it is written, the more chaotic it is).

Here are the following good practices:

- Improve the log in the process of locating the problem. If it takes a long time to locate the problem, it means that there are still problems in the system log and need to be further improved and optimized;
- Need to think about whether it is possible to predict whether the problem may occur in advance by optimizing the log (such as an error caused by a certain resource exhaustion, you can record the resource usage)
- Define the logging specifications of the entire team to ensure that the log format of each development record is uniform; in particular, it is necessary to define a clear format for DEBUG/TRACE-level logs, rather than letting developers play freely;
- The entire team (including development, operation and maintenance and testing) regularly reviews the recorded log content;
- Develop do operation and maintenance, and optimize the way of logging through the process of checking problems;
- The problems found in the log during operation and maintenance or testing must be reported to the developer in time;

## 6. Log Monitoring

By monitoring the keywords in the log, system failures can be discovered and alarmed in time, which is essential to ensure the SLA of the service.

Service monitoring and alarming is a big topic. Here are just some issues that need to be paid attention to in log monitoring and alarming:

1. The alarms will not be sent if they can. Only the error that needs to be handled immediately by the operation and maintenance needs to be sent to the alarm. The reason for this is to avoid long-term alarm harassment so that the operation and maintenance personnel are no longer sensitive to the alarm. In the end, when the real alarm comes, it becomes the legend of the wolf coming;
2. Clear alarm keywords, for example, use ERROR as the alarm keyword instead of various complicated rules. The reason for this is that log monitoring is essentially a continuous string matching operation. If the rules are too many and too complicated, it may have an impact on online services;
3. For some early warning operations, such as a certain service that needs to be retried many times to succeed, or a user’s quota is almost used up, etc., you can provide feedback through an alarm email every day;
4. Every time the system fails, you need to check whether the log alarm is sensitive, whether the description of the log alarm is accurate, etc., and continuously optimize the log alarm;

## 7. Other considerations

**Log observation after going online**

After each deployment, in addition to the complete regression test of the system, it is also necessary to observe the log. Especially when a new feature is launched, the log can be used to confirm whether the new feature is working properly.

**Log output to different files**

Another problem encountered during performance testing is that when the amount of concurrency is large, some request processing may fail (such as 0.5%). In order to analyze these errors, you need to check the logs of these wrong requests. And because of the huge amount of logs in this case, it becomes difficult to analyze the error log.

In this case, all error logs can be output to a single file at the same time.

**Log file size**

The log file should not be too large. If a log file is too large, it will bring inconvenience to log monitoring and problem location. Therefore, log files need to be segmented. Whether log files should be divided by day or hour should be determined according to the log volume. The principle is to facilitate the development or operation and maintenance personnel to quickly find the logs.

In order to prevent log files from filling up the entire disk space, log files need to be deleted regularly. For example, when you receive a disk alarm, you can delete the log files two months ago. The best practice here is:

- Collect all log files so that even if they are deleted on the machine that records the logs, the previous problems can be located through the collected logs;
- Delete expired logs through scheduled tasks every day, such as deleting logs 60 days ago at 4 am every day



## 8. Summary

A summary of all the suggestions made in the article is as follows:

- Fully recognize the key role of logs for a reliable back-end system
- The entire team (including operation and maintenance personnel) needs to have clear regulations on the log level, what log output is at what level, what level of error occurs and how to deal with it
- The log content needs to be optimized and updated regularly, the purpose is to locate the problem quickly and accurately through the log
- To clarify the purpose of different logs, classify the content of the logs
- Never print useless logs to prevent them from overwhelming important information
- The log information should be accurate and comprehensive, and strive to locate the problem with the log alone
- The optimization of the log is a matter that requires continuous investment and continuous learning from mistakes
- Generate RequestID according to different purposes, and encode as much information as possible in RequestID when necessary
- Associate the entire processing flow of a request with a unique RequestID
- Support dynamic log output to facilitate online problem location
- Be sure to observe the log after the new server is online. In particular, developers can confirm whether the new function is working properly by observing the log
- Discover potential problems by increasing the log level
- Monitor and alarm the log, find system problems before customers
- Determine the operating status of the system through keywords in the log
- The log format should be standardized
- Output the error log to a separate file for analysis
- The size of the log, how to split, how to delete, etc. should be established as specifications

## References:

https://studygolang.com/articles/10321

https://stackoverflow.com/questions/2031163/when-to-use-the-different-log-levels

