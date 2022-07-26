
## 错误处理



### 错误码

- 常见的错误码设计方式
  - 1.不论请求成功或失败，始终返回200 http status code，在 HTTP Body 中包含所有的错误信息
  ```json
    {
        "error": {
            "message": "Syntax error \"Field picture specified more than once. This is only possible before version 2.1\" at character 23: id,name,picture,picture",
            "type": "OAuthException",
            "code": 2500,
            "fbtrace_id": "xxxxxxxxxxx"
        }
    }
  ```

  
  
  - 2.根据不同的请求返回 http 错误码，并在 Body 中返回简单的错误信息
  ```http
  HTTP/1.1 400 Bad Request
  x-connection-hash: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  set-cookie: guest_id=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
  Date: Thu, 01 Jun 2017 03:04:23 GMT
  Content-Length: 62
  x-response-time: 5
  strict-transport-security: max-age=631138519
  Connection: keep-alive
  Content-Type: application/json; charset=utf-8
  Server: tsa_b
  
  {"errors":[{"code":215,"message":"Bad Authentication data."}]}
  ```
  
  
  
  - 3.根据不同的请求返回 http 错误码，并在 Body 中返回详细的错误信息（建议使用）
  ```http
  HTTP/1.1 400
  Date: Thu, 01 Jun 2017 03:40:55 GMT
  Content-Length: 276
  Connection: keep-alive
  Content-Type: application/json; charset=utf-8
  Server: Microsoft-IIS/10.0
  X-Content-Type-Options: nosniff
  
    {
        "SearchResponse": {
            "Version": "2.2",
            "Query": {
                "SearchTerms": "api error codes"
            },
            "Errors": [
                {
                    "Code": 1001,
                    "Message": "Required parameter is missing.",
                    "Parameter": "SearchRequest.AppId",
                    "HelpUrl": "http\u003a\u002f\u002fmsdn.microsoft.com\u002fen-us\u002flibrary\u002fdd251042.aspx"
                }
            ]
        }
    }
  ```



- 错误码设计建议
    - 1.有区别于http status code的业务码，业务码需要有一定规则，可以通过业务码判断出是哪类错误
    - 2.请求出错时，可以通过http status code直接感知到请求出错
    - 3.需要在请求出错时，返回详细的信息，通常包括 3 类信息：业务 Code 码、错误信息和参考文档（可选）
    - 4.返回的错误信息，需要是可以直接展示给用户的安全信息，也就是说不能包含敏感信息；同时也要有内部更详细的错误信息，方便 debug。
    - 5.返回的数据格式应该是固定的、规范的、简洁的、有用的



- 业务 Code 码设计
    - **Code 码设计规范：纯数字表示，不同部位代表不同的服务，不同的模块**
    - 错误代码说明：100101
        - **10**：服务
        - **01**：某个服务下面的某个模块
        - **01**：模块下的错误码序号
    - 按照这个设计，每个类别最多100个实例；超过100个实例，证明模块或服务过于庞大，建议拆分；实在无法错误，可以使用3位的错误码
    - **10** 设计为通用服务，即所有模块都可以使用，再细化子模块，避免重复造轮子
    
        | 服务 | 模块 | 说明（服务 - 模块）          |
        | ---- | ---- | ---------------------------- |
        | 10   | 00   | 通用 - 基本错误              |
        | 10   | 01   | 通用 - 数据库类错误          |
        | 11   | 00   | apiserver服务 - 用户模块错误 |
        | 11   | 01   | apiserver服务 - 密钥模块错误 |
    
        
    
- HTTP Status Code 设置
    - Go net/http 包提供了 60 个错误码，大致分为如下 5 类：
        - 1XX - （指示信息）表示请求已接收，继续处理。（没用过）
        - 2XX - （请求成功）表示成功处理了请求的状态代码。
        - 3XX - （请求被重定向）表示要完成请求，需要进一步操作。通常，这些状态代码用来重定向。（就用过302）
        - 4XX - （请求错误）这些状态代码表示请求可能出错，妨碍了服务器的处理，通常是客户端出错，需要客户端做进一步的处理。
        - 5XX - （服务器错误）这些状态代码表示服务器在尝试处理请求时发生内部错误。这些错误可能是服务器本身的错误，而不是客户端的问题。
    - 需要的http status code
        - 200 - 表示请求成功执行（201，204统一为200）
        - 400 - 表示客户端出问题（4xx的错误用的比较多，通过 HTTP Status Code 进一步区分的话，可以适当扩充，也可以统一定义为400，之后在业务 code 中再进行区分）
            - 401 - 表示认证失败
            - 403 - 表示授权失败
            - 404 - 表示资源找不到，这里的资源可以是 URL 或者 RESTful 资源
            - 405 - 方法不允许
            - 422 - 请求格式错误，语义错误
        - 500 - 表示服务端出问题



- response 返回值
    ```json
        {
            "code": 100001,
            "message": "error message",
            "data": {
                "field1": "xxxx1",
                "field2": "xxxx2",
            }
        }
    ```



- 错误码定义参考

  | Identifier                  | Code   | HTTP code | Description                                                 |
  | --------------------------- | ------ | --------- | ----------------------------------------------------------- |
  | ErrSuccess                  | 100001 | 200       | OK                                                          |
  | ErrUnknown                  | 100002 | 500       | Internal server error                                       |
  | ErrBind                     | 100003 | 400       | Error occurred while binding the request body to the struct |
  | ErrValidation               | 100004 | 400       | Validation failed                                           |
  | ErrTokenInvalid             | 100005 | 401       | Token invalid                                               |
  | ErrPageNotFound             | 100006 | 404       | Page not found                                              |
  | ErrDatabase                 | 100101 | 500       | Database error                                              |
  | ErrEncrypt                  | 100201 | 401       | Error occurred while encrypting the user password           |
  | ErrSignatureInvaild         | 100202 | 401       | Signature is invalid                                        |
  | ErrTokenExpired             | 100203 | 401       | Token expired                                               |
  | ErrInvalidAuthHeader        | 100204 | 401       | Invaild authorization header                                |
  | ErrMissingHeader            | 100205 | 401       | The Authorization header was empty                          |
  | ErrAccountPasswordIncorrect | 100206 | 401       | Account or password was incorrect                           |
  | ErrPermissionDenied         | 100207 | 403       | Permission denied                                           |
  | ErrSecretNotFound           | 110102 | 404       | Secret not found                                            |
  | ErrPolicyNotFound           | 110201 | 404       | User not found                                              |

  

### 错误包

- [**错误包**](https://github.com/pkg/errors): git 上最受欢迎的golang错误包
    - 支持错误堆栈
    - 支持 Wrap/Unwrap 功能
- [**二次封装添加了业务code的错误包**](https://github.com/marmotedu/errors)