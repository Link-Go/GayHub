* fatal: could not read Username for 'https://github.com': terminal prompts disabled   获取私有仓库
    * go env -w GOPRIVATE=github.com/mycompany/*

* git clone 
    * fatal: could not read Username for 'https://git.duowan.com': terminal prompts disabled
    * export GIT_TERMINAL_PROMPT=1

* git 存储账号密码
    * ~/.git-credentials  文件路径
    * git config --global credential.helper store
    * 调用一次 git push，输入账号密码，git就会自动生成一个.git-credentials文件，然后将本次的账号密码信息保存进文件中，后续操作不再需要输入账号密码