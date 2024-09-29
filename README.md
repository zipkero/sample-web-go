## Sample Web Application

### 1. Gin Web Framework 사용 예제
### 2. Redis, Mongodb 사용 예제
### 3. 패키지 역할
- handler: http request handler
- entity: db schema
- dto: data transfer object
- db: db client provider
- service: business logic
- middleware: gin middleware
- router: route 설정
- config: 환경변수 설정
- server: server 실행

### 고민거리
- dao 패키지를 만들어야 할까?
- service 패키지를 어느 시점에 어디에서 주입해야 할까?