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
- ~~dao 패키지를 만들어야 할까?~~
  - 규모가 크고 복잡하면 모를까 service 에서 대신 처리하는 방향으로 진행 
- provider 는 entity 를 알아야 할까?
  - 각 db 의 결과를 service 가 모두 알아야 하는 문제가 발생
- ~~service 패키지를 어느 시점에 어디에서 주입해야 할까?~~
  - handler 에서 주입하는 방향으로 진행
 