## Govulncheck

Go漏洞数据库的数据库客户端和工具

默认情况下 Govulncheck 向位于 https://vuln.go.dev/ 的漏洞数据库发送请求（离线无法使用）
不建议使用 `x/vulndb` 项目中的文件（https://go.dev/security/vuln/database）

更多详细内容：https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck

#### 用法
```bash
# 需要 1.18 版本及以上

go install golang.org/x/vuln/cmd/govulncheck@latest
cd /your-project
govulncheck ./...

```

#### 信息
```txt
govulncheck is an experimental tool. Share feedback at https://go.dev/s/govulncheck-feedback.

Using go1.19.4 and govulncheck@v0.0.0 with
vulnerability data from https://vuln.go.dev (last modified 2023-04-18 21:32:26 +0000 UTC).

Scanning your code and 236 packages across 7 dependent modules for known vulnerabilities...
Your code is affected by 4 vulnerabilities from the Go standard library.

Vulnerability #1: GO-2023-1705
  Multipart form parsing can consume large amounts of CPU and
  memory when processing form inputs containing very large numbers      
  of parts. This stems from several causes: 1.
  mime/multipart.Reader.ReadForm limits the total memory a parsed       
  multipart form can consume. ReadForm can undercount the amount        
  of memory consumed, leading it to accept larger inputs than
  intended. 2. Limiting total memory does not account for
  increased pressure on the garbage collector from large numbers        
  of small allocations in forms with many parts. 3. ReadForm can        
  allocate a large number of short-lived buffers, further
  increasing pressure on the garbage collector. The combination of      
  these factors can permit an attacker to cause an program that
  parses multipart forms to consume large amounts of CPU and
  memory, potentially resulting in a denial of service. This
  affects programs that use mime/multipart.Reader.ReadForm, as
  well as form parsing in the net/http package with the Request
  methods FormFile, FormValue, ParseMultipartForm, and
  PostFormValue. With fix, ReadForm now does a better job of
  estimating the memory consumption of parsed forms, and performs
  many fewer short-lived allocations. In addition, the fixed
  mime/multipart.Reader imposes the following limits on the size
  of parsed forms: 1. Forms parsed with ReadForm may contain no
  more than 1000 parts. This limit may be adjusted with the
  environment variable GODEBUG=multipartmaxparts=. 2. Form parts
  parsed with NextPart and NextRawPart may contain no more than
  10,000 header fields. In addition, forms parsed with ReadForm
  may contain no more than 10,000 header fields across all parts.
  This limit may be adjusted with the environment variable
  GODEBUG=multipartmaxheaders=.

  More info: https://pkg.go.dev/vuln/GO-2023-1705

  Standard library
    Found in: net/textproto@go1.19.4
    Fixed in: net/textproto@go1.20.3

    Call stacks in your code:
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls net/textproto.Reader.ReadMIMEHeader

Vulnerability #2: GO-2023-1704
  HTTP and MIME header parsing can allocate large amounts of
  memory, even when parsing small inputs, potentially leading to a
  denial of service. Certain unusual patterns of input data can
  cause the common function used to parse HTTP and MIME headers to
  allocate substantially more memory than required to hold the
  parsed headers. An attacker can exploit this behavior to cause
  an HTTP server to allocate large amounts of memory from a small
  request, potentially leading to memory exhaustion and a denial
  of service. With fix, header parsing now correctly allocates
  only the memory required to hold parsed headers.

  More info: https://pkg.go.dev/vuln/GO-2023-1704

  Standard library
    Found in: net/textproto@go1.19.4
    Fixed in: net/textproto@go1.20.3

    Call stacks in your code:
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls net/textproto.Reader.ReadMIMEHeader

Vulnerability #3: GO-2023-1621
  The ScalarMult and ScalarBaseMult methods of the P256 Curve may
  return an incorrect result if called with some specific
  unreduced scalars (a scalar larger than the order of the curve).
  This does not impact usages of crypto/ecdsa or crypto/ecdh.

  More info: https://pkg.go.dev/vuln/GO-2023-1621

  Standard library
    Found in: crypto/internal/nistec@go1.19.4
    Fixed in: crypto/internal/nistec@go1.20.2

    Call stacks in your code:
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls crypto/internal/nistec.P256OrdInverse
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls crypto/internal/nistec.P256Point.ScalarBaseMult
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls crypto/internal/nistec.P256Point.ScalarMult

Vulnerability #4: GO-2023-1570
  Large handshake records may cause panics in crypto/tls. Both
  clients and servers may send large TLS handshake records which
  cause servers and clients, respectively, to panic when
  attempting to construct responses. This affects all TLS 1.3
  clients, TLS 1.2 clients which explicitly enable session
  resumption (by setting Config.ClientSessionCache to a non-nil
  value), and TLS 1.3 servers which request client certificates
  (by setting Config.ClientAuth >= RequestClientCert).

  More info: https://pkg.go.dev/vuln/GO-2023-1570

  Standard library
    Found in: crypto/tls@go1.19.4
    Fixed in: crypto/tls@go1.20.1

    Call stacks in your code:
      main.go:58:27: client.main calls client/example.calculatorComputeAverageClient.Recv, which eventually calls crypto/tls.Conn.Read
      main.go:63:13: client.main calls fmt.Printf, which eventually calls crypto/tls.Conn.Write

=== Informational ===

Found 4 vulnerabilities in packages that you import, but there are no call
stacks leading to the use of these vulnerabilities. You may not need to
take any action. See https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck
for details.

Vulnerability #1: GO-2023-1703
  Templates do not properly consider backticks (`) as Javascript
  string delimiters, and do not escape them as expected. Backticks
  are used, since ES6, for JS template literals. If a template
  contains a Go template action within a Javascript template
  literal, the contents of the action can be used to terminate the
  literal, injecting arbitrary Javascript code into the Go
  template. As ES6 template literals are rather complex, and
  themselves can do string interpolation, the decision was made to
  simply disallow Go template actions from being used inside of
  them (e.g. "var a = {{.}}"), since there is no obviously safe
  way to allow this behavior. This takes the same approach as
  github.com/google/safehtml. With fix, Template.Parse returns an
  Error when it encounters templates like this, with an ErrorCode
  of value 12. This ErrorCode is currently unexported, but will be
  exported in the release of Go 1.21. Users who rely on the
  previous behavior can re-enable it using the GODEBUG flag
  jstmpllitinterp=1, with the caveat that backticks will now be
  escaped. This should be used with caution.
  More info: https://pkg.go.dev/vuln/GO-2023-1703
  Found in: html/template@go1.19.4
  Fixed in: html/template@go1.20.3

Vulnerability #2: GO-2023-1571
  A maliciously crafted HTTP/2 stream could cause excessive CPU
  consumption in the HPACK decoder, sufficient to cause a denial
  of service from a small number of small requests.
  More info: https://pkg.go.dev/vuln/GO-2023-1571
  Found in: net/http@go1.19.4
  Fixed in: net/http@go1.20.1

Vulnerability #3: GO-2023-1569
  A denial of service is possible from excessive resource
  consumption in net/http and mime/multipart. Multipart form
  parsing with mime/multipart.Reader.ReadForm can consume largely
  unlimited amounts of memory and disk files. This also affects
  form parsing in the net/http package with the Request methods
  FormFile, FormValue, ParseMultipartForm, and PostFormValue.
  ReadForm takes a maxMemory parameter, and is documented as
  storing "up to maxMemory bytes +10MB (reserved for non-file
  parts) in memory". File parts which cannot be stored in memory
  are stored on disk in temporary files. The unconfigurable 10MB
  reserved for non-file parts is excessively large and can
  potentially open a denial of service vector on its own. However,
  ReadForm did not properly account for all memory consumed by a
  parsed form, such as map entry overhead, part names, and MIME
  headers, permitting a maliciously crafted form to consume well
  over 10MB. In addition, ReadForm contained no limit on the
  number of disk files created, permitting a relatively small
  request body to create a large number of disk temporary files.
  With fix, ReadForm now properly accounts for various forms of
  memory overhead, and should now stay within its documented limit
  of 10MB + maxMemory bytes of memory consumption. Users should
  still be aware that this limit is high and may still be
  hazardous. In addition, ReadForm now creates at most one on-disk
  temporary file, combining multiple form parts into a single
  temporary file. The mime/multipart.File interface type's
  documentation states, "If stored on disk, the File's underlying
  concrete type will be an *os.File.". This is no longer the case
  when a form contains more than one file part, due to this
  coalescing of parts into a single file. The previous behavior of
  using distinct files for each form part may be reenabled with
  the environment variable GODEBUG=multipartfiles=distinct. Users
  should be aware that multipart.ReadForm and the http.Request
  methods that call it do not limit the amount of disk consumed by
  temporary files. Callers can limit the size of form data with
  http.MaxBytesReader.
  More info: https://pkg.go.dev/vuln/GO-2023-1569
  Found in: mime/multipart@go1.19.4
  Fixed in: mime/multipart@go1.20.1

Vulnerability #4: GO-2023-1568
  A path traversal vulnerability exists in filepath.Clean on
  Windows. On Windows, the filepath.Clean function could transform
  an invalid path such as "a/../c:/b" into the valid path "c:\b".
  This transformation of a relative (if invalid) path into an
  absolute path could enable a directory traversal attack. After
  fix, the filepath.Clean function transforms this path into the
  relative (but still invalid) path ".\c:\b".
  More info: https://pkg.go.dev/vuln/GO-2023-1568
  Found in: path/filepath@go1.19.4
  Fixed in: path/filepath@go1.20.1
  Platforms: windows
```

以 ``=== Informational ===`` 为分界线
* 上：部分引用的包有漏洞，直接按照要求升级新版本即可
    * 特殊情况，部分包多级依赖，并非单独升级可以解决，需要慎重分析
* 下：部分引用的包有漏洞，但是你的项目代码没有直接引用，无需关注
