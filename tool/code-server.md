## code-server
VS Code in the browser
通过浏览器启动 VS Code


**[code-server](https://github.com/coder/code-server)**
* [Documentation](https://coder.com/docs/code-server/latest)
* 简单使用说明
    * 安装
        * 脚本安装（最方便）
        ```shell
        预览：curl -fsSL https://code-server.dev/install.sh | sh -s -- --dry-run
        实际安装：curl -fsSL https://code-server.dev/install.sh | sh

        脚本内容：

        Installing v4.6.0 of the amd64 rpm package from GitHub.

        + mkdir -p ~/.cache/code-server
        + curl -#fL -o ~/.cache/code-server/code-server-4.6.0-amd64.rpm.incomplete -C - https://github.com/coder/code-server/releases/download/v4.6.0/code-server-4.6.0-amd64.rpm
        ######################################################################## 100.0%
        + mv ~/.cache/code-server/code-server-4.6.0-amd64.rpm.incomplete ~/.cache/code-server/code-server-4.6.0-amd64.rpm
        + rpm -i ~/.cache/code-server/code-server-4.6.0-amd64.rpm

        rpm package has been installed.

        To have systemd start code-server now and restart on boot:
        sudo systemctl enable --now code-server@$USER
        Or, if you don't want/need a background service you can run:
        code-server

        ```

        * 安装包，路径：https://github.com/coder/code-server/releases

    * 启动
        ```shell
        code-server
    
        // ~/.config/code-server/config.yaml 配置文件
        [2022-08-24T02:58:50.784Z] info  Wrote default config file to ~/.config/code-server/config.yaml
        [2022-08-24T02:58:51.190Z] info  code-server 4.6.0 6d3f9ca6a6df30a1bfad6f073f6fa33c0e63abdb
        [2022-08-24T02:58:51.191Z] info  Using user-data-dir ~/.local/share/code-server
        [2022-08-24T02:58:51.201Z] info  Using config file ~/.config/code-server/config.yaml
        [2022-08-24T02:58:51.202Z] info  HTTP server listening on http://127.0.0.1:8080/
        [2022-08-24T02:58:51.202Z] info    - Authentication is enabled
        [2022-08-24T02:58:51.202Z] info      - Using password from ~/.config/code-server/config.yaml
        [2022-08-24T02:58:51.202Z] info    - Not serving HTTPS

        ```

    * 浏览器访问： http://127.0.0.1:8080；之后的操作就跟在本机使用 vs code 一样

    * 使用docker启动
        ```shell
        docker pull linuxserver/code-server
        ```