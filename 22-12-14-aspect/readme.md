## 出现报错，但是不加@Aspect注解就不会走到AOP切面
```text
Error starting ApplicationContext. To display the conditions report re-run your application with 'debug' enabled.
2022-12-14 21:35:54.796 ERROR 78709 --- [           main] o.s.boot.SpringApplication               : Application run failed

org.springframework.context.ApplicationContextException: Unable to start web server; nested exception is org.springframework.beans.factory.BeanCreationException: Error creating bean with name 'org.springframework.boot.autoconfigure.web.servlet.ServletWebServerFactoryConfiguration$EmbeddedTomcat': Initialization of bean failed; nested exception is java.lang.IllegalArgumentException: error Type referred to is not an annotation type: com$example$demo$aop
	at org.springframework.boot.web.servlet.context.ServletWebServerApplicationContext.onRefresh(ServletWebServerApplicationContext.java:165) ~[spring-boot-2.7.6.jar:2.7.6]
	at org.springframework.context.support.AbstractApplicationContext.refresh(AbstractApplicationContext.java:577) ~[spring-context-5.3.24.jar:5.3.24]
	at org.springframework.boot.web.servlet.context.ServletWebServerApplicationContext.refresh(ServletWebServerApplicationContext.java:147) ~[spring-boot-2.7.6.jar:2.7.6]
	at org.springframework.boot.SpringApplication.refresh(SpringApplication.java:731) ~[spring-boot-2.7.6.jar:2.7.6]
	at org.springframework.boot.SpringApplication.refreshContext(SpringApplication.java:408) ~[spring-boot-2.7.6.jar:2.7.6]
	at org.springframework.boot.SpringApplication.run(SpringApplication.java:307) ~[spring-boot-2.7.6.jar:2.7.6]
	at com.example.demo.Application.main(Application.java:14) ~[classes/:na]
```
